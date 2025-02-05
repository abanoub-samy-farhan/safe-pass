/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "safe-pass",
	Short: "A CLI app for securly storing passwords, keys and tokens for you :)",
	Long: `safe-pass is a tool for managing sensitive data such as passwords, keys
and tokens. It uses a redis database for storing the data, and provides a
simple and secure way to interact with the data.

The tool is designed to be used from the command line, and provides a number of
commands to add, list, retrieve and delete data from the database.

The tool also provides a setup command, which will prompt the user to enter a
master password. This password is used to encrypt the data in the database,
and is the only way to access the data.

The tool is written in Go, and is designed to be cross-platform and
extensible.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.safe-pass.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


