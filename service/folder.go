package service

import (
	"errors"
	"fmt"

	"vfs"
)

type FolderManageService struct {
	folderRepo vfs.FolderRepository
	userRepo   vfs.UserRepository
}

func NewFolderManageService(folderRepo vfs.FolderRepository, userRepo vfs.UserRepository) *FolderManageService {
	return &FolderManageService{folderRepo: folderRepo, userRepo: userRepo}
}

func (fms *FolderManageService) AddFolder(folder *vfs.Folder) error {
	user := fms.userRepo.GetUser(folder.UserName)
	if user == nil {
		errMsg := fmt.Sprintf("The %s doesn't exist.", folder.UserName)
		return errors.New(errMsg)
	}

	folder.UserID = user.ID

	if err := fms.folderRepo.AddFolder(folder); err != nil {
		return err
	}
	return nil
}

func (fms *FolderManageService) DeleteFolder(name string) error {
	return nil
}

func (fms *FolderManageService) GetFolders(req *vfs.GetFoldersRequest) (*vfs.Folder, error) {
	return nil, nil
}

func (fms *FolderManageService) UpdateFolder(oldName, newName string) error {
	return nil
}
