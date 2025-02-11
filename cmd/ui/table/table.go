package table

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	table      table.Model
	page       int
	totalPages int
}

func NewModel(data []table.Row, headers []table.Column, totalItems int) model {
	t := table.New(
		table.WithColumns(headers),
		table.WithRows(data),
		table.WithFocused(true),
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

	return model{
		table:      t,
		page:       1,
		totalPages: (totalItems + len(data) - 1) / len(data), // Calculate total pages
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\nPage %d of %d\n\nPress arrow keys to navigate, q to quit.",
		m.table.View(),
		m.page,
		m.totalPages,
	)
}