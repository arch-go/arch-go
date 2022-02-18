package cmd

import (
	"github.com/fdaines/arch-go/internal/impl"
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
	impl.DescribeArchitectureGuidelines(cmd.OutOrStdout())
}
