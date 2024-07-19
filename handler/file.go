package handler

import (
	"fmt"

	"vfs"
)

func (hm *HandlerManager) HandleCreateFile(args []string) {
	if len(args) < 3 || len(args) > 4 {
		fmt.Println("Usage: create-file [username] [foldername] [filename] [description]?")
		return
	}

	if !isValidFolderFileName(args[1]) {
		fmt.Printf("The %s contains invalid chars.\n", args[1])
		return
	}

	file := &vfs.File{
		Name:       args[2],
		UserName:   args[0],
		FolderName: args[1],
	}
	if len(args) == 4 {
		file.Description = args[3]
	}

	if err := hm.fileService.AddFile(file); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Create %s in %s/%s successfully.\n", args[2], args[0], args[1])
	}
}

func (hm *HandlerManager) HandleDeleteFile(args []string) {
	if len(args) != 3 {
		fmt.Println("Usage: delete-file [username] [foldername] [filename]")
		return
	}

	if !isValidFolderFileName(args[1]) || !isValidFolderFileName(args[2]) {
		fmt.Printf("The %s contains invalid chars.\n", args[1])
		return
	}

	if !isValidFolderFileName(args[2]) {
		fmt.Printf("The %s contains invalid chars.\n", args[2])
		return
	}

	key := vfs.FileKeySet{
		UserName:   args[0],
		FolderName: args[1],
		FileName:   args[2],
	}

	if err := hm.fileService.DeleteFile(key); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Delete %s successfully.\n", args[2])
	}
}

func (hm *HandlerManager) HandleListFiles(args []string) {
	if len(args) < 2 || len(args) > 4 {
		fmt.Println("Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]")
		return
	}

	req := &vfs.GetFilesRequest{
		UserName:   args[0],
		FolderName: args[1],
	}

	if len(args) == 4 {
		sortBy, err := parseSortType(args[2])
		if err != nil {
			fmt.Println(err)
			return
		}

		orderBy, err := parseOrderType(args[3])
		if err != nil {
			fmt.Println(err)
			return
		}

		req.SortBy = sortBy
		req.OrderBy = orderBy
	}

	files, err := hm.fileService.GetFiles(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(files) == 0 {
		fmt.Println("Warning: The folder is empty.")
		return
	}

	for _, file := range files {
		fmt.Printf("%s\t%s\t%s\t%s\t%s\n", file.Name, file.Description, file.CreatedAt.Format("2006-01-02 15:04:05"), file.FolderName, file.UserName)
	}
}
