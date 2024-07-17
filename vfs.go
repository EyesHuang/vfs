package vfs

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID
	Name string
}

type Folder struct {
	ID          uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
	Files       map[string]*File
	UserID      uuid.UUID
}

type File struct {
	ID          uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
	FolderID    uuid.UUID
}

type UserService interface {
	Register(name string) error
}

type UserRepository interface {
	AddUser(name string) error
}
