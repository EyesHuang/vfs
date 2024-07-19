package mock

import (
	"time"

	"vfs"
)

var (
	now = time.Now()

	File1 = &vfs.File{
		UserName:   "user1",
		FolderName: "folder1",
		Name:       "file1",
		CreatedAt:  now.Add(-2 * time.Hour),
	}

	File2 = &vfs.File{
		UserName:   "user1",
		FolderName: "folder1",
		Name:       "file2",
		CreatedAt:  now.Add(-1 * time.Hour),
	}

	File3 = &vfs.File{
		UserName:   "user2",
		FolderName: "folder2",
		Name:       "file3",
		CreatedAt:  now,
	}

	FileKeySet = vfs.FileKeySet{UserName: "user1", FolderName: "folder1", FileName: "file1"}

	GetFilesRequest = &vfs.GetFilesRequest{
		UserName:   "user1",
		FolderName: "folder1",
		SortBy:     vfs.Name,
		OrderBy:    vfs.Asc,
	}
)
