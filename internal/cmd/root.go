package cmd

import (
	"github.com/fatih/color"
	"github.com/fdaines/arch-go/internal/common"
	"github.com/fdaines/arch-go/internal/impl"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:     "arch-go",
	Version: common.Version,
	Short:   "Architecture checks for Go",
	Long: `Architecture checks for Go:
* Dependencies
* Package contents
* Cyclic dependencies
* Function rules
* Naming rules`,
	Run: runCommand,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&common.Verbose, "verbose", "v", false, "Verbose Output")
	rootCmd.PersistentFlags().BoolVar(&common.Html, "html", false, "Generate HTML Report")
	rootCmd.PersistentFlags().StringVar(&common.Color, "color", "auto", "Print colors (auto, yes, no)")
}

func runCommand(cmd *cobra.Command, args []string) {
	// Skip when set to auto or anything else
	if strings.ToLower(common.Color) == "yes" {
		color.NoColor = false
	}
	if strings.ToLower(common.Color) == "no" {
		color.NoColor = true
	}

	success := impl.CheckArchitecture()
	if !success {
		os.Exit(1)
	}
}
