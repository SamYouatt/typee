package main

import (
	"fmt"
	"math/rand"
	"slices"

	"github.com/NimbleMarkets/ntcharts/canvas"
	"github.com/NimbleMarkets/ntcharts/linechart"
	"github.com/charmbracelet/lipgloss"
)

func renderGraph(m Model) string {
	axisStyle := lipgloss.NewStyle().Foreground(colours.Fg)
	lineStyle := lipgloss.NewStyle().Foreground(colours.Primary)
	backgroundStyle := lipgloss.NewStyle().Foreground(colours.Fg)

	// seed some random values
	yCoords := make([]int, 50)
	for i := 0; i < len(yCoords); i++ {
		yCoords[i] = rand.Intn(30) + 90
	}

	// set up the chart
	chart := linechart.New(80, 10,
		0, float64(len(yCoords)),
		0, float64(slices.Max(yCoords)),
		linechart.WithXYSteps(1, 1),
		linechart.WithStyles(axisStyle, axisStyle, backgroundStyle))
	chart.AxisStyle = axisStyle
	chart.LabelStyle = axisStyle
	chart.DrawXYAxisAndLabel()

	// draw the points to the chart
	for i := 0; i < len(yCoords)-1; i++ {
		y1 := yCoords[i]
		y2 := yCoords[i+1]
		coord1 := canvas.Float64Point{X: float64(i), Y: float64(y1)}
		coord2 := canvas.Float64Point{X: float64(i + 1), Y: float64(y2)}
		chart.DrawBrailleLineWithStyle(coord1, coord2, lineStyle)
	}

	// render the chart
	titleStyle := lipgloss.NewStyle().Foreground(colours.Primary).PaddingBottom(1)
	axisLabelStyle := lipgloss.NewStyle().Foreground(colours.FgSubtle).PaddingLeft(3)

	return lipgloss.JoinVertical(lipgloss.Left, titleStyle.Render("wpm"), chart.View(), axisLabelStyle.Render("words"))
}

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

	labelStyle := lipgloss.NewStyle().
		Background(colours.Bg).
		Foreground(colours.FgSubtle).
		PaddingRight(1)

	statStyle := lipgloss.NewStyle().
		Background(colours.Bg).
		Foreground(colours.Primary)

	wpmLabel := labelStyle.Render("wpm")
	wpmStat := statStyle.Render(fmt.Sprint(m.completedTest.wpm))
	wpm := lipgloss.JoinHorizontal(lipgloss.Top, wpmLabel, wpmStat)

	timeTakenLabel := labelStyle.Render("time")
	timeTakenStat := statStyle.Render(fmt.Sprintf("%.2fs", m.completedTest.timeTaken.Seconds()))
	timeTaken := lipgloss.JoinHorizontal(lipgloss.Top, timeTakenLabel, timeTakenStat)

	stats := lipgloss.JoinHorizontal(lipgloss.Top, wpm, timeTaken)

	testCompleteStyle := lipgloss.NewStyle().
		Background(colours.Bg).
		Foreground(colours.Fg).
		Align(lipgloss.Center, lipgloss.Center)
	testComplete := testCompleteStyle.Render("Test complete! 🎉")

	view := lipgloss.JoinVertical(lipgloss.Center, stats, testComplete, renderGraph(m))

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, view, lipgloss.WithWhitespaceBackground(colours.Bg))
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
