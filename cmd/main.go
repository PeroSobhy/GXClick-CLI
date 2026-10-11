package main

import (
	"github.com/PeroSobhy/Portfolio-CLI/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(ui.NewMenuModel())
	p.Start()
}
