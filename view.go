package main

import (
	"github.com/SamYouatt/typee/features/stats"
	"github.com/SamYouatt/typee/util"
	"github.com/charmbracelet/lipgloss"
)

var colours = util.AppColours

func readyToStartView(m Model) string {
	headerTextStyle := lipgloss.NewStyle().
		Width(m.width).
		Background(colours.Bg).
		Foreground(colours.Primary).
		Align(lipgloss.Center, lipgloss.Center)
	headerTextRaw := `
░▒▓████████▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓███████▓▒░░▒▓████████▓▒░▒▓████████▓▒░ 
   ░▒▓█▓▒░   ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░        
   ░▒▓█▓▒░    ░▒▓██████▓▒░░▒▓███████▓▒░░▒▓██████▓▒░ ░▒▓██████▓▒░   
   ░▒▓█▓▒░      ░▒▓█▓▒░   ░▒▓█▓▒░      ░▒▓█▓▒░      ░▒▓█▓▒░        
   ░▒▓█▓▒░      ░▒▓█▓▒░   ░▒▓█▓▒░      ░▒▓████████▓▒░▒▓████████▓▒░ 
`
	headerText := headerTextStyle.Render(headerTextRaw)

	instructionTextStyle := lipgloss.NewStyle().
		Width(m.width).
		PaddingTop(1).
		Background(colours.Bg).
		Foreground(colours.Fg).
		Align(lipgloss.Center, lipgloss.Center)
	instructionText := instructionTextStyle.Render("Press Enter to start")

	joined := lipgloss.JoinVertical(lipgloss.Center, headerText, instructionText)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, joined, lipgloss.WithWhitespaceBackground(colours.Bg))
}

func (m Model) View() string {
	switch m.state {
	case Ready:
		return readyToStartView(m)
	case InTest:
		return m.test.View(m.width, m.height)
	case TestComplete:
		return stats.View(m.result, m.width, m.height)
	}

	return ""
}
