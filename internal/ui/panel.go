package ui

import "github.com/2brackets/gocommander/internal/filemanager"

type panel struct {
	fileManager   *filemanager.FileManager
	entries       []filemanager.FileEntry
	selectedIndex int
}
