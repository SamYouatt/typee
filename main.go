package main

import (
	"fmt"
	"os"

	"github.com/SamYouatt/typee/domain"
	"github.com/SamYouatt/typee/features/practice"

	tea "github.com/charmbracelet/bubbletea"
)

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
