/*
Copyright 2025 Nscale.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
	//c := mock.MockOpenStackComputeClient{}
	//i := mock.MockOpenStackImageClient{}
	//n := mock.MockOpenStackNetworkClient{}
	//ss3 := mock.MockS3Interface{}
	//s := NewOpenStackScanner(&c, &i, &n, ss3, trivy.HIGH, &images.Image{})
}

func TestUploadResultsToS3(t *testing.T) {

}
