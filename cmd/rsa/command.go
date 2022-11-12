/*
Copyright © 2022 Robert Tezli robert.tezli+github@gmail.com

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

package rsa

import (
	"github.com/spf13/cobra"
)

// Handler handles RSA commands
type Handler interface {
	Invoke(bits int, secret []byte, fileName string, createPublicKey bool) error
}

// Command provides RSA commands for keygen
func Command(handler Handler) *cobra.Command {
	var bits int
	// var secret string
	var command = &cobra.Command{
		Use:     "rsa",
		Short:   "Creates RSA keypair",
		Example: "  keygen rsa -b 2048 -o id_rsa",
		Args:    func(cmd *cobra.Command, args []string) error { return nil },
		RunE: func(cmd *cobra.Command, args []string) error {

			file, _ := cmd.Flags().GetString("out")
			public, _ := cmd.Flags().GetBool("public")
			secret, _ := cmd.Flags().GetString("secret")

			return handler.Invoke(bits, []byte(secret), file, public)
		},
	}
	command.Flags().IntVarP(&bits, "bits", "b", 0, "keygen rsa -b 2048")
	command.MarkFlagRequired("bits")
	// command.Flags().StringVarP(&secret, "secret", "s", "", "A secret protecting your private key.")
	return command
}
