package ui

import (
	"fmt"

	"github.com/2brackets/gocommander/internal/filemanager"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	leftPanel   panel
	rightPanel  panel
	activePanel *panel
	err         error
}

func newModel() *model {
	leftFM := filemanager.New(".")
	rightFM := filemanager.New(".")

	m := &model{
		leftPanel: panel{
			fileManager: leftFM,
		},
		rightPanel: panel{
			fileManager: rightFM,
		},
	}
	m.activePanel = &m.leftPanel
	return m
}

func (m *model) Init() tea.Cmd {
	leftEntries, err := m.leftPanel.fileManager.Read()
	if err != nil {
		m.err = err
		return nil
	}
	rightEntries, err := m.rightPanel.fileManager.Read()
	if err != nil {
		m.err = err
		return nil
	}
	m.leftPanel.entries = leftEntries
	m.rightPanel.entries = rightEntries
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up":
			m.moveUp()

		case "down":
			m.moveDown()

		case "enter":
			m.enterDirectory()

		case "backspace":
			m.parentDirectory()

		case "tab":
			m.switchPanel()
		}
	}
	return m, nil
}

func (m *model) errorView() string {
	if m.err == nil {
		return ""
	}

	return fmt.Sprintf("Error: %v\n", m.err)
}
