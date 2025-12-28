package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/abanoub-samy-farhan/safe-pass/client"
	"github.com/abanoub-samy-farhan/safe-pass/utils"
	"github.com/atotto/clipboard"

	"github.com/spf13/cobra"

	"golang.org/x/exp/maps"
)

var showCmd = &cobra.Command{
	Use: "show",
	Example: "safe-pass show",
	Short: "Get a specific entry from the database",
	Run: showData,
}

func ParseKey(key string) (string, string, string, string) {
	splitted := strings.Split(key, "-")
	cat, specifier := splitted[0], splitted[1]
	dom, tag := strings.Split(specifier, ":")[0], strings.Split(specifier, ":")[1]
	displayed := cat + " -> " + utils.MakeColored("Green", dom) + 
	":" + utils.MakeColored("Green", tag)
	return cat, dom, tag, displayed
}

func showData(cmd *cobra.Command, args []string){

	client := client.InitiateClient(0)
	defer client.Close()

	ctx := context.Background()
	keys := client.Keys(ctx, "*").Val()
	if len(keys) == 0 {
		fmt.Println("There are no data found matching your request")
	}

	parsed := make(map[string][]string)

	for _, key := range keys {
		cat, _, _, _ := ParseKey(key)
		parsed[cat] = append(parsed[cat], key)
	}

	selectedCategory, err := utils.PromptSelect(utils.PromptOpts{
		Message: "Select a category to show: ",
		Items: maps.Keys(parsed),
	})
	if err != nil {
		return
	}
	selectedKeys := parsed[selectedCategory]

	// make a list of displayed keys for the selected category
	displayedKeys := make(map[string]string)
	for _, key := range selectedKeys {
		_, _, _, dispKey := ParseKey(key)
		displayedKeys[dispKey] = key
	}

	if len(selectedKeys) == 0 {
		fmt.Println("No data found in the selected category")
		return
	}

	selectedDisplayedKey, err := utils.PromptSelect(utils.PromptOpts{
		Message: "Select an entry to show: ",
		Items: maps.Keys(displayedKeys),
		UseSearcher: true,
	})

	if err != nil {
		return
	}

	selectedKey := displayedKeys[selectedDisplayedKey]

	encryptedData, err := client.Get(ctx, selectedKey).Result()
	if err != nil {
		panic(err)
	}
	data := utils.DecryptData(encryptedData)
	clipboard.WriteAll(data)
	fmt.Println(utils.MakeColored("Green", "Data is copied to your clipboard"))
}

func init(){
	rootCmd.AddCommand(showCmd)
}