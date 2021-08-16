package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sharran-murali/apibot/src/botfactory"
	"github.com/sharran-murali/apibot/src/constants"
	"gopkg.in/yaml.v2"
)

func GetConfigFile() string {
	homeDir, err := os.UserHomeDir()
	botfactory.CheckErr(err)

	return homeDir + "/" + constants.ApiBotDir + "/config.yaml"
}

func UpdateProfile(name string, profile Profile) {
	profiles := make(map[string]Profile)
	statusMsg := "created"

	if botfactory.IsFileExist(GetConfigFile()) {
		existingConfigData, err := ioutil.ReadFile(GetConfigFile())
		botfactory.CheckErr(err)

		err = yaml.Unmarshal(existingConfigData, &profiles)
		botfactory.CheckErr(err)
		statusMsg = "updated"
	}

	profiles[name] = profile
	newConfigData, err := yaml.Marshal(&profiles)
	botfactory.CheckErr(err)

	err = ioutil.WriteFile(GetConfigFile(), newConfigData, constants.FilePerm)
	botfactory.CheckErr(err)

	fmt.Printf("ApiBot profile %v successfully!\n", statusMsg)
}

func GetProfile(name string) Profile {
	profiles := make(map[string]Profile)

	existingConfigData, err := ioutil.ReadFile(GetConfigFile())
	botfactory.CheckErr(err)

	err = yaml.Unmarshal(existingConfigData, &profiles)
	botfactory.CheckErr(err)

	if _, ok := profiles[name]; !ok {
		fmt.Println()
		botfactory.LogFatalErrorln(fmt.Sprintf("No profile named '%v' found. Pls provide a valid profile name or create a new Profile using init command", name))
	}

	return profiles[name]
}
