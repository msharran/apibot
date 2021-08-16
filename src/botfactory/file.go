package botfactory

import (
	"os"

	"github.com/sharran-murali/apibot/src/constants"
)

func GetFileName(fileName string) string {
	homeDir, err := os.UserHomeDir()
	CheckErr(err)

	return homeDir + "/" + constants.ApiBotDir + "/" + fileName
}

func GetApiBotDir() string {
	homeDir, err := os.UserHomeDir()
	CheckErr(err)

	return homeDir + "/" + constants.ApiBotDir
}

func IsFileExist(fileOrFolder string) bool {
	_, err := os.Stat(fileOrFolder)
	return !os.IsNotExist(err)
}

func IsFileNotExist(fileOrFolder string) bool {
	_, err := os.Stat(fileOrFolder)
	return os.IsNotExist(err)
}
