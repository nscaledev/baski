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

package sign

import (
	"github.com/nscaledev/baski/pkg/provisoner"
	"github.com/nscaledev/baski/pkg/util/flags"
	"github.com/nscaledev/baski/pkg/util/sign"
	"github.com/spf13/cobra"
	"os"
)

// NewSignValidateCommand creates a command that allows the signing of an image.
func NewSignValidateCommand() *cobra.Command {

	o := &flags.SignOptions{}

	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate digital signature",
		Long: `Validate digital signature.

This just validates a signature. It's useful for verifying a signed image.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			o.SetOptionsFromViper()

			var key []byte
			var err error

			vaultClient := sign.VaultClient{
				Endpoint: o.VaultURL,
				Token:    o.VaultToken,
			}
			if len(o.PublicKey) != 0 {
				key, err = os.ReadFile(o.PublicKey)
				if err != nil {
					return err
				}
			} else if len(o.VaultURL) != 0 {
				key, err = vaultClient.Fetch(o.VaultMountPath, o.VaultSecretPath, "public-key")
				if err != nil {
					return err
				}
			}

			signer := provisoner.NewSigner(o)
			err = signer.ValidateImage(key)
			if err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
