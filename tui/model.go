package tui

import (
	"github.com/SamYouatt/typee/domain"
	"github.com/SamYouatt/typee/features/practice"
	tea "github.com/charmbracelet/bubbletea"
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
