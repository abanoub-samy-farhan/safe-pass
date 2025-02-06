package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/abanoub-samy-farhan/safe-pass/client"
	"strings"
	"context"
)

var deleteCmd = &cobra.Command{
	Use: "delete -c <category> -d <domain> -t <tag>",
	Short: "Delete a password, key or token from the database",
	Long: `Delete a password, key or token from the database for a specific domain and tag.
	Example:
		safe-pass delete password -c passwords -d google -t work

	Take Care, this action is IRREVOCABLE
	`,
	Run: deleteData,
}

func deleteData(cmd *cobra.Command, args []string) {
	category, _ := cmd.Flags().GetString("category")
	domain, _ := cmd.Flags().GetString("domain")
	tag, _ := cmd.Flags().GetString("tag")

	category = strings.ToLower(category)
	domain = strings.ToLower(domain)
	tag = strings.ToLower(tag)

	var assertion string
	fmt.Printf("Are you sure you want to delete the data for \nCategory " + Red + category + Reset +
	", domain " + Red + domain + Reset + " and tag " + Red + tag + Reset + " ? \n\nProceed (Y/n) ")
	fmt.Scanln(&assertion)
	if strings.ToUpper(assertion) != "Y" {
		fmt.Println("Action cancelled")
		return
	}

	client := client.InitiateClient()
	defer client.Close()

	ctx := context.Background()
	lookup := "*"
	if category != ""{
		lookup += category + "-"
	}
	if domain != ""{
		lookup += domain + ":"
	}
	if tag != ""{
		lookup += tag
	}
	lookup = lookup + "*";

	keys := client.Keys(ctx, lookup).Val()
	if len(keys) == 0 {
		fmt.Println("There are no data found matching your request")
	}

	_, err := client.Del(ctx, keys...).Result()
	if err != nil {
		fmt.Println("An error occurred while deleting the data: ", err)
		return
	}

	fmt.Println(Green + "Your Data is deleted successfully!" + Reset)
}

func init(){
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("category", "c", "passwords", "Category of the data")
	deleteCmd.Flags().StringP("domain", "d", "default", "Domain of the data")
	deleteCmd.Flags().StringP("tag", "t", "default", "Tag of the data")
}