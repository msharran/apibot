package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sharran-murali/apibot/src/constants"
	"github.com/sharran-murali/apibot/src/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func GetConfigFile() string {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)

	return homeDir + "/" + constants.ApiBotDir + "/config.yaml"
}

func UpdateProfile(name string, profile Profile) {
	profiles := make(map[string]Profile)
	statusMsg := "created"

	if utils.IsFileExist(GetConfigFile()) {
		existingConfigData, err := ioutil.ReadFile(GetConfigFile())
		cobra.CheckErr(err)

		err = yaml.Unmarshal(existingConfigData, &profiles)
		cobra.CheckErr(err)
		statusMsg = "updated"
	}

	profiles[name] = profile
	newConfigData, err := yaml.Marshal(&profiles)
	cobra.CheckErr(err)

	err = ioutil.WriteFile(GetConfigFile(), newConfigData, constants.FilePerm)
	cobra.CheckErr(err)

	fmt.Printf("ApiBot profile %v successfully!\n", statusMsg)
}

func GetProfile(name string) Profile {
	profiles := make(map[string]Profile)

	existingConfigData, err := ioutil.ReadFile(GetConfigFile())
	cobra.CheckErr(err)

	err = yaml.Unmarshal(existingConfigData, &profiles)
	cobra.CheckErr(err)

	if _, ok := profiles[name]; !ok {
		utils.ExitWithMsg("Invalid Profile")
	}

	return profiles[name]
}
