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

type KeySet struct {
	UserName   string
	FolderName string
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
	AddFolder(folder *Folder) error
	DeleteFolder(key KeySet) error
	GetFolders(req *GetFoldersRequest) (*Folder, error)
	UpdateFolder(oldName, newName string) error
}

type FolderRepository interface {
	GetFolder(key KeySet) *Folder
	AddFolder(folder *Folder) error
	UpdateFolder(folder *Folder) error
	DeleteFolder(key KeySet) error
}
