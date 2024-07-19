package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	handler := SetupTestEnvironment()

	// Capture output for user registration
	out, _ := captureOutput(func() {
		args := []string{"testuser"}
		handler.HandleRegister(args)
	})
	assert.Contains(t, out, "Add testuser successfully.", "Expected success message for user registration")
}

func TestDuplicateUserRegistration(t *testing.T) {
	handler := SetupTestEnvironment()

	// Register a user
	captureOutput(func() {
		args := []string{"testuser"}
		handler.HandleRegister(args)
	})

	// Try registering the same user again and capture output
	out, _ := captureOutput(func() {
		args := []string{"testuser"}
		handler.HandleRegister(args)
	})
	assert.Contains(t, out, "The testuser has already existed.", "Expected error message for duplicate user registration")
}
