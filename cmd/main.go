package main

import (
	"vfs/handler"
	"vfs/repl"
	"vfs/repo"
	"vfs/service"
)

func main() {
	userRepo := repo.NewMemoUserRepo()
	userService := service.NewUserManageService(userRepo)

	folderRepo := repo.NewMemoFolderRepo()
	folderService := service.NewFolderManageService(folderRepo, userRepo)

	handler := handler.NewHandlerManager(userService, folderService)

	repl.StartREPL(handler)
}
