package cmd

import (
	"fmt"
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
	fmt.Printf("Describe architecture rules\n")
}
