package handler

import (
	"errors"
	"regexp"
	"strings"
)

const (
	MaxNameLength  = 255
	ValidNameRegex = `^[a-zA-Z0-9_\- ]+$`
)

func validateName(name string) error {
	if len(name) == 0 || len(name) > MaxNameLength {
		return errors.New("name length must be between 1 and 255 characters")
	}

	validName := regexp.MustCompile(ValidNameRegex)
	if !validName.MatchString(name) {
		return errors.New("name contains invalid characters")
	}
	return nil
}

// isValidFolderFileName checks if the given folder/file name is valid
func isValidFolderFileName(name string) bool {
	// Check if the first character is a whitespace
	if strings.HasPrefix(name, " ") {
		return false
	}

	// Define the invalid characters for a folder name using regex
	// Invalid characters are \ / : * ? " < > |
	invalidFolderName := regexp.MustCompile(`[\\/:*?"<>|]`)
	return !invalidFolderName.MatchString(name)
}
