package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/Soli0222/spn-cli/modules"
	"github.com/Soli0222/spn-cli/spotify"
	"github.com/fatih/color"
	pkce "github.com/nirasan/go-oauth-pkce-code-verifier"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var (
	conf         *oauth2.Config
	ctx          context.Context
	srv          *http.Server
	wg           sync.WaitGroup
	codeVerifier *pkce.CodeVerifier
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Spotify",
	Long:  "Authenticate with Spotify and fetch user profile information.",
	Run: func(cmd *cobra.Command, args []string) {
		ctx = context.Background()
		conf = &oauth2.Config{
			ClientID: "6622cea4fe86432287916c67bc3cc4e5",
			Scopes:   []string{"user-read-currently-playing", "user-read-private", "user-read-playback-state"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://accounts.spotify.com/authorize",
				TokenURL: "https://accounts.spotify.com/api/token",
			},
			RedirectURL: "http://127.0.0.1:9999/oauth/callback",
		}

		codeVerifier, _ = pkce.CreateCodeVerifier()
		codeChallenge := codeVerifier.CodeChallengeS256()

		authURL := conf.AuthCodeURL("state", oauth2.AccessTypeOffline,
			oauth2.SetAuthURLParam("code_challenge", codeChallenge),
			oauth2.SetAuthURLParam("code_challenge_method", "S256"))

		log.Println(color.CyanString("You will now be taken to your browser for authentication"))
		time.Sleep(1 * time.Second)
		open.Run(authURL)
		time.Sleep(1 * time.Second)

		wg.Add(1)
		srv = &http.Server{Addr: ":9999"}
		http.HandleFunc("/oauth/callback", callbackHandler)
		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("HTTP server error: %v", err)
			}
		}()
		wg.Wait()
	},
}

func authenticate(r *http.Request) (*oauth2.Token, error) {
	queryParts, _ := url.ParseQuery(r.URL.RawQuery)
	code := queryParts["code"][0]

	tok, err := conf.Exchange(ctx, code, oauth2.SetAuthURLParam("code_verifier", codeVerifier.String()))
	if err != nil {
		return nil, fmt.Errorf("error exchanging code for token: %w", err)
	}

	tokenFilePath, err := modules.GetTokenFilePath("token.json")
	if err != nil {
		return nil, err
	}
	if err := saveToken(tokenFilePath, tok); err != nil {
		return nil, fmt.Errorf("error saving token: %w", err)
	}

	return tok, nil
}

func saveToken(filePath string, token *oauth2.Token) error {
	os.MkdirAll(filepath.Dir(filePath), 0700)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(token)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	tok, err := authenticate(r)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Authentication failed", http.StatusInternalServerError)
		return
	}

	client := conf.Client(ctx, tok)

	userProfile, err := spotify.FetchUserProfile(client)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Failed to fetch user profile", http.StatusInternalServerError)
		return
	}

	log.Println("Login successful! Welcome,", userProfile)

	msg := "<p><strong>Success!</strong></p>"
	msg += "<p>You are authenticated and your Spotify profile information has been retrieved.</p>"
	fmt.Fprint(w, msg)

	wg.Done()
	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Server Shutdown Failed:%+v", err)
		}
	}()
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
