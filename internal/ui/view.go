package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const panelWidth = 38

var (
	panelStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			Padding(0, 1).
			Width(panelWidth)

	activePanelStyle = lipgloss.NewStyle().
				Border(lipgloss.ThickBorder()).
				Padding(0, 1).
				Width(panelWidth)
)

func (m *model) View() string {
	if m.err != nil {
		return m.errorView()
	}

	left := renderPanel(m.leftPanel.Path(), &m.leftPanel, m.activePanel == &m.leftPanel)
	right := renderPanel(m.rightPanel.Path(), &m.rightPanel, m.activePanel == &m.rightPanel)

	panels := lipgloss.JoinHorizontal(lipgloss.Top, left, right)

	footer := "\nF3 View  F4 Edit  F5 Copy  F6 Move  F7 MkDir  F8 Delete\nTab switch panel   q quit"

	return "GoCommander\n\n" + panels + footer
}

func renderPanel(title string, p *panel, active bool) string {
	var builder strings.Builder

	builder.WriteString(title)
	builder.WriteString("\n\n")

	for i, entry := range p.entries {
		if i == p.selectedIndex {
			builder.WriteString("> ")
		} else {
			builder.WriteString("  ")
		}

		if entry.IsDir {
			builder.WriteString("[DIR] ")
		} else {
			builder.WriteString("      ")
		}

		builder.WriteString(entry.Name)
		builder.WriteString("\n")
	}

	if active {
		return activePanelStyle.Render(builder.String())
	}

	return panelStyle.Render(builder.String())
}
