package cmd

import (
	"fmt"
	"os"

	migrate_config "github.com/fdaines/arch-go/pkg/commands/migrate-config"
	"github.com/spf13/cobra"
)

func NewMigrateConfigCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate-config",
		Short: "Migrate architecture configuration (arch-go.yml) to current schema",
		Run:   migrateConfig,
	}
}

func init() {
	migrateCmd := NewMigrateConfigCommand()
	rootCmd.AddCommand(migrateCmd)
}

func migrateConfig(cmd *cobra.Command, args []string) {
	migrate_config.NewCommand(cmd.OutOrStdout(), getWorkingDirectory()).Run()
}

func getWorkingDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}
	return wd
}
