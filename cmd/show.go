package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/abanoub-samy-farhan/safe-pass/client"
	"github.com/abanoub-samy-farhan/safe-pass/utils"
	"github.com/atotto/clipboard"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"golang.org/x/exp/maps"
)

var showCmd = &cobra.Command{
	Use: "show",
	Example: "safe-pass show",
	Run: showData,
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
		cat := strings.Split(key, "-")[0]
		parsed[cat] = append(parsed[cat], key)
	}

	prompt := promptui.Select{
		Label: "Select a category to show: ",
		Items: maps.Keys(parsed),
	}
	_, selectedCategory, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed:", err)
		return
	}

	selectedKeys := parsed[selectedCategory]
	if len(selectedKeys) == 0 {
		fmt.Println("No data found in the selected category")
		return
	}

	searcher := func (curr string, ind int) bool {
		curr = strings.ToLower(curr)
		selected := strings.ToLower(selectedKeys[ind])
		return strings.Contains(selected, curr)
	}

	prompt = promptui.Select{
		Label: "Select a key to show: ",
		Items: selectedKeys,
		Searcher: searcher,
		StartInSearchMode: true,
	}
	_, selectedKey, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt failed:", err)
		return
	}

	encryptedData, err := client.Get(ctx, selectedKey).Result()
	if err != nil {
		panic(err)
	}
	data := utils.DecryptData(encryptedData)
	clipboard.WriteAll(data)
	fmt.Println(Green + "Data is copied to your clipboard")
}

func init(){
	rootCmd.AddCommand(showCmd)
}