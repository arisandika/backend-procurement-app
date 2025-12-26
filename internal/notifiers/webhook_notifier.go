package notifiers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"procurement-app/config"
)

type WebhookNotifier struct{}

func NewWebhookNotifier() *WebhookNotifier {
	return &WebhookNotifier{}
}

func (w *WebhookNotifier) Send(payload interface{}) error {
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest(
		"POST",
		config.WebhookURL(),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: config.WebhookTimeout(),
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
