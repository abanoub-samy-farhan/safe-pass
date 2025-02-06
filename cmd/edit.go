package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"context"
	"github.com/abanoub-samy-farhan/safe-pass/client"
	"github.com/abanoub-samy-farhan/safe-pass/utils"
	"github.com/redis/go-redis/v9"
)

var editCmd = &cobra.Command{
	Use: "edit [password|key|token] -c <category> -d <domain> -t <tag>",
	Short: "Edit a password, key or token",
	Long: `Edit a password, key or token from the database for a specific domain and tag.
	Example:
		safe-pass edit password -c passwords -d google -t work
	`,
	Run: editData,
}

func editData(cmd *cobra.Command, args []string) {

	if len(args) < 1 {
		cmd.Help()
		return
	}

	newdata := args[0]
	category, _:= cmd.Flags().GetString("category")
	domain, _:= cmd.Flags().GetString("domain")
	tag, _ := cmd.Flags().GetString("tag")

	category = strings.ToLower(category)
	domain = strings.ToLower(domain)
	tag = strings.ToLower(tag)

	client := client.InitiateClient()
	defer client.Close()

	ctx := context.Background()

	data, err := client.Get(ctx, fmt.Sprintf("%s-%s:%s", category, domain, tag)).Result()
	if err != redis.Nil {
		fmt.Println(`Data matching this inputs are not found, 
		Make sure you have this in your database
		
		Run:
			safe-pass show -c <category> -d <domain> -t <tag>
		`)
		return
	}

	fmt.Println(Green + "Your Data is: " + utils.DecryptData(data) + Reset)
	fmt.Println("Are you sure you want to change it to" + Red + newdata + Reset + " ? (Y/n) ")

	var input string
	fmt.Scanln(&input)

	if strings.ToUpper(input) == "Y" {
		encryptedData := utils.EncryptData(newdata)
		_, err := client.Set(ctx, fmt.Sprintf("%s-%s:%s", category, domain, tag), encryptedData, 0).Result()
		if err != nil {
			fmt.Printf(`An error occurred while adding the data: %s\n`, err)
			return
		}
		fmt.Println(Green + "Your Data is updated successfully!\nRun `safe-pass show -c", category, " -d", domain,  "-t", tag, "to view it`" + Reset)
	} else {
		fmt.Println(Red + "Action cancelled" + Reset)
	}
}

func init(){
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().StringP("category", "c", "passwords", "Category of the data")
	editCmd.Flags().StringP("domain", "d", "default", "Domain of the data")
	editCmd.Flags().StringP("tag", "t", "default", "Tag of the data")
}