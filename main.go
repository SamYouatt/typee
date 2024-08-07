package main

import (
	"fmt"
	"os"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Colours struct {
	Bg       lipgloss.Color
	Fg       lipgloss.Color
	FgSubtle lipgloss.Color
	Error    lipgloss.Color
}

var colours = Colours{
	Bg:       lipgloss.Color("#E1E1E3"),
	Fg:       lipgloss.Color("#313437"),
	FgSubtle: lipgloss.Color("#AAAEB3"),
	Error:    lipgloss.Color("#DA3333"),
}


type model struct {
	width  int
	height int

	test *Test
}

func initModel() model {
	return model{
		test: NewTest("Hello world"),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}

		if len(msg.String()) == 1 {
			m.test.Play(msg.String()[0])
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

func (test Test) renderTest() string {
	untypedStyle := lipgloss.NewStyle().
		Foreground(colours.FgSubtle).
		Background(colours.Bg)
	correctlyTypedStyle := lipgloss.NewStyle().
		Foreground(colours.Fg).
		Background(colours.Bg)
	errorStyle := lipgloss.NewStyle().
		Foreground(colours.Error).
		Background(colours.Bg)

	testLetters := []string{}

	for index, letter := range test.text {
		if slices.Contains(test.errorIndices, index) {
			testLetters = append(testLetters, errorStyle.Render(string(letter)))
			continue
		}
		if test.currentIndex > index {
			testLetters = append(testLetters, correctlyTypedStyle.Render(string(letter)))
			continue
		}
		testLetters = append(testLetters, untypedStyle.Render(string(letter)))
	}

	testRendered := lipgloss.JoinHorizontal(lipgloss.Center, testLetters...)

	viewStyle := lipgloss.NewStyle().
		Width(80).
		AlignHorizontal(lipgloss.Center).
		Background(colours.Bg)

	return viewStyle.Render(testRendered)
}

func (m model) View() string {
	currentIndexStyle := lipgloss.NewStyle().
		Width(80).
		Foreground(colours.Fg).
		Background(colours.Bg).
		AlignHorizontal(lipgloss.Center)
	currentIndex := currentIndexStyle.Render(fmt.Sprintf("%d", m.test.currentIndex))

	numErrorsStyle := lipgloss.NewStyle().
		Width(80).
		Foreground(colours.Error).
		Background(colours.Bg).
		AlignHorizontal(lipgloss.Center)
	numErrors := numErrorsStyle.Render(fmt.Sprintf("%d", len(m.test.errorIndices)))

	view := lipgloss.JoinVertical(lipgloss.Center, m.test.renderTest(), currentIndex, numErrors)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, view, lipgloss.WithWhitespaceBackground(colours.Bg))
}

func main() {
	program := tea.NewProgram(initModel(), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
