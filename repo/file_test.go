package repo

import (
	"testing"

	"vfs/mock"

	"vfs"

	"github.com/stretchr/testify/assert"
)

func TestGetFile(t *testing.T) {
	repo := NewMemoryFileRepo()

	_ = repo.AddFile(mock.File1)

	retrievedFile := repo.GetFile(mock.FileKeySet)
	assert.NotNil(t, retrievedFile)
	assert.Equal(t, "file1", retrievedFile.Name)
}

func TestGetFiles_SortyByName(t *testing.T) {
	repo := NewMemoryFileRepo()

	_ = repo.AddFile(mock.File1)
	_ = repo.AddFile(mock.File2)
	_ = repo.AddFile(mock.File3)

	files := repo.GetFiles(mock.GetFilesRequest)
	assert.Len(t, files, 2)
	assert.Equal(t, "file1", files[0].Name)
	assert.Equal(t, "file2", files[1].Name)
}

func TestGetFiles_SortyByCreatedTime(t *testing.T) {
	repo := NewMemoryFileRepo()

	_ = repo.AddFile(mock.File1)
	_ = repo.AddFile(mock.File2)
	_ = repo.AddFile(mock.File3)

	mock.GetFilesRequest.SortBy = vfs.Created

	files := repo.GetFiles(mock.GetFilesRequest)
	assert.Len(t, files, 2)
	assert.Equal(t, "file1", files[0].Name)
	assert.Equal(t, "file2", files[1].Name)
}

func TestGetFiles_SortyByCreatedTimeWithDescOrder(t *testing.T) {
	repo := NewMemoryFileRepo()

	_ = repo.AddFile(mock.File1)
	_ = repo.AddFile(mock.File2)
	_ = repo.AddFile(mock.File3)

	mock.GetFilesRequest.SortBy = vfs.Created
	mock.GetFilesRequest.OrderBy = vfs.Desc

	files := repo.GetFiles(mock.GetFilesRequest)
	assert.Len(t, files, 2)
	assert.Equal(t, "file2", files[0].Name)
	assert.Equal(t, "file1", files[1].Name)
}

func TestAddFile_Success(t *testing.T) {
	repo := NewMemoryFileRepo()

	err := repo.AddFile(mock.File1)
	assert.Nil(t, err)
	assert.NotNil(t, repo.GetFile(mock.FileKeySet))
}

func TestAddFile_ExistingFile(t *testing.T) {
	repo := NewMemoryFileRepo()

	_ = repo.AddFile(mock.File1)

	err := repo.AddFile(mock.File1)
	assert.NotNil(t, err)
	assert.Equal(t, "The file1 has already existed.", err.Error())
}

func TestDeleteFile_Success(t *testing.T) {
	repo := NewMemoryFileRepo()

	_ = repo.AddFile(mock.File1)

	err := repo.DeleteFile(mock.FileKeySet)
	assert.Nil(t, err)
	assert.Nil(t, repo.GetFile(mock.FileKeySet))
}

func TestDeleteFile_NotFound(t *testing.T) {
	repo := NewMemoryFileRepo()

	mock.FileKeySet.FileName = "nonexistent"

	err := repo.DeleteFile(mock.FileKeySet)
	assert.NotNil(t, err)
	assert.Equal(t, "The nonexistent doesn't exist.", err.Error())
}
