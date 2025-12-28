package cmd

import (
	"fmt"
	"regexp"


	"github.com/spf13/cobra"
	"github.com/abanoub-samy-farhan/safe-pass/client"
	"github.com/abanoub-samy-farhan/safe-pass/utils"
	"strings"
	"context"
)

var addCmd = &cobra.Command{
	Use: "add '[password|key|token]' -c [category] -d [domain] -t [tag]",
	Short: "Add a password, key or token to the database, you can specify category, domain and tag",
	Example: "safe-pass add 'secret-password'",
	Run: addData,
}

func addData(cmd *cobra.Command, args []string){
	if len(args) < 1 {
		cmd.Help()
		return
	}

	data := args[0]
	category, _:= cmd.Flags().GetString("category")
	domain, _:= cmd.Flags().GetString("domain")
	tag, _ := cmd.Flags().GetString("tag")

	key := Field{
		category: category,
		domain:   domain,
		tag:      tag,
	}
	re := regexp.MustCompile(`[-:]`)
	for _, item := range []string{key.category, key.domain, key.tag} {
		if re.MatchString(item) || item == "" {
			fmt.Println("Category, Domain and Tag cannot contain '-' or ':' characters and cannot be empty")
			return
		}
	}

	client := client.InitiateClient(0)
	defer client.Close()

	ctx := context.Background()
	value, _ := client.Get(ctx, fmt.Sprintf("%s-%s:%s", key.category, key.domain, key.tag)).Result()


	if value != "" {
		fmt.Printf("The domain and tag provided are already found, tag will be overrided. \n\nProceed? (Y/n) ")
		var proceed string
		fmt.Scanln(&proceed)
		if strings.ToUpper(proceed) == "N" || strings.ToUpper(proceed) != "Y" {
			fmt.Println("Add cancelled, no changes made")
			return
		}
	}

	encryptedData := utils.EncryptData(data)
	_, err := client.Set(ctx, fmt.Sprintf("%s-%s:%s", key.category, key.domain, key.tag), encryptedData, 0).Result()
	if err != nil {
		fmt.Printf(`An error occurred while adding the data: %s\n`, err)
		return
	}

	fmt.Println("Your Data is saved successfully!\nRun", Red + 
	"`safe-pass show` to view it" + Reset)
}

func init(){
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("category", "c", "passwords", "Category of the data")
	addCmd.Flags().StringP("domain", "d", "default", "Domain of the data")
	addCmd.Flags().StringP("tag", "t", "default", "Tag of the data")
}