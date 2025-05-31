package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Discord interface {
	SendMessage(message string) error
}

type discord struct {
	configs []DiscordConfig
}

type DiscordConfig struct {
	ChannelID string
	Token     string
}

func NewDiscord(configs []DiscordConfig) Discord {
	return &discord{
		configs: configs,
	}
}

func (d *discord) SendMessage(message string) error {
	payload := map[string]string{
		"content": message,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	for _, config := range d.configs {
		webhookURL := fmt.Sprintf("https://discord.com/api/webhooks/%s/%s", config.ChannelID, config.Token)

		_, err = http.Post(webhookURL, "application/json", bytes.NewBuffer(body))
		if err != nil {
			return err
		}
	}

	return nil
}
