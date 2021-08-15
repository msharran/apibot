package utils

import (
	"fmt"
	"os"
)

func ExitWithError(err error) {
	fmt.Println("[error] " + err.Error())
	os.Exit(1)
}

func ExitWithMsg(msg string) {
	fmt.Println("[error] " + msg)
	os.Exit(1)
}
