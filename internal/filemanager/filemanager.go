package filemanager

import "os"

type FileEntry struct {
	Name  string
	IsDir bool
}

type FileManager struct {
	currentPath string
}

func New(path string) *FileManager {
	return &FileManager{
		currentPath: path,
	}
}

func (fm *FileManager) ListDirectory() ([]FileEntry, error) {
	dirEntries, err := os.ReadDir(fm.currentPath)
	if err != nil {
		return nil, err
	}
	entries := make([]FileEntry, 0, len(dirEntries))

	for _, dirEntry := range dirEntries {
		entry := FileEntry{
			Name:  dirEntry.Name(),
			IsDir: dirEntry.IsDir(),
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
