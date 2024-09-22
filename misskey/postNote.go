package misskey

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CreatedNoteResponse struct {
	CreatedNote struct {
		Id string `json:"id"`
	} `json:"createdNote"`
}

func PostNote(host string, token string, currentPlayingInfo string) (string, error) {
	url := fmt.Sprintf("https://%s/api/notes/create", host)

	payload := map[string]string{
		"i":    token,
		"text": currentPlayingInfo,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error marshalling JSON: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	var createdNoteResponse CreatedNoteResponse
	if err := json.Unmarshal(body, &createdNoteResponse); err != nil {
		return "", fmt.Errorf("error unmarshalling response: %w", err)
	}

	return createdNoteResponse.CreatedNote.Id, nil
}
