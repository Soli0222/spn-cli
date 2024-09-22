/*
Copyright © 2024 Soli
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the current version of the SpotifyNowPlaying CLI tool",
	Long: `The "version" command outputs the current version of this SpotifyNowPlaying CLI tool.
It can be useful for checking the version installed and verifying compatibility.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version 1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
