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
	"time"

	spinnerLib "github.com/briandowns/spinner"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var profileName = "default"
var client = resty.New()
var spinner = spinnerLib.New(spinnerLib.CharSets[11], 100*time.Millisecond)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "apibot",
	Short: "An API client which makes life easier for developers",
	Long:  `An API client which makes life easier for developers`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hi, from ApiBot")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.apibot.yaml)")
	rootCmd.PersistentFlags().StringVarP(&profileName, "profile", "p", "default", "You can use multiple profiles for your client. By default the profile name is default")
	spinner.Color("blue")
	spinner.Prefix = "Please wait... "

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
