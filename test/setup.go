package test

import (
	"bytes"
	"io"
	"os"

	"vfs/handler"
	"vfs/repo"
	"vfs/service"
)

func SetupTestEnvironment() *handler.HandlerManager {
	// Create repositories
	userRepo := repo.NewMemoUserRepo()
	folderRepo := repo.NewMemoFolderRepo()
	fileRepo := repo.NewMemoFileRepo()

	// Create services
	userService := service.NewUserManageService(userRepo)
	folderService := service.NewFolderManageService(folderRepo, userRepo)
	fileService := service.NewFileManageService(fileRepo, userRepo, folderRepo)

	// Create handler manager
	handlerManager := handler.NewHandlerManager(userService, folderService, fileService)

	return handlerManager
}

// captureOutput captures standard output and standard error for a given function.
func captureOutput(f func()) (string, string) {
	oldStdout := os.Stdout
	oldStderr := os.Stderr

	rOut, wOut, _ := os.Pipe()
	rErr, wErr, _ := os.Pipe()

	os.Stdout = wOut
	os.Stderr = wErr

	outC := make(chan string)
	errC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, rOut)
		outC <- buf.String()
	}()

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, rErr)
		errC <- buf.String()
	}()

	// Execute the function
	f()

	// Close the writers
	wOut.Close()
	wErr.Close()

	// Restore the original stdout and stderr
	os.Stdout = oldStdout
	os.Stderr = oldStderr

	// Read the captured output
	out := <-outC
	err := <-errC

	return out, err
}
