package mock

import (
	"time"

	"vfs"
)

var (
	now = time.Now()

	User1 = &vfs.User{Name: "user1"}

	Folder1 = &vfs.Folder{UserName: User1.Name, Name: "folder1"}

	FolderKeySet = vfs.FolderKeySet{UserName: User1.Name, FolderName: Folder1.Name}

	File1 = &vfs.File{
		UserName:   User1.Name,
		FolderName: Folder1.Name,
		Name:       "file1",
		CreatedAt:  now.Add(-2 * time.Hour),
	}

	File2 = &vfs.File{
		UserName:   User1.Name,
		FolderName: Folder1.Name,
		Name:       "file2",
		CreatedAt:  now.Add(-1 * time.Hour),
	}

	File3 = &vfs.File{
		UserName:   "user2",
		FolderName: "folder2",
		Name:       "file3",
		CreatedAt:  now,
	}

	FileKeySet = vfs.FileKeySet{UserName: User1.Name, FolderName: Folder1.Name, FileName: File1.Name}

	GetFilesRequest = &vfs.GetFilesRequest{
		UserName:   User1.Name,
		FolderName: Folder1.Name,
		SortBy:     vfs.Name,
		OrderBy:    vfs.Asc,
	}
)
