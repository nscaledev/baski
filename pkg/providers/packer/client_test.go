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

import (
	"encoding/json"
	"github.com/nscaledev/baski/pkg/util/flags"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"testing"
)

var packerConf = `{
  "builders": [
	{}
  ],
  "post-processors": [],
  "provisioners": [],
  "variables": {}
}`

func TestNewCoreBuildconfig(t *testing.T) {
	tc := []struct {
		name     string
		options  flags.BuildOptions
		expected GlobalBuildConfig
	}{
		{
			name: "Test basic config with no additions",
			options: flags.BuildOptions{
				ImagePrefix:       "kmi",
				ContainerdVersion: "1.7.13",
				ContainerdSHA256:  "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:        "1.2.0",
				CniDebVersion:     "1.2.0-2.1",
				CrictlVersion:     "1.26.0",
				KubeVersion:       "1.28.2",
				KubeRpmVersion:    "1.28.2",
				KubeDebVersion:    "1.28.2-1.1",
				ExtraDebs:         "nfs-common",
			},
			expected: GlobalBuildConfig{
				ContainerdVersion:    "1.7.13",
				ContainerdSHA256:     "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				ExtraDebs:            "nfs-common",
			},
		},
		{
			name: "Test config with Falco",
			options: flags.BuildOptions{
				ImagePrefix:       "kmi",
				ContainerdVersion: "1.7.13",
				ContainerdSHA256:  "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:        "1.2.0",
				CniDebVersion:     "1.2.0-2.1",
				CrictlVersion:     "1.26.0",
				KubeVersion:       "1.28.2",
				KubeRpmVersion:    "1.28.2",
				KubeDebVersion:    "1.28.2-1.1",
				ExtraDebs:         "nfs-common",
				AddFalco:          true,
			},
			expected: GlobalBuildConfig{
				ContainerdVersion:    "1.7.13",
				ContainerdSHA256:     "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				NodeCustomRolesPre:   "security",
				AnsibleUserVars:      "security_install_falco=true",
				ExtraDebs:            "nfs-common",
			},
		},
		{
			name: "Test config with Trivy",
			options: flags.BuildOptions{
				ImagePrefix:       "kmi",
				ContainerdVersion: "1.7.13",
				ContainerdSHA256:  "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:        "1.2.0",
				CniDebVersion:     "1.2.0-2.1",
				CrictlVersion:     "1.26.0",
				KubeVersion:       "1.28.2",
				KubeRpmVersion:    "1.28.2",
				KubeDebVersion:    "1.28.2-1.1",
				ExtraDebs:         "nfs-common",
				AddTrivy:          true,
			},
			expected: GlobalBuildConfig{
				ContainerdVersion:    "1.7.13",
				ContainerdSHA256:     "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				NodeCustomRolesPre:   "security",
				AnsibleUserVars:      "security_install_trivy=true",
				ExtraDebs:            "nfs-common",
			},
		},
		{
			name: "Test config with Falco & trivy",
			options: flags.BuildOptions{
				ImagePrefix:       "kmi",
				ContainerdVersion: "1.7.13",
				ContainerdSHA256:  "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:        "1.2.0",
				CniDebVersion:     "1.2.0-2.1",
				CrictlVersion:     "1.26.0",
				KubeVersion:       "1.28.2",
				KubeRpmVersion:    "1.28.2",
				KubeDebVersion:    "1.28.2-1.1",
				ExtraDebs:         "nfs-common",
				AddFalco:          true,
				AddTrivy:          true,
			},
			expected: GlobalBuildConfig{
				ContainerdVersion:    "1.7.13",
				ContainerdSHA256:     "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				NodeCustomRolesPre:   "security",
				AnsibleUserVars:      "security_install_falco=true security_install_trivy=true",
				ExtraDebs:            "nfs-common",
			},
		},
		{
			name: "Test config with additional images",
			options: flags.BuildOptions{
				ImagePrefix:       "kmi",
				ContainerdVersion: "1.7.13",
				ContainerdSHA256:  "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:        "1.2.0",
				CniDebVersion:     "1.2.0-2.1",
				CrictlVersion:     "1.26.0",
				KubeVersion:       "1.28.2",
				KubeRpmVersion:    "1.28.2",
				KubeDebVersion:    "1.28.2-1.1",
				ExtraDebs:         "nfs-common",
				AdditionalImages:  []string{"image1", "image2"},
			},
			expected: GlobalBuildConfig{
				ContainerdVersion:    "1.7.13",
				ContainerdSHA256:     "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				AnsibleUserVars:      "load_additional_components=true additional_registry_images=true additional_registry_images_list=image1,image2",
				ExtraDebs:            "nfs-common",
			},
		},
		{
			name: "Test config with additional images & security",
			options: flags.BuildOptions{
				ImagePrefix:       "kmi",
				ContainerdVersion: "1.7.13",
				ContainerdSHA256:  "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:        "1.2.0",
				CniDebVersion:     "1.2.0-2.1",
				CrictlVersion:     "1.26.0",
				KubeVersion:       "1.28.2",
				KubeRpmVersion:    "1.28.2",
				KubeDebVersion:    "1.28.2-1.1",
				ExtraDebs:         "nfs-common",
				AdditionalImages:  []string{"image1", "image2"},
				AddFalco:          true,
				AddTrivy:          true,
			},
			expected: GlobalBuildConfig{
				ContainerdVersion:    "1.7.13",
				ContainerdSHA256:     "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				NodeCustomRolesPre:   "security",
				AnsibleUserVars:      "load_additional_components=true additional_registry_images=true additional_registry_images_list=image1,image2 security_install_falco=true security_install_trivy=true",
				ExtraDebs:            "nfs-common",
			},
		},
		{
			name: "Test config with AMD GPU",
			options: flags.BuildOptions{
				ImagePrefix:       "kmi",
				ContainerdVersion: "1.7.13",
				ContainerdSHA256:  "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:        "1.2.0",
				CniDebVersion:     "1.2.0-2.1",
				CrictlVersion:     "1.26.0",
				KubeVersion:       "1.28.2",
				KubeRpmVersion:    "1.28.2",
				KubeDebVersion:    "1.28.2-1.1",
				ExtraDebs:         "nfs-common",
				AddGpuSupport:     true,
				GpuVendor:         "amd",
				AMDVersion:        "6.0.2",
				AMDDebVersion:     "6.0.60002-1",
				AMDUseCase:        "dkms",
			},
			expected: GlobalBuildConfig{
				ContainerdVersion:    "1.7.13",
				ContainerdSHA256:     "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				NodeCustomRolesPre:   "gpu",
				AnsibleUserVars:      "gpu_vendor=amd amd_version=6.0.2 amd_deb_version=6.0.60002-1 gpu_amd_usecase=dkms",
				ExtraDebs:            "nfs-common",
			},
		},
		{
			name: "Test config with AMD GPU & security",
			options: flags.BuildOptions{
				ImagePrefix:       "kmi",
				ContainerdVersion: "1.7.13",
				ContainerdSHA256:  "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:        "1.2.0",
				CniDebVersion:     "1.2.0-2.1",
				CrictlVersion:     "1.26.0",
				KubeVersion:       "1.28.2",
				KubeRpmVersion:    "1.28.2",
				KubeDebVersion:    "1.28.2-1.1",
				ExtraDebs:         "nfs-common",
				AddFalco:          true,
				AddTrivy:          true,
				AddGpuSupport:     true,
				GpuVendor:         "amd",
				AMDVersion:        "6.0.2",
				AMDDebVersion:     "6.0.60002-1",
				AMDUseCase:        "dkms",
			},
			expected: GlobalBuildConfig{
				ContainerdVersion:    "1.7.13",
				ContainerdSHA256:     "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				NodeCustomRolesPre:   "gpu security",
				AnsibleUserVars:      "gpu_vendor=amd amd_version=6.0.2 amd_deb_version=6.0.60002-1 gpu_amd_usecase=dkms security_install_falco=true security_install_trivy=true",
				ExtraDebs:            "nfs-common",
			},
		},
		{
			name: "Test config with AMD GPU, additional images & security",
			options: flags.BuildOptions{
				ImagePrefix:       "kmi",
				ContainerdVersion: "1.7.13",
				ContainerdSHA256:  "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:        "1.2.0",
				CniDebVersion:     "1.2.0-2.1",
				CrictlVersion:     "1.26.0",
				KubeVersion:       "1.28.2",
				KubeRpmVersion:    "1.28.2",
				KubeDebVersion:    "1.28.2-1.1",
				ExtraDebs:         "nfs-common",
				AdditionalImages:  []string{"image1", "image2"},
				AddFalco:          true,
				AddTrivy:          true,
				AddGpuSupport:     true,
				GpuVendor:         "amd",
				AMDVersion:        "6.0.2",
				AMDDebVersion:     "6.0.60002-1",
				AMDUseCase:        "dkms",
			},
			expected: GlobalBuildConfig{
				ContainerdVersion:    "1.7.13",
				ContainerdSHA256:     "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				NodeCustomRolesPre:   "gpu security",
				AnsibleUserVars:      "gpu_vendor=amd amd_version=6.0.2 amd_deb_version=6.0.60002-1 gpu_amd_usecase=dkms load_additional_components=true additional_registry_images=true additional_registry_images_list=image1,image2 security_install_falco=true security_install_trivy=true",
				ExtraDebs:            "nfs-common",
			},
		},
		{
			name: "Test config with NVIDIA GPU",
			options: flags.BuildOptions{
				ImagePrefix:             "kmi",
				ContainerdVersion:       "1.7.13",
				ContainerdSHA256:        "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:              "1.2.0",
				CniDebVersion:           "1.2.0-2.1",
				CrictlVersion:           "1.26.0",
				KubeVersion:             "1.28.2",
				KubeRpmVersion:          "1.28.2",
				KubeDebVersion:          "1.28.2-1.1",
				ExtraDebs:               "nfs-common",
				AddGpuSupport:           true,
				GpuVendor:               "nvidia",
				NvidiaVersion:           "535.129.03",
				NvidiaBucket:            "nvidia",
				NvidiaInstallerLocation: "NVIDIA-Linux-x86_64-535.129.03-grid.run",
				NvidiaTOKLocation:       "client_configuration_token.tok",
				NvidiaGriddFeatureType:  4,
				S3Flags: flags.S3Flags{
					Endpoint:  "https://example.com",
					AccessKey: "123456",
					SecretKey: "987654",
					Region:    "us-east-1",
					IsCeph:    true,
				},
			},
			expected: GlobalBuildConfig{
				ContainerdVersion:    "1.7.13",
				ContainerdSHA256:     "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				NodeCustomRolesPre:   "gpu",
				AnsibleUserVars:      "gpu_vendor=nvidia nvidia_s3_url=https://example.com nvidia_bucket=nvidia nvidia_bucket_access=123456 nvidia_bucket_secret=987654 nvidia_ceph=true nvidia_installer_location=NVIDIA-Linux-x86_64-535.129.03-grid.run nvidia_tok_location=client_configuration_token.tok gridd_feature_type=4",
				ExtraDebs:            "nfs-common",
			},
		},
		{
			name: "Test config with NVIDIA GPU & security",
			options: flags.BuildOptions{
				ImagePrefix:             "kmi",
				ContainerdVersion:       "1.7.13",
				ContainerdSHA256:        "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:              "1.2.0",
				CniDebVersion:           "1.2.0-2.1",
				CrictlVersion:           "1.26.0",
				KubeVersion:             "1.28.2",
				KubeRpmVersion:          "1.28.2",
				KubeDebVersion:          "1.28.2-1.1",
				ExtraDebs:               "nfs-common",
				AddFalco:                true,
				AddTrivy:                true,
				AddGpuSupport:           true,
				GpuVendor:               "nvidia",
				NvidiaVersion:           "535.129.03",
				NvidiaBucket:            "nvidia",
				NvidiaInstallerLocation: "NVIDIA-Linux-x86_64-535.129.03-grid.run",
				NvidiaTOKLocation:       "client_configuration_token.tok",
				NvidiaGriddFeatureType:  4,
				S3Flags: flags.S3Flags{
					Endpoint:  "https://example.com",
					AccessKey: "123456",
					SecretKey: "987654",
					Region:    "us-east-1",
					IsCeph:    true,
				},
			},
			expected: GlobalBuildConfig{
				ContainerdVersion:    "1.7.13",
				ContainerdSHA256:     "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				NodeCustomRolesPre:   "gpu security",
				AnsibleUserVars:      "gpu_vendor=nvidia nvidia_s3_url=https://example.com nvidia_bucket=nvidia nvidia_bucket_access=123456 nvidia_bucket_secret=987654 nvidia_ceph=true nvidia_installer_location=NVIDIA-Linux-x86_64-535.129.03-grid.run nvidia_tok_location=client_configuration_token.tok gridd_feature_type=4 security_install_falco=true security_install_trivy=true",
				ExtraDebs:            "nfs-common",
			},
		},
		{
			name: "Test config with NVIDIA GPU, additional images & security",
			options: flags.BuildOptions{
				ImagePrefix:             "kmi",
				ContainerdVersion:       "1.7.13",
				ContainerdSHA256:        "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:              "1.2.0",
				CniDebVersion:           "1.2.0-2.1",
				CrictlVersion:           "1.26.0",
				KubeVersion:             "1.28.2",
				KubeRpmVersion:          "1.28.2",
				KubeDebVersion:          "1.28.2-1.1",
				ExtraDebs:               "nfs-common",
				AdditionalImages:        []string{"image1", "image2"},
				AddFalco:                true,
				AddTrivy:                true,
				AddGpuSupport:           true,
				GpuVendor:               "nvidia",
				NvidiaVersion:           "535.129.03",
				NvidiaBucket:            "nvidia",
				NvidiaInstallerLocation: "NVIDIA-Linux-x86_64-535.129.03-grid.run",
				NvidiaTOKLocation:       "client_configuration_token.tok",
				NvidiaGriddFeatureType:  4,
				S3Flags: flags.S3Flags{
					Endpoint:  "https://example.com",
					AccessKey: "123456",
					SecretKey: "987654",
					Region:    "us-east-1",
					IsCeph:    true,
				},
			},
			expected: GlobalBuildConfig{
				ContainerdVersion:    "1.7.13",
				ContainerdSHA256:     "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19",
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				NodeCustomRolesPre:   "gpu security",
				AnsibleUserVars:      "gpu_vendor=nvidia nvidia_s3_url=https://example.com nvidia_bucket=nvidia nvidia_bucket_access=123456 nvidia_bucket_secret=987654 nvidia_ceph=true nvidia_installer_location=NVIDIA-Linux-x86_64-535.129.03-grid.run nvidia_tok_location=client_configuration_token.tok gridd_feature_type=4 load_additional_components=true additional_registry_images=true additional_registry_images_list=image1,image2 security_install_falco=true security_install_trivy=true",
				ExtraDebs:            "nfs-common",
			},
		},
	}

	for _, v := range tc {
		t.Run(v.name, func(t *testing.T) {
			conf, name, err := NewCoreBuildconfig(&v.options)
			if err != nil {
				t.Error(err.Error())
			}

			re := regexp.MustCompile(`kmi-\d{6}-\w{8}`)
			if !re.MatchString(name) {
				t.Errorf("expected name %s, but it didn't match the regex\n", name)
			}

			if !reflect.DeepEqual(&v.expected, conf) {
				t.Errorf("expected: %v, got %v\n", &v.expected, conf)
			}
		})
	}

}

