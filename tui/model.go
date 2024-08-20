package tui

import (
	"github.com/SamYouatt/typee/domain"
	"github.com/SamYouatt/typee/features/practice"
	tea "github.com/charmbracelet/bubbletea"
)

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

	test   *practice.Test
	result *domain.Result
}

func InitModel() Model {
	return Model{
		state: Ready,
		test:  practice.NewTest(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
