package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/kareem717/k7-cbo/internal/entities/lead"
	"github.com/spf13/cobra"
)

var leadgenCmd = &cobra.Command{
	Use:   "leadgen",
	Short: "Generate leads from LinkedIn.",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		err := DB.Lead().Create(ctx, lead.CreateLeadParams{
			LinkedInURL: "https://www.linkedin.com/in/kareem717",
		})

		if err != nil {
			return err
		}

		leads, err := DB.Lead().GetMany(ctx)
		if err != nil {
			return err
		}

		jsonLeads, err := json.MarshalIndent(leads, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println(string(jsonLeads))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(leadgenCmd)
}
