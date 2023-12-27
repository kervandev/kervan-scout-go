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

type issue struct {
	Title       string       `json:"title"`
	Message     string       `json:"message"`
	Payload     *interface{} `json:"payload,omitempty"`
	Type        IssueType    `json:"type"`
	MaskPayload *bool        `json:"mask_payload"`
}

type Options struct {
	Payload     *interface{} `json:"payload,omitempty"`
	Type        IssueType    `json:"type"`
	MaskPayload *bool        `json:"mask_payload"`
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

func (c *Client) request(iss *issue) error {
	data, err := json.Marshal(&iss)
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

func (c *Client) SendIssue(title, message string, additional ...Options) {
	var opts Options
	if len(additional) > 0 {
		opts = additional[0]
	}

	iss := &issue{
		Title:       title,
		Message:     message,
		Payload:     opts.Payload,
		Type:        opts.Type,
		MaskPayload: opts.MaskPayload,
	}

	c.request(iss)
}

func (c *Client) CatchPanicError(title ...string) {
	var t string

	if len(title) > 0 {
		t = title[0]
	} else {
		t = "panic error"
	}

	if r := recover(); r != nil {
		c.request(
			&issue{
				Title:   t,
				Message: r.(string),
				Type:    IssueTypeExecution,
			})
	}
}
