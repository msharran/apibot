package botfactory

import (
	"os"

	"github.com/fatih/color"
)

var redPrintf = color.New(color.Bold, color.FgHiRed).PrintfFunc()
var bluePrintf = color.New(color.Bold, color.FgBlue).PrintfFunc()
var blueSprintf = color.New(color.Bold, color.FgBlue).SprintfFunc()
var greenPrintf = color.New(color.Bold, color.FgGreen).PrintfFunc()
var yellowPrintf = color.New(color.Bold, color.FgYellow).PrintfFunc()
var yellowSprintf = color.New(color.Bold, color.FgYellow).SprintfFunc()

func LogFatalErrorln(err interface{}) {
	redPrintf("[Error] %v\n", err)
	os.Exit(1)
}

func CheckErr(err interface{}) {
	if err != nil {
		LogFatalErrorln(err)
	}
}

func LogInfoln(info interface{}) {
	yellowPrintf("%v\n", info)
}

func LogInfo(info interface{}) {
	yellowPrintf("%v", info)
}

func LogInfoSprintf(info interface{}) string {
	return yellowSprintf("%v", info)
}

func LogBlueln(info interface{}) {
	bluePrintf("%v\n", info)
}

func LogBlue(info interface{}) {
	bluePrintf("%v", info)
}

func LogBlueSprintf(info interface{}) string {
	return blueSprintf("%v", info)
}

func LogSuccessln(success interface{}) {
	greenPrintf("%v\n", success)
}

func LogSuccess(success interface{}) {
	greenPrintf("%v", success)
}

func LogErrorln(err interface{}) {
	redPrintf("%v\n", err)
}

func LogError(err interface{}) {
	redPrintf("%v", err)
}
