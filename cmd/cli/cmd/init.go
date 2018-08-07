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

	"github.com/dkoshkin/gofer/pkg/dependency/manager"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize an empty 'config.yaml' file",
	Long:  `Initialize an empty 'config.yaml' file with the current API version and an empty list of dependencies.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		mngr := manager.NewFileManager(cfgFile)
		if _, err := mngr.Init(apiVersion); err != nil {
			return err
		}
		fmt.Fprintf(out, "Wrote new config file to %q\n", cfgFile)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
