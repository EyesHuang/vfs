package repo

import (
	"errors"
	"fmt"
	"sort"

	"vfs"
)

type MemoryFileRepo struct {
	Files map[vfs.FileKeySet]*vfs.File
}

func NewMemoryFileRepo() *MemoryFileRepo {
	return &MemoryFileRepo{
		Files: make(map[vfs.FileKeySet]*vfs.File),
	}
}

func (mfr *MemoryFileRepo) GetFile(key vfs.FileKeySet) *vfs.File {
	if file, exists := mfr.Files[key]; exists {
		return file
	}
	return nil
}

func (mfr *MemoryFileRepo) GetFiles(req *vfs.GetFilesRequest) []*vfs.File {
	var files []*vfs.File
	for key, file := range mfr.Files {
		if key.UserName == req.UserName && key.FolderName == req.FolderName {
			files = append(files, file)
		}
	}

	// Sorting
	switch req.SortBy {
	case vfs.Name:
		if req.OrderBy == vfs.Asc {
			sort.Slice(files, func(i, j int) bool {
				return files[i].Name < files[j].Name
			})
		} else {
			sort.Slice(files, func(i, j int) bool {
				return files[i].Name > files[j].Name
			})
		}
	case vfs.Created:
		if req.OrderBy == vfs.Asc {
			sort.Slice(files, func(i, j int) bool {
				return files[i].CreatedAt.Before(files[j].CreatedAt)
			})
		} else {
			sort.Slice(files, func(i, j int) bool {
				return files[i].CreatedAt.After(files[j].CreatedAt)
			})
		}
	default:
		sort.Slice(files, func(i, j int) bool {
			return files[i].Name < files[j].Name
		})
	}

	return files
}

func (mfr *MemoryFileRepo) AddFile(file *vfs.File) error {
	key := vfs.FileKeySet{
		UserName:   file.UserName,
		FolderName: file.FolderName,
		FileName:   file.Name,
	}
	if mfr.GetFile(key) != nil {
		errMsg := fmt.Sprintf("The %s has already existed.", file.Name)
		return errors.New(errMsg)
	}
	mfr.Files[key] = file
	return nil
}

func (mfr *MemoryFileRepo) DeleteFile(key vfs.FileKeySet) error {
	if mfr.GetFile(key) == nil {
		errMsg := fmt.Sprintf("The %s doesn't exist.", key.FileName)
		return errors.New(errMsg)
	}
	delete(mfr.Files, key)
	return nil
}
