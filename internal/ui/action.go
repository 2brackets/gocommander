package ui

func (m *model) moveUp() {
	if m.activePanel.selectedIndex > 0 {
		m.activePanel.selectedIndex--
	}
}

func (m *model) moveDown() {
	if m.activePanel.selectedIndex < len(m.activePanel.entries)-1 {
		m.activePanel.selectedIndex++
	}
}

func (m *model) switchPanel() {
	if m.activePanel == &m.leftPanel {
		m.activePanel = &m.rightPanel
	} else {
		m.activePanel = &m.leftPanel
	}
}

func (m *model) enterDirectory() {
	if err := m.activePanel.enterDirectory(); err != nil {
		m.err = err
	}
}

func (m *model) parentDirectory() {
	if err := m.activePanel.parentDirectory(); err != nil {
		m.err = err
	}
}
