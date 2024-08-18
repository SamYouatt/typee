package stats

import (
	"fmt"
	"slices"

	"github.com/NimbleMarkets/ntcharts/canvas"
	"github.com/NimbleMarkets/ntcharts/linechart"
	"github.com/SamYouatt/typee/domain"
	"github.com/SamYouatt/typee/util"
	"github.com/charmbracelet/lipgloss"
)

var colours = util.AppColours

func renderGraph(result *domain.Result) string {
	axisStyle := lipgloss.NewStyle().Foreground(colours.Fg)
	lineStyle := lipgloss.NewStyle().Foreground(colours.Primary)
	backgroundStyle := lipgloss.NewStyle().Foreground(colours.Fg)

	width, height := 80, 10
	yCoords := result.RunningWpm
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

func View(result *domain.Result, width, height int) string {
	if result == nil {
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
	wpmStat := statStyle.Render(fmt.Sprint(result.Wpm))
	wpm := lipgloss.JoinHorizontal(lipgloss.Top, wpmLabel, wpmStat)

	timeTakenLabel := labelStyle.Render("time")
	timeTakenStat := statStyle.Render(fmt.Sprintf("%.2fs", result.TimeTaken.Seconds()))
	timeTaken := lipgloss.JoinHorizontal(lipgloss.Top, timeTakenLabel, timeTakenStat)

	testStats := lipgloss.JoinHorizontal(lipgloss.Top, wpm, timeTaken)

	testCompleteStyle := lipgloss.NewStyle().
		Background(colours.Bg).
		Foreground(colours.Fg).
		Align(lipgloss.Center, lipgloss.Center)
	testComplete := testCompleteStyle.Render("Test complete! 🎉")

	view := lipgloss.JoinVertical(lipgloss.Center, testStats, testComplete, renderGraph(result))

	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, view, lipgloss.WithWhitespaceBackground(colours.Bg))
}
