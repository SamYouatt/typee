package main

import (
	"fmt"
	"slices"

	"github.com/charmbracelet/lipgloss"
)

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
	if m.completedTest == nil {
		panic("Shouldn't be trying to render result screen without a completed test")
	}

	testCompleteStyle := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Background(colours.Bg).
		Foreground(colours.Fg).
		Align(lipgloss.Center, lipgloss.Center)

	return testCompleteStyle.Render("Test complete! 🎉")
}

func (test Test) renderTest(width int) string {
	untypedStyle := lipgloss.NewStyle().
		Foreground(colours.FgSubtle).
		Background(colours.Bg)
	correctlyTypedStyle := lipgloss.NewStyle().
		Foreground(colours.Fg).
		Background(colours.Bg)
	previousErrorStyle := lipgloss.NewStyle().
		Foreground(colours.Error).
		Background(colours.Bg)
	currentPendingStyle := lipgloss.NewStyle().
		Foreground(colours.FgSubtle).
		Background(colours.BgSubtle)
	currentlyInvalidStyle := lipgloss.NewStyle().
		Foreground(colours.Error).
		Background(colours.BgSubtle)

	testLetters := []string{}

	for index, letter := range test.text {
		// Current character is wrong
		if test.currentIndex == index && test.currentlyInvalid {
			testLetters = append(testLetters, currentlyInvalidStyle.Render(string(letter)))
			continue
		}

		// Current character pending
		if test.currentIndex == index {
			testLetters = append(testLetters, currentPendingStyle.Render(string(letter)))
			continue
		}

		// Previously entered incorrectly
		if slices.Contains(test.errorIndices, index) {
			testLetters = append(testLetters, previousErrorStyle.Render(string(letter)))
			continue
		}

		// Previously entered correctly
		if test.currentIndex > index {
			testLetters = append(testLetters, correctlyTypedStyle.Render(string(letter)))
			continue
		}

		// Upcoming letters
		testLetters = append(testLetters, untypedStyle.Render(string(letter)))
	}

	testRendered := lipgloss.JoinHorizontal(lipgloss.Center, testLetters...)

	viewStyle := lipgloss.NewStyle().
		Width(width).
		AlignHorizontal(lipgloss.Center).
		Background(colours.Bg)

	return viewStyle.Render(testRendered)
}

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
		return readyToStartView(m)
	case InTest:
		return testView(m)
	case TestComplete:
		return finishedTestView(m)
	}

	return ""
}
