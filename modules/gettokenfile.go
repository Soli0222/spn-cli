package modules

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetTokenFilePath(tokenFileName string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory: %w", err)
	}
	return filepath.Join(home, ".spn-cli", tokenFileName), nil
}
