package jdoodle

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type Client interface {
	ExecuteCode(ctx context.Context, code string) (*ExecuteResponse, error)
}

type client struct {
	url          string
	clientID     string
	clientSecret string
}

func NewClient(url, clientID, clientSecret string) Client {
	return &client{
		url:          url,
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

func (c *client) ExecuteCode(ctx context.Context, code string) (*ExecuteResponse, error) {
	req := ExecuteRequest{
		ClientID:     c.clientID,
		ClientSecret: c.clientSecret,
		Script:       code,
		Language:     "go",
		VersionIndex: "3",
		CompileOnly:  false,
	}
	bb, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(c.url, "application/json", bytes.NewBuffer(bb))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var out ExecuteResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	return &out, nil
}
