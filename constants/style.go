package constants

import "github.com/charmbracelet/lipgloss"

var BaseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

var TableFocusStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.DoubleBorder()).
	BorderForeground(lipgloss.Color("240"))

var (
	FocusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	NoStyle      = lipgloss.NewStyle()
)
