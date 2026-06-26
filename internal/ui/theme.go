package ui

import "github.com/charmbracelet/lipgloss"

type theme struct {
	PanelBorder       lipgloss.Color
	ActivePanelBorder lipgloss.Color
	Directory         lipgloss.Color
	File              lipgloss.Color
	Selected          lipgloss.Color
	Title             lipgloss.Color
}

var Theme = theme{
	PanelBorder:       lipgloss.Color("240"),
	ActivePanelBorder: lipgloss.Color("39"),
	Directory:         lipgloss.Color("45"),
	File:              lipgloss.Color("252"),
	Selected:          lipgloss.Color("238"),
	Title:             lipgloss.Color("81"),
}
