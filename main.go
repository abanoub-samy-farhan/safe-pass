/*
Copyright 2025 abanoub-samy-farhan <abanoub-samy-farhan@gmail.com>

*/

package main

import (
	"github.com/abanoub-samy-farhan/safe-pass/cmd"
	"github.com/abanoub-samy-farhan/safe-pass/utils"
	"os"
)

func main() {
	if os.Getuid() == 0 && os.Args[1] == "init" {
		utils.Setup()
		return
	}
	if !utils.Auth() {
		os.Exit(1)
	}
	cmd.Execute()
}