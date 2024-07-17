package repo

import (
	"errors"
	"fmt"

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

func (mfr *MemoryFolderRepo) UpdateFolder(folder *vfs.Folder) error {
	key := vfs.KeySet{
		UserName:   folder.UserName,
		FolderName: folder.Name,
	}
	if _, exists := mfr.Folders[key]; !exists {
		errMsg := fmt.Sprintf("The %s doesn't exist.", folder.Name)
		return errors.New(errMsg)
	}
	mfr.Folders[key] = folder
	return nil
}

func (mfr *MemoryFolderRepo) DeleteFolder(key vfs.KeySet) error {
	if mfr.GetFolder(key) != nil {
		errMsg := fmt.Sprintf("The %s doesn't exist.", key.FolderName)
		return errors.New(errMsg)
	}
	delete(mfr.Folders, key)
	return nil
}

func (mfr *MemoryFolderRepo) GetFolders(req *vfs.GetFoldersRequest) ([]*vfs.Folder, error) {
	return nil, nil
}
