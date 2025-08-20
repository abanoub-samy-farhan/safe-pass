package cmd

import (
	"context"
	"fmt"
	"os"
	"encoding/json"

	"github.com/abanoub-samy-farhan/safe-pass/client"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use: "safe-pass backup",
	Short: "for making backup snapshot of the current state database",
	Run: backupData,
}

func backupData(cmd *cobra.Command, args []string){

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


	tmpFile, err := os.Create("tmp.json")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	defer tmpFile.Close()

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

	if err := json.NewEncoder(tmpFile).Encode(jsonData); err != nil {
		fmt.Println("Error writing to temporary file:", err)
		return
	}
}

func init(){
	rootCmd.AddCommand(backupCmd)
}