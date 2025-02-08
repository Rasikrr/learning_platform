package jdoodle

import (
	"context"
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
	c := NewClient(url, clientID, clientSecret)
	code := `num = int(input())
if num % 2 == 0:
    print("Четное")
else:
    print("Нечетное")`

	res, err := c.ExecuteCode(context.Background(), code)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err)
	require.Equal(t, 200, res.StatusCode)
	require.Equal(t, "Четное", res.Output)
}
