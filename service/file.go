package service

import (
	"errors"
	"fmt"
	"time"

	"vfs"

	"github.com/google/uuid"
)

type FileManageService struct {
	fileRepo   vfs.FileRepository
	userRepo   vfs.UserRepository
	folderRepo vfs.FolderRepository
}

func NewFileManageService(fileRepo vfs.FileRepository, userRepo vfs.UserRepository,
	folderRepo vfs.FolderRepository,
) *FileManageService {
	return &FileManageService{fileRepo: fileRepo, userRepo: userRepo, folderRepo: folderRepo}
}

func (fms *FileManageService) AddFile(file *vfs.File) error {
	user := fms.userRepo.GetUser(file.UserName)
	if user == nil {
		errMsg := fmt.Sprintf("The %s doesn't exist.", file.UserName)
		return errors.New(errMsg)
	}

	folder := fms.folderRepo.GetFolder(vfs.FolderKeySet{UserName: file.UserName, FolderName: file.FolderName})
	if folder == nil {
		errMsg := fmt.Sprintf("The %s doesn't exist.", file.FolderName)
		return errors.New(errMsg)
	}

	file.UserID = user.ID
	file.FolderID = folder.ID
	file.ID = uuid.New()
	file.CreatedAt = time.Now()

	if err := fms.fileRepo.AddFile(file); err != nil {
		return err
	}

	return nil
}

func (fms *FileManageService) DeleteFile(key vfs.FileKeySet) error {
	user := fms.userRepo.GetUser(key.UserName)
	if user == nil {
		errMsg := fmt.Sprintf("The %s doesn't exist.", key.UserName)
		return errors.New(errMsg)
	}

	folder := fms.folderRepo.GetFolder(vfs.FolderKeySet{UserName: key.UserName, FolderName: key.FolderName})
	if folder == nil {
		errMsg := fmt.Sprintf("The %s doesn't exist.", key.FolderName)
		return errors.New(errMsg)
	}

	if err := fms.fileRepo.DeleteFile(key); err != nil {
		return err
	}

	return nil
}

func (fms *FileManageService) GetFiles(req *vfs.GetFilesRequest) ([]*vfs.File, error) {
	user := fms.userRepo.GetUser(req.UserName)
	if user == nil {
		errMsg := fmt.Sprintf("The %s doesn't exist.", req.UserName)
		return nil, errors.New(errMsg)
	}

	folder := fms.folderRepo.GetFolder(vfs.FolderKeySet{UserName: req.UserName, FolderName: req.FolderName})
	if folder == nil {
		errMsg := fmt.Sprintf("The %s doesn't exist.", req.FolderName)
		return nil, errors.New(errMsg)
	}

	return fms.fileRepo.GetFiles(req), nil
}
