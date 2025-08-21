package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

type Field struct {
	category string
	domain   string
	tag      string
}

type JSONBack struct {
	Data []JSONEntry `json:"data"`
}

type JSONEntry struct {
	Key string `json:"key"`
	Val string `json:"value"`
}

var Reset = "\033[0m" 
var Red = "\033[31m" 
var Green = "\033[32m" 
var Yellow = "\033[33m" 
var Blue = "\033[34m" 
var Magenta = "\033[35m" 
var Cyan = "\033[36m" 
var Gray = "\033[37m" 
var White = "\033[97m"


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "safe-pass",
	Short: "A CLI app for securly storing passwords, keys and tokens for you :)",
	Long: `safe-pass is a tool for managing sensitive data such as passwords, keys
and tokens. It uses a redis database for storing the data, and provides a
simple and secure way to interact with the data.`,
	ValidArgs: []string{"add", "delete", "edit", "passgen", "show", "backup", "restore"},
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


