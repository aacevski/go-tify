package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/aacevski/go-tify/constants"
	"github.com/aacevski/go-tify/fetchers"
	utils "github.com/aacevski/go-tify/utils/files"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type UserInterface struct {
	textInput  textinput.Model
	table      table.Model
	rows       []table.Row
	focusIndex int
}

func (userInterface UserInterface) Init() tea.Cmd { return nil }

func (userInterface UserInterface) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return userInterface, tea.Quit
		case "enter":
			if userInterface.focusIndex == 1 {
				selectedRow := userInterface.table.SelectedRow()
				fetchers.PlayUri(utils.ReadAccessToken(), selectedRow[3])
			}
			if userInterface.focusIndex == 0 {
				if userInterface.textInput.Value() == "" {
					return userInterface, nil
				}

				res := fetchers.Search_Song_Rows(utils.ReadAccessToken(), userInterface.textInput.Value())
				userInterface.rows = res
				userInterface.focusIndex = 1
			}
		case "tab", "shift+tab":
			s := msg.String()
			if s == "shift+tab" {
				userInterface.focusIndex--
			} else {
				userInterface.focusIndex++
			}

			if userInterface.focusIndex > 1 {
				userInterface.focusIndex = 0
			} else if userInterface.focusIndex < 0 {
				userInterface.focusIndex = 1
			}
		}
	}

	cmds := make([]tea.Cmd, 2)

	userInterface.textInput.Value()
	userInterface.textInput.PromptStyle = constants.NoStyle
	userInterface.textInput.TextStyle = constants.NoStyle

	userInterface.table.SetRows(userInterface.rows)
	if userInterface.focusIndex == 1 {
		userInterface.table.Focus()
		userInterface.table, cmds[1] = userInterface.table.Update(msg)
	} else {
		cmds[1] = userInterface.textInput.Focus()
		userInterface.textInput.PromptStyle = constants.FocusedStyle
		userInterface.textInput.TextStyle = constants.FocusedStyle

		userInterface.textInput, cmds[0] = userInterface.textInput.Update(msg)
	}

	return userInterface, tea.Batch(cmds...)
}

func (userInterface UserInterface) View() string {
	var builder strings.Builder
	builder.WriteString(constants.BaseStyle.Render(userInterface.textInput.View()) + "\n")

	if userInterface.table.Focused() {
		builder.WriteString(constants.TableFocusStyle.Render(userInterface.table.View()) + "\n")
	} else {
		builder.WriteString(constants.BaseStyle.Render(userInterface.table.View()) + "\n")
	}

	return builder.String()
}

func Build() {
	columns := []table.Column{
		{Title: "#", Width: 4},
		{Title: "Track", Width: 30},
		{Title: "Artist", Width: 10},
		{Title: "Track#", Width: 40},
	}

	rows := fetchers.Search_Song_Rows(utils.ReadAccessToken(), "Killer Queen")
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(7),
	)

	styles := table.DefaultStyles()
	styles.Header = styles.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	styles.Selected = styles.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(styles)

	textInput := textinput.New()
	textInput.Placeholder = "Killer Queen"
	textInput.Focus()
	textInput.CharLimit = 156
	textInput.Width = 20

	userInterface := UserInterface{
		table:      t,
		textInput:  textInput,
		focusIndex: 0,
		rows:       rows,
	}

	if _, err := tea.NewProgram(userInterface).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
