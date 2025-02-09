package jdoodle

import (
	"context"
	"github.com/Rasikrr/learning_platform/configs"
	"github.com/Rasikrr/learning_platform/internal/domain/enum"
	"github.com/stretchr/testify/require"
	"testing"
)

// nolint
const (
	url          = "https://api.jdoodle.com/v1/execute"
	clientID     = "444e54affc8c6da6723dbfdbce63cf5d"
	clientSecret = "11208b5cf61686bb22cdbe1ed0c52e0c0b0b1e74edd110abd9f6655e30a0b284"
)

func TestClient(t *testing.T) {
	t.Skip()
	ctx := context.Background()
	cfg := &configs.Config{
		JDoodle: configs.JDoodleConfig{
			URL:          url,
			ClientID:     clientID,
			ClientSecret: clientSecret,
		},
	}
	c := NewClient(cfg)
	code := `package main

import (
    "fmt"
    "time"
)

func first() {
    fmt.Println("First")
}

func second() {
    fmt.Println("Second")
}

func main() {
    go first()
    go second()
    time.Sleep(time.Second)
}`

	res, err := c.ExecuteCode(ctx, code, enum.ProgrammingLanguageGo)
	require.NoError(t, err)
	require.Equal(t, "First\nSecond", res)
}
