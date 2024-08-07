package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func testView(m Model) string {
	testViewWidth := min(80, len(m.test.text))
	testProgressStyle := lipgloss.NewStyle().
		Width(testViewWidth).
		Foreground(colours.Fg).
		Background(colours.Bg).
		AlignHorizontal(lipgloss.Left)
	testProgress := testProgressStyle.Render(fmt.Sprintf("%d/%d", m.test.completedWords, m.test.numWords))

	view := lipgloss.JoinVertical(lipgloss.Center, testProgress, m.test.renderTest(testViewWidth))

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, view, lipgloss.WithWhitespaceBackground(colours.Bg))
}

func (m Model) View() string {
	switch m.state {
	case Ready:
		return "Press enter to start"
	case InTest:
		return testView(m)
	case TestComplete:
		return "not done yet"
	}

	return ""
}
