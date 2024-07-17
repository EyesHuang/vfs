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
	UserID      uuid.UUID
	UserName    string
}

type File struct {
	ID          uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
	FolderID    uuid.UUID
}

type GetFoldersRequest struct {
	UserName string
	SortBy   string
	OrderBy  OrderType
}

type OrderType string

const (
	Asc  OrderType = "asc"
	Desc OrderType = "desc"
)

type UserService interface {
	Register(name string) error
}

type UserRepository interface {
	AddUser(name string) error
	GetUser(name string) *User
}

type FolderService interface {
	AddFolder(req *GetFoldersRequest) error
	DeleteFolder(name string) error
	GetFolders(name string) (*Folder, error)
	UpdateFolder(oldName, newName string) error
}

type FolderRepository interface {
	GetFolder(name string) (*Folder, error)
	AddFolder(folder *Folder) error
	UpdateFolder(folder *Folder) error
	DeleteFolder(name string) error
}
