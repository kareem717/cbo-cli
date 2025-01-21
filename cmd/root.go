package cmd

import (
	"context"
	"os"

	"github.com/kareem717/k7-cbo/internal/service"
	"github.com/spf13/cobra"
)

// Service is the package global service instance.
var Service *service.Service

var rootCmd = &cobra.Command{
	Use:   "k7-cbo",
	Short: "A personal CBO AI agent.",
	Long: `k7-CBO is a personal CBO AI agent that helps you with your business goals.
	It handles lead prospecting, cold outreach and more. It is open source and free.

	We might provide a hosted version of this tool in the future.`,
}

func Execute(ctx context.Context, service *service.Service) {
	Service = service

	err := rootCmd.ExecuteContext(ctx)

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(companyCmd)
}
