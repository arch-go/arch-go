package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/fdaines/arch-go/api/configuration"
	"github.com/fdaines/arch-go/internal/commands/describe"
)

func NewDescribeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "describe",
		Short: "Describe architecture rules",
		Run:   describeRules,
	}
}

//nolint:gochecknoinits
func init() {
	describeCmd := NewDescribeCommand()
	rootCmd.AddCommand(describeCmd)
}

func describeRules(cmd *cobra.Command, _ []string) {
	conf, err := configuration.LoadConfig(viper.ConfigFileUsed())
	if err != nil {
		fmt.Fprintf(cmd.OutOrStdout(), "Error: %+v\n", err)
		os.Exit(1)

		return
	}

	describe.NewCommand(conf, cmd.OutOrStdout()).Run()
}
