package utils

import (
	"os"

	"github.com/sharran-murali/apibot/src/constants"
	"github.com/spf13/cobra"
)

func GetFileName(fileName string) string {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)

	return homeDir + "/" + constants.ApiBotDir + "/" + fileName
}

func GetApiBotDir() string {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)

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
