package handler

import (
	"fmt"

	"vfs"
)

type HandlerManager struct {
	userService   vfs.UserService
	folderService vfs.FolderService
}

func NewHandlerManager(us vfs.UserService, fs vfs.FolderService) *HandlerManager {
	return &HandlerManager{userService: us, folderService: fs}
}

func (hm *HandlerManager) HandleRegister(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: register [username]")
		return
	}
	if err := hm.userService.Register(args[0]); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Add %s successfully.\n", args[0])
	}
}
