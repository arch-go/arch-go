package cmd

import (
	"github.com/fdaines/arch-go/pkg/commands/describe"
	"github.com/spf13/cobra"
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
	_, _ = rootCmd.AddCommand, describeCmd
}

func describeRules(cmd *cobra.Command, args []string) {
	describe.NewCommand(cmd.OutOrStdout()).Run()
}
