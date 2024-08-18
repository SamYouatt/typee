package main

import (
	"fmt"
	"os"

	"github.com/SamYouatt/typee/domain"
	"github.com/SamYouatt/typee/features/practice"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)


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

	test   *practice.Test
	result *domain.Result
}

func initModel() Model {
	return Model{
		state: Ready,
		test:  practice.NewTest("blue red fast slow kind strong quick brown light dark happy smart quiet loud sleep"),
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
			m.result = m.test.CompleteTest()
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
