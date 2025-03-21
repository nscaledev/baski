package provisoner

import (
	"errors"
	"github.com/nscaledev/baski/pkg/util/flags"
	"os"
	"testing"
)

func TestNewBuilder(t *testing.T) {
	tests := []struct {
		name    string
		options *flags.BuildOptions
		wantNil bool
	}{
		{"openstack", &flags.BuildOptions{InfraType: "openstack"}, false},
		{"kubevirt", &flags.BuildOptions{InfraType: "kubevirt"}, false},
		{"unsupported", &flags.BuildOptions{InfraType: "unknown"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBuilder(tt.options)
			if (got == nil) != tt.wantNil {
				t.Errorf("NewBuilder() = %v, want nil = %v", got, tt.wantNil)
			}
		})
	}
}

func TestNewScanner(t *testing.T) {
	tests := []struct {
		name    string
		options *flags.ScanOptions
		wantNil bool
	}{
		{"openstack", &flags.ScanOptions{InfraType: "openstack"}, false},
		{"kubevirt", &flags.ScanOptions{InfraType: "kubevirt"}, false},
		{"unsupported", &flags.ScanOptions{InfraType: "unknown"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewScanner(tt.options)
			if (got == nil) != tt.wantNil {
				t.Errorf("NewScanner() = %v, want nil = %v", got, tt.wantNil)
			}
		})
	}
}

func TestNewSigner(t *testing.T) {
	tests := []struct {
		name    string
		options *flags.SignOptions
		wantNil bool
	}{
		{"openstack", &flags.SignOptions{InfraType: "openstack"}, false},
		{"kubevirt", &flags.SignOptions{InfraType: "kubevirt"}, false},
		{"unsupported", &flags.SignOptions{InfraType: "unknown"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSigner(tt.options)
			if (got == nil) != tt.wantNil {
				t.Errorf("NewSigner() = %v, want nil = %v", got, tt.wantNil)
			}
		})
	}
}

func TestSaveImageIDToFile(t *testing.T) {
	tests := []struct {
		name      string
		imageID   string
		wantError bool
	}{
		{"valid image id", "image123", false},
		{"empty image id", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := saveImageIDToFile(tt.imageID)
			if (err != nil) != tt.wantError {
				t.Errorf("saveImageIDToFile() error = %v, wantError %v", err, tt.wantError)
			}
			if !tt.wantError {
				defer os.Remove("/tmp/imgid.out")
				content, _ := os.ReadFile("/tmp/imgid.out")
				if string(content) != tt.imageID {
					t.Errorf("saveImageIDToFile() wrote %v, want %v", string(content), tt.imageID)
				}
			}
		})
	}
}

func TestGenerateBuilderMetadata(t *testing.T) {
	tests := []struct {
		name     string
		options  *flags.BuildOptions
		validate func(map[string]string) error
	}{
		{
			"basic metadata",
			&flags.BuildOptions{
				BuildOS:     "linux",
				KubeVersion: "1.20",
				OpenStackCoreFlags: flags.OpenStackCoreFlags{
					MetadataPrefix: "test",
				},
			},
			func(meta map[string]string) error {
				if meta["os"] != "linux" {
					return errors.New("missing 'os' metadata")
				}
				if _, ok := meta["test:kubernetes_version"]; !ok {
					return errors.New("missing 'test:kubernetes_version' metadata")
				}
				return nil
			},
		},
		{
			"gpu metadata",
			&flags.BuildOptions{
				AddGpuSupport:      true,
				GpuVendor:          "nvidia",
				NvidiaVersion:      "470",
				GpuModelSupport:    "RTX",
				GpuInstanceSupport: "shared",
				OpenStackCoreFlags: flags.OpenStackCoreFlags{
					MetadataPrefix: "test",
				},
			},
			func(meta map[string]string) error {
				if meta["test:gpu_vendor"] != "NVIDIA" {
					return errors.New("incorrect 'gpu_vendor' metadata")
				}
				if meta["test:gpu_driver_version"] != "v470" {
					return errors.New("incorrect 'gpu_driver_version' metadata")
				}
				if meta["test:gpu_models"] != "RTX" {
					return errors.New("incorrect 'gpu_models' metadata")
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateBuilderMetadata(tt.options)
			if err := tt.validate(got); err != nil {
				t.Errorf("Validation failed: %v", err)
			}
		})
	}
}
