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

package ecdsa

import (
	"github.com/spf13/cobra"
)

// Handler handles ECDSA commands
type Handler interface {
	// Handle creates ECDSA key(s)
	Invoke(curve string, secret []byte, fileName string, createPublicKey bool) error
}

// Command provides ECDSA commands for keygen
func Command(handler Handler) *cobra.Command {
	var curve string
	// var secret string
	var command = &cobra.Command{
		Use:     "ecdsa",
		Short:   "Creates an ECDSA key(pair).",
		Long:    "Creates an ECDSA key(pair). Creates PEM-encoded key(s) in SEC 1, ASN.1 DER form.",
		Example: "  keygen ecdsa -c P521 -o id_ecdsa",
		RunE: func(cmd *cobra.Command, args []string) error {

			file, _ := cmd.Flags().GetString("out")
			public, _ := cmd.Flags().GetBool("public")
			secret, _ := cmd.Flags().GetString("secret")

			return handler.Invoke(curve, []byte(secret), file, public)
		},
	}
	command.Flags().StringVarP(&curve, "curve", "c", "", "One of P224|P256|P384|P521.")
	command.MarkFlagRequired("curve")
	// command.Flags().StringVarP(&secret, "secret", "s", "", "A secret protecting your private key.")
	return command
}
