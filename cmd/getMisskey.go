/*
Copyright © 2024 Soli
*/
package cmd

import (
	"log"

	"github.com/Soli0222/spn-cli/misskey"
	"github.com/Soli0222/spn-cli/modules"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var getMisskeyCmd = &cobra.Command{
	Use:   "getMisskey",
	Short: "Retrieve your Misskey user profile information",
	Long: `Fetch detailed information about your Misskey user profile, 
including username, display name, and other relevant account details.`,
	Run: func(cmd *cobra.Command, args []string) {
		tokenFilePath, err := modules.GetTokenFilePath("token_misskey.json")
		if err != nil {
			log.Fatal(err)
		}

		if credentials, err := modules.LoadMisskey(tokenFilePath); err == nil {
			log.Println(color.GreenString("Using saved token"))
			if err != nil {
				log.Fatal(err)
			}

			name, username, err := misskey.FetchUserProfile(credentials.Hostname, credentials.Token)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("Successfully logged in Misskey\n%s (@%s@%s)", name, username, credentials.Hostname)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(getMisskeyCmd)
}
