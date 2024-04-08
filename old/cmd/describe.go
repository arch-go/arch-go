package cmd

import (
	"github.com/fdaines/arch-go/old/impl/describe"
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
	rootCmd.AddCommand(describeCmd)
}

func describeRules(cmd *cobra.Command, args []string) {
	describe.DescribeArchitectureGuidelines(cmd.OutOrStdout())
}
