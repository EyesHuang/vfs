package handler

import (
	"fmt"
	"regexp"

	"vfs"
)

func (hm *HandlerManager) HandleCreateFolder(args []string) {
	if len(args) < 2 || len(args) > 3 {
		fmt.Println("Usage: create-folder [username] [foldername] [description]?")
		return
	}

	if !isValidFolderFileName(args[1]) {
		fmt.Printf("The %s contain invalid chars.\n", args[1])
		return
	}

	folder := &vfs.Folder{
		Name:     args[1],
		UserName: args[0],
	}
	if len(args) == 3 {
		folder.Description = args[2]
	}

	if err := hm.folderService.AddFolder(folder); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Create %s successfully.\n", args[1])
	}
}

func (hm *HandlerManager) HandleDeleteFolder(args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: delete-folder [username] [foldername]")
		return
	}
}

func (hm *HandlerManager) HandleListFolders(args []string) {
	if len(args) < 1 || len(args) > 3 {
		fmt.Println("Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]")
		return
	}
}

func (hm *HandlerManager) HandleRenameFolder(args []string) {
	if len(args) != 3 {
		fmt.Println("Usage: rename-folder [username] [foldername] [new-folder-name]")
		return
	}
}

// isValidFolderFileName checks if the given folder name is valid
func isValidFolderFileName(name string) bool {
	// Define the valid characters for a folder name using regex
	// Here we allow letters (a-z, A-Z), digits (0-9), underscores (_), hyphens (-), and spaces
	validFolderName := regexp.MustCompile(`^[a-zA-Z0-9_\- ]+$`)
	return validFolderName.MatchString(name)
}
