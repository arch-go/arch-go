package cmd

import (
	"fmt"
	"os"

	"github.com/fdaines/arch-go/api/configuration"

	"github.com/fdaines/arch-go/internal/commands/describe"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewDescribeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "describe",
		Short: "Describe architecture rules",
		Run:   describeRules,
	}
}

func init() {
	describeCmd := NewDescribeCommand()
	rootCmd.AddCommand(describeCmd)
}

func describeRules(cmd *cobra.Command, args []string) {
	configuration, err := configuration.LoadConfig(viper.ConfigFileUsed())
	if err != nil {
		fmt.Fprintf(cmd.OutOrStdout(), "Error: %+v\n", err)
		os.Exit(1)
		return
	}
	describe.NewCommand(configuration, cmd.OutOrStdout()).Run()
}
