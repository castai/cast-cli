/*
Copyright © 2020 CAST AI

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
	"fmt"

	"github.com/spf13/cobra"

	"github.com/castai/cast-cli/pkg/version"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"-v"},
	Short:   "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version: %s, commit: %s\n", version.Version, version.Commit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
