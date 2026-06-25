package ui

import "github.com/2brackets/gocommander/internal/filemanager"

type panel struct {
	fileManager   *filemanager.FileManager
	entries       []filemanager.FileEntry
	selectedIndex int
}

func (p *panel) selectedEntry() (filemanager.FileEntry, bool) {
	if len(p.entries) == 0 {
		return filemanager.FileEntry{}, false
	}

	return p.entries[p.selectedIndex], true
}

func (p *panel) refresh() error {
	entries, err := p.fileManager.Read()
	if err != nil {
		return err
	}

	p.entries = entries
	p.selectedIndex = 0
	return nil
}

func (p *panel) enterDirectory() error {
	entry, ok := p.selectedEntry()
	if !ok {
		return nil
	}
	if entry.IsParent {
		if err := p.fileManager.ParentDirectory(); err != nil {
			return err
		}
	}
	if !entry.IsDir {
		return nil
	}
	if err := p.fileManager.ChangeDirectory(entry.Path); err != nil {
		return err
	}
	return p.refresh()
}

func (p *panel) parentDirectory() error {
	if err := p.fileManager.ParentDirectory(); err != nil {
		return err
	}
	return p.refresh()
}

func (p *panel) Path() string {
	return p.fileManager.Path()
}
