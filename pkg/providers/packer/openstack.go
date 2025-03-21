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

package packer

// OpenStackBuildconfig adds additional packer vars for OpenStack builds
type OpenStackBuildconfig struct {
	ImageName             string `json:"image_name,omitempty"`
	SourceImage           string `json:"source_image"`
	Networks              string `json:"networks"`
	Flavor                string `json:"flavor"`
	AttachConfigDrive     string `json:"attach_config_drive,omitempty"`
	SSHPrivateKeyFile     string `json:"ssh_private_key_file,omitempty"`
	SSHKeypairName        string `json:"ssh_keypair_name,omitempty"`
	UseFloatingIp         string `json:"use_floating_ip,omitempty"`
	FloatingIpNetwork     string `json:"floating_ip_network,omitempty"`
	SecurityGroup         string `json:"security_groups,omitempty"`
	ImageVisibility       string `json:"image_visibility,omitempty"`
	ImageDiskFormat       string `json:"image_disk_format"`
	UseBlockStorageVolume string `json:"use_blockstorage_volume,omitempty"`
	VolumeType            string `json:"volume_type"`
	VolumeSize            string `json:"volume_size"`
}
