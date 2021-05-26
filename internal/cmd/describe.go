package cmd

import (
	"github.com/fdaines/arch-go/internal/impl"
	"github.com/spf13/cobra"
)

var (
	describeCmd = &cobra.Command{
		Use:   "describe",
		Short: "Describe architecture rules",
		Run:   describeRules,
	}
)

func init() {
	rootCmd.AddCommand(describeCmd)
}

func describeRules(cmd *cobra.Command, args []string) {
	impl.DescribeArchitectureGuidelines()
}
