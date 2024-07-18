package service

import (
	"errors"
	"fmt"
	"time"

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
	folder.CreatedAt = time.Now()

	if err := fms.folderRepo.AddFolder(folder); err != nil {
		return err
	}
	return nil
}

func (fms *FolderManageService) DeleteFolder(key vfs.KeySet) error {
	user := fms.userRepo.GetUser(key.UserName)
	if user == nil {
		errMsg := fmt.Sprintf("The %s doesn't exist.", key.UserName)
		return errors.New(errMsg)
	}

	if err := fms.folderRepo.DeleteFolder(key); err != nil {
		return err
	}

	return nil
}

func (fms *FolderManageService) GetFolders(req *vfs.GetFoldersRequest) ([]*vfs.Folder, error) {
	user := fms.userRepo.GetUser(req.UserName)
	if user == nil {
		errMsg := fmt.Sprintf("The %s doesn't exist.", req.UserName)
		return nil, errors.New(errMsg)
	}

	return fms.folderRepo.GetFolders(req), nil
}

func (fms *FolderManageService) UpdateFolder(req *vfs.UpdateFolderRequest) error {
	user := fms.userRepo.GetUser(req.UserName)
	if user == nil {
		errMsg := fmt.Sprintf("The %s doesn't exist.", req.UserName)
		return errors.New(errMsg)
	}

	if err := fms.folderRepo.UpdateFolder(req); err != nil {
		return err
	}

	return nil
}
