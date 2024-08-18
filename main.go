package main

import (
	"fmt"
	"os"

	"github.com/SamYouatt/typee/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	program := tea.NewProgram(tui.InitModel(), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
