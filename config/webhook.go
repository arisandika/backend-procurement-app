package config

import (
	"os"
	"time"
)

func WebhookURL() string {
	return os.Getenv("WEBHOOK_URL")
}

func WebhookTimeout() time.Duration {
	return time.Second * 5
}
