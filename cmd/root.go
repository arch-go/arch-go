package cmd

import (
	"fmt"
	"github.com/fdaines/arch-go/common"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "arch-go",
	Version: common.Version,
	Short:   "Architecture checks for Go",
	Long: `Architecture checks for Go:
* Dependency checks`,
	Run: runCommand,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&common.Verbose, "verbose", "v", false, "Verbose Output")
}

func runCommand(cmd *cobra.Command, args []string) {
	fmt.Printf("Not Implemented.\n")
}