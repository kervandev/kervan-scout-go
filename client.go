package kervanscout

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	config *Config
}

type Config struct {
	Host         string `json:"host"`
	ProjectToken string `json:"project_token"`
}

type body struct {
	Title   string       `json:"title"`
	Message string       `json:"message"`
	Payload *interface{} `json:"payload,omitempty"`
}

func New(cfg *Config) *Client {
	return &Client{
		config: cfg,
	}
}

func (c *Client) SendIssue(title string, message string, payload ...interface{}) error {
	b := body{
		Title:   title,
		Message: message,
	}
	if len(payload) > 0 {
		b.Payload = &payload[0]
	}

	data, err := json.Marshal(&b)
	if err != nil {
		return err
	}

	_, err = http.Post(c.config.Host, "application/json", bytes.NewBuffer(data))

	return err
}
