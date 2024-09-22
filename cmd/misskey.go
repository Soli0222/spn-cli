/*
Copyright © 2024 Soli
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/Soli0222/spn-cli/misskey"
	"github.com/Soli0222/spn-cli/modules"
	"github.com/Soli0222/spn-cli/spotify"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// misskeyCmd represents the misskey command
var misskeyCmd = &cobra.Command{
	Use:   "misskey",
	Short: "Share the currently playing Spotify track to Misskey",
	Long: `The misskey command allows you to share the currently playing track on Spotify to your Misskey account.
You can specify the track ID to share a specific track or share the currently playing track automatically.`,
	Run: func(cmd *cobra.Command, args []string) {
		spotifytokenFilePath, err := modules.GetTokenFilePath("token.json")
		if err != nil {
			log.Fatal(err)
		}

		misskeytokenFilePath, err := modules.GetTokenFilePath("token_misskey.json")
		if err != nil {
			log.Fatal(err)
		}

		if tok, err := modules.LoadToken(spotifytokenFilePath); err == nil {
			if !tok.Valid() {
				log.Println("Token is expired or invalid")
				return
			}
			log.Println(color.GreenString("Using saved token"))
			client := conf.Client(ctx, tok)
			trackName, artistNames, trackURL, err := spotify.FetchCurrentryPlaying(client)
			if err != nil {
				log.Fatal(err)
			}
			currentPlayingInfo := fmt.Sprintf("\n%s / %s\n%s\n#NowPlaying #PsrPlaying\n",
				trackName, artistNames, trackURL)
			log.Println("Currentry Playing:", currentPlayingInfo)

			if credentials, err := modules.LoadMisskey(misskeytokenFilePath); err == nil {
				log.Println(color.GreenString("Using saved token"))
				if err != nil {
					log.Fatal(err)
				}

				id, err := misskey.PostNote(credentials.Hostname, credentials.Token, currentPlayingInfo)
				if err != nil {
					log.Fatal(err)
				}

				log.Printf("Successfully note to Misskey\nNote URL: https://%s/notes/%s", credentials.Hostname, id)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(misskeyCmd)
}
