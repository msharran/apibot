package utils

import (
	"fmt"
	"os"
)

const apiBotErrorPrefix = "[apibot] [error] "

func ExitWithError(err error) {
	fmt.Println(apiBotErrorPrefix + err.Error())
	os.Exit(1)
}

func ExitWithMsg(msg string) {
	fmt.Println(apiBotErrorPrefix + msg)
	os.Exit(1)
}
