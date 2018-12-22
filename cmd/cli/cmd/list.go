// Copyright © 2018 Dimitri Koshkin
//
// Licensed under the Apache License, String 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"strings"

	"github.com/dkoshkin/gofer/pkg/dependency"
	"github.com/dkoshkin/gofer/pkg/dependency/manager"
	"github.com/spf13/cobra"
)

var outputTypes = []string{"table", "yaml", "json"}
var output string
var outdated bool
var types []string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the dependencies from the 'config.yaml' file",
	RunE: func(cmd *cobra.Command, args []string) error {
		var validOutputType bool
		for _, t := range outputTypes {
			if t == output {
				validOutputType = true
			}
		}
		if !validOutputType {
			return fmt.Errorf("output %q is not valid", output)
		}

		// read config file and print dependencies
		mngr := manager.NewFileManager(cfgFile)
		manifest, err := mngr.Read()
		if err != nil {
			return err
		}
		writeManifest(manifest, output, outdated, types)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listCmd.Flags().StringVarP(&output, "output", "o", "yaml", "output format (options \"table\"|\"yaml\"|\"json\")")
	listCmd.Flags().BoolVar(&outdated, "outdated", false, "only list the dependencies that have outdated versions")
	listCmd.Flags().StringSliceVar(&types, "types", []string{}, "source type(s), leave empty to select all (options \"github\"|\"docker\"|\"manual\")")
}

func writeManifest(manifest *dependency.Manifest, outputType string, onlyOutdated bool, types []string) {
	mw := dependency.ManifestWriter{
		Writer:        out,
		FilterOptions: dependency.FilterOptions{Outdated: onlyOutdated, Types: types},
	}
	switch output {
	case "table":
		fmt.Fprintln(out, strings.Repeat("-", 120))
		mw.WriteTable(*manifest)
	case "yaml":
		mw.WriteYAML(*manifest)
	case "json":
		mw.WriteJSON(*manifest)
	}

}
