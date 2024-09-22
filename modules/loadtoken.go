package modules

import (
	"encoding/json"
	"os"

	"golang.org/x/oauth2"
)

func LoadToken(filePath string) (*oauth2.Token, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var token oauth2.Token
	err = json.NewDecoder(file).Decode(&token)
	return &token, err
}
