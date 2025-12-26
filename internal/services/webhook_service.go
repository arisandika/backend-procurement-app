package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func TriggerWebhook(purchasingID uint) {
	payload := map[string]interface{}{
		"event": "purchasing_created",
		"id":    purchasingID,
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(
		"POST",
		os.Getenv("WEBHOOK_URL"),
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	client.Do(req)
}
