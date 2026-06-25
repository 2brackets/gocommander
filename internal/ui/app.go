package ui

import tea "github.com/charmbracelet/bubbletea"

func Run() error {
	program := tea.NewProgram(newModel())
	_, err := program.Run()
	return err
}
