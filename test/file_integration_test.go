package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFile(t *testing.T) {
	handler := SetupTestEnvironment()

	// Register a user
	out, _ := captureOutput(func() {
		args := []string{"testuser"}
		handler.HandleRegister(args)
	})
	assert.Contains(t, out, "Add testuser successfully.", "Expected success message for user registration")

	// Create a folder
	out, _ = captureOutput(func() {
		args := []string{"testuser", "testfolder"}
		handler.HandleCreateFolder(args)
	})
	assert.Contains(t, out, "Create testfolder successfully.", "Expected success message for folder creation")

	// Create a file
	out, _ = captureOutput(func() {
		args := []string{"testuser", "testfolder", "testfile", "This is a test file"}
		handler.HandleCreateFile(args)
	})
	assert.Contains(t, out, "Create testfile in testuser/testfolder successfully.", "Expected success message for file creation")
}

func TestListFiles(t *testing.T) {
	handler := SetupTestEnvironment()

	// Register a user
	captureOutput(func() {
		args := []string{"testuser"}
		handler.HandleRegister(args)
	})

	// Create a folder
	captureOutput(func() {
		args := []string{"testuser", "testfolder"}
		handler.HandleCreateFolder(args)
	})

	// Create a file
	captureOutput(func() {
		args := []string{"testuser", "testfolder", "testfile", "This is a test file"}
		handler.HandleCreateFile(args)
	})

	// List files
	out, _ := captureOutput(func() {
		args := []string{"testuser", "testfolder"}
		handler.HandleListFiles(args)
	})
	assert.Contains(t, out, "testfile", "Expected file to be listed")
}

func TestDeleteFile(t *testing.T) {
	handler := SetupTestEnvironment()

	// Register a user
	captureOutput(func() {
		args := []string{"testuser"}
		handler.HandleRegister(args)
	})

	// Create a folder
	captureOutput(func() {
		args := []string{"testuser", "testfolder"}
		handler.HandleCreateFolder(args)
	})

	// Create a file
	captureOutput(func() {
		args := []string{"testuser", "testfolder", "testfile", "This is a test file"}
		handler.HandleCreateFile(args)
	})

	// Delete the file
	out, _ := captureOutput(func() {
		args := []string{"testuser", "testfolder", "testfile"}
		handler.HandleDeleteFile(args)
	})
	assert.Contains(t, out, "Delete testfile successfully.", "Expected success message for file deletion")

	// List files to verify deletion
	out, _ = captureOutput(func() {
		args := []string{"testuser", "testfolder"}
		handler.HandleListFiles(args)
	})
	assert.Contains(t, out, "Warning: The folder is empty.", "Expected warning message for no files")
}
