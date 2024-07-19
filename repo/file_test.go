package repo

import (
	"testing"
	"time"

	"vfs"

	"github.com/stretchr/testify/assert"
)

func TestAddFile_Success(t *testing.T) {
	repo := NewMemoryFileRepo()

	file := &vfs.File{
		UserName:   "user1",
		FolderName: "folder1",
		Name:       "file1",
	}

	err := repo.AddFile(file)
	assert.Nil(t, err)
	assert.NotNil(t, repo.GetFile(vfs.FileKeySet{UserName: "user1", FolderName: "folder1", FileName: "file1"}))
}

func TestAddFile_ExistingFile(t *testing.T) {
	repo := NewMemoryFileRepo()

	file := &vfs.File{
		UserName:   "user1",
		FolderName: "folder1",
		Name:       "file1",
	}
	_ = repo.AddFile(file)

	err := repo.AddFile(file)
	assert.NotNil(t, err)
	assert.Equal(t, "The file1 has already existed.", err.Error())
}

func TestGetFile(t *testing.T) {
	repo := NewMemoryFileRepo()

	file := &vfs.File{
		UserName:   "user1",
		FolderName: "folder1",
		Name:       "file1",
	}
	_ = repo.AddFile(file)

	retrievedFile := repo.GetFile(vfs.FileKeySet{UserName: "user1", FolderName: "folder1", FileName: "file1"})
	assert.NotNil(t, retrievedFile)
	assert.Equal(t, "file1", retrievedFile.Name)
}

func TestGetFiles_SortyByName(t *testing.T) {
	repo := NewMemoryFileRepo()

	now := time.Now()
	file1 := &vfs.File{
		UserName:   "user1",
		FolderName: "folder1",
		Name:       "file1",
		CreatedAt:  now.Add(-2 * time.Hour),
	}
	file2 := &vfs.File{
		UserName:   "user1",
		FolderName: "folder1",
		Name:       "file2",
		CreatedAt:  now.Add(-1 * time.Hour),
	}
	file3 := &vfs.File{
		UserName:   "user2",
		FolderName: "folder2",
		Name:       "file3",
		CreatedAt:  now,
	}
	_ = repo.AddFile(file1)
	_ = repo.AddFile(file2)
	_ = repo.AddFile(file3)

	req := &vfs.GetFilesRequest{
		UserName:   "user1",
		FolderName: "folder1",
		SortBy:     vfs.Name,
		OrderBy:    vfs.Asc,
	}

	files := repo.GetFiles(req)
	assert.Len(t, files, 2)
	assert.Equal(t, "file1", files[0].Name)
	assert.Equal(t, "file2", files[1].Name)
}

func TestGetFiles_SortyByCreatedTime(t *testing.T) {
	repo := NewMemoryFileRepo()

	now := time.Now()
	file1 := &vfs.File{
		UserName:   "user1",
		FolderName: "folder1",
		Name:       "file1",
		CreatedAt:  now.Add(-2 * time.Hour),
	}
	file2 := &vfs.File{
		UserName:   "user1",
		FolderName: "folder1",
		Name:       "file2",
		CreatedAt:  now.Add(-1 * time.Hour),
	}
	file3 := &vfs.File{
		UserName:   "user2",
		FolderName: "folder2",
		Name:       "file3",
		CreatedAt:  now,
	}
	_ = repo.AddFile(file1)
	_ = repo.AddFile(file2)
	_ = repo.AddFile(file3)

	req := &vfs.GetFilesRequest{
		UserName:   "user1",
		FolderName: "folder1",
		SortBy:     vfs.Created,
		OrderBy:    vfs.Asc,
	}

	files := repo.GetFiles(req)
	assert.Len(t, files, 2)
	assert.Equal(t, "file1", files[0].Name)
	assert.Equal(t, "file2", files[1].Name)
}

func TestGetFiles_SortyByCreatedTimeWithDescOrder(t *testing.T) {
	repo := NewMemoryFileRepo()

	now := time.Now()
	file1 := &vfs.File{
		UserName:   "user1",
		FolderName: "folder1",
		Name:       "file1",
		CreatedAt:  now.Add(-2 * time.Hour),
	}
	file2 := &vfs.File{
		UserName:   "user1",
		FolderName: "folder1",
		Name:       "file2",
		CreatedAt:  now.Add(-1 * time.Hour),
	}
	file3 := &vfs.File{
		UserName:   "user2",
		FolderName: "folder2",
		Name:       "file3",
		CreatedAt:  now,
	}
	_ = repo.AddFile(file1)
	_ = repo.AddFile(file2)
	_ = repo.AddFile(file3)

	req := &vfs.GetFilesRequest{
		UserName:   "user1",
		FolderName: "folder1",
		SortBy:     vfs.Created,
		OrderBy:    vfs.Desc,
	}

	files := repo.GetFiles(req)
	assert.Len(t, files, 2)
	assert.Equal(t, "file2", files[0].Name)
	assert.Equal(t, "file1", files[1].Name)
}

func TestDeleteFile_Success(t *testing.T) {
	repo := NewMemoryFileRepo()

	file := &vfs.File{
		UserName:   "user1",
		FolderName: "folder1",
		Name:       "file1",
	}
	_ = repo.AddFile(file)

	err := repo.DeleteFile(vfs.FileKeySet{UserName: "user1", FolderName: "folder1", FileName: "file1"})
	assert.Nil(t, err)
	assert.Nil(t, repo.GetFile(vfs.FileKeySet{UserName: "user1", FolderName: "folder1", FileName: "file1"}))
}

func TestDeleteFile_NotFound(t *testing.T) {
	repo := NewMemoryFileRepo()

	err := repo.DeleteFile(vfs.FileKeySet{UserName: "user1", FolderName: "folder1", FileName: "nonexistent"})
	assert.NotNil(t, err)
	assert.Equal(t, "The nonexistent doesn't exist.", err.Error())
}
