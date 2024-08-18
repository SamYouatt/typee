package stats

import (
	"slices"

	"github.com/NimbleMarkets/ntcharts/canvas"
	"github.com/NimbleMarkets/ntcharts/linechart"
	"github.com/SamYouatt/typee/domain"
	"github.com/SamYouatt/typee/util"
	"github.com/charmbracelet/lipgloss"
)

var colours = util.AppColours

func RenderGraph(result *domain.Result) string {
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
