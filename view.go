package main

import (
	"fmt"

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

func finishedTestView(m Model) string {
	if m.result == nil {
		panic("Shouldn't be trying to render result screen without a completed test")
	}

	labelStyle := lipgloss.NewStyle().
		Background(colours.Bg).
		Foreground(colours.FgSubtle).
		PaddingRight(1)

	statStyle := lipgloss.NewStyle().
		Background(colours.Bg).
		Foreground(colours.Primary)

	wpmLabel := labelStyle.Render("wpm")
	wpmStat := statStyle.Render(fmt.Sprint(m.result.Wpm))
	wpm := lipgloss.JoinHorizontal(lipgloss.Top, wpmLabel, wpmStat)

	timeTakenLabel := labelStyle.Render("time")
	timeTakenStat := statStyle.Render(fmt.Sprintf("%.2fs", m.result.TimeTaken.Seconds()))
	timeTaken := lipgloss.JoinHorizontal(lipgloss.Top, timeTakenLabel, timeTakenStat)

	testStats := lipgloss.JoinHorizontal(lipgloss.Top, wpm, timeTaken)

	testCompleteStyle := lipgloss.NewStyle().
		Background(colours.Bg).
		Foreground(colours.Fg).
		Align(lipgloss.Center, lipgloss.Center)
	testComplete := testCompleteStyle.Render("Test complete! 🎉")

	view := lipgloss.JoinVertical(lipgloss.Center, testStats, testComplete, stats.RenderGraph(m.result))

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, view, lipgloss.WithWhitespaceBackground(colours.Bg))
}

func (m Model) View() string {
	switch m.state {
	case Ready:
		return readyToStartView(m)
	case InTest:
		return m.test.View(m.width, m.height)
	case TestComplete:
		return finishedTestView(m)
	}

	return ""
}