func TestUpdatePackerBuildersJson(t *testing.T) {
	infra := "test"
	dir := "/tmp/test"
	expected := `{"builders":[{"metadata":{"test":"test"}}],"post-processors":[],"provisioners":[],"variables":{}}`

	fullPath := filepath.Join(dir, "images", "capi", "packer", infra)

	err := os.MkdirAll(fullPath, 0755)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = os.WriteFile(filepath.Join(fullPath, "packer.json"), []byte(packerConf), 0644)
	if err != nil {
		t.Fatal(err.Error())
	}

	modifierFunc := BuildersModifier{
		Function: func(metadata map[string]string, data []byte) []byte {
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
		},
		Metadata: map[string]string{"test": "test"},
	}

	err = UpdatePackerBuildersJson(dir, infra, modifierFunc)
	if err != nil {
		t.Fatal(err.Error())
	}

	file, err := os.ReadFile(filepath.Join(fullPath, "packer.json"))
	if err != nil {
		t.Fatal(err.Error())
	}

	if expected != string(file) {
		t.Errorf("expected: %s, got: %s\n", expected, string(file))
	}
}

func TestGenerateVariablesFile(t *testing.T) {
	tc := []struct {
		name     string
		conf     GlobalBuildConfig
		expected string
	}{
		{
			name: "Test basic config with no additions",
			conf: GlobalBuildConfig{
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				ExtraDebs:            "nfs-common",
			},
			expected: `{"kubernetes_cni_semver":"v1.2.0","kubernetes_cni_deb_version":"1.2.0-2.1","crictl_version":"1.26.0","kubernetes_semver":"v1.28.2","kubernetes_rpm_version":"1.28.2","kubernetes_series":"v1.28","kubernetes_deb_version":"1.28.2-1.1","extra_debs":"nfs-common","source_image":"","networks":"","flavor":"","image_disk_format":"","volume_type":"","volume_size":"","qemu_binary":"","disk_size":"","output_directory":""}`,
		},
		{
			name: "Test config with Falco",
			conf: GlobalBuildConfig{
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				NodeCustomRolesPre:   "security",
				AnsibleUserVars:      "security_install_falco=true",
				ExtraDebs:            "nfs-common",
			},
			expected: `{"kubernetes_cni_semver":"v1.2.0","kubernetes_cni_deb_version":"1.2.0-2.1","crictl_version":"1.26.0","kubernetes_semver":"v1.28.2","kubernetes_rpm_version":"1.28.2","kubernetes_series":"v1.28","kubernetes_deb_version":"1.28.2-1.1","node_custom_roles_pre":"security","ansible_user_vars":"security_install_falco=true","extra_debs":"nfs-common","source_image":"","networks":"","flavor":"","image_disk_format":"","volume_type":"","volume_size":"","qemu_binary":"","disk_size":"","output_directory":""}`,
		},
		{
			name: "Test config with Trivy",
			conf: GlobalBuildConfig{
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				NodeCustomRolesPre:   "security",
				AnsibleUserVars:      "security_install_trivy=true",
				ExtraDebs:            "nfs-common",
			},
			expected: `{"kubernetes_cni_semver":"v1.2.0","kubernetes_cni_deb_version":"1.2.0-2.1","crictl_version":"1.26.0","kubernetes_semver":"v1.28.2","kubernetes_rpm_version":"1.28.2","kubernetes_series":"v1.28","kubernetes_deb_version":"1.28.2-1.1","node_custom_roles_pre":"security","ansible_user_vars":"security_install_trivy=true","extra_debs":"nfs-common","source_image":"","networks":"","flavor":"","image_disk_format":"","volume_type":"","volume_size":"","qemu_binary":"","disk_size":"","output_directory":""}`,
		},
		{
			name: "Test config with Falco & trivy",
			conf: GlobalBuildConfig{
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				NodeCustomRolesPre:   "security",
				AnsibleUserVars:      "security_install_falco=true security_install_trivy=true",
				ExtraDebs:            "nfs-common",
			},
			expected: `{"kubernetes_cni_semver":"v1.2.0","kubernetes_cni_deb_version":"1.2.0-2.1","crictl_version":"1.26.0","kubernetes_semver":"v1.28.2","kubernetes_rpm_version":"1.28.2","kubernetes_series":"v1.28","kubernetes_deb_version":"1.28.2-1.1","node_custom_roles_pre":"security","ansible_user_vars":"security_install_falco=true security_install_trivy=true","extra_debs":"nfs-common","source_image":"","networks":"","flavor":"","image_disk_format":"","volume_type":"","volume_size":"","qemu_binary":"","disk_size":"","output_directory":""}`,
		},
		{
			name: "Test config with additional images",
			conf: GlobalBuildConfig{
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				AnsibleUserVars:      "load_additional_components=true additional_registry_images=true additional_registry_images_list=image1,image2",
				ExtraDebs:            "nfs-common",
			},
			expected: `{"kubernetes_cni_semver":"v1.2.0","kubernetes_cni_deb_version":"1.2.0-2.1","crictl_version":"1.26.0","kubernetes_semver":"v1.28.2","kubernetes_rpm_version":"1.28.2","kubernetes_series":"v1.28","kubernetes_deb_version":"1.28.2-1.1","ansible_user_vars":"load_additional_components=true additional_registry_images=true additional_registry_images_list=image1,image2","extra_debs":"nfs-common","source_image":"","networks":"","flavor":"","image_disk_format":"","volume_type":"","volume_size":"","qemu_binary":"","disk_size":"","output_directory":""}`,
		},
		{
			name: "Test config with additional images & security",
			conf: GlobalBuildConfig{
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				NodeCustomRolesPre:   "security",
				AnsibleUserVars:      "load_additional_components=true additional_registry_images=true additional_registry_images_list=image1,image2 security_install_falco=true security_install_trivy=true",
				ExtraDebs:            "nfs-common",
			},
			expected: `{"kubernetes_cni_semver":"v1.2.0","kubernetes_cni_deb_version":"1.2.0-2.1","crictl_version":"1.26.0","kubernetes_semver":"v1.28.2","kubernetes_rpm_version":"1.28.2","kubernetes_series":"v1.28","kubernetes_deb_version":"1.28.2-1.1","node_custom_roles_pre":"security","ansible_user_vars":"load_additional_components=true additional_registry_images=true additional_registry_images_list=image1,image2 security_install_falco=true security_install_trivy=true","extra_debs":"nfs-common","source_image":"","networks":"","flavor":"","image_disk_format":"","volume_type":"","volume_size":"","qemu_binary":"","disk_size":"","output_directory":""}`,
		},
		{
			name: "Test config with AMD GPU",
			conf: GlobalBuildConfig{
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				NodeCustomRolesPre:   "gpu",
				AnsibleUserVars:      "gpu_vendor=amd amd_version=6.0.2 amd_deb_version=6.0.60002-1 gpu_amd_usecase=dkms",
				ExtraDebs:            "nfs-common",
			},
			expected: `{"kubernetes_cni_semver":"v1.2.0","kubernetes_cni_deb_version":"1.2.0-2.1","crictl_version":"1.26.0","kubernetes_semver":"v1.28.2","kubernetes_rpm_version":"1.28.2","kubernetes_series":"v1.28","kubernetes_deb_version":"1.28.2-1.1","node_custom_roles_pre":"gpu","ansible_user_vars":"gpu_vendor=amd amd_version=6.0.2 amd_deb_version=6.0.60002-1 gpu_amd_usecase=dkms","extra_debs":"nfs-common","source_image":"","networks":"","flavor":"","image_disk_format":"","volume_type":"","volume_size":"","qemu_binary":"","disk_size":"","output_directory":""}`,
		},
		{
			name: "Test config with AMD GPU & security",
			conf: GlobalBuildConfig{
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				NodeCustomRolesPre:   "gpu security",
				AnsibleUserVars:      "gpu_vendor=amd amd_version=6.0.2 amd_deb_version=6.0.60002-1 gpu_amd_usecase=dkms security_install_falco=true security_install_trivy=true",
				ExtraDebs:            "nfs-common",
			},
			expected: `{"kubernetes_cni_semver":"v1.2.0","kubernetes_cni_deb_version":"1.2.0-2.1","crictl_version":"1.26.0","kubernetes_semver":"v1.28.2","kubernetes_rpm_version":"1.28.2","kubernetes_series":"v1.28","kubernetes_deb_version":"1.28.2-1.1","node_custom_roles_pre":"gpu security","ansible_user_vars":"gpu_vendor=amd amd_version=6.0.2 amd_deb_version=6.0.60002-1 gpu_amd_usecase=dkms security_install_falco=true security_install_trivy=true","extra_debs":"nfs-common","source_image":"","networks":"","flavor":"","image_disk_format":"","volume_type":"","volume_size":"","qemu_binary":"","disk_size":"","output_directory":""}`,
		},
		{
			name: "Test config with AMD GPU, additional images & security",
			conf: GlobalBuildConfig{
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				NodeCustomRolesPre:   "gpu security",
				AnsibleUserVars:      "gpu_vendor=amd amd_version=6.0.2 amd_deb_version=6.0.60002-1 gpu_amd_usecase=dkms load_additional_components=true additional_registry_images=true additional_registry_images_list=image1,image2 security_install_falco=true security_install_trivy=true",
				ExtraDebs:            "nfs-common",
			},
			expected: `{"kubernetes_cni_semver":"v1.2.0","kubernetes_cni_deb_version":"1.2.0-2.1","crictl_version":"1.26.0","kubernetes_semver":"v1.28.2","kubernetes_rpm_version":"1.28.2","kubernetes_series":"v1.28","kubernetes_deb_version":"1.28.2-1.1","node_custom_roles_pre":"gpu security","ansible_user_vars":"gpu_vendor=amd amd_version=6.0.2 amd_deb_version=6.0.60002-1 gpu_amd_usecase=dkms load_additional_components=true additional_registry_images=true additional_registry_images_list=image1,image2 security_install_falco=true security_install_trivy=true","extra_debs":"nfs-common","source_image":"","networks":"","flavor":"","image_disk_format":"","volume_type":"","volume_size":"","qemu_binary":"","disk_size":"","output_directory":""}`,
		},
		{
			name: "Test config with NVIDIA GPU",
			conf: GlobalBuildConfig{
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				NodeCustomRolesPre:   "gpu",
				AnsibleUserVars:      "gpu_vendor=nvidia nvidia_s3_url=https://example.com nvidia_bucket=nvidia nvidia_bucket_access=123456 nvidia_bucket_secret=987654 nvidia_ceph=true nvidia_installer_location=NVIDIA-Linux-x86_64-535.129.03-grid.run nvidia_tok_location=client_configuration_token.tok gridd_feature_type=4",
				ExtraDebs:            "nfs-common",
			},
			expected: `{"kubernetes_cni_semver":"v1.2.0","kubernetes_cni_deb_version":"1.2.0-2.1","crictl_version":"1.26.0","kubernetes_semver":"v1.28.2","kubernetes_rpm_version":"1.28.2","kubernetes_series":"v1.28","kubernetes_deb_version":"1.28.2-1.1","node_custom_roles_pre":"gpu","ansible_user_vars":"gpu_vendor=nvidia nvidia_s3_url=https://example.com nvidia_bucket=nvidia nvidia_bucket_access=123456 nvidia_bucket_secret=987654 nvidia_ceph=true nvidia_installer_location=NVIDIA-Linux-x86_64-535.129.03-grid.run nvidia_tok_location=client_configuration_token.tok gridd_feature_type=4","extra_debs":"nfs-common","source_image":"","networks":"","flavor":"","image_disk_format":"","volume_type":"","volume_size":"","qemu_binary":"","disk_size":"","output_directory":""}`,
		},
		{
			name: "Test config with NVIDIA GPU & security",
			conf: GlobalBuildConfig{
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				NodeCustomRolesPre:   "gpu security",
				AnsibleUserVars:      "gpu_vendor=nvidia nvidia_s3_url=https://example.com nvidia_bucket=nvidia nvidia_bucket_access=123456 nvidia_bucket_secret=987654 nvidia_ceph=true nvidia_installer_location=NVIDIA-Linux-x86_64-535.129.03-grid.run nvidia_tok_location=client_configuration_token.tok gridd_feature_type=4 security_install_falco=true security_install_trivy=true",
				ExtraDebs:            "nfs-common",
			},
			expected: `{"kubernetes_cni_semver":"v1.2.0","kubernetes_cni_deb_version":"1.2.0-2.1","crictl_version":"1.26.0","kubernetes_semver":"v1.28.2","kubernetes_rpm_version":"1.28.2","kubernetes_series":"v1.28","kubernetes_deb_version":"1.28.2-1.1","node_custom_roles_pre":"gpu security","ansible_user_vars":"gpu_vendor=nvidia nvidia_s3_url=https://example.com nvidia_bucket=nvidia nvidia_bucket_access=123456 nvidia_bucket_secret=987654 nvidia_ceph=true nvidia_installer_location=NVIDIA-Linux-x86_64-535.129.03-grid.run nvidia_tok_location=client_configuration_token.tok gridd_feature_type=4 security_install_falco=true security_install_trivy=true","extra_debs":"nfs-common","source_image":"","networks":"","flavor":"","image_disk_format":"","volume_type":"","volume_size":"","qemu_binary":"","disk_size":"","output_directory":""}`,
		},
		{
			name: "Test config with NVIDIA GPU, additional images & security",
			conf: GlobalBuildConfig{
				CniVersion:           "v1.2.0",
				CniDebVersion:        "1.2.0-2.1",
				CrictlVersion:        "1.26.0",
				KubernetesSemver:     "v1.28.2",
				KubernetesRpmVersion: "1.28.2",
				KubernetesSeries:     "v1.28",
				KubernetesDebVersion: "1.28.2-1.1",
				NodeCustomRolesPre:   "gpu security",
				AnsibleUserVars:      "gpu_vendor=nvidia nvidia_s3_url=https://example.com nvidia_bucket=nvidia nvidia_bucket_access=123456 nvidia_bucket_secret=987654 nvidia_ceph=true nvidia_installer_location=NVIDIA-Linux-x86_64-535.129.03-grid.run nvidia_tok_location=client_configuration_token.tok gridd_feature_type=4 load_additional_components=true additional_registry_images=true additional_registry_images_list=image1,image2 security_install_falco=true security_install_trivy=true",
				ExtraDebs:            "nfs-common",
			},
			expected: `{"kubernetes_cni_semver":"v1.2.0","kubernetes_cni_deb_version":"1.2.0-2.1","crictl_version":"1.26.0","kubernetes_semver":"v1.28.2","kubernetes_rpm_version":"1.28.2","kubernetes_series":"v1.28","kubernetes_deb_version":"1.28.2-1.1","node_custom_roles_pre":"gpu security","ansible_user_vars":"gpu_vendor=nvidia nvidia_s3_url=https://example.com nvidia_bucket=nvidia nvidia_bucket_access=123456 nvidia_bucket_secret=987654 nvidia_ceph=true nvidia_installer_location=NVIDIA-Linux-x86_64-535.129.03-grid.run nvidia_tok_location=client_configuration_token.tok gridd_feature_type=4 load_additional_components=true additional_registry_images=true additional_registry_images_list=image1,image2 security_install_falco=true security_install_trivy=true","extra_debs":"nfs-common","source_image":"","networks":"","flavor":"","image_disk_format":"","volume_type":"","volume_size":"","qemu_binary":"","disk_size":"","output_directory":""}`,
		},
	}

	for _, v := range tc {
		t.Run(v.name, func(t *testing.T) {
			v.conf.GenerateVariablesFile("/tmp/")

			res, err := os.ReadFile("/tmp/tmp.json")
			if err != nil {
				t.Errorf("%s", err.Error())
			}

			if v.expected != string(res) {
				t.Errorf("expected: %s, got %s\n", v.expected, string(res))
			}

			err = os.Truncate("/tmp/tmp.json", 0)
			if err != nil {
				t.Errorf("%s", err.Error())
			}
		})
	}

}
