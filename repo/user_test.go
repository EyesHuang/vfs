package repo

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddUser_Success(t *testing.T) {
	repo := NewMemoUserRepo()

	err := repo.AddUser("user1")
	assert.Nil(t, err)
	assert.Contains(t, repo.Users, "user1")
	assert.Equal(t, "user1", repo.Users["user1"].Name)
	_, err = uuid.Parse(repo.Users["user1"].ID.String()) // Check if ID is a valid UUID
	assert.Nil(t, err)
}

func TestAddUser_ExistingUser(t *testing.T) {
	repo := NewMemoUserRepo()

	// Add a user first
	err := repo.AddUser("user1")
	assert.Nil(t, err)

	// Test adding an existing user
	err = repo.AddUser("user1")
	assert.NotNil(t, err)
	assert.Equal(t, "The user1 has already existed.", err.Error())
}

func TestGetUser_Success(t *testing.T) {
	repo := NewMemoUserRepo()

	// Add a user first
	err := repo.AddUser("user1")
	assert.Nil(t, err)

	// Retrieve the user
	user := repo.GetUser("user1")
	assert.NotNil(t, user)
	assert.Equal(t, "user1", user.Name)
}

func TestGetUser_NotFound(t *testing.T) {
	repo := NewMemoUserRepo()

	// Try to retrieve a non-existing user
	user := repo.GetUser("nonExistentUser")
	assert.Nil(t, user)
}
