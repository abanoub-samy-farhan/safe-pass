package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/abanoub-samy-farhan/safe-pass/client"
	"github.com/abanoub-samy-farhan/safe-pass/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use: "restore",
	Short: "Restoring a backup snapshot of the current state database",
	Run: restoreData,
}

func restoreData(cmd *cobra.Command, args []string) {
	dir := string(os.Getenv("BACKUP"))
	content, _ := os.ReadDir(dir)

	if len(content) == 0 {
		fmt.Println("No backup files found")
		return
	}

	var backups []string

	for _, item := range content {
		if !item.IsDir() {
			backups = append(backups, item.Name())
		}
	}

	searcher := func (input string, index int) bool {
		return strings.Contains(strings.ToLower(backups[index]), strings.ToLower(input))
	}
	prompt := promptui.Select{
		Label: "Select a backup file to restore: ",
		Items: backups,
		Searcher: searcher,
		StartInSearchMode: true,
	}

	_ , chosenBackup, err := prompt.Run()
	if err != nil {
		fmt.Println("Error while choosing the backup")
		return
	}

	fmt.Println("Restoring backup: ", chosenBackup)

	// unzip the backup file
	backupPath := dir + chosenBackup
	decompress := exec.Command("gzip", "-d", backupPath)
	decompress.Run()
	backupPath = strings.TrimSuffix(backupPath, ".gz")

	cipheredData, _ := os.ReadFile(backupPath)
	plainDataString := utils.DecryptData(string(cipheredData))

	var snapshot JSONBack
	er := json.Unmarshal([]byte(plainDataString), &snapshot)
	if er != nil {
		fmt.Println("Error while unmarshalling JSON: ", er)
		return
	}
	auth := client.InitiateClient(0)
	defer auth.Close()

	ctx := context.Background()

	for _, entry := range snapshot.Data {
		val := 	auth.Get(ctx, entry.Key).Val()
		if val != "" {
			overridePrompt := promptui.Prompt{
				Label:     fmt.Sprintf("Key '%s' already exists. Override", entry.Key),
				IsConfirm: true,
			}
			_, err := overridePrompt.Run()
			if err != nil {
				continue
			}
		}
		auth.Set(ctx, entry.Key, entry.Val, 0)
		fmt.Printf("Key '%s' restored\n", entry.Key)
	}

	// recompress the data
	compress := exec.Command("gzip", backupPath)
	compress.Run()
}

func init(){
	rootCmd.AddCommand(restoreCmd)
}