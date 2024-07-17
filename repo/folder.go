package repo

import (
	"errors"
	"fmt"

	"vfs"
)

type MemoryFolderRepo struct {
	Folders map[keySet]*vfs.Folder
}

type keySet struct {
	userName   string
	folderName string
}

func NewMemoFolderRepo() *MemoryFolderRepo {
	return &MemoryFolderRepo{
		make(map[keySet]*vfs.Folder),
	}
}

func (mfr *MemoryFolderRepo) GetFolder(name string) (*vfs.Folder, error) {
	key := keySet{
		folderName: name,
	}
	if folder, exists := mfr.Folders[key]; exists {
		return folder, nil
	}
	return nil, nil
}

func (mfr *MemoryFolderRepo) AddFolder(folder *vfs.Folder) error {
	key := keySet{
		userName:   folder.UserName,
		folderName: folder.Name,
	}
	if _, exists := mfr.Folders[key]; exists {
		errMsg := fmt.Sprintf("The %s has already existed.", folder.Name)
		return errors.New(errMsg)
	}
	mfr.Folders[key] = folder
	return nil
}

func (mfr *MemoryFolderRepo) UpdateFolder(folder *vfs.Folder) error {
	key := keySet{
		userName:   folder.UserName,
		folderName: folder.Name,
	}
	if _, exists := mfr.Folders[key]; !exists {
		errMsg := fmt.Sprintf("The %s doesn't exist.", folder.Name)
		return errors.New(errMsg)
	}
	mfr.Folders[key] = folder
	return nil
}

func (mfr *MemoryFolderRepo) DeleteFolder(name string) error {
	key := keySet{
		folderName: name,
	}
	if _, exists := mfr.Folders[key]; !exists {
		errMsg := fmt.Sprintf("The %s doesn't exist.", name)
		return errors.New(errMsg)
	}
	delete(mfr.Folders, key)
	return nil
}

func (mfr *MemoryFolderRepo) GetFolders(req *vfs.GetFoldersRequest) ([]*vfs.Folder, error) {
	return nil, nil
}
