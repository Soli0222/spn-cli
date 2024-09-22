package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CurrentlyPlayingResponse struct {
	Item struct {
		Name    string `json:"name"`
		Artists []struct {
			Name string `json:"name"`
		} `json:"artists"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
	} `json:"item"`
}

func FetchCurrentryPlaying(client *http.Client) (string, string, string, error) {
	apiURL := "https://api.spotify.com/v1/me/player?market=JP"
	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Set("Accept-Language", "ja")
	resp, err := client.Do(req)

	if err != nil {
		return "", "", "", fmt.Errorf("error making API request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", fmt.Errorf("error reading response body: %w", err)
	}

	var result CurrentlyPlayingResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", "", "", fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	artistNames := ""
	for i, artist := range result.Item.Artists {
		if i > 0 {
			artistNames += ", "
		}
		artistNames += artist.Name
	}

	return result.Item.Name, artistNames, result.Item.ExternalUrls.Spotify, nil
}
