// cmd/main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"vfs"

	"vfs/repo"
	"vfs/service"
)

func main() {
	userRepo := repo.NewMemoUserRepo()
	userService := service.NewUserManageService(userRepo)

	folderRepo := repo.NewMemoFolderRepo()
	folderService := service.NewFolderManageService(folderRepo, userRepo)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Virtual File System REPL")
	for {
		fmt.Print("# ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Split(input, " ")

		if len(args) == 0 {
			continue
		}

		command := args[0]
		args = args[1:]

		switch command {
		case "register":
			if len(args) != 1 {
				fmt.Println("Usage: register [username]")
				continue
			}
			if err := userService.Register(args[0]); err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Add %s successfully.\n", args[0])
			}
		case "create-folder":
			if len(args) < 2 || len(args) > 3 {
				fmt.Println("Usage: create-folder [username] [foldername] [description]?")
				continue
			}

			if !isValidFolderFileName(args[1]) {
				fmt.Printf("The %s contain invalid chars.\n", args[1])
				continue
			}

			var folder *vfs.Folder

			if len(args) == 2 {
				folder = &vfs.Folder{
					Name:     args[1],
					UserName: args[0],
				}
			} else {
				folder = &vfs.Folder{
					Name:        args[1],
					Description: args[2],
					UserName:    args[0],
				}
			}

			if err := folderService.AddFolder(folder); err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Create %s successfully.\n", args[1])
			}
		case "delete-folder":
			if len(args) != 2 {
				fmt.Println("Usage: delete-folder [username] [foldername]")
				continue
			}
		case "list-folders":
			if len(args) < 1 || len(args) > 3 {
				fmt.Println("Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]")
				continue
			}
		case "rename-folder":
			if len(args) != 3 {
				fmt.Println("Usage: rename-folder [username] [foldername] [new-folder-name]")
				continue
			}
		case "create-file":
			if len(args) < 3 || len(args) > 4 {
				fmt.Println("Usage: create-file [username] [foldername] [filename] [description]?")
				continue
			}
		case "delete-file":
			if len(args) != 3 {
				fmt.Println("Usage: delete-file [username] [foldername] [filename]")
				continue
			}
		case "list-files":
			if len(args) != 4 {
				fmt.Println("Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]")
				continue
			}
		default:
			fmt.Println("Unrecognized command")
		}
	}
}

// isValidFolderFileName checks if the given folder name is valid
func isValidFolderFileName(name string) bool {
	// Define the valid characters for a folder name using regex
	// Here we allow letters (a-z, A-Z), digits (0-9), underscores (_), hyphens (-), and spaces
	validFolderName := regexp.MustCompile(`^[a-zA-Z0-9_\- ]+$`)
	return validFolderName.MatchString(name)
}
