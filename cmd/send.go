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
	"fmt"
	"strings"

	"github.com/sharran-murali/apibot/src/botfactory"
	"github.com/sharran-murali/apibot/src/config"
	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send request defined in ~/.apibot/config.yaml",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			botfactory.LogFatalErrorln("Request name is missing. Eg., apibot send `request_name`")
		}
		reqName := args[0]
		profile := config.GetProfile(profileName)
		request := profile.Requests[reqName]

		fmt.Printf("Using profile: %s\n", profileName)
		fmt.Printf("Using request: %s\n", reqName)

		spinner.Start()
		resp, err := client.Request(profileName).
			SetBody(request.Body).
			Execute(strings.ToUpper(request.Method), request.Endpoint)
		spinner.Stop()
		botfactory.CheckErr(err)
		client.Println(resp)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
