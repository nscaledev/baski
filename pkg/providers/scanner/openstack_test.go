package scanner

import (
	"testing"
)

func TestNewOpenStackScanner(t *testing.T) {
	//TODO a test needs writing for this
}

func TestRunScan(t *testing.T) {
	//c := gomock.NewController(t)
	//defer c.Finish()
	//
	//m := mock.NewMockOpenStackScannerInterface(c)
	//m.EXPECT().RunScan(&flags.ScanOptions{
	//	OpenStackFlags: flags.OpenStackFlags{
	//		OpenStackCoreFlags: flags.OpenStackCoreFlags{
	//			CloudsPath: "",
	//			CloudName:  "",
	//		},
	//		OpenStackInstanceFlags: flags.OpenStackInstanceFlags{
	//			AttachConfigDrive: false,
	//			NetworkID:         "",
	//			FlavorName:        "",
	//		},
	//		SourceImageID:         "",
	//		SSHPrivateKeyFile:     "",
	//		SSHKeypairName:        "",
	//		UseFloatingIP:         false,
	//		FloatingIPNetworkName: "",
	//		SecurityGroup:         "",
	//		ImageVisibility:       "",
	//		ImageDiskFormat:       "",
	//		UseBlockStorageVolume: "",
	//		VolumeType:            "",
	//		VolumeSize:            0,
	//	},
	//	ScanSingleOptions:   flags.ScanSingleOptions{},
	//	ScanMultipleOptions: flags.ScanMultipleOptions{},
	//	AutoDeleteImage:     false,
	//	SkipCVECheck:        false,
	//	MaxSeverityScore:    0,
	//	MaxSeverityType:     "",
	//	ScanBucket:          "",
	//	TrivyignorePath:     "",
	//	TrivyignoreFilename: "",
	//	TrivyignoreList:     nil,
	//})
}

func TestFetchScanResults(t *testing.T) {

}

func TestCheckResults(t *testing.T) {

}

func TestTagImage(t *testing.T) {

}

func TestUploadResultsToS3(t *testing.T) {

}
