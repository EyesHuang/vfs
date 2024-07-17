package mock

import (
	"vfs"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) AddUser(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func (m *MockUserService) GetUser(name string) *vfs.User {
	args := m.Called(name)
	return args.Get(0).(*vfs.User)
}
