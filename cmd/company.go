package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
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

var createCmd = &cobra.Command{
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

func init() {
	// Add company command to root command
	rootCmd.AddCommand(companyCmd)

	// Add subcommands
	companyCmd.AddCommand(createCmd)
}
