package repl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    `create-file "user name" foldername "file name"`,
			expected: []string{"create-file", "user name", "foldername", "file name"},
		},
		{
			input:    `delete-file username foldername filename`,
			expected: []string{"delete-file", "username", "foldername", "filename"},
		},
		{
			input:    `list-files username foldername --sort-name asc`,
			expected: []string{"list-files", "username", "foldername", "--sort-name", "asc"},
		},
		{
			input:    `update-file "user name" "folder name" "file name"`,
			expected: []string{"update-file", "user name", "folder name", "file name"},
		},
		{
			input:    `search-files username "folder name" --keyword "test keyword"`,
			expected: []string{"search-files", "username", "folder name", "--keyword", "test keyword"},
		},
	}

	for _, tc := range testCases {
		tokens := parseInput(tc.input)
		assert.Equal(t, tc.expected, tokens, "Expected %v but got %v for input %s", tc.expected, tokens, tc.input)
	}
}
