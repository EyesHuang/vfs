package mock

import (
	"vfs"

	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) AddUser(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func (m *MockUserRepo) GetUser(name string) *vfs.User {
	args := m.Called(name)
	return args.Get(0).(*vfs.User)
}

type MockFolderRepo struct {
	mock.Mock
}

func (m *MockFolderRepo) GetFolder(key vfs.FolderKeySet) *vfs.Folder {
	args := m.Called(key)
	return args.Get(0).(*vfs.Folder)
}

func (m *MockFolderRepo) GetFolders(req *vfs.GetFoldersRequest) []*vfs.Folder {
	args := m.Called(req)
	return args.Get(0).([]*vfs.Folder)
}

func (m *MockFolderRepo) AddFolder(folder *vfs.Folder) error {
	args := m.Called(folder)
	return args.Error(0)
}

func (m *MockFolderRepo) UpdateFolder(req *vfs.UpdateFolderRequest) error {
	args := m.Called(req)
	return args.Error(0)
}

func (m *MockFolderRepo) DeleteFolder(key vfs.FolderKeySet) error {
	args := m.Called(key)
	return args.Error(0)
}

type MockFileRepo struct {
	mock.Mock
}

func (m *MockFileRepo) GetFile(key vfs.FileKeySet) *vfs.File {
	args := m.Called(key)
	return args.Get(0).(*vfs.File)
}

func (m *MockFileRepo) GetFiles(req *vfs.GetFilesRequest) []*vfs.File {
	args := m.Called(req)
	return args.Get(0).([]*vfs.File)
}

func (m *MockFileRepo) AddFile(file *vfs.File) error {
	args := m.Called(file)
	return args.Error(0)
}

func (m *MockFileRepo) DeleteFile(key vfs.FileKeySet) error {
	args := m.Called(key)
	return args.Error(0)
}
