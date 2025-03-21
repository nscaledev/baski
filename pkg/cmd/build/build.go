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

package build

import (
	"fmt"
	"path/filepath"

	"github.com/nscaledev/baski/pkg/constants"
	"github.com/nscaledev/baski/pkg/providers/packer"
	"github.com/nscaledev/baski/pkg/provisoner"
	"github.com/nscaledev/baski/pkg/util/flags"
	"github.com/spf13/cobra"
)

// NewBuildCommand creates a command that allows the building of an image.
func NewBuildCommand() *cobra.Command {
	o := &flags.BuildOptions{}

	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build image",
		Long: `Build image.

Building images requires a set of commands to be run on the terminal however this is tedious and time consuming.
By using this, the time is reduced and automation can be enabled.`,
		TraverseChildren: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			o.SetOptionsFromViper()

			if !checkValidOSSelected(o.BuildOS) {
				return fmt.Errorf("an unsupported OS has been entered. Please select a valid OS: %s\n", constants.SupportedOS)
			}

			builder := provisoner.NewBuilder(o)

			// Init the provisioner
			err := builder.Init()
			if err != nil {
				return err
			}

			// Either use a local dir or fetch image-builder from a git repo
			buildGitDir := o.ImageRepoDir
			if buildGitDir == "" {
				buildGitDir = createRepoDirectory()
				if err := fetchBuildRepo(buildGitDir, o); err != nil {
					return err
				}
			}

			// Generate a packer config
			packerBuildConfig, err := builder.GeneratePackerConfig()
			if err != nil {
				return err
			}

			// If the builder requires it, modify it directly here.
			modifierFunc := packer.BuildersModifier{
				Function: builder.UpdatePackerBuilders,
				Metadata: packerBuildConfig.Metadata,
			}
			err = packer.UpdatePackerBuildersJson(buildGitDir, o.BaseOptions.InfraType, modifierFunc)
			if err != nil {
				return err
			}

			// Generate a tmp.json file to be consumed by the image-builder for variables.
			capiPath := filepath.Join(buildGitDir, "images", "capi")
			packerBuildConfig.GenerateVariablesFile(capiPath)

			// Install any dependencies
			installDependencies(capiPath, o.InfraType, o.Verbose)

			// Build the image
			err = buildImage(capiPath, o.InfraType, o.BuildOS, o.Verbose)
			if err != nil {
				return err
			}

			err = builder.PostBuildAction()
			if err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}

// checkValidOSSelected checks that the build os provided is a valid one.
func checkValidOSSelected(buildOS string) bool {
	for _, v := range constants.SupportedOS {
		if buildOS == v {
			return true
		}
	}
	return false
}
