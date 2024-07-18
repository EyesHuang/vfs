package repo

import (
	"errors"
	"fmt"
	"sort"

	"vfs"
)

type MemoryFolderRepo struct {
	Folders map[vfs.KeySet]*vfs.Folder
}

func NewMemoFolderRepo() *MemoryFolderRepo {
	return &MemoryFolderRepo{
		make(map[vfs.KeySet]*vfs.Folder),
	}
}

func (mfr *MemoryFolderRepo) GetFolder(key vfs.KeySet) *vfs.Folder {
	if folder, exists := mfr.Folders[key]; exists {
		return folder
	}
	return nil
}

func (mfr *MemoryFolderRepo) AddFolder(folder *vfs.Folder) error {
	key := vfs.KeySet{
		UserName:   folder.UserName,
		FolderName: folder.Name,
	}
	if mfr.GetFolder(key) != nil {
		errMsg := fmt.Sprintf("The %s has already existed.", folder.Name)
		return errors.New(errMsg)
	}
	mfr.Folders[key] = folder
	return nil
}

func (mfr *MemoryFolderRepo) UpdateFolder(req *vfs.UpdateFolderRequest) error {
	oldKey := vfs.KeySet{
		UserName:   req.UserName,
		FolderName: req.OldName,
	}
	newKey := vfs.KeySet{
		UserName:   req.UserName,
		FolderName: req.NewName,
	}

	if mfr.GetFolder(oldKey) == nil {
		errMsg := fmt.Sprintf("The folder %s doesn't exist.", req.OldName)
		return errors.New(errMsg)
	}

	if mfr.GetFolder(newKey) != nil {
		errMsg := fmt.Sprintf("The folder %s already exists.", req.NewName)
		return errors.New(errMsg)
	}

	// Perform the renaming
	folder := mfr.Folders[oldKey]
	delete(mfr.Folders, oldKey)
	folder.Name = req.NewName
	mfr.Folders[newKey] = folder

	return nil
}

func (mfr *MemoryFolderRepo) DeleteFolder(key vfs.KeySet) error {
	if mfr.GetFolder(key) == nil {
		errMsg := fmt.Sprintf("The %s doesn't exist.", key.FolderName)
		return errors.New(errMsg)
	}
	delete(mfr.Folders, key)
	return nil
}

func (mfr *MemoryFolderRepo) GetFolders(req *vfs.GetFoldersRequest) []*vfs.Folder {
	var folders []*vfs.Folder
	for key, folder := range mfr.Folders {
		if key.UserName == req.UserName {
			folders = append(folders, folder)
		}
	}

	// Sorting
	switch req.SortBy {
	case vfs.FolderName:
		if req.OrderBy == vfs.Asc {
			sort.Slice(folders, func(i, j int) bool {
				return folders[i].Name < folders[j].Name
			})
		} else {
			sort.Slice(folders, func(i, j int) bool {
				return folders[i].Name > folders[j].Name
			})
		}
	case vfs.Created:
		if req.OrderBy == vfs.Asc {
			sort.Slice(folders, func(i, j int) bool {
				return folders[i].CreatedAt.Before(folders[j].CreatedAt)
			})
		} else {
			sort.Slice(folders, func(i, j int) bool {
				return folders[i].CreatedAt.After(folders[j].CreatedAt)
			})
		}
	default:
		sort.Slice(folders, func(i, j int) bool {
			return folders[i].Name < folders[j].Name
		})
	}

	return folders
}
