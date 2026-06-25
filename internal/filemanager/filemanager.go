package filemanager

import (
	"os"
	"path/filepath"
)

type FileEntry struct {
	Name     string
	Path     string
	IsDir    bool
	IsParent bool
}

type FileManager struct {
	currentPath string
}

func New(path string) *FileManager {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		absolutePath = path
	}
	return &FileManager{
		currentPath: absolutePath,
	}
}

func (fm *FileManager) Read() ([]FileEntry, error) {
	dirEntries, err := os.ReadDir(fm.currentPath)
	if err != nil {
		return nil, err
	}

	entries := make([]FileEntry, 0, len(dirEntries)+1)

	entries = append(entries, FileEntry{
		Name:     "..",
		Path:     filepath.Dir(fm.currentPath),
		IsDir:    true,
		IsParent: true,
	})

	for _, dirEntry := range dirEntries {
		entries = append(entries, FileEntry{
			Name:  dirEntry.Name(),
			Path:  filepath.Join(fm.currentPath, dirEntry.Name()),
			IsDir: dirEntry.IsDir(),
		})
	}

	return entries, nil
}

func (fm *FileManager) ChangeDirectory(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return nil
	}
	fm.currentPath = path
	return nil
}

func (fm *FileManager) ParentDirectory() error {
	parent := filepath.Dir(fm.currentPath)
	return fm.ChangeDirectory(parent)
}

func (fm *FileManager) Path() string {
	return fm.currentPath
}
