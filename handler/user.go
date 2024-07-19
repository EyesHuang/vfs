package handler

import (
	"fmt"
	"os"

	"vfs"
)

type HandlerManager struct {
	userService   vfs.UserService
	folderService vfs.FolderService
	fileService   vfs.FileService
}

func NewHandlerManager(userService vfs.UserService, folderService vfs.FolderService,
	fileService vfs.FileService,
) *HandlerManager {
	return &HandlerManager{userService: userService, folderService: folderService, fileService: fileService}
}

func (hm *HandlerManager) HandleRegister(args []string) {
	if len(args) != 1 {
		_, _ = fmt.Fprintln(os.Stderr, "Usage: register [username]")
		return
	}
	if err := hm.userService.Register(args[0]); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	} else {
		_, _ = fmt.Fprintf(os.Stdout, "Add %s successfully.\n", args[0])
	}
}
