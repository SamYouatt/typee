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
	BgSubtle lipgloss.Color
	Fg       lipgloss.Color
	FgSubtle lipgloss.Color
	Error    lipgloss.Color
}

var colours = Colours{
	Bg:       lipgloss.Color("#E1E1E3"),
	BgSubtle: lipgloss.Color("#D1D3D8"),
	Fg:       lipgloss.Color("#313437"),
	FgSubtle: lipgloss.Color("#AAAEB3"),
	Error:    lipgloss.Color("#DA3333"),
}

type modelState = int

const (
	Ready modelState = iota
	InTest
	TestComplete
)

type Model struct {
	state modelState

	width  int
	height int

	test *Test
}

func initModel() Model {
	return Model{
		state: Ready,
		test:  NewTest("Hello world, this is a really really cool typing test"),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

type TestCompleteMsg struct{}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case " ":
			m.test.PlaySpace()
			return m, nil
		case "backspace":
			m.test.PlayBackspace()
			return m, nil
		case "enter":
			if m.state == Ready {
				m.state = InTest
				return m, nil
			}
		}

		if len(msg.String()) == 1 {
			cmd := m.test.PlayCharacter(msg.String()[0])
			return m, cmd
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case TestCompleteMsg:
		fmt.Println("test complete")
	}

	return m, nil
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

func main() {
	program := tea.NewProgram(initModel(), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
