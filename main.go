package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Colours struct {
	Bg       lipgloss.Color
	BgSubtle lipgloss.Color
	Fg       lipgloss.Color
	FgSubtle lipgloss.Color
	Error    lipgloss.Color
	Primary  lipgloss.Color
}

var colours = Colours{
	Bg:       lipgloss.Color("#E1E1E3"),
	BgSubtle: lipgloss.Color("#D1D3D8"),
	Fg:       lipgloss.Color("#313437"),
	FgSubtle: lipgloss.Color("#AAAEB3"),
	Error:    lipgloss.Color("#DA3333"),
	Primary:  lipgloss.Color("#0369a1"),
}

type modelState = int

const (
	Ready        modelState = 1
	InTest                  = 2
	TestComplete            = 3
)

type Model struct {
	state modelState

	width  int
	height int

	test          *Test
	completedTest *CompletedTest
}

func initModel() Model {
	return Model{
		state: Ready,
		test:  NewTest("blue red fast slow kind strong quick brown light dark happy smart quiet loud sleep"),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.state == Ready {
				m.state = InTest
				return m, nil
			}
		}

		if m.test.PlayInput(msg.String()) {
			m.completedTest = m.test.CompleteTest()
			m.test = nil
			m.state = TestComplete
		}

		return m, nil

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

func main() {
	program := tea.NewProgram(initModel(), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
