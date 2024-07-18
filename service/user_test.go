package service

import (
	"errors"
	"testing"

	"vfs/mock"

	"github.com/stretchr/testify/assert"
)

func TestRegister_Success(t *testing.T) {
	mockRepo := new(mock.MockUserRepo)
	service := NewUserManageService(mockRepo)
	mockRepo.On("AddUser", "user1").Return(nil)

	err := service.Register("user1")
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRegister_ExistingUser(t *testing.T) {
	mockRepo := new(mock.MockUserRepo)
	service := NewUserManageService(mockRepo)
	mockRepo.On("AddUser", "user1").Return(nil).Once()

	// Add a user first
	err := service.Register("user1")
	assert.Nil(t, err)

	mockRepo.On("AddUser", "user1").Return(errors.New("The user1 has already existed.")).Once()

	// Attempt to add the same user again
	err = service.Register("user1")
	assert.NotNil(t, err)
	assert.Equal(t, "The user1 has already existed.", err.Error())
	mockRepo.AssertExpectations(t)
}
