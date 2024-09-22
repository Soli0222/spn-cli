package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserProfile struct {
	DisplayName string `json:"display_name"`
}

func FetchUserProfile(client *http.Client) (string, error) {
	apiURL := "https://api.spotify.com/v1/me"
	resp, err := client.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("error making API request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	var profile UserProfile
	if err := json.Unmarshal(body, &profile); err != nil {
		return "", fmt.Errorf("error unmarshalling response: %w", err)
	}

	return profile.DisplayName, nil
}
