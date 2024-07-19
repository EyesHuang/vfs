package handler

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"vfs"
)

func (hm *HandlerManager) HandleCreateFolder(args []string) {
	if len(args) < 2 || len(args) > 3 {
		_, _ = fmt.Fprintln(os.Stderr, "Usage: create-folder [username] [foldername] [description]?")
		return
	}

	if !isValidFolderFileName(args[1]) {
		_, _ = fmt.Fprintf(os.Stderr, "The %s contains invalid chars.\n", args[1])
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
		_, _ = fmt.Fprintln(os.Stderr, err)
	} else {
		_, _ = fmt.Fprintf(os.Stdout, "Create %s successfully.\n", args[1])
	}
}

func (hm *HandlerManager) HandleDeleteFolder(args []string) {
	if len(args) != 2 {
		_, _ = fmt.Fprintln(os.Stderr, "Usage: delete-folder [username] [foldername]")
		return
	}

	if !isValidFolderFileName(args[1]) {
		_, _ = fmt.Fprintf(os.Stderr, "The %s contains invalid chars.\n", args[1])
		return
	}

	key := vfs.FolderKeySet{
		UserName:   args[0],
		FolderName: args[1],
	}

	if err := hm.folderService.DeleteFolder(key); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	} else {
		_, _ = fmt.Fprintf(os.Stdout, "Delete %s successfully.\n", args[1])
	}
}

func (hm *HandlerManager) HandleListFolders(args []string) {
	if len(args) < 1 || len(args) > 3 {
		_, _ = fmt.Fprintln(os.Stderr, "Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]")
		return
	}

	req := &vfs.GetFoldersRequest{
		UserName: args[0],
	}

	if len(args) == 3 {
		sortBy, err := parseSortType(args[1])
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			return
		}

		orderBy, err := parseOrderType(args[2])
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			return
		}

		req.SortBy = sortBy
		req.OrderBy = orderBy
	}

	folders, err := hm.folderService.GetFolders(req)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	if len(folders) == 0 {
		_, _ = fmt.Fprintf(os.Stdout, "Warning: The %s doesn't have any folders.\n", req.UserName)
		return
	}

	for _, folder := range folders {
		_, _ = fmt.Fprintf(os.Stdout, "%s\t%s\t%s\t%s\n", folder.Name, folder.Description, folder.CreatedAt.Format("2006-01-02 15:04:05"), folder.UserName)
	}
}

func (hm *HandlerManager) HandleRenameFolder(args []string) {
	if len(args) != 3 {
		_, _ = fmt.Fprintln(os.Stderr, "Usage: rename-folder [username] [foldername] [new-folder-name]")
		return
	}

	if !isValidFolderFileName(args[1]) || !isValidFolderFileName(args[2]) {
		_, _ = fmt.Fprintln(os.Stderr, "Folder names contains invalid chars.")
		return
	}

	req := &vfs.UpdateFolderRequest{
		OldName:  args[1],
		NewName:  args[2],
		UserName: args[0],
	}

	if err := hm.folderService.UpdateFolder(req); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	} else {
		_, _ = fmt.Fprintf(os.Stdout, "Rename %s to %s successfully.\n", args[1], args[2])
	}
}

// isValidFolderFileName checks if the given folder/file name is valid
func isValidFolderFileName(name string) bool {
	// Check if the first character is a whitespace
	if strings.HasPrefix(name, " ") {
		return false
	}

	// Define the invalid characters for a folder name using regex
	// Invalid characters are \ / : * ? " < > |
	invalidFolderName := regexp.MustCompile(`[\\/:*?"<>|]`)
	return !invalidFolderName.MatchString(name)
}

func parseSortType(sortType string) (vfs.SortType, error) {
	var sortBy vfs.SortType

	switch sortType {
	case "--sort-name":
		sortBy = vfs.Name
	case "--sort-created":
		sortBy = vfs.Created
	default:
		return sortBy, fmt.Errorf("invalid sort type. use '--sort-name' or '--sort-created'")
	}

	return sortBy, nil
}

func parseOrderType(orderType string) (vfs.OrderType, error) {
	var orderBy vfs.OrderType

	switch orderType {
	case "asc":
		orderBy = vfs.Asc
	case "desc":
		orderBy = vfs.Desc
	default:
		return orderBy, fmt.Errorf("invalid order type. use 'asc' or 'desc'")
	}

	return orderBy, nil
}
