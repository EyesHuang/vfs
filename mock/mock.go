package mock

import "github.com/stretchr/testify/mock"

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) AddUser(name string) error {
	args := m.Called(name)
	return args.Error(0)
}
