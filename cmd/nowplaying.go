/*
Copyright © 2024 Soli
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/Soli0222/spn-cli/modules"
	"github.com/Soli0222/spn-cli/spotify"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var nowplayingCmd = &cobra.Command{
	Use:   "nowplaying",
	Short: "Get currently playing song along with the artist and track URL from Spotify.",
	Long: `This command retrieves information about the currently playing song from the Spotify API, including the song title, artist(s), and a direct URL to the track on Spotify.
It formats the data as “Song Title / Artist(s)” followed by the track’s Spotify URL, allowing users to easily access and share the song they are listening to.`,
	Run: func(cmd *cobra.Command, args []string) {
		tokenFilePath, err := modules.GetTokenFilePath("token.json")
		if err != nil {
			log.Fatal(err)
		}

		if tok, err := modules.LoadToken(tokenFilePath); err == nil {
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
			currentPlayingInfo := fmt.Sprintf("\n%s / %s\n%s",
				trackName, artistNames, trackURL)
			log.Println("Currentry Playing:", currentPlayingInfo)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(nowplayingCmd)
}
