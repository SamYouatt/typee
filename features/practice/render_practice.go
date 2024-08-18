package practice

import (
	"fmt"
	"slices"

	"github.com/SamYouatt/typee/util"
	"github.com/charmbracelet/lipgloss"
)

var colours = util.AppColours

func (test Test) renderTestProgress() string {
	testViewWidth := min(80, len(test.text))
	testProgressStyle := lipgloss.NewStyle().
		Width(testViewWidth).
		Foreground(colours.Fg).
		Background(colours.Bg).
		AlignHorizontal(lipgloss.Left)

	return testProgressStyle.Render(fmt.Sprintf("%d/%d", test.completedWords, test.numWords))
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

func (test Test) View(width, height int) string {
    testProgress := test.renderTestProgress()
	testBody := test.renderTest(80)

	view := lipgloss.JoinVertical(lipgloss.Center, testProgress, testBody)
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, view, lipgloss.WithWhitespaceBackground(colours.Bg))
}
