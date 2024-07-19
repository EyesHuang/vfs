package handler

import (
	"fmt"
	"os"

	"vfs"
)

func (hm *HandlerManager) HandleCreateFile(args []string) {
	if len(args) < 3 || len(args) > 4 {
		_, _ = fmt.Fprintln(os.Stderr, "Usage: create-file [username] [foldername] [filename] [description]?")
		return
	}

	if !isValidFolderFileName(args[1]) {
		_, _ = fmt.Fprintf(os.Stderr, "The %s contains invalid chars.\n", args[1])
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
		_, _ = fmt.Fprintln(os.Stderr, err)
	} else {
		_, _ = fmt.Fprintf(os.Stdout, "Create %s in %s/%s successfully.\n", args[2], args[0], args[1])
	}
}

func (hm *HandlerManager) HandleDeleteFile(args []string) {
	if len(args) != 3 {
		_, _ = fmt.Fprintln(os.Stderr, "Usage: delete-file [username] [foldername] [filename]")
		return
	}

	if !isValidFolderFileName(args[1]) || !isValidFolderFileName(args[2]) {
		_, _ = fmt.Fprintf(os.Stderr, "The %s contains invalid chars.\n", args[1])
		return
	}

	if !isValidFolderFileName(args[2]) {
		_, _ = fmt.Fprintf(os.Stderr, "The %s contains invalid chars.\n", args[2])
		return
	}

	key := vfs.FileKeySet{
		UserName:   args[0],
		FolderName: args[1],
		FileName:   args[2],
	}

	if err := hm.fileService.DeleteFile(key); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	} else {
		_, _ = fmt.Fprintf(os.Stdout, "Delete %s successfully.\n", args[2])
	}
}

func (hm *HandlerManager) HandleListFiles(args []string) {
	if len(args) < 2 || len(args) > 4 {
		_, _ = fmt.Fprintln(os.Stderr, "Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]")
		return
	}

	req := &vfs.GetFilesRequest{
		UserName:   args[0],
		FolderName: args[1],
	}

	if len(args) == 4 {
		sortBy, err := parseSortType(args[2])
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			return
		}

		orderBy, err := parseOrderType(args[3])
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			return
		}

		req.SortBy = sortBy
		req.OrderBy = orderBy
	}

	files, err := hm.fileService.GetFiles(req)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	if len(files) == 0 {
		_, _ = fmt.Fprintln(os.Stdout, "Warning: The folder is empty.")
		return
	}

	for _, file := range files {
		_, _ = fmt.Fprintf(os.Stdout, "%s\t%s\t%s\t%s\t%s\n", file.Name, file.Description, file.CreatedAt.Format("2006-01-02 15:04:05"), file.FolderName, file.UserName)
	}
}
