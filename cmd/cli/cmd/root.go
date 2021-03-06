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
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

const (
	apiVersion = "v0.1"
)

type Version struct {
	Version   string `json:"Version"`
	BuildDate string `json:"BuildDate"`
}

var cfgFile string
var out = os.Stdout
var errOut = os.Stderr

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gofer",
	Short: "A CLI utility to help you keep your project's ever-changing dependency versions up to date",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version Version) {
	// set version string
	bytes, _ := json.MarshalIndent(version, "", "    ")
	rootCmd.SetVersionTemplate(string(bytes) + "\n")
	// also need to set String to get cobra to print it
	rootCmd.Version = version.Version
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "f", "./.gofer/config.yaml", "config file containing the list of dependencies")
}
