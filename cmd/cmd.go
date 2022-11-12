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

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tezli/keygen/cmd/ecdsa"
	"github.com/tezli/keygen/cmd/rsa"
	"github.com/tezli/keygen/cmd/version"
	"github.com/tezli/keygen/pkg/io"
)

var rootCmd = &cobra.Command{
	Use:              "keygen",
	Short:            "Creates key(pairs) for cryptographic purposes.",
	TraverseChildren: true,
}

// Execute executes the root command
func Execute(silent bool) error {
	if silent {
		rootCmd.SetHelpFunc(func(c *cobra.Command, s []string) {})
	}
	return rootCmd.Execute()
}

func init() {
	writer := io.OSWriter{}

	ecdsaHandler := ecdsa.NewHandler(writer)
	ecdsaCommand := ecdsa.Command(ecdsaHandler)
	rootCmd.AddCommand(ecdsaCommand)

	rsaHandler := rsa.NewHandler(writer)
	rsaCommand := rsa.Command(rsaHandler)
	rootCmd.AddCommand(rsaCommand)

	rootCmd.AddCommand(version.Command())

	rootCmd.Example = `
  keygen rsa -p -b 4096 -o id_rsa
  keygen ecdsa -p -c P521 -o id_ecdsa
	`
	rootCmd.PersistentFlags().StringP("out", "o", "", "Writes the key(s) to a file.")
	rootCmd.PersistentFlags().BoolP("public", "p", false, "Creates a public key as well.")
	rootCmd.TraverseChildren = true
}
