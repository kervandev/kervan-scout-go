package kervanscout

import (
	"bytes"
	"encoding/json"
	"io"
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
	Title       string      `json:"title"`
	Message     string      `json:"message"`
	Payload     interface{} `json:"payload,omitempty"`
	Type        IssueType   `json:"type"`
	MaskPayload bool        `json:"mask_payload"`
}

type Options struct {
	Payload     interface{} `json:"payload,omitempty"`
	Type        IssueType   `json:"type"`
	MaskPayload bool        `json:"mask_payload"`
}

type IssueResponseData struct {
	ID      string `json:"id"`
	Message string `json:"message"`
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

func (c *Client) request(iss *issue) (*IssueResponseData, error) {
	if iss.Type == "" {
		iss.Type = IssueTypeExecution
	}

	issueType, err := ParseIssueType(iss.Type.String())
	if err != nil {
		return nil, err
	}
	iss.Type = issueType

	data, err := json.Marshal(&iss)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	req, err := http.NewRequest("POST", c.config.Host+"/api/v1/issues", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	host, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Origin", host)

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"project_token": {c.config.ProjectToken},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respData *IssueResponseData
	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		return nil, err
	}

	return respData, nil
}

func (c *Client) GetHost() string {
	return c.config.Host
}

func (c *Client) GetProjectToken() string {
	return c.config.ProjectToken
}

func (c *Client) SendIssue(title, message string, additional ...Options) (*IssueResponseData, error) {
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

	return c.request(iss)
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
