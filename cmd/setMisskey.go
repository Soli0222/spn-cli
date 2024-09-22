/*
Copyright © 2024 Soli
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Soli0222/spn-cli/modules"
	"github.com/spf13/cobra"
)

// setMisskeyCmd represents the setMisskey command
var setMisskeyCmd = &cobra.Command{
	Use:   "setMisskey",
	Short: "Set Misskey hostname and token",
	Long: `Set the Misskey instance hostname and authentication token to be used by the application.

Example usage:
  spn-cli setMisskey --hostname example.com --token your_token_here`,
	Run: func(cmd *cobra.Command, args []string) {
		hostname, _ := cmd.Flags().GetString("hostname")
		token, _ := cmd.Flags().GetString("token")

		if hostname == "" || token == "" {
			log.Fatal("Both hostname and token are required")
		}

		tokenFilePath, err := modules.GetTokenFilePath("token_misskey.json")
		if err != nil {
			log.Fatal(err)
		}

		if err := saveMisskey(tokenFilePath, hostname, token); err != nil {
			log.Fatal(fmt.Errorf("error saving token: %w", err))
		}

		fmt.Println("Misskey credentials saved successfully!")
	},
}

func saveMisskey(filePath, hostname, token string) error {
	os.MkdirAll(filepath.Dir(filePath), 0700)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data := map[string]string{
		"hostname": hostname,
		"token":    token,
	}

	return json.NewEncoder(file).Encode(data)
}

func init() {
	setMisskeyCmd.Flags().StringP("hostname", "H", "", "Misskey instance hostname")
	setMisskeyCmd.Flags().StringP("token", "T", "", "Misskey authentication token")

	rootCmd.AddCommand(setMisskeyCmd)
}
