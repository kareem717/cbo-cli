package cmd

import (
	"context"
	"os"

	"github.com/kareem717/k7-cbo/internal/storage"
	"github.com/spf13/cobra"
)

// DB is the package global database instance.
var DB storage.Repository

var rootCmd = &cobra.Command{
	Use:   "k7-cbo",
	Short: "A personal CBO AI agent.",
	Long: `k7-CBO is a personal CBO AI agent that helps you with your business goals.
	It handles lead prospecting, cold outreach and more. It is open source and free.

	We might provide a hosted version of this tool in the future.`,
}

func Execute(ctx context.Context, db storage.Repository) {
	DB = db

	err := rootCmd.ExecuteContext(ctx)

	if err != nil {
		os.Exit(1)
	}
}
