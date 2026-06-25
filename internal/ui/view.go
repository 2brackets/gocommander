package ui

import "strings"

func (m *model) View() string {
	var builder strings.Builder

	builder.WriteString("GoCommander\n\n")

	leftTitle := "Left Panel"
	rightTitle := "Right Panel"

	if m.activePanel == &m.leftPanel {
		leftTitle = "> " + leftTitle
	} else {
		leftTitle = "  " + leftTitle
	}

	if m.activePanel == &m.rightPanel {
		rightTitle = "> " + rightTitle
	} else {
		rightTitle = "  " + rightTitle
	}

	builder.WriteString(leftTitle)
	builder.WriteString("                    ")
	builder.WriteString(rightTitle)
	builder.WriteString("\n")

	builder.WriteString("-----------------------------------------\n")

	maxRows := len(m.leftPanel.entries)
	if len(m.rightPanel.entries) > maxRows {
		maxRows = len(m.rightPanel.entries)
	}

	for i := 0; i < maxRows; i++ {
		if i < len(m.leftPanel.entries) {
			writeEntry(&builder, &m.leftPanel, i)
		} else {
			builder.WriteString(strings.Repeat(" ", 24))
		}

		builder.WriteString("   ")

		if i < len(m.rightPanel.entries) {
			writeEntry(&builder, &m.rightPanel, i)
		}

		builder.WriteString("\n")
	}

	builder.WriteString("\nF3 View   F4 Edit   F5 Copy   F6 Move   F7 MkDir   F8 Delete\n")
	builder.WriteString("Tab switch panel   q quit\n")

	return builder.String()
}

func writeEntry(builder *strings.Builder, p *panel, index int) {
	entry := p.entries[index]

	if index == p.selectedIndex {
		builder.WriteString("> ")
	} else {
		builder.WriteString("  ")
	}

	if entry.IsDir {
		builder.WriteString("[DIR] ")
	} else {
		builder.WriteString("      ")
	}

	name := entry.Name
	if len(name) > 16 {
		name = name[:16]
	}

	builder.WriteString(name)

	padding := 24 - len(name)
	if padding > 0 {
		builder.WriteString(strings.Repeat(" ", padding))
	}
}
