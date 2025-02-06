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
	if !utils.Auth() {
		os.Exit(1)
	}
	cmd.Execute()
}