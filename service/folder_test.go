package service

import (
	"testing"

	"vfs"
	"vfs/mock"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddFolder_Success(t *testing.T) {
	mockFolderRepo := new(mock.MockFolderRepo)
	mockUserRepo := new(mock.MockUserRepo)
	service := NewFolderManageService(mockFolderRepo, mockUserRepo)

	user := &vfs.User{ID: uuid.New(), Name: "user1"}
	folder := &vfs.Folder{UserName: "user1", Name: "folder1"}

	mockUserRepo.On("GetUser", "user1").Return(user)
	mockFolderRepo.On("AddFolder", folder).Return(nil)

	err := service.AddFolder(folder)
	assert.Nil(t, err)
	mockUserRepo.AssertExpectations(t)
	mockFolderRepo.AssertExpectations(t)
}

func TestAddFolder_UserNotFound(t *testing.T) {
	mockFolderRepo := new(mock.MockFolderRepo)
	mockUserRepo := new(mock.MockUserRepo)
	service := NewFolderManageService(mockFolderRepo, mockUserRepo)

	folder := &vfs.Folder{UserName: "nonexistent", Name: "folder1"}

	mockUserRepo.On("GetUser", "nonexistent").Return((*vfs.User)(nil))

	err := service.AddFolder(folder)
	assert.NotNil(t, err)
	assert.Equal(t, "The nonexistent doesn't exist.", err.Error())
	mockUserRepo.AssertExpectations(t)
}

func TestDeleteFolder_Success(t *testing.T) {
	mockFolderRepo := new(mock.MockFolderRepo)
	mockUserRepo := new(mock.MockUserRepo)
	service := NewFolderManageService(mockFolderRepo, mockUserRepo)

	key := vfs.KeySet{UserName: "user1", FolderName: "folder1"}
	user := &vfs.User{ID: uuid.New(), Name: "user1"}

	mockUserRepo.On("GetUser", "user1").Return(user)
	mockFolderRepo.On("DeleteFolder", key).Return(nil)

	err := service.DeleteFolder(key)
	assert.Nil(t, err)
	mockUserRepo.AssertExpectations(t)
	mockFolderRepo.AssertExpectations(t)
}

func TestDeleteFolder_UserNotFound(t *testing.T) {
	mockFolderRepo := new(mock.MockFolderRepo)
	mockUserRepo := new(mock.MockUserRepo)
	service := NewFolderManageService(mockFolderRepo, mockUserRepo)

	key := vfs.KeySet{UserName: "nonexistent", FolderName: "folder1"}

	mockUserRepo.On("GetUser", "nonexistent").Return((*vfs.User)(nil))

	err := service.DeleteFolder(key)
	assert.NotNil(t, err)
	assert.Equal(t, "The nonexistent doesn't exist.", err.Error())
	mockUserRepo.AssertExpectations(t)
}

func TestGetFolders_Success(t *testing.T) {
	mockFolderRepo := new(mock.MockFolderRepo)
	mockUserRepo := new(mock.MockUserRepo)
	service := NewFolderManageService(mockFolderRepo, mockUserRepo)

	user := &vfs.User{ID: uuid.New(), Name: "user1"}
	req := &vfs.GetFoldersRequest{UserName: "user1"}
	folders := []*vfs.Folder{
		{UserName: "user1", Name: "folder1"},
		{UserName: "user1", Name: "folder2"},
	}

	mockUserRepo.On("GetUser", "user1").Return(user)
	mockFolderRepo.On("GetFolders", req).Return(folders)

	result, err := service.GetFolders(req)
	assert.Nil(t, err)
	assert.Equal(t, folders, result)
	mockUserRepo.AssertExpectations(t)
	mockFolderRepo.AssertExpectations(t)
}

func TestGetFolders_UserNotFound(t *testing.T) {
	mockFolderRepo := new(mock.MockFolderRepo)
	mockUserRepo := new(mock.MockUserRepo)
	service := NewFolderManageService(mockFolderRepo, mockUserRepo)

	req := &vfs.GetFoldersRequest{UserName: "nonexistent"}

	mockUserRepo.On("GetUser", "nonexistent").Return((*vfs.User)(nil))

	result, err := service.GetFolders(req)
	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "The nonexistent doesn't exist.", err.Error())
	mockUserRepo.AssertExpectations(t)
}

func TestUpdateFolder_Success(t *testing.T) {
	mockFolderRepo := new(mock.MockFolderRepo)
	mockUserRepo := new(mock.MockUserRepo)
	service := NewFolderManageService(mockFolderRepo, mockUserRepo)

	user := &vfs.User{ID: uuid.New(), Name: "user1"}
	req := &vfs.UpdateFolderRequest{UserName: "user1", OldName: "folder1", NewName: "folder2"}

	mockUserRepo.On("GetUser", "user1").Return(user)
	mockFolderRepo.On("UpdateFolder", req).Return(nil)

	err := service.UpdateFolder(req)
	assert.Nil(t, err)
	mockUserRepo.AssertExpectations(t)
	mockFolderRepo.AssertExpectations(t)
}

func TestUpdateFolder_UserNotFound(t *testing.T) {
	mockFolderRepo := new(mock.MockFolderRepo)
	mockUserRepo := new(mock.MockUserRepo)
	service := NewFolderManageService(mockFolderRepo, mockUserRepo)

	req := &vfs.UpdateFolderRequest{UserName: "nonexistent", OldName: "folder1", NewName: "folder2"}

	mockUserRepo.On("GetUser", "nonexistent").Return((*vfs.User)(nil))

	err := service.UpdateFolder(req)
	assert.NotNil(t, err)
	assert.Equal(t, "The nonexistent doesn't exist.", err.Error())
	mockUserRepo.AssertExpectations(t)
}
