package telegramnotifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Notifier represents the client for sending Telegram notifications
type Notifier struct {
	botToken string
	client   *http.Client
}

// Message represents a message to be sent
type Message struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

// New creates a new instance of the notifier
func New(botToken string) *Notifier {
	return &Notifier{
		botToken: botToken,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Send sends a message to a specific chat
func (n *Notifier) Send(chatID int64, message string) error {
	msg := Message{
		ChatID: chatID,
		Text:   message,
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("error serializing message: %w", err)
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", n.botToken)
	resp, err := n.client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error sending message: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error in Telegram response: status %d", resp.StatusCode)
	}

	return nil
}