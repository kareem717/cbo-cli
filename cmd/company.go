package cmd

import (
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	tableui "github.com/kareem717/k7-cbo/cmd/ui/table"
	"github.com/kareem717/k7-cbo/cmd/ui/textinput"
	"github.com/kareem717/k7-cbo/internal/entities/company"
	"github.com/spf13/cobra"
)

var companyCmd = &cobra.Command{
	Use:   "company",
	Short: "Manage companies.",
}

type CompanyOptions struct {
	Name        *textinput.Output
	Description *textinput.Output
}

var createCompanyCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a company.",
	RunE: func(cmd *cobra.Command, args []string) error {
		options := CompanyOptions{
			Name:        &textinput.Output{},
			Description: &textinput.Output{},
		}

		var shouldExit bool

		// Prompt for company name
		nameProgram := tea.NewProgram(textinput.InitialTextInputModel(
			options.Name,
			"What is the name of the company?",
			&shouldExit,
		))
		if _, err := nameProgram.Run(); err != nil {
			return err
		}

		if shouldExit {
			return nameProgram.ReleaseTerminal()
		}

		// Prompt for company description
		descProgram := tea.NewProgram(textinput.InitialTextInputModel(
			options.Description,
			"Enter the company description:",
			&shouldExit,
		))
		if _, err := descProgram.Run(); err != nil {
			return err
		}

		if shouldExit {
			return descProgram.ReleaseTerminal()
		}

		company := company.CreateCompanyParams{
			Name:        options.Name.Output,
			Description: options.Description.Output,
		}

		err := Service.Company.Create(cmd.Context(), company)
		if err != nil {
			return err
		}

		return nil
	},
}

var listCompanyCmd = &cobra.Command{
	Use:   "list",
	Short: "List all companies.",
	RunE: func(cmd *cobra.Command, args []string) error {
		companies, err := Service.Company.GetMany(cmd.Context())
		if err != nil {
			return err
		}

		tableRows := make([]table.Row, len(companies))
		for i, company := range companies {
			tableRows[i] = table.Row{strconv.Itoa(company.ID), company.Name, company.Description}
		}

		t := tableui.NewModel(tableRows, []table.Column{
			{Title: "ID", Width: 10},
			{Title: "Name", Width: 10},
			{Title: "Description", Width: 50},
		}, len(companies))

		// Create and run the Bubble Tea program
		p := tea.NewProgram(t)
		if _, err := p.Run(); err != nil {
			return err
		}

		return nil
	},
}
func init() {
	// Add company command to root command
	rootCmd.AddCommand(companyCmd)

	// Add subcommands
	companyCmd.AddCommand(createCompanyCmd)
	companyCmd.AddCommand(listCompanyCmd)
}
