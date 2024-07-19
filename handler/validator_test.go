package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateName_ValidNames(t *testing.T) {
	data := make([]rune, MaxNameLength)

	for i := range data {
		data[i] = ' '
	}

	validNames := []string{
		"validName",
		"Valid_Name123",
		"Valid-Name 123",
		"v",
		string(data), // Name with max length
	}

	for _, name := range validNames {
		err := validateName(name)
		assert.Nil(t, err, "Expected no error for valid name: %s", name)
	}
}

func TestValidateName_InvalidNames(t *testing.T) {
	invalidNames := map[string]string{
		"":                                    "name length must be between 1 and 255 characters",
		string(make([]byte, MaxNameLength+1)): "name length must be between 1 and 255 characters", // Name exceeding max length
		"Invalid@Name!":                       "name contains invalid characters",
		"Invalid/Name":                        "name contains invalid characters",
	}

	for name, expectedErr := range invalidNames {
		err := validateName(name)
		assert.NotNil(t, err, "Expected error for invalid name: %s", name)
		assert.Equal(t, expectedErr, err.Error(), "Unexpected error message for invalid name: %s", name)
	}
}

func TestIsValidFolderFileName_ValidNames(t *testing.T) {
	validNames := []string{
		"validName",
		"Valid_Name123",
		"Valid-Name 123",
		"v",
		"valid.name",
	}

	for _, name := range validNames {
		valid := isValidFolderFileName(name)
		assert.True(t, valid, "Expected name to be valid: %s", name)
	}
}

func TestIsValidFolderFileName_InvalidNames(t *testing.T) {
	invalidNames := []string{
		" InvalidName", // Starts with a whitespace
		"Invalid/Name",
		"Invalid\\Name",
		"Invalid:Name",
		"Invalid*Name",
		"Invalid?Name",
		"Invalid\"Name",
		"Invalid<Name",
		"Invalid>Name",
		"Invalid|Name",
	}

	for _, name := range invalidNames {
		valid := isValidFolderFileName(name)
		assert.False(t, valid, "Expected name to be invalid: %s", name)
	}
}
