package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Colours struct {
	Bg       lipgloss.Color
	Fg       lipgloss.Color
	FgSubtle lipgloss.Color
}

var colours = Colours {
	Bg:       lipgloss.Color("#E1E1E3"),
	Fg:       lipgloss.Color("#313437"),
	FgSubtle: lipgloss.Color("#AAAEB3"),

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
		if test.currentIndex == len(test.text) - 1 {
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

func (m model) View() string {
	textStyle := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Align(lipgloss.Center, lipgloss.Center).
		Foreground(colours.FgSubtle).
		Background(colours.Bg)

	return textStyle.Render(m.runText)
}

func main() {
	program := tea.NewProgram(initModel(), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
