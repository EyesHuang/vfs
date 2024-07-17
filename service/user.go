package service

import "vfs"

type UserManageService struct {
	repo vfs.UserRepository
}

func NewUserManageService(repo vfs.UserRepository) *UserManageService {
	return &UserManageService{repo: repo}
}

func (ums *UserManageService) Register(name string) error {
	if err := ums.repo.AddUser(name); err != nil {
		return err
	}
	return nil
}
