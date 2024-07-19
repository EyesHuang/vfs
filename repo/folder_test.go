package repo

import (
	"testing"
	"time"

	"vfs"

	"github.com/stretchr/testify/assert"
)

func TestAddFolder_Success(t *testing.T) {
	repo := NewMemoFolderRepo()

	folder := &vfs.Folder{
		UserName: "user1",
		Name:     "folder1",
	}

	err := repo.AddFolder(folder)
	assert.Nil(t, err)
	assert.NotNil(t, repo.GetFolder(vfs.FolderKeySet{UserName: "user1", FolderName: "folder1"}))
}

func TestAddFolder_ExistingFolder(t *testing.T) {
	repo := NewMemoFolderRepo()

	folder := &vfs.Folder{
		UserName: "user1",
		Name:     "folder1",
	}
	_ = repo.AddFolder(folder)

	err := repo.AddFolder(folder)
	assert.NotNil(t, err)
	assert.Equal(t, "The folder1 has already existed.", err.Error())
}

func TestGetFolder(t *testing.T) {
	repo := NewMemoFolderRepo()

	folder := &vfs.Folder{
		UserName: "user1",
		Name:     "folder1",
	}
	_ = repo.AddFolder(folder)

	retrievedFolder := repo.GetFolder(vfs.FolderKeySet{UserName: "user1", FolderName: "folder1"})
	assert.NotNil(t, retrievedFolder)
	assert.Equal(t, "folder1", retrievedFolder.Name)

	retrievedFolder = repo.GetFolder(vfs.FolderKeySet{UserName: "user1", FolderName: "nonexistent"})
	assert.Nil(t, retrievedFolder)
}

func TestUpdateFolder_Success(t *testing.T) {
	repo := NewMemoFolderRepo()

	folder := &vfs.Folder{
		UserName: "user1",
		Name:     "folder1",
	}
	_ = repo.AddFolder(folder)

	req := &vfs.UpdateFolderRequest{
		UserName: "user1",
		OldName:  "folder1",
		NewName:  "folder2",
	}

	err := repo.UpdateFolder(req)
	assert.Nil(t, err)
	assert.Nil(t, repo.GetFolder(vfs.FolderKeySet{UserName: "user1", FolderName: "folder1"}))
	assert.NotNil(t, repo.GetFolder(vfs.FolderKeySet{UserName: "user1", FolderName: "folder2"}))
}

func TestUpdateFolder_OldFolderNotFound(t *testing.T) {
	repo := NewMemoFolderRepo()

	req := &vfs.UpdateFolderRequest{
		UserName: "user1",
		OldName:  "nonexistent",
		NewName:  "folder2",
	}

	err := repo.UpdateFolder(req)
	assert.NotNil(t, err)
	assert.Equal(t, "The folder nonexistent doesn't exist.", err.Error())
}

func TestUpdateFolder_NewFolderExists(t *testing.T) {
	repo := NewMemoFolderRepo()

	folder1 := &vfs.Folder{
		UserName: "user1",
		Name:     "folder1",
	}
	folder2 := &vfs.Folder{
		UserName: "user1",
		Name:     "folder2",
	}
	_ = repo.AddFolder(folder1)
	_ = repo.AddFolder(folder2)

	req := &vfs.UpdateFolderRequest{
		UserName: "user1",
		OldName:  "folder1",
		NewName:  "folder2",
	}

	err := repo.UpdateFolder(req)
	assert.NotNil(t, err)
	assert.Equal(t, "The folder folder2 already exists.", err.Error())
}

func TestDeleteFolder_Success(t *testing.T) {
	repo := NewMemoFolderRepo()

	folder := &vfs.Folder{
		UserName: "user1",
		Name:     "folder1",
	}
	_ = repo.AddFolder(folder)

	err := repo.DeleteFolder(vfs.FolderKeySet{UserName: "user1", FolderName: "folder1"})
	assert.Nil(t, err)
	assert.Nil(t, repo.GetFolder(vfs.FolderKeySet{UserName: "user1", FolderName: "folder1"}))
}

func TestDeleteFolder_NotFound(t *testing.T) {
	repo := NewMemoFolderRepo()

	err := repo.DeleteFolder(vfs.FolderKeySet{UserName: "user1", FolderName: "nonexistent"})
	assert.NotNil(t, err)
	assert.Equal(t, "The nonexistent doesn't exist.", err.Error())
}

func TestGetFolders_SortyByFolderName(t *testing.T) {
	repo := NewMemoFolderRepo()

	now := time.Now()
	folder1 := &vfs.Folder{
		UserName:  "user1",
		Name:      "folder1",
		CreatedAt: now.Add(-2 * time.Hour),
	}
	folder2 := &vfs.Folder{
		UserName:  "user1",
		Name:      "folder2",
		CreatedAt: now.Add(-1 * time.Hour),
	}
	folder3 := &vfs.Folder{
		UserName:  "user2",
		Name:      "folder3",
		CreatedAt: now,
	}
	_ = repo.AddFolder(folder1)
	_ = repo.AddFolder(folder2)
	_ = repo.AddFolder(folder3)

	req := &vfs.GetFoldersRequest{
		UserName: "user1",
		SortBy:   vfs.Name,
		OrderBy:  vfs.Asc,
	}

	folders := repo.GetFolders(req)
	assert.Len(t, folders, 2)
	assert.Equal(t, "folder1", folders[0].Name)
	assert.Equal(t, "folder2", folders[1].Name)
}

func TestGetFolders_SortyByCreatedTime(t *testing.T) {
	repo := NewMemoFolderRepo()

	now := time.Now()
	folder1 := &vfs.Folder{
		UserName:  "user1",
		Name:      "folder1",
		CreatedAt: now.Add(-2 * time.Hour),
	}
	folder2 := &vfs.Folder{
		UserName:  "user1",
		Name:      "folder2",
		CreatedAt: now.Add(-1 * time.Hour),
	}
	folder3 := &vfs.Folder{
		UserName:  "user2",
		Name:      "folder3",
		CreatedAt: now,
	}
	_ = repo.AddFolder(folder1)
	_ = repo.AddFolder(folder2)
	_ = repo.AddFolder(folder3)

	req := &vfs.GetFoldersRequest{
		UserName: "user1",
		SortBy:   vfs.Created,
		OrderBy:  vfs.Asc,
	}

	folders := repo.GetFolders(req)
	assert.Len(t, folders, 2)
	assert.Equal(t, "folder1", folders[0].Name)
	assert.Equal(t, "folder2", folders[1].Name)
}

func TestGetFolders_SortyByCreatedTimeWithDescOrder(t *testing.T) {
	repo := NewMemoFolderRepo()

	now := time.Now()
	folder1 := &vfs.Folder{
		UserName:  "user1",
		Name:      "folder1",
		CreatedAt: now.Add(-2 * time.Hour),
	}
	folder2 := &vfs.Folder{
		UserName:  "user1",
		Name:      "folder2",
		CreatedAt: now.Add(-1 * time.Hour),
	}
	folder3 := &vfs.Folder{
		UserName:  "user2",
		Name:      "folder3",
		CreatedAt: now,
	}
	_ = repo.AddFolder(folder1)
	_ = repo.AddFolder(folder2)
	_ = repo.AddFolder(folder3)

	req := &vfs.GetFoldersRequest{
		UserName: "user1",
		SortBy:   vfs.Created,
		OrderBy:  vfs.Desc,
	}

	folders := repo.GetFolders(req)
	assert.Len(t, folders, 2)
	assert.Equal(t, "folder2", folders[0].Name)
	assert.Equal(t, "folder1", folders[1].Name)
}
