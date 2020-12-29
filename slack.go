package slackbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// SlackTextMessage is a container for the text message
type SlackTextMessage struct {
	Text string `json:"text,omitempty"`
}

// SendSlackMessage is for sending a response text message to a slack response url
func SendSlackMessage(message string, responseURL string) error {

	delayedResponseBytes, err := json.Marshal(SlackTextMessage{
		Text: message,
	})
	if err != nil {
		return fmt.Errorf("json.Marshal error: %v", err)
	}
	byteReader := bytes.NewReader(delayedResponseBytes)

	// Send request
	resp, err := http.Post(responseURL, "application/json", byteReader)
	if err != nil {
		return fmt.Errorf("http.Post error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Post statusCode: %v, error: %v", resp.StatusCode, resp.Body)
	}

	return nil
}
