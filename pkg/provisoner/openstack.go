package provisoner

import (
	"bufio"
	"encoding/json"
	"fmt"
	ostack "github.com/drewbernetes/baski/pkg/providers/openstack"
	"github.com/drewbernetes/baski/pkg/providers/packer"
	"github.com/drewbernetes/baski/pkg/providers/scanner"
	"github.com/drewbernetes/baski/pkg/trivy"
	"github.com/drewbernetes/baski/pkg/util/data"
	"github.com/drewbernetes/baski/pkg/util/flags"
	"github.com/drewbernetes/baski/pkg/util/interfaces"
	"github.com/drewbernetes/baski/pkg/util/sign"
	simple_s3 "github.com/drewbernetes/simple-s3"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

// OpenStackBuildProvisioner contains the options for the build.
type OpenStackBuildProvisioner struct {
	Opts *flags.BuildOptions
}

// newOpenStackBuilder returns a new OpenStackBuildProvisioner.
func newOpenStackBuilder(o *flags.BuildOptions) *OpenStackBuildProvisioner {
	p := &OpenStackBuildProvisioner{
		Opts: o,
	}

	return p
}

// Init will set an ENV VAR so that the OpenStack builder knows which cloud to use.
func (p *OpenStackBuildProvisioner) Init() error {
	err := os.Setenv("OS_CLOUD", p.Opts.OpenStackFlags.CloudName)
	if err != nil {
		return err
	}
	return nil
}

// GeneratePackerConfig generates a packer vars file for OpenStack builds.
func (p *OpenStackBuildProvisioner) GeneratePackerConfig() (*packer.GlobalBuildConfig, error) {
	o := p.Opts
	b, imgName, err := packer.NewCoreBuildconfig(o)
	if err != nil {
		return nil, err
	}

	b.OpenStackBuildconfig = packer.OpenStackBuildconfig{
		AttachConfigDrive:     strconv.FormatBool(o.OpenStackFlags.AttachConfigDrive),
		Flavor:                o.OpenStackFlags.FlavorName,
		FloatingIpNetwork:     o.OpenStackFlags.FloatingIPNetworkName,
		ImageDiskFormat:       o.OpenStackFlags.ImageDiskFormat,
		ImageVisibility:       o.OpenStackFlags.ImageVisibility,
		ImageName:             imgName,
		Networks:              o.OpenStackFlags.NetworkID,
		SecurityGroup:         o.OpenStackFlags.SecurityGroup,
		SourceImage:           o.OpenStackFlags.SourceImageID,
		UseBlockStorageVolume: o.OpenStackFlags.UseBlockStorageVolume,
		UseFloatingIp:         strconv.FormatBool(o.OpenStackFlags.UseFloatingIP),
		VolumeType:            o.OpenStackFlags.VolumeType,
		VolumeSize:            strconv.Itoa(o.OpenStackFlags.VolumeSize),
	}

	if len(o.OpenStackFlags.SSHPrivateKeyFile) > 0 && len(o.OpenStackFlags.SSHKeypairName) > 0 {
		b.OpenStackBuildconfig.SSHPrivateKeyFile = o.OpenStackFlags.SSHPrivateKeyFile
		b.OpenStackBuildconfig.SSHKeypairName = o.OpenStackFlags.SSHKeypairName
	}

	b.Metadata = generateBuilderMetadata(o)

	return b, nil
}

// UpdatePackerBuilders will update the builders field with the metadata values as required. This is done this way as passing it in via Packer vars is prone to error or just complete failures.
func (p *OpenStackBuildProvisioner) UpdatePackerBuilders(metadata map[string]string, data []byte) []byte {
	jsonStruct := struct {
		Builders       []map[string]interface{} `json:"builders"`
		PostProcessors []map[string]interface{} `json:"post-processors"`
		Provisioners   []map[string]interface{} `json:"provisioners"`
		Variables      map[string]interface{}   `json:"variables"`
	}{}

	err := json.Unmarshal(data, &jsonStruct)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	jsonStruct.Builders[0]["metadata"] = metadata

	res, err := json.Marshal(jsonStruct)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return res
}

// PostBuildAction retrieves the image ID from the output and stores it into a file.
func (p *OpenStackBuildProvisioner) PostBuildAction() error {

	imgID, err := retrieveNewOpenStackImageID()
	if err != nil {
		return err
	}

	err = saveImageIDToFile(imgID)
	if err != nil {
		return err
	}

	return nil
}

// retrieveNewOpenStackImageID identifies the new ImageID from the output text so that it can be used/retrieved later.
func retrieveNewOpenStackImageID() (string, error) {
	var i string
	//TODO: The output goes to STDOUT during buildImage,
	// we need to figure out if we can pull this from the openstack instance instead
	// to remove the requirement of parsing STDOUT.
	f, err := os.Open("/tmp/out-build.txt")
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewScanner(f)
	re := regexp.MustCompile("An image was created: [0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}")
	for r.Scan() {
		m := re.MatchString(string(r.Bytes()))
		if m {
			//There is likely two outputs here due to how packer outputs, so we need to break on the first find.
			i = strings.Split(r.Text(), ": ")[2]
			break
		}
	}

	return i, nil
}

// OpenStackScanProvisioner contains the parameters required for scanning images.
type OpenStackScanProvisioner struct {
	imageClient           interfaces.OpenStackImageClient
	computeClient         interfaces.OpenStackComputeClient
	networkClient         interfaces.OpenStackNetworkClient
	cloudName             string
	flavorName            string
	floatingIPNetworkName string
	metadataPrefix        string
	networkID             string
	attachConfigDrive     bool
	autoDeleteImage       bool
	skipCVECheck          bool
	maxSeverity           string
	imageVisibility       string
	securityGroup         string
	imageID               string
	imageWildCard         string
	scanConcurrency       int
	trivyignorePath       string
	trivyignoreFilename   string
	trivyignoreList       []string
	s3Endpoint            string
	s3AccessKey           string
	s3SecretKey           string
	s3Region              string
	s3Bucket              string
}

// newOpenStackScanner returns a new instance of OpenStackScanProvisioner
func newOpenStackScanner(o *flags.ScanOptions) *OpenStackScanProvisioner {
	var flavorName string
	if o.ScanFlavorName != "" {
		flavorName = o.ScanFlavorName
	} else {
		flavorName = o.OpenStackFlags.OpenStackInstanceFlags.FlavorName
	}

	concurrency := 1
	if o.Concurrency != 0 {
		concurrency = o.Concurrency
	}

	p := &OpenStackScanProvisioner{
		cloudName:             o.OpenStackFlags.OpenStackCoreFlags.CloudName,
		flavorName:            flavorName,
		floatingIPNetworkName: o.OpenStackFlags.FloatingIPNetworkName,
		networkID:             o.OpenStackInstanceFlags.NetworkID,
		attachConfigDrive:     o.OpenStackInstanceFlags.AttachConfigDrive,
		metadataPrefix:        o.OpenStackCoreFlags.MetadataPrefix,
		autoDeleteImage:       o.AutoDeleteImage,
		skipCVECheck:          o.SkipCVECheck,
		maxSeverity:           o.MaxSeverityType,
		imageVisibility:       o.OpenStackFlags.ImageVisibility,
		securityGroup:         o.OpenStackFlags.SecurityGroup,
		scanConcurrency:       concurrency,
		s3Endpoint:            o.S3Flags.Endpoint,
		s3AccessKey:           o.S3Flags.AccessKey,
		s3SecretKey:           o.S3Flags.SecretKey,
		s3Region:              o.S3Flags.Region,
		s3Bucket:              o.ScanBucket,
		trivyignorePath:       o.TrivyignorePath,
		trivyignoreFilename:   o.TrivyignoreFilename,
		trivyignoreList:       o.TrivyignoreList,
	}

	if o.ScanSingleOptions.ImageID != "" {
		p.imageID = o.ScanSingleOptions.ImageID
	}
	if o.ScanMultipleOptions.ImageSearch != "" {
		p.imageWildCard = o.ScanMultipleOptions.ImageSearch
	}

	return p
}

// Prepare the requirements for scanning images. this includes setting up the OpenStack clients so that communication with OpenStack is successful.
func (s *OpenStackScanProvisioner) Prepare() error {
	var err error

	cloudProvider := ostack.NewCloudsProvider(s.cloudName)

	s.imageClient, err = ostack.NewImageClient(cloudProvider)
	if err != nil {
		return err
	}

	s.computeClient, err = ostack.NewComputeClient(cloudProvider)
	if err != nil {
		return err
	}

	s.networkClient, err = ostack.NewNetworkClient(cloudProvider)
	if err != nil {
		return err
	}

	return nil
}

// ScanImages will parse the input of imageID or wildcard to determine whether a single or multiple scan needs to occur.
// The image is then scanned and the results uploaded to S3.
func (s *OpenStackScanProvisioner) ScanImages() error {
	var err error

	imgs := []images.Image{}

	// Parse the image ID or wildcard and load the images from OpenStack
	if s.imageID != "" {
		var img *images.Image

		img, err = s.imageClient.FetchImage(s.imageID)
		if err != nil {
			return err
		}

		imgs = append(imgs, *img)
	} else if s.imageWildCard != "" {
		imgs, err = s.imageClient.FetchAllImages(s.imageWildCard)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no scan image ID or WILDCARD specified")
	}

	severity := trivy.Severity(strings.ToUpper(s.maxSeverity))

	var s3Conn *simple_s3.S3

	s3Conn, err = simple_s3.New(s.s3Endpoint, s.s3AccessKey, s.s3SecretKey, s.s3Bucket, s.s3Region)
	if err != nil {
		log.Println(err)
		return err
	}

	// Let's scan a bunch of images based on the concurrency

	// Error collection channel
	errChan := make(chan error, len(imgs))
	imageChan := make(chan images.Image)

	// Worker Pool
	var wg sync.WaitGroup
	wg.Add(s.scanConcurrency)

	for i := 0; i < s.scanConcurrency; i++ {
		go func() {
			defer wg.Done()

			for img := range imageChan {
				sc := scanner.NewOpenStackScanner(s.computeClient, s.imageClient, s.networkClient, s3Conn, severity, &img)
				err = s.scanServer(sc, &wg)

				if err != nil {
					errChan <- fmt.Errorf("failed to scan image %s with ID %s. error: %s", img.Name, img.ID, err.Error())
				}
			}
		}()
	}

	go func() {
		for _, img := range imgs {
			wg.Add(1)
			imageChan <- img
		}
		close(imageChan)
	}()
	wg.Wait()
	close(errChan)

	//Collect Errors
	var errs []error
	for err := range errChan {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return fmt.Errorf("encounted errors during image scanning: %v", errs)
	}

	return nil
}

// scanServer will scan, parse the results and upload them to S3. It's in its own function for the purpose of threading.
func (s *OpenStackScanProvisioner) scanServer(sc *scanner.OpenStackScannerClient, wg *sync.WaitGroup) error {
	defer wg.Done()

	log.Printf("Processing Image with ID: %s\n", sc.Img.ID)

	// Run the scan.
	err := sc.RunScan(s.trivyignorePath, s.trivyignoreFilename, s.trivyignoreList, s.floatingIPNetworkName, s.flavorName, s.networkID, s.securityGroup, s.attachConfigDrive)
	if err != nil {
		return err
	}

	// Fetch the results and write them to a file locally.
	err = sc.FetchScanResults()
	if err != nil {
		return err
	}

	// Read the local results file and parse them into a more consumable json format, then write out to file.
	err = sc.CheckResults()
	if err != nil {
		return err
	}

	// Check if a visibility option has been supplied.
	// If so, we'll leave things as they are otherwise we'll control it.
	if s.imageVisibility == "" {
		// If the image has no vulnerabilities, we can set the image to public otherwise default to private.
		visibility := images.ImageVisibilityPrivate
		if len(sc.Vulns) == 0 {
			visibility = images.ImageVisibilityPublic
		}

		err = s.imageClient.ChangeImageVisibility(sc.Img.ID, visibility)
		if err != nil {
			return err
		}
	}

	// If the image is not set to auto delete, tag the image with the check result.
	if !s.autoDeleteImage {
		err = sc.TagImage(s.metadataPrefix)
		if err != nil {
			return err
		}
	} else {
		if len(sc.Vulns) != 0 {
			// Remove the image if vulns are found and the flag/config item is set.
			err = s.imageClient.RemoveImage(sc.Img.ID)
			if err != nil {
				return err
			}
		}
	}

	// Upload the parsed results file to S3
	err = sc.UploadResultsToS3()
	if err != nil {
		return err
	}

	log.Printf("Finished processing Image ID: %s\n", sc.Img.ID)

	// Check if the CVE checking is being skipped, if not then bail out here.
	if !s.skipCVECheck {
		errMsg := "vulnerabilities detected above threshold. Please see the possible fixes located at '/tmp/results.json' for further information on this"
		if s.autoDeleteImage {
			errMsg = fmt.Sprintf("%s - %s", errMsg, ". The image has been removed from the infra.")
		}
		return fmt.Errorf("%s", errMsg)
	}
	return nil
}

// OpenStackSignProvisioner contains the parameters required for signing images.
type OpenStackSignProvisioner struct {
	cloudName      string
	imageID        string
	metadataPrefix string
}

// newOpenStackSigner returns a new instance of OpenStackSignProvisioner.
func newOpenStackSigner(o *flags.SignOptions) *OpenStackSignProvisioner {
	p := &OpenStackSignProvisioner{
		cloudName:      o.OpenStackCoreFlags.CloudName,
		imageID:        o.ImageID,
		metadataPrefix: o.OpenStackCoreFlags.MetadataPrefix,
	}

	return p
}

// SignImage will take the digest of the signing process and tag the image with the appropriate metadata field.
func (s *OpenStackSignProvisioner) SignImage(digest string) error {
	cloudProvider := ostack.NewCloudsProvider(s.cloudName)

	i, err := ostack.NewImageClient(cloudProvider)
	if err != nil {
		return err
	}

	img, err := i.FetchImage(s.imageID)
	if err != nil {
		return err
	}

	digestPropertyName := "digest"
	if s.metadataPrefix != "" {
		digestPropertyName = strings.Join([]string{s.metadataPrefix, digestPropertyName}, ":")
	}
	err = i.TagImage(img.Properties, s.imageID, digest, digestPropertyName)
	if err != nil {
		return err
	}

	return nil
}

// ValidateImage can validate the signing of an image using the supplied key.
// It will search for the digest metadata/property on the image and attempt to validate it.
func (s *OpenStackSignProvisioner) ValidateImage(key []byte) error {
	cloudProvider := ostack.NewCloudsProvider(s.cloudName)

	i, err := ostack.NewImageClient(cloudProvider)
	if err != nil {
		return err
	}

	img, err := i.FetchImage(s.imageID)
	if err != nil {
		return err
	}

	field, err := data.GetNestedField(img.Properties, "digest")
	if err != nil {
		return err
	}
	if field == nil {
		return fmt.Errorf("the digest field was empty")
	}

	digest := field.(string)

	valid, err := sign.Validate(s.imageID, key, digest)
	if err != nil {
		return err
	}

	log.Printf("The validation result was: %t", valid)

	return nil
}
