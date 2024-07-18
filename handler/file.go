package handler

import "fmt"

func (hm *HandlerManager) HandleCreateFile(args []string) {
	if len(args) < 3 || len(args) > 4 {
		fmt.Println("Usage: create-file [username] [foldername] [filename] [description]?")
		return
	}
}

func (hm *HandlerManager) HandleDeleteFile(args []string) {
	if len(args) != 3 {
		fmt.Println("Usage: delete-file [username] [foldername] [filename]")
		return
	}
}

func (hm *HandlerManager) HandleListFiles(args []string) {
	if len(args) != 4 {
		fmt.Println("Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]")
		return
	}
}
