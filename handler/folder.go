package handler

import (
	"fmt"

	"vfs"
)

func (hm *HandlerManager) HandleCreateFolder(args []string) {
	if len(args) < 2 || len(args) > 3 {
		fmt.Println("Usage: create-folder [username] [foldername] [description]?")
		return
	}

	if !isValidFolderFileName(args[1]) {
		fmt.Printf("The %s contains invalid chars.\n", args[1])
		return
	}

	for _, name := range args[:2] {
		if err := validateName(name); err != nil {
			fmt.Println(err)
			return
		}
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

	if !isValidFolderFileName(args[1]) {
		fmt.Printf("The %s contains invalid chars.\n", args[1])
		return
	}

	for _, name := range args {
		if err := validateName(name); err != nil {
			fmt.Println(err)
			return
		}
	}

	key := vfs.FolderKeySet{
		UserName:   args[0],
		FolderName: args[1],
	}

	if err := hm.folderService.DeleteFolder(key); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Delete %s successfully.\n", args[1])
	}
}

func (hm *HandlerManager) HandleListFolders(args []string) {
	if len(args) < 1 || len(args) > 3 {
		fmt.Println("Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]")
		return
	}

	if err := validateName(args[0]); err != nil {
		fmt.Println(err)
		return
	}

	req := &vfs.GetFoldersRequest{
		UserName: args[0],
	}

	if len(args) == 3 {
		sortBy, err := parseSortType(args[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		orderBy, err := parseOrderType(args[2])
		if err != nil {
			fmt.Println(err)
			return
		}

		req.SortBy = sortBy
		req.OrderBy = orderBy
	}

	folders, err := hm.folderService.GetFolders(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(folders) == 0 {
		fmt.Printf("Warning: The %s doesn't have any folders.\n", req.UserName)
		return
	}

	for _, folder := range folders {
		fmt.Printf("%s\t%s\t%s\t%s\n", folder.Name, folder.Description, folder.CreatedAt.Format("2006-01-02 15:04:05"), folder.UserName)
	}
}

func (hm *HandlerManager) HandleRenameFolder(args []string) {
	if len(args) != 3 {
		fmt.Println("Usage: rename-folder [username] [foldername] [new-folder-name]")
		return
	}

	for _, name := range args[1:] {
		if !isValidFolderFileName(name) {
			fmt.Printf("Folder name %s contains invalid chars.\n", name)
			return
		}
	}

	for _, name := range args {
		if err := validateName(name); err != nil {
			fmt.Println(err)
			return
		}
	}

	req := &vfs.UpdateFolderRequest{
		OldName:  args[1],
		NewName:  args[2],
		UserName: args[0],
	}

	if err := hm.folderService.UpdateFolder(req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Rename %s to %s successfully.\n", args[1], args[2])
	}
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
