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

func (c *Client) request(title, message string, payload ...interface{}) error {
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

	client := http.Client{}
	req, err := http.NewRequest("POST", c.config.Host, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"project_token": {c.config.ProjectToken},
	}

	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SendIssue(title string, message string, payload ...interface{}) error {
	if err := c.request(title, message, payload...); err != nil {
		return err
	}

	defer func() {
		if err := recover(); err != nil {
			c.request(title, err.(string))
		}
	}()

	return nil
}
