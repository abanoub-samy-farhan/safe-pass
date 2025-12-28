package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/abanoub-samy-farhan/safe-pass/client"
	"github.com/abanoub-samy-farhan/safe-pass/utils"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use: "delete -c <category> -d <domain> -t <tag>",
	Short: "Delete a password, key or token from the database",
	Long: `Delete a password, key or token from the database for a specific domain and tag.
	Take Care, this action is IRREVOCABLE
	`,
	Example: "safe-pass delete password -c passwords -d google -t work",
	Run: deleteData,
}

func deleteData(cmd *cobra.Command, args []string) {
	category, _ := cmd.Flags().GetString("category")
	domain, _ := cmd.Flags().GetString("domain")
	tag, _ := cmd.Flags().GetString("tag")

	category = strings.ToLower(category)
	domain = strings.ToLower(domain)
	tag = strings.ToLower(tag)

	err := utils.PromptConfirm(
		fmt.Sprintf("You are about to delete the data for Category %s, domain %s and tag %s, Proceed?", 
		utils.MakeColored("Green",category), 
		utils.MakeColored("Green",domain), 
		utils.MakeColored("Green",tag)),
	)
	if err != nil {
		fmt.Println("Action cancelled")
		return
	}

	client := client.InitiateClient(0)
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

	_, err2 := client.Del(ctx, keys...).Result()
	if err2 != nil {
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