package repo

import (
	"errors"
	"fmt"

	"vfs"
)

type MemoryFolderRepo struct {
	Folders map[string]*vfs.Folder
}

func NewMemoFolderRepo() *MemoryFolderRepo {
	return &MemoryFolderRepo{
		make(map[string]*vfs.Folder),
	}
}

func (mfr *MemoryFolderRepo) GetFolder(name string) (*vfs.Folder, error) {
	if folder, exists := mfr.Folders[name]; exists {
		return folder, nil
	}
	return nil, nil
}

func (mfr *MemoryFolderRepo) AddFolder(folder *vfs.Folder) error {
	if _, exists := mfr.Folders[folder.Name]; exists {
		errMsg := fmt.Sprintf("The %s has already existed.", folder.Name)
		return errors.New(errMsg)
	}
	mfr.Folders[folder.Name] = folder
	return nil
}

func (mfr *MemoryFolderRepo) UpdateFolder(folder *vfs.Folder) error {
	if _, exists := mfr.Folders[folder.Name]; !exists {
		errMsg := fmt.Sprintf("The %s doesn't exist.", folder.Name)
		return errors.New(errMsg)
	}
	mfr.Folders[folder.Name] = folder
	return nil
}

func (mfr *MemoryFolderRepo) DeleteFolder(name string) error {
	if _, exists := mfr.Folders[name]; !exists {
		errMsg := fmt.Sprintf("The %s doesn't exist.", name)
		return errors.New(errMsg)
	}
	delete(mfr.Folders, name)
	return nil
}

func (mfr *MemoryFolderRepo) GetFolders(req *vfs.GetFoldersRequest) ([]*vfs.Folder, error) {
	return nil, nil
}
