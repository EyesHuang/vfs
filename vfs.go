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

type FolderKeySet struct {
	UserName   string
	FolderName string
}

type GetFoldersRequest struct {
	UserName string
	SortBy   SortType
	OrderBy  OrderType
}

type UpdateFolderRequest struct {
	OldName  string
	NewName  string
	UserName string
}

type FileKeySet struct {
	UserName   string
	FolderName string
	FileName   string
}

type GetFilesRequest struct {
	UserName   string
	FolderName string
	SortBy     SortType
	OrderBy    OrderType
}

type SortType string

type OrderType string

const (
	Asc        OrderType = "asc"
	Desc       OrderType = "desc"
	FolderName SortType  = "--sort-name"
	Created    SortType  = "--sort-created"
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
	DeleteFolder(key FolderKeySet) error
	GetFolders(req *GetFoldersRequest) ([]*Folder, error)
	UpdateFolder(req *UpdateFolderRequest) error
}

type FolderRepository interface {
	GetFolder(key FolderKeySet) *Folder
	GetFolders(req *GetFoldersRequest) []*Folder
	AddFolder(folder *Folder) error
	UpdateFolder(req *UpdateFolderRequest) error
	DeleteFolder(key FolderKeySet) error
}

type FileService interface {
	AddFile(file *File) error
	DeleteFile(key FileKeySet) error
	GetFiles(req *GetFilesRequest) ([]*File, error)
}

type FileRepository interface {
	GetFile(key FileKeySet) *File
	GetFiles(req *GetFilesRequest) []*File
	AddFile(file *File) error
	DeleteFile(key FileKeySet) error
}
