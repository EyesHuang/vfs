package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFolder(t *testing.T) {
	handler := SetupTestEnvironment()

	// Register a user
	captureOutput(func() {
		args := []string{"testuser"}
		handler.HandleRegister(args)
	})

	// Create a folder and capture output
	out, _ := captureOutput(func() {
		args := []string{"testuser", "testfolder"}
		handler.HandleCreateFolder(args)
	})
	assert.Contains(t, out, "Create testfolder successfully.", "Expected success message for folder creation")
}

func TestDeleteFolder(t *testing.T) {
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

	// Delete the folder and capture output
	out, _ := captureOutput(func() {
		args := []string{"testuser", "testfolder"}
		handler.HandleDeleteFolder(args)
	})
	assert.Contains(t, out, "Delete testfolder successfully.", "Expected success message for folder deletion")

	// List folders to verify deletion
	out, _ = captureOutput(func() {
		args := []string{"testuser"}
		handler.HandleListFolders(args)
	})
	assert.Contains(t, out, "Warning: The testuser doesn't have any folders.", "Expected warning message for no folders")
}

func TestListFolders(t *testing.T) {
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

	// List folders and capture output
	out, _ := captureOutput(func() {
		args := []string{"testuser"}
		handler.HandleListFolders(args)
	})
	assert.Contains(t, out, "testfolder", "Expected folder to be listed")
}

func TestRenameFolder(t *testing.T) {
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

	// Rename the folder and capture output
	out, _ := captureOutput(func() {
		args := []string{"testuser", "testfolder", "newfolder"}
		handler.HandleRenameFolder(args)
	})
	assert.Contains(t, out, "Rename testfolder to newfolder successfully.", "Expected success message for folder rename")

	// List folders to verify rename
	out, _ = captureOutput(func() {
		args := []string{"testuser"}
		handler.HandleListFolders(args)
	})
	assert.Contains(t, out, "newfolder", "Expected new folder name to be listed")
}
