package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/fdaines/arch-go/internal/commands/migrate"
)

func NewMigrateConfigCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate-configuration",
		Short: "Migrate architecture configuration (arch-go.yml) to current schema",
		Run:   migrateConfig,
	}
}

//nolint:gochecknoinits
func init() {
	migrateCmd := NewMigrateConfigCommand()
	rootCmd.AddCommand(migrateCmd)
}

func migrateConfig(cmd *cobra.Command, _ []string) {
	migrate.NewCommand(cmd.OutOrStdout(), getWorkingDirectory()).Run()
}

func getWorkingDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}

	return wd
}
