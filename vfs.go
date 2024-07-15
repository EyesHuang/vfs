package vfs

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username string
}

type UserManager struct {
	Users map[string]*User
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
