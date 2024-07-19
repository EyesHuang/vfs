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

	fileRepo := repo.NewMemoFileRepo()
	fileService := service.NewFileManageService(fileRepo, userRepo, folderRepo)

	handler := handler.NewHandlerManager(userService, folderService, fileService)

	repl.StartREPL(handler)
}
