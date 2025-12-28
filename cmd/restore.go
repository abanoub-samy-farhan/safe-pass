package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
	"regexp"

	"github.com/abanoub-samy-farhan/safe-pass/client"
	"github.com/abanoub-samy-farhan/safe-pass/utils"
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

	re := regexp.MustCompile(`safe-pass-.*\.bin\.gz`)
	for _, item := range content {
		if !item.IsDir() &&  re.MatchString(item.Name()) {
			backups = append(backups, item.Name())
		}
	}

	// sort backups decendingly
	sort.Slice(backups, func(i, j int) bool {
		return backups[i] > backups[j]
	})

	chosenBackup, err := utils.PromptSelect(utils.PromptOpts{
		Message: "Select a backup file to restore: ",
		Items: backups,
		UseSearcher: true,
	})

	if err != nil {
		fmt.Println("Error while choosing the backup")
		return
	}

	fmt.Println("Restoring backup: ", chosenBackup)

	// unzip the backup file
	backupPath := dir + "/" + chosenBackup
	decompress := exec.Command("gzip", "-dk", backupPath)
	decompress.Run()
	backupPath = strings.TrimSuffix(backupPath, ".gz")

	fmt.Println(backupPath)

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
			err := utils.PromptConfirm(
				fmt.Sprintf("Key '%s' already exists. Override", entry.Key))
			if err != nil {
				continue
			}
		}
		auth.Set(ctx, entry.Key, entry.Val, 0)
		fmt.Println(utils.MakeColored("Green", "Key '" + entry.Key + "' restored"))
	}

	// remove the decompressed file
	os.Remove(backupPath)
	fmt.Println(utils.MakeColored("Green", "Backup restored successfully!"))
}

func init(){
	rootCmd.AddCommand(restoreCmd)
}