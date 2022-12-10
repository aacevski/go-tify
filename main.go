package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/aacevski/go-tify/fetchers"
	utils "github.com/aacevski/go-tify/utils/files"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

var tableFocusStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.DoubleBorder()).
	BorderForeground(lipgloss.Color("240"))

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	noStyle      = lipgloss.NewStyle()
)

type model struct {
	textInput  textinput.Model
	table      table.Model
	rows       []table.Row
	focusIndex int
}

func Search_Song_Rows(access_token string, song string) []table.Row {
	result := fetchers.Search_Song_Raw(access_token, url.QueryEscape(song))
	toReturn := []table.Row{}
	for index, item := range result.Tracks.Items {
		artist := item.Artists[0].Name
		name := item.Name
		uri := item.URI
		toReturn = append(toReturn, table.Row{strconv.Itoa(index), name, artist, uri})
	}
	return toReturn
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.focusIndex == 1 {
				selectedRow := m.table.SelectedRow()
				fetchers.PlayUri(utils.Read_Access_Token(), selectedRow[3])
			}
			if m.focusIndex == 0 {
				res := Search_Song_Rows(utils.Read_Access_Token(), m.textInput.Value())
				m.rows = res
				m.focusIndex = 1
			}
		case "tab", "shift+tab":
			s := msg.String()
			if s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > 1 {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = 1
			}
		}
	}

	cmds := make([]tea.Cmd, 2)

	m.textInput.Value()
	m.textInput.PromptStyle = noStyle
	m.textInput.TextStyle = noStyle

	m.table.SetRows(m.rows)
	if m.focusIndex == 1 {
		m.table.Focus()
		m.table, cmds[1] = m.table.Update(msg)
	} else {
		cmds[1] = m.textInput.Focus()
		m.textInput.PromptStyle = focusedStyle
		m.textInput.TextStyle = focusedStyle

		m.textInput, cmds[0] = m.textInput.Update(msg)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder
	b.WriteString(baseStyle.Render(m.textInput.View()) + "\n")
	if m.table.Focused() {
		b.WriteString(tableFocusStyle.Render(m.table.View()) + "\n")
	} else {
		b.WriteString(baseStyle.Render(m.table.View()) + "\n")
	}
	return b.String()
}

func main() {
	columns := []table.Column{
		{Title: "#", Width: 4},
		{Title: "Track", Width: 30},
		{Title: "Artist", Width: 10},
		{Title: "Track#", Width: 40},
	}

	rows := Search_Song_Rows(utils.Read_Access_Token(), "Killer Queen")
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	ti := textinput.New()
	ti.Placeholder = "Pikachu"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	m := model{
		table:      t,
		textInput:  ti,
		focusIndex: 0,
		rows:       rows,
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
