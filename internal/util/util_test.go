package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlicesEqual(t *testing.T) {
	testCases := []struct {
		Name     string
		Slice1   []string
		Slice2   []string
		Expected bool
	}{
		{
			Name:     "Test equal slices",
			Slice1:   []string{"a", "b", "c"},
			Slice2:   []string{"a", "b", "c"},
			Expected: true,
		},
		{
			Name:     "Test not equal slices",
			Slice1:   []string{"a", "b", "c"},
			Slice2:   []string{"a", "b", "d"},
			Expected: false,
		},
		{
			Name:     "Test not equal slices",
			Slice1:   []string{"a", "b", "c"},
			Slice2:   []string{"a", "b"},
			Expected: false,
		},
		{
			Name:     "Test empty slices",
			Slice1:   []string{},
			Slice2:   []string{},
			Expected: true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()
			res := SlicesEqual(testCase.Slice1, testCase.Slice2)
			assert.Equal(t, testCase.Expected, res)
		})
	}
}
