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

// Represents a single run of a typing test
// `test` is the contents for the test
// `currentIndex` is the index of the character expected to be typed next
// `currentlyInvalid` indicates that the current location was entered incorrectly
// `errorIndices` is an array of indices into the text where errors occurred. There can be mulitple occurrences of any index
type test struct {
	text             string
	currentIndex     int
	currentlyInvalid bool
	errorIndices     []int
	complete         bool
}

// Plays the passed in character on the test
func (test *test) play(char byte) {
	if test.complete {
		return
	}

	if test.text[test.currentIndex] == char {
		if test.currentIndex == len(test.text)-1 {
			test.complete = true
		} else {
			test.currentIndex++
		}
	} else {
		test.errorIndices = append(test.errorIndices, test.currentIndex)
	}
}

type model struct {
	width  int
	height int

	test test
}

func initModel() model {
	return model{
		test: test{"Hello world", 0, false, []int{}, false},
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
			m.test.play(msg.String()[0])
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

func (test test) renderTest() string {
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
	return testRendered
}

func (m model) View() string {
	currentIndexStyle := lipgloss.NewStyle().
		Foreground(colours.Fg)
	currentIndex := currentIndexStyle.Render(fmt.Sprintf("%d", m.test.currentIndex))

	numErrorsStyle := lipgloss.NewStyle().
		Foreground(colours.Error)
	numErrors := numErrorsStyle.Render(fmt.Sprintf("%d", len(m.test.errorIndices)))

	view := lipgloss.JoinVertical(lipgloss.Center, m.test.renderTest(), currentIndex, numErrors)

	viewStyle := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Background(colours.Bg).
		Align(lipgloss.Center, lipgloss.Center)

	return viewStyle.Render(view)
}

func main() {
	program := tea.NewProgram(initModel(), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
