package ui

import "github.com/charmbracelet/lipgloss"

const panelWidth = 42

var (
	panelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Theme.PanelBorder).
			Padding(0, 1)

	activePanelStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(Theme.ActivePanelBorder).
				Padding(0, 1)

	titleStyle = lipgloss.NewStyle().
			Foreground(Theme.Title).
			Bold(true)

	directoryStyle = lipgloss.NewStyle().
			Foreground(Theme.Directory)

	fileStyle = lipgloss.NewStyle().
			Foreground(Theme.File)

	selectedStyle = lipgloss.NewStyle().
			Background(Theme.Selected).
			Foreground(lipgloss.Color("255")).
			Bold(true)

	menuStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("255")).
			Background(Theme.ActivePanelBorder).
			Bold(true).
			Padding(0, 1)

	footerStyle = lipgloss.NewStyle().
			Foreground(Theme.Title)

	entryStyle = lipgloss.NewStyle().
			Width(panelWidth - 4)

	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(Theme.Title)

	separatorStyle = lipgloss.NewStyle().
			Foreground(Theme.PanelBorder)
)
