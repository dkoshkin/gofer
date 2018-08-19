// Copyright © 2018 Dimitri Koshkin
//
// Licensed under the Apache License, Version 2.0 (the "License");
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

	"github.com/dkoshkin/gofer/pkg/dependency"
	"github.com/dkoshkin/gofer/pkg/dependency/manager"
	"github.com/spf13/cobra"
)

var mask string
var sourceType string

var validTypes = []string{"github", "docker", "manual"}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add name version",
	Args:  cobra.MinimumNArgs(2),
	Short: "Add a dependency to your config file",
	RunE: func(cmd *cobra.Command, args []string) error {
		// read config file and add the dependency
		mngr := manager.NewFileManager(cfgFile)
		manifest, err := mngr.Read()
		if err != nil {
			return err
		}
		dep := dependency.Spec{
			Type:    sourceType,
			Name:    args[0],
			Version: args[1],
			Mask:    mask,
		}
		if sourceType != "" {
			if !stringInSlice(sourceType, validTypes) {
				return fmt.Errorf("dependency not added, %q is not a valid type", sourceType)
			}
			dep.Type = sourceType
		} else {
			dep.Type = dependency.DetermineType(args[1])
		}
		if dep.Type == dependency.UnknownType {
			fmt.Fprintf(out, "Could not determine source type, setting as %q", dependency.UnknownType)
		}
		added := manifest.Append(dep)
		if !added {
			return fmt.Errorf("dependency not added, %q is already in the config file", dep.Name)
		}
		err = mngr.Write(*manifest)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringVar(&mask, "mask", "", "a regex to match 'version', leave blank to match any version")
	addCmd.Flags().StringVar(&sourceType, "type", "", "source type, leave empty to autodetect (options \"github\"|\"docker\"|\"manual\")")
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
