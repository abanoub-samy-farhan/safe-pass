/*
Copyright 2025 abanoub-samy-farhan <abanoub-samy-farhan@gmail.com>

*/

package main

import (
	"os"
	"fmt"

	"github.com/abanoub-samy-farhan/safe-pass/cmd"
	"github.com/abanoub-samy-farhan/safe-pass/utils"
	"github.com/joho/godotenv"
)

func main() {
	envFile := os.Getenv("HOME") + "/.config/safe-pass/.env"
	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		fmt.Println("Please run `sudo /go/bin/safe-pass init` to configure the environment")
		return
	}
	if os.Getuid() == 0 && os.Args[1] == "init" {
		utils.Setup()
		return
	}
	if !utils.Auth() {
		os.Exit(1)
	}
	cmd.Execute()
}