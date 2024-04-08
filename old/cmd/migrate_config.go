package cmd

import (
	"github.com/fdaines/arch-go/old/impl/migrate_config"
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
	migrate_config.MigrateConfiguration(cmd.OutOrStdout())
}
