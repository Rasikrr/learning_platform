package jdoodle

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	url          = "https://api.jdoodle.com/v1/execute"
	clientID     = "444e54affc8c6da6723dbfdbce63cf5d"
	clientSecret = "11208b5cf61686bb22cdbe1ed0c52e0c0b0b1e74edd110abd9f6655e30a0b284"
)

func TestClient(t *testing.T) {
	c := NewClient(url, clientID, clientSecret)
	res, err := c.ExecuteCode(context.Background(), "package main\n import \"fmt\"\nfunc main() {\n fmt.Println(3)\n}")
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err)
	require.Equal(t, 200, res.StatusCode)
	require.Equal(t, "3", res.Output)
}
