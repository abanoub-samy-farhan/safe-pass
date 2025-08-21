package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/abanoub-samy-farhan/safe-pass/client"
	"github.com/abanoub-samy-farhan/safe-pass/utils"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use: "backup",
	Short: "Making backup snapshot of the current state database",
	Run: backupData,
}

func backupData(cmd *cobra.Command, args []string){
	filename := utils.GetSnapshotName()
	// open a new file in the location at the backup
	auth := client.InitiateClient(0)
	if auth == nil {
		fmt.Println("Redis client is not configured properly, please check the redis database status")
		return
	}

	defer auth.Close()
	ctx := context.Background()

	currentBackup := auth.Keys(ctx, "*").Val()
	if len(currentBackup) == 0 {
		fmt.Println("No keys found to backup")
		return
	}

	jsonData := JSONBack{
		Data: make([]JSONEntry, 0),
	}

	for _, key := range currentBackup {
		val := auth.Get(ctx, key).Val()
		entry := JSONEntry{
			Key: key,
			Val: val,
		}
		jsonData.Data = append(jsonData.Data, entry)
	}

	plainText, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("Error happened while converting data to json", err.Error())
	}
	cipheredText := utils.EncryptData(string(plainText))

	snapshotPath :=  string(os.Getenv("BACKUP")) + "/" + filename

	backupFile, err := os.Create(snapshotPath)
	if err != nil {
		fmt.Println("Error happened while opening the backup file")
	}
	defer backupFile.Close()

	backupFile.WriteString(cipheredText)
	compress := exec.Command("gzip", snapshotPath)
	compress.Run()

	fmt.Println(Green + "Backup is created at: " + snapshotPath)
}

func init(){
	rootCmd.AddCommand(backupCmd)
}