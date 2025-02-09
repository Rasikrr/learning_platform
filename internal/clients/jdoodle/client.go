package jdoodle

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/Rasikrr/learning_platform/configs"
	"github.com/Rasikrr/learning_platform/internal/domain/enum"
	"log"
	"net/http"
)

var (
	ErrInvalidStatusCode = errors.New("invalid status code")
	ErrCompilationError  = errors.New("compilation error")
)

type Client interface {
	ExecuteCode(_ context.Context, code string, lang enum.ProgrammingLanguage) (string, error)
	ExecuteTestCase(_ context.Context, code, stdin string, lang enum.ProgrammingLanguage) (string, error)
}

type client struct {
	url          string
	clientID     string
	clientSecret string
}

func NewClient(cfg *configs.Config) Client {
	return &client{
		url:          cfg.JDoodle.URL,
		clientID:     cfg.JDoodle.ClientID,
		clientSecret: cfg.JDoodle.ClientSecret,
	}
}

func (c *client) ExecuteCode(_ context.Context, code string, lang enum.ProgrammingLanguage) (string, error) {
	req := ExecuteRequest{
		ClientID:     c.clientID,
		ClientSecret: c.clientSecret,
		Script:       code,
		Language:     lang.String(),
		VersionIndex: "3",
		CompileOnly:  false,
	}

	bb, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(c.url, "application/json", bytes.NewBuffer(bb))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var out ExecuteResponse
	if err = json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	log.Printf("Request: %+v\n", req)
	log.Printf("Response: %+v\n", out)
	if out.StatusCode != 200 {
		return "", ErrInvalidStatusCode
	}
	return out.Output, nil
}

func (c *client) ExecuteTestCase(_ context.Context, code, stdin string, lang enum.ProgrammingLanguage) (string, error) {
	req := ExecuteRequest{
		ClientID:     c.clientID,
		ClientSecret: c.clientSecret,
		Script:       code,
		Stdin:        stdin,
		Language:     lang.String(),
		VersionIndex: "3",
		CompileOnly:  false,
	}

	bb, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(c.url, "application/json", bytes.NewBuffer(bb))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var out ExecuteResponse
	if err = json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	log.Printf("Request: %+v\n", req)
	log.Printf("Response: %+v\n", out)
	if out.StatusCode != 200 {
		return "", ErrInvalidStatusCode
	}
	return out.Output, nil
}
