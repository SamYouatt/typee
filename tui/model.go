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
		test:  practice.NewTest("Apple Banana Carrot Door Elephant Flower Garden House Island Jacket Kite Lemon Mountain Night Ocean People Queen River Street Table Umbrella Victory Window Xylophone Yellow Zebra Animal Basket Cloud Dream Eagle Fire Ghost Honey Igloo Jungle Kitten Ladder Music Notebook Orange Piano"),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
