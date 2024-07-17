package repo

import (
	"errors"
	"fmt"

	"vfs"

	"github.com/google/uuid"
)

type MemoryUserRepo struct {
	Users map[string]*vfs.User
}

func NewMemoUserRepo() *MemoryUserRepo {
	return &MemoryUserRepo{
		make(map[string]*vfs.User),
	}
}

func (mur *MemoryUserRepo) AddUser(name string) error {
	if _, exists := mur.Users[name]; exists {
		errMsg := fmt.Sprintf("The %s has already existed.", name)
		return errors.New(errMsg)
	}
	id := uuid.New()
	mur.Users[name] = &vfs.User{
		ID:   id,
		Name: name,
	}
	return nil
}
