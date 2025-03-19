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

package provisoner

import (
	"fmt"
	"github.com/nscaledev/baski/pkg/providers/packer"
	"github.com/nscaledev/baski/pkg/util/flags"
	"os"
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
	date := "date"
	os := "os"
	k8s := "kubernetes_version"

	if o.OpenStackCoreFlags.MetadataPrefix != "" {
		date = fmt.Sprintf("%s:%s", o.OpenStackCoreFlags.MetadataPrefix, date)
		os = fmt.Sprintf("%s:%s", o.OpenStackCoreFlags.MetadataPrefix, os)
		k8s = fmt.Sprintf("%s:%s", o.OpenStackCoreFlags.MetadataPrefix, k8s)
	}
	meta := map[string]string{
		date: time.Now().Format(time.RFC3339),
		os:   o.BuildOS,
		k8s:  fmt.Sprintf("v%s", o.KubeVersion),
	}

	if len(o.AdditionalMetadata) > 0 {
		for k, v := range o.AdditionalMetadata {
			meta[k] = v
		}
	}
	return meta
}
