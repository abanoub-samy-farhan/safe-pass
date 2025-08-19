package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/abanoub-samy-farhan/safe-pass/client"
	"github.com/abanoub-samy-farhan/safe-pass/utils"
	"context"
	"strings"
	"time"
)

var showCmd = &cobra.Command{
	Use: "show",
	Short: "Show all data or data by category, domain and tag",
	Long: `Show all data or data by category, domain and tag. Running it without any flags would 
	return the database of all the values found
	Example:
		safe-pass show
	or you can be specifying the category:
		safe-pass show -c <category>
	or specifiying the domain and/or tag:
		safe-pass show -c <category> -d <domain> -t <tag>
	`,
	Run: showData,
}

func showData(cmd *cobra.Command, args []string){
	category,_ := cmd.Flags().GetString("category")
	domain, _ := cmd.Flags().GetString("domain")
	tag, _ := cmd.Flags().GetString("tag")

	category = strings.ToLower(category)
	domain = strings.ToLower(domain)
	tag = strings.ToLower(tag)

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

	parsed := make(map[string][]string)

	for _, key := range keys {
		cat := strings.Split(key, "-")[0]
		parsed[cat] = append(parsed[cat], key)
	}

	starttime := time.Now()
	for key, values := range parsed {
		fmt.Println(Green + "Category: " + key + Reset)
		for _, value := range values {
			encryptedData, err := client.Get(ctx, value).Result()
			if err != nil {
				panic(err)
			}

			parsedKeys := strings.Split(value, "-")
			domtagpair := parsedKeys[len(parsedKeys)-1]
			dom := strings.Split(domtagpair, ":")[0]
			tag := strings.Split(domtagpair, ":")[1]

			fmt.Println(Red + "\tDomain: " + dom + "\tTag: " + tag +
			":" + Reset, utils.DecryptData(encryptedData))
		}
	}

	fmt.Println("Time elapsed: ", time.Since(starttime))
}

func init(){
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().StringP("category", "c", "", "Category of the data")
	showCmd.Flags().StringP("domain", "d", "", "Domain of the data")
	showCmd.Flags().StringP("tag", "t", "", "Tag of the data")
}