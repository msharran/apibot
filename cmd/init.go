/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"github.com/sharran-murali/apibot/src/config"
	"github.com/sharran-murali/apibot/src/constants"
	"github.com/sharran-murali/apibot/src/utils"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes apibot config file",
	Long:  `Creates or updates an existing config file in ~/.apibot/config.yaml with user provided environment profile.`,
	Run: func(cmd *cobra.Command, args []string) {
		if utils.IsFileNotExist(utils.GetApiBotDir()) {
			err := os.Mkdir(utils.GetApiBotDir(), constants.DirPerm)
			cobra.CheckErr(err)
		}

		profileName, _ := rootCmd.PersistentFlags().GetString("profile")
		baseURL, _ := cmd.Flags().GetString("base-url")
		authorizationHeader, _ := cmd.Flags().GetString("authorization-header")

		profile := config.Profile{
			BaseUrl:             baseURL,
			AuthorizationHeader: authorizationHeader,
		}

		config.UpdateProfile(profileName, profile)
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
	rootCmd.PersistentFlags().StringP("profile", "p", "default", "You can use multiple profiles for your client. By default the profile name is default")
	initCmd.Flags().StringP("base-url", "b", "", "The base url for the API")
	initCmd.Flags().StringP("authorization-header", "a", "", `Authorization header for the API. eg., -a="Bearer *token*"`)

	initCmd.MarkFlagRequired("base-url")
	initCmd.MarkFlagRequired("authorization-header")
}
