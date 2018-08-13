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
	"github.com/dkoshkin/gofer/pkg/dependency/docker"
	"github.com/dkoshkin/gofer/pkg/dependency/manager"
	"github.com/spf13/cobra"
)

var dryRun bool

// digCmd represents the dig command
var digCmd = &cobra.Command{
	Use:   "dig",
	Short: "Fetch the latest versions of all dependency from the 'config.yaml' file",
	RunE: func(cmd *cobra.Command, args []string) error {
		// read config file and print dependencies with latest versions
		mngr := manager.NewFileManager(cfgFile)
		manifest, err := mngr.Read()
		if err != nil {
			return err
		}

		updatedManifest := &dependency.Manifest{APIVersion: manifest.APIVersion}
		dc := docker.New()
		for _, dep := range manifest.Dependencies {
			depType := dep.GetType()
			switch depType {
			case dependency.DockerType:
				latest, err := dc.LatestVersion(dep.Name, dep.Mask)
				if err != nil {
					if err == docker.NoTags {
						dep.Notes = fmt.Sprintf("could not find latest tag")
					} else {
						dep.Notes = fmt.Sprintf("error retrieving latest tag: %v", err)
					}
				}
				dep.LatestVersion = latest
			case dependency.ManualType:
				continue
			case dependency.UnknownType:
				dep.Notes = fmt.Sprintf("could not determine type")
			default:
				dep.Notes = fmt.Sprintf("unhandled type %q", depType)
			}
			dep.Type = depType
			updatedManifest.Dependencies = append(updatedManifest.Dependencies, dep)
		}
		writeManifest(updatedManifest, output, false, []string{})

		if !dryRun {
			if err := mngr.Write(*updatedManifest); err != nil {
				return fmt.Errorf("error trying to write out config file: %v", err)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(digCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// digCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	digCmd.Flags().StringVarP(&output, "output", "o", "yaml", "output format to print to stdout (options \"table\"|\"yaml\"|\"json\")")
	digCmd.Flags().BoolVar(&dryRun, "dry-run", false, "don't overwrite the config file, just print to stdout")
}
