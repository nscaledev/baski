package provisoner

import (
	"fmt"
	"github.com/drewbernetes/baski/pkg/providers/packer"
	"github.com/drewbernetes/baski/pkg/util/flags"
	"os"
	"strings"
	"time"
)

type BuilderProvisioner interface {
	Init() error
	GeneratePackerConfig() (*packer.GlobalBuildConfig, error)
	UpdatePackerBuilders(metadata map[string]string, data []byte) []byte
	PostBuildAction() error
}

// NewBuilder returns a new provisioner based on the infra type that is used for building images.
func NewBuilder(o *flags.BuildOptions) BuilderProvisioner {
	switch o.InfraType {
	case "openstack":
		return newOpenStackBuilder(o)
	case "kubevirt":
		return newKubeVirtBuilder(o)

	}

	return nil
}

type ScannerProvisioner interface {
	Prepare() error
	ScanImages() error
}

func NewScanner(o *flags.ScanOptions) ScannerProvisioner {
	switch o.InfraType {
	case "openstack":
		return newOpenStackScanner(o)
	case "kubevirt":
		return newKubeVirtScanner(o)
	}
	return nil
}

type SignerProvisioner interface {
	SignImage(digest string) error
	ValidateImage(key []byte) error
}

func NewSigner(o *flags.SignOptions) SignerProvisioner {
	switch o.InfraType {
	case "openstack":
		return newOpenStackSigner(o)
	case "kubevirt":
		return newKubeVirtSigner(o)

	}

	return nil
}

// saveImageIDToFile exports the image ID to a file so that it can be read later by the scan system.
func saveImageIDToFile(imgID string) error {
	f, err := os.Create("/tmp/imgid.out")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(imgID))
	if err != nil {
		return err
	}

	return nil
}

// generateBuilderMetadata generates some glance metadata for the image.
func generateBuilderMetadata(o *flags.BuildOptions) map[string]string {
	var gpuVendor string
	gpuVersion := "no_gpu"
	gpuVendor = strings.ToUpper(o.GpuVendor)
	if gpuVendor == "NVIDIA" {
		gpuVersion = fmt.Sprintf("v%s", o.NvidiaVersion)
	} else if gpuVendor == "AMD" {
		gpuVersion = fmt.Sprintf("v%s", o.AMDVersion)
	}

	metaPrefix := o.OpenStackCoreFlags.MetadataPrefix
	meta := map[string]string{
		"date": time.Now().Format(time.RFC3339),
		"os":   o.BuildOS,
		fmt.Sprintf("%s:kubernetes_version", metaPrefix): fmt.Sprintf("v%s", o.KubeVersion),
	}

	if o.AddGpuSupport {
		gpuMeta := map[string]string{
			fmt.Sprintf("%s:gpu_vendor", metaPrefix):         gpuVendor,
			fmt.Sprintf("%s:gpu_models", metaPrefix):         strings.ToUpper(o.GpuModelSupport),
			fmt.Sprintf("%s:gpu_driver_version", metaPrefix): gpuVersion,
			fmt.Sprintf("%s:virtualization", metaPrefix):     o.GpuInstanceSupport,
		}
		for k, v := range gpuMeta {
			meta[k] = v
		}
	}
	if len(o.AdditionalMetadata) > 0 {
		for k, v := range o.AdditionalMetadata {
			meta[k] = v
		}
	}
	return meta
}
