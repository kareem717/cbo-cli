package cmd

import (
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kareem717/k7-cbo/cmd/steps"
	"github.com/kareem717/k7-cbo/cmd/ui/selectinput"
	tableui "github.com/kareem717/k7-cbo/cmd/ui/table"
	"github.com/kareem717/k7-cbo/cmd/ui/textinput"
	"github.com/kareem717/k7-cbo/internal/entities/mom"
	"github.com/spf13/cobra"
)

var momCmd = &cobra.Command{
	Use:   "mom",
	Short: "Manage mom tests.",
}

type MomOptions struct {
	CompanyID  *selectinput.Selection
	Hypothesis *textinput.Output
}

var createMomCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a mom test.",
	RunE: func(cmd *cobra.Command, args []string) error {
		options := MomOptions{
			CompanyID:  &selectinput.Selection{},
			Hypothesis: &textinput.Output{},
		}

		var shouldExit bool

		companies, err := Service.Company.GetMany(cmd.Context())
		if err != nil {
			return err
		}

		items := make([]steps.Item, len(companies))
		for i, company := range companies {
			items[i] = steps.Item{Flag: strconv.Itoa(company.ID), Title: company.Name}
		}

		// Prompt for company name
		nameProgram := tea.NewProgram(selectinput.InitialModelSelect(
			items,
			options.CompanyID,
			"What company is this test for?",
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
			options.Hypothesis,
			"Enter the hypothesis for this test:",
			&shouldExit,
		))
		if _, err := descProgram.Run(); err != nil {
			return err
		}

		if shouldExit {
			return descProgram.ReleaseTerminal()
		}

		companyID, err := strconv.Atoi(options.CompanyID.Choice)
		if err != nil {
			return err
		}

		mom := mom.CreateMomTestParams{
			CompanyID:  companyID,
			Hypothesis: options.Hypothesis.Output,
		}

		err = Service.Mom.Create(cmd.Context(), mom)
		if err != nil {
			return err
		}

		return nil
	},
}

var listMomCmd = &cobra.Command{
	Use:   "list",
	Short: "List all mom tests.",
	RunE: func(cmd *cobra.Command, args []string) error {
		moms, err := Service.Mom.GetMany(cmd.Context())
		if err != nil {
			return err
		}

		tableRows := make([]table.Row, len(moms))
		for i, mom := range moms {
			tableRows[i] = table.Row{strconv.Itoa(mom.ID), strconv.Itoa(mom.CompanyID), mom.Hypothesis}
		}

		t := tableui.NewModel(tableRows, []table.Column{
			{Title: "ID", Width: 10},
			{Title: "Company ID", Width: 10},
			{Title: "Hypothesis", Width: 50},
		}, len(moms))

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
	rootCmd.AddCommand(momCmd)

	// Add subcommands
	momCmd.AddCommand(createMomCmd)
	momCmd.AddCommand(listMomCmd)
}
