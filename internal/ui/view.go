package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/2brackets/gocommander/internal/filemanager"
	"github.com/charmbracelet/lipgloss"
)

func (m *model) View() string {
	if m.err != nil {
		return m.errorView()
	}

	width := m.width
	height := m.height

	if width == 0 {
		width = 120
	}
	if height == 0 {
		height = 30
	}

	menu := menuStyle.
		Width(width).
		Render("  Left   File   Command   Options   Help  ")

	footerText := "  F3 View   F4 Edit   F5 Copy   F6 Move   F7 MkDir   F8 Delete   F10 Quit  "
	footer := menuStyle.
		Width(width).
		Render(footerText)

	panelHeight := height - 10
	panelWidth := (width - 4) / 2

	left := renderPanel(&m.leftPanel, m.activePanel == &m.leftPanel, panelWidth, panelHeight)
	right := renderPanel(&m.rightPanel, m.activePanel == &m.rightPanel, panelWidth, panelHeight)

	panels := lipgloss.JoinHorizontal(lipgloss.Top, left, right)

	return menu + "\n" + panels + "\n" + footer
}

func renderPanel(p *panel, active bool, width int, height int) string {
	var builder strings.Builder
	tableWidth := NameColumnWidth + 1 + SizeColumnWidth + 1 + DateColumnWidth
	header := fmt.Sprintf(
		"%-*s %*s %-*s",
		NameColumnWidth, "Name",
		SizeColumnWidth, "Size",
		DateColumnWidth, "Modified",
	)

	builder.WriteString(headerStyle.Render(header))
	builder.WriteString("\n")

	builder.WriteString(separatorStyle.Render(strings.Repeat("─", tableWidth)))
	builder.WriteString("\n")

	for i, entry := range p.entries {
		line := renderEntry(entry, i == p.selectedIndex, width-4)

		builder.WriteString(line)
		builder.WriteString("\n")
	}

	style := panelStyle
	if active {
		style = activePanelStyle
	}

	return style.
		Width(width).
		Height(height).
		Render(builder.String())
}

func formatSize(size int64, isDir bool, isParent bool) string {
	if isParent {
		return ""
	}
	if isDir {
		return "<DIR>"
	}
	if size < 1024 {
		return fmt.Sprintf("%d B", size)
	}
	if size < 1024*1024 {
		return fmt.Sprintf("%.1f KB", float64(size)/1024)
	}
	return fmt.Sprintf("%.1f MB", float64(size)/(1024*1024))
}

func formatDate(t time.Time, isParent bool) string {
	if isParent || t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02")
}

func renderEntry(entry filemanager.FileEntry, selected bool, width int) string {
	icon := Icons.File

	if entry.IsParent {
		icon = Icons.Parent
	} else if entry.IsDir {
		icon = Icons.Folder
	}

	nameWidth := width - 22
	if nameWidth < 10 {
		nameWidth = 10
	}

	sizeWidth := 8
	dateWidth := 12

	name := icon + " " + truncate(entry.Name, nameWidth-2)
	size := formatSize(entry.Size, entry.IsDir, entry.IsParent)
	date := formatDate(entry.Modified, entry.IsParent)

	namePart := lipgloss.NewStyle().Width(nameWidth).Render(name)
	sizePart := lipgloss.NewStyle().Width(sizeWidth).Align(lipgloss.Right).Render(size)
	datePart := lipgloss.NewStyle().Width(dateWidth).Render(date)

	line := lipgloss.JoinHorizontal(lipgloss.Top, namePart, sizePart, datePart)

	if entry.IsDir {
		line = directoryStyle.Render(line)
	} else {
		line = fileStyle.Render(line)
	}

	if selected {
		line = selectedStyle.Width(width).Render(line)
	}

	return line
}

func truncate(value string, max int) string {
	if len(value) <= max {
		return value
	}

	if max <= 3 {
		return value[:max]
	}

	return value[:max-3] + "..."
}
