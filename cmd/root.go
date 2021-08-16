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
	"time"

	spinnerLib "github.com/briandowns/spinner"
	"github.com/sharran-murali/apibot/src/api"
	"github.com/sharran-murali/apibot/src/utils"
	"github.com/spf13/cobra"
)

var profileName = "default"
var headersString = ""
var headers map[string]string
var spinner = spinnerLib.New(spinnerLib.CharSets[39], 80*time.Millisecond)
var client = api.NewClient()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "apibot",
	Short: "An API client which makes life easier for developers",
	Long:  `An API client which makes life easier for developers`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hi, I am ApiBot. I am here to ease your API testing by managing all your different profiles")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	spinner.Prefix = utils.LogBlueSprintf("=> Please wait... ")
	cobra.OnInitialize(setHeaders)

	rootCmd.PersistentFlags().StringVarP(&profileName, "profile", "p", "default", "You can use multiple profiles for your client. By default the profile name is default")
	rootCmd.PersistentFlags().StringVarP(&headersString, "headers", "H", "", `Add custom request headers separated by comma. Eg., -h="key1:value1,key2:value2"`)
}

func setHeaders() {
	headersSlice := strings.Split(headersString, ",")
	headers = make(map[string]string)
	for _, header := range headersSlice {
		headerKVPair := strings.Split(header, ":")
		if len(headerKVPair) == 2 {
			headers[headerKVPair[0]] = headerKVPair[1]
		}
	}
	client.SetHeaders(headers)
}
