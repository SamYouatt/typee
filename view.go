package main

import (
	"fmt"
	"slices"

	"github.com/NimbleMarkets/ntcharts/canvas"
	"github.com/NimbleMarkets/ntcharts/linechart"
	"github.com/SamYouatt/typee/features/practice"
	"github.com/charmbracelet/lipgloss"
)

func renderGraph(m Model) string {
	axisStyle := lipgloss.NewStyle().Foreground(colours.Fg)
	lineStyle := lipgloss.NewStyle().Foreground(colours.Primary)
	backgroundStyle := lipgloss.NewStyle().Foreground(colours.Fg)

	width, height := 80, 10
	yCoords := m.completedTest.runningWpm
	minX, maxX := 0.0, float64(len(yCoords))
	minY, maxY := 0.0, float64(slices.Max(yCoords))

	chart := linechart.New(width, height,
		minX, maxX,
		minY, maxY,
		linechart.WithXYSteps(1, 1),
		linechart.WithStyles(axisStyle, axisStyle, backgroundStyle))
	chart.AxisStyle = axisStyle
	chart.LabelStyle = axisStyle
	chart.DrawXYAxisAndLabel()

	for i := 0; i < len(yCoords)-1; i++ {
		y1 := yCoords[i]
		y2 := yCoords[i+1]
		coord1 := canvas.Float64Point{X: float64(i), Y: float64(y1)}
		coord2 := canvas.Float64Point{X: float64(i + 1), Y: float64(y2)}
		chart.DrawBrailleLineWithStyle(coord1, coord2, lineStyle)
	}

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

func testView(m Model) string {
	testViewWidth := min(80, len(m.test.text))
	testProgressStyle := lipgloss.NewStyle().
		Width(testViewWidth).
		Foreground(colours.Fg).
		Background(colours.Bg).
		AlignHorizontal(lipgloss.Left)
	testProgress := testProgressStyle.Render(fmt.Sprintf("%d/%d", m.test.completedWords, m.test.numWords))

	view := lipgloss.JoinVertical(lipgloss.Center, testProgress, m.test.RenderTest(testViewWidth))

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
