package selectinput

import (
	"fmt"

	"github.com/kareem717/k7-cbo/cmd/steps"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
	titleStyle        = lipgloss.NewStyle().Background(lipgloss.Color("#01FAC6")).Foreground(lipgloss.Color("#030303")).Bold(true).Padding(0, 1, 0)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170")).Bold(true)
	descriptionStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#40BDA3"))
)

type Selection struct {
	Choice string
}

func (s *Selection) Update(optionName string) {
	s.Choice = optionName
}

type model struct {
	cursor  int
	options []steps.Item
	choice  *Selection
	header  string
	exit    *bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func InitialModelSelect(options []steps.Item, selection *Selection, header string, exit *bool) model {
	return model{
		options: options,
		choice:  selection,
		header:  titleStyle.Render(header),
		exit:    exit,
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			*m.exit = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.choice.Update(m.options[m.cursor].Flag)
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := m.header + "\n\n"

	for i, option := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = focusedStyle.Render(">")
			option.Title = selectedItemStyle.Render(option.Title)
		}

		title := focusedStyle.Render(option.Title)
		description := descriptionStyle.Render(option.Desc)

		s += fmt.Sprintf("%s %s\n%s\n\n", cursor, title, description)
	}

	s += fmt.Sprintf("Press %s to confirm choice.\n", focusedStyle.Render("enter"))
	return s
}
