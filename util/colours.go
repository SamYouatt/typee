package util

import "github.com/charmbracelet/lipgloss"

type Colours struct {
	Bg       lipgloss.Color
	BgSubtle lipgloss.Color
	Fg       lipgloss.Color
	FgSubtle lipgloss.Color
	Error    lipgloss.Color
	Primary  lipgloss.Color
}

var AppColours = Colours{
	Bg:       lipgloss.Color("#E1E1E3"),
	BgSubtle: lipgloss.Color("#D1D3D8"),
	Fg:       lipgloss.Color("#313437"),
	FgSubtle: lipgloss.Color("#AAAEB3"),
	Error:    lipgloss.Color("#DA3333"),
	Primary:  lipgloss.Color("#0369a1"),
}
