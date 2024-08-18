package menu

import (
	"github.com/SamYouatt/typee/util"
	"github.com/charmbracelet/lipgloss"
)

var colours = util.AppColours

func View(width, height int) string {
	headerTextStyle := lipgloss.NewStyle().
		Width(width).
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
		Width(width).
		PaddingTop(1).
		Background(colours.Bg).
		Foreground(colours.Fg).
		Align(lipgloss.Center, lipgloss.Center)
	instructionText := instructionTextStyle.Render("Press Enter to start")

	joined := lipgloss.JoinVertical(lipgloss.Center, headerText, instructionText)

	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, joined, lipgloss.WithWhitespaceBackground(colours.Bg))
}
