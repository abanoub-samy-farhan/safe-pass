package cmd

import (
	"fmt"
	"math/rand"
	"github.com/spf13/cobra"
)

var passgenCmd = &cobra.Command{
	Use: "passgen",
	Short: "Password Generator",
	Long: `Password Generator for making the user's life easier 
	by simply hitting the generation, specifing some flags like:
	-l --length - determine the length of the password
	-s --special-characters - determine wether the password contains special characters or not
	-n --numbers - determine wether the password contains numbers or not

	Example:
		safe-pass passgen -l 20 -s -n

	You could also use it to be redirected to your databases:
		safe-pass passgen -l 20 -s -n | safe-pass add -c passwords -d <domain> -t <tag>
	`,
	Run: generatePassword,
}

func generatePassword(cmd *cobra.Command, args []string){
	length, _ := cmd.Flags().GetInt("length")
	isSpecial, _ := cmd.Flags().GetBool("special-characters")
	isNumber, _ := cmd.Flags().GetBool("numbers")

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if isNumber {
		charset += "0123456789"
	}
	if isSpecial {
		charset += "!@#$%^&*()"
	}

	var newpass string
	for i := 0; i < length; i++ {
		newpass += string(charset[rand.Intn(len(charset))])
	}

	fmt.Println(newpass)
}

func init(){
	rootCmd.AddCommand(passgenCmd)
	passgenCmd.Flags().IntP("length", "l", 20, "Length of the password")
	passgenCmd.Flags().BoolP("special-characters", "s", false, "Include special characters")
	passgenCmd.Flags().BoolP("numbers", "n", false, "Include numbers")
}