package service

import (
	"testing"
	"vfs"
	"vfs/mock"

	"github.com/stretchr/testify/assert"
)

func TestAddFile_Success(t *testing.T) {
	mockFileRepo := new(mock.MockFileRepo)
	mockUserRepo := new(mock.MockUserRepo)
	mockFolderRepo := new(mock.MockFolderRepo)
	service := NewFileManageService(mockFileRepo, mockUserRepo, mockFolderRepo)

	user := &vfs.User{Name: "user1"}
	key := vfs.FolderKeySet{UserName: "user1", FolderName: "folder1"}
	folder := &vfs.Folder{UserName: "user1", Name: "folder1"}
	file := &vfs.File{
		UserName:   "user1",
		FolderName: "folder1",
		Name:       "file1",
	}

	mockUserRepo.On("GetUser", "user1").Return(user)
	mockFolderRepo.On("GetFolder", key).Return(folder)
	mockFileRepo.On("AddFile", file).Return(nil)

	err := service.AddFile(file)
	assert.Nil(t, err)

	mockUserRepo.AssertExpectations(t)
	mockFolderRepo.AssertExpectations(t)
	mockFileRepo.AssertExpectations(t)
}

func TestAddFile_UserNotFound(t *testing.T) {
	mockFileRepo := new(mock.MockFileRepo)
	mockUserRepo := new(mock.MockUserRepo)
	mockFolderRepo := new(mock.MockFolderRepo)
	service := NewFileManageService(mockFileRepo, mockUserRepo, mockFolderRepo)

	file := &vfs.File{
		UserName:   "nonexistent",
		FolderName: "folder1",
		Name:       "file1",
	}

	mockUserRepo.On("GetUser", "nonexistent").Return((*vfs.User)(nil))

	err := service.AddFile(file)
	assert.NotNil(t, err)
	assert.Equal(t, "The nonexistent doesn't exist.", err.Error())

	mockUserRepo.AssertExpectations(t)
}

func TestAddFile_FolderNotFound(t *testing.T) {
	mockFileRepo := new(mock.MockFileRepo)
	mockUserRepo := new(mock.MockUserRepo)
	mockFolderRepo := new(mock.MockFolderRepo)
	service := NewFileManageService(mockFileRepo, mockUserRepo, mockFolderRepo)

	user := &vfs.User{Name: "user1"}
	key := vfs.FolderKeySet{UserName: "user1", FolderName: "nonexistent"}
	file := &vfs.File{
		UserName:   "user1",
		FolderName: "nonexistent",
		Name:       "file1",
	}

	mockUserRepo.On("GetUser", "user1").Return(user)
	mockFolderRepo.On("GetFolder", key).Return((*vfs.Folder)(nil))

	err := service.AddFile(file)
	assert.NotNil(t, err)
	assert.Equal(t, "The nonexistent doesn't exist.", err.Error())

	mockUserRepo.AssertExpectations(t)
	mockFolderRepo.AssertExpectations(t)
}

func TestDeleteFile_Success(t *testing.T) {
	mockFileRepo := new(mock.MockFileRepo)
	mockUserRepo := new(mock.MockUserRepo)
	mockFolderRepo := new(mock.MockFolderRepo)
	service := NewFileManageService(mockFileRepo, mockUserRepo, mockFolderRepo)

	user := &vfs.User{Name: "user1"}
	folderKeySet := vfs.FolderKeySet{UserName: "user1", FolderName: "folder1"}
	fileKeySet := vfs.FileKeySet{UserName: "user1", FolderName: "folder1", FileName: "file1"}
	folder := &vfs.Folder{UserName: "user1", Name: "folder1"}

	mockUserRepo.On("GetUser", "user1").Return(user)
	mockFolderRepo.On("GetFolder", folderKeySet).Return(folder)
	mockFileRepo.On("DeleteFile", fileKeySet).Return(nil)

	err := service.DeleteFile(fileKeySet)
	assert.Nil(t, err)

	mockUserRepo.AssertExpectations(t)
	mockFolderRepo.AssertExpectations(t)
	mockFileRepo.AssertExpectations(t)
}

func TestDeleteFile_UserNotFound(t *testing.T) {
	mockFileRepo := new(mock.MockFileRepo)
	mockUserRepo := new(mock.MockUserRepo)
	mockFolderRepo := new(mock.MockFolderRepo)
	service := NewFileManageService(mockFileRepo, mockUserRepo, mockFolderRepo)

	fileKeySet := vfs.FileKeySet{UserName: "nonexistent", FolderName: "folder1", FileName: "file1"}

	mockUserRepo.On("GetUser", "nonexistent").Return((*vfs.User)(nil))

	err := service.DeleteFile(fileKeySet)
	assert.NotNil(t, err)
	assert.Equal(t, "The nonexistent doesn't exist.", err.Error())

	mockUserRepo.AssertExpectations(t)
}

func TestDeleteFile_FolderNotFound(t *testing.T) {
	mockFileRepo := new(mock.MockFileRepo)
	mockUserRepo := new(mock.MockUserRepo)
	mockFolderRepo := new(mock.MockFolderRepo)
	service := NewFileManageService(mockFileRepo, mockUserRepo, mockFolderRepo)

	user := &vfs.User{Name: "user1"}
	folderKeySet := vfs.FolderKeySet{UserName: "user1", FolderName: "nonexistent"}
	fileKeySet := vfs.FileKeySet{UserName: "user1", FolderName: "nonexistent", FileName: "file1"}

	mockUserRepo.On("GetUser", "user1").Return(user)
	mockFolderRepo.On("GetFolder", folderKeySet).Return((*vfs.Folder)(nil))

	err := service.DeleteFile(fileKeySet)
	assert.NotNil(t, err)
	assert.Equal(t, "The nonexistent doesn't exist.", err.Error())

	mockUserRepo.AssertExpectations(t)
	mockFolderRepo.AssertExpectations(t)
}

func TestGetFiles_Success(t *testing.T) {
	mockFileRepo := new(mock.MockFileRepo)
	mockUserRepo := new(mock.MockUserRepo)
	mockFolderRepo := new(mock.MockFolderRepo)
	service := NewFileManageService(mockFileRepo, mockUserRepo, mockFolderRepo)

	user := &vfs.User{Name: "user1"}
	folderKeySet := vfs.FolderKeySet{UserName: "user1", FolderName: "folder1"}
	folder := &vfs.Folder{UserName: "user1", Name: "folder1"}
	req := &vfs.GetFilesRequest{UserName: "user1", FolderName: "folder1", SortBy: vfs.Name, OrderBy: vfs.Asc}

	files := []*vfs.File{
		{UserName: "user1", FolderName: "folder1", Name: "file1"},
		{UserName: "user1", FolderName: "folder1", Name: "file2"},
	}

	mockUserRepo.On("GetUser", "user1").Return(user)
	mockFolderRepo.On("GetFolder", folderKeySet).Return(folder)
	mockFileRepo.On("GetFiles", req).Return(files)

	result, err := service.GetFiles(req)
	assert.Nil(t, err)
	assert.Equal(t, files, result)

	mockUserRepo.AssertExpectations(t)
	mockFolderRepo.AssertExpectations(t)
	mockFileRepo.AssertExpectations(t)
}

func TestGetFiles_UserNotFound(t *testing.T) {
	mockFileRepo := new(mock.MockFileRepo)
	mockUserRepo := new(mock.MockUserRepo)
	mockFolderRepo := new(mock.MockFolderRepo)
	service := NewFileManageService(mockFileRepo, mockUserRepo, mockFolderRepo)

	req := &vfs.GetFilesRequest{UserName: "nonexistent", FolderName: "folder1", SortBy: vfs.Name, OrderBy: vfs.Asc}

	mockUserRepo.On("GetUser", "nonexistent").Return((*vfs.User)(nil))

	result, err := service.GetFiles(req)
	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "The nonexistent doesn't exist.", err.Error())

	mockUserRepo.AssertExpectations(t)
}

func TestGetFiles_FolderNotFound(t *testing.T) {
	mockFileRepo := new(mock.MockFileRepo)
	mockUserRepo := new(mock.MockUserRepo)
	mockFolderRepo := new(mock.MockFolderRepo)
	service := NewFileManageService(mockFileRepo, mockUserRepo, mockFolderRepo)

	user := &vfs.User{Name: "user1"}
	key := vfs.FolderKeySet{UserName: "user1", FolderName: "nonexistent"}
	req := &vfs.GetFilesRequest{UserName: "user1", FolderName: "nonexistent", SortBy: vfs.Name, OrderBy: vfs.Asc}

	mockUserRepo.On("GetUser", "user1").Return(user)
	mockFolderRepo.On("GetFolder", key).Return((*vfs.Folder)(nil))

	result, err := service.GetFiles(req)
	assert.NotNil(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "The nonexistent doesn't exist.", err.Error())

	mockUserRepo.AssertExpectations(t)
	mockFolderRepo.AssertExpectations(t)
}
