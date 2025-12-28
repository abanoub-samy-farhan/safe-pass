/*
Copyright 2025 abanoub-samy-farhan <abanoub-samy-farhan@gmail.com>

*/

package main

import (
	"os"
	"os/user"
	"fmt"

	"github.com/abanoub-samy-farhan/safe-pass/cmd"
	"github.com/abanoub-samy-farhan/safe-pass/utils"
	"github.com/joho/godotenv"
)

const envFilePath string = "/.config/safe-pass/.env"

func main() {
	if os.Getuid() == 0 && os.Args[1] == "init" {
		utils.Setup()
		return
	}
	homeDir := os.Getenv("HOME")
	if os.Getuid() == 0 {
		usr, _ := user.Lookup(os.Getenv("SUDO_USER"))
		homeDir = usr.HomeDir
	}

	envFile := homeDir + envFilePath
	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		fmt.Println("Please run", utils.MakeColored("Green", "sudo safe-pass init"), "to configure the environment")
		return
	}
	
	if !utils.Auth() {
		os.Exit(1)
	}
	cmd.Execute()
}