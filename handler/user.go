package handler

import (
	"fmt"

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
		fmt.Println("Usage: register [username]")
		return
	}

	if err := validateName(args[0]); err != nil {
		fmt.Println(err)
		return
	}

	if err := hm.userService.Register(args[0]); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Add %s successfully.\n", args[0])
	}
}
