/*
Copyright © 2024 Soli
*/
package cmd

import (
	"fmt"
	"log"
	"net/url"
	"os/exec"
	"runtime"

	"github.com/Soli0222/spn-cli/modules"
	"github.com/Soli0222/spn-cli/spotify"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var tweetCmd = &cobra.Command{
	Use:   "tweet",
	Short: "Share the currently playing Spotify track to Twitter",
	Long: `The tweet command allows you to share the currently playing track on Spotify to your Twitter account.
You can specify the track ID to share a specific track or share the currently playing track automatically.`,
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
			currentPlayingInfo := fmt.Sprintf("%s / %s\n%s",
				trackName, artistNames, trackURL)
			log.Println("Currentry Playing:", currentPlayingInfo)
			openbrowser("https://x.com/intent/tweet?url=" + trackURL + "&text=" + url.QueryEscape(fmt.Sprintf("%s / %s\n#NowPlaying #PsrPlaying", trackName, artistNames)))

			return
		}
	},
}

func openbrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Success Open Twitter")
}

func init() {
	rootCmd.AddCommand(tweetCmd)
}
