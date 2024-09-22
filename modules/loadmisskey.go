package modules

import (
	"encoding/json"
	"fmt"
	"os"
)

type MisskeyCredentials struct {
	Hostname string `json:"hostname"`
	Token    string `json:"token"`
}

// LoadMisskey loads the Misskey credentials from a JSON file and returns the hostname and token.
func LoadMisskey(filePath string) (*MisskeyCredentials, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var credentials MisskeyCredentials
	if err := json.NewDecoder(file).Decode(&credentials); err != nil {
		return nil, fmt.Errorf("failed to decode credentials: %w", err)
	}

	if credentials.Hostname == "" || credentials.Token == "" {
		return nil, fmt.Errorf("incomplete credentials: hostname or token is missing")
	}

	return &credentials, nil
}
