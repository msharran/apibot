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
	"github.com/sharran-murali/apibot/src/botfactory"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Http GET method",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			botfactory.LogFatalErrorln("Missing argument endpoint: Eg, apibot get /example-endpoint")
		}
		endpoint := args[0]

		spinner.Start()
		resp, err := client.Request(profileName).Get(endpoint)
		spinner.Stop()
		client.Println(resp)

		botfactory.CheckErr(err)

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
