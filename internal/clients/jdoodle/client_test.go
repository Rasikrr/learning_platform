package jdoodle

import (
	"context"
	"github.com/Rasikrr/learning_platform/configs"
	"github.com/Rasikrr/learning_platform/internal/domain/enum"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// nolint: gosec
const (
	url          = "https://api.jdoodle.com/v1/execute"
	clientID     = "a229ba42922780b16291ccc4d1d3efd7"
	clientSecret = "d09ec5c306cb82a55f5ca29dc207644b35a3ce26feb058c6265c08b6a484799d"
)

func TestClient(t *testing.T) {
	cfg := &configs.Config{
		JDoodle: configs.JDoodleConfig{
			URL:          url,
			ClientID:     clientID,
			ClientSecret: clientSecret,
		},
	}
	ctx := context.Background()

	client := NewClient(cfg)

	testCases := []struct {
		Name           string
		Code           string
		Lang           enum.ProgrammingLanguage
		ExpectedOutput string
	}{
		{
			Name:           "Test go code",
			Code:           goCode,
			Lang:           enum.ProgrammingLanguageGo,
			ExpectedOutput: "Hello, world!",
		},
		{
			Name:           "Test python code",
			Code:           pythonCode,
			Lang:           enum.ProgrammingLanguagePython3,
			ExpectedOutput: "5",
		},
		{
			Name:           "Test java code",
			Code:           javaCode,
			Lang:           enum.ProgrammingLanguageJava,
			ExpectedOutput: "Hello, world!",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()
			res, err := client.ExecuteCode(ctx, testCase.Code, testCase.Lang)
			require.NoError(t, err)
			assert.Equal(t, testCase.ExpectedOutput, res)
		})
	}
}

const (
	pythonCode = `a = 2
b = 3
c = a + b
print(c)`
	javaCode = `public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello, world!");
    }
}`
	goCode = `package main
import (
    "fmt"
)

func main() {
	fmt.Println("Hello, world!")
}`
)
