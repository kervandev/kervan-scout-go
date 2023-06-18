package kervanscout

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type Client struct {
	config Config
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

func New(cfg Config) *Client {
	client := &Client{
		config: cfg,
	}

	if cfg.Host == "" {
		client.config.Host = "https://scout-api.tapsilat.dev"
	}

	return client
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
	req, err := http.NewRequest("POST", c.config.Host+"/api/v1/issues", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	host, err := os.Hostname()
	if err != nil {
		return err
	}
	req.Header.Set("Origin", host)

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

func (c *Client) GetHost() string {
	return c.config.Host
}

func (c *Client) GetProjectToken() string {
	return c.config.ProjectToken
}

func (c *Client) SendIssue(title string, message string, payload ...interface{}) {
	c.request(title, message, payload...)
}

func (c *Client) CatchPanicError(title ...string) {
	var t string

	if len(title) > 0 {
		t = title[0]
	} else {
		t = "panic error"
	}

	if r := recover(); r != nil {
		c.request(t, r.(string))
	}
}
