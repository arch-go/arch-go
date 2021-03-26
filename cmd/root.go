package cmd

import (
	"fmt"
	"github.com/fdaines/arch-go/common"
	"github.com/fdaines/arch-go/config"
	packages "github.com/fdaines/arch-go/utils"
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
	configuration, err := config.LoadConfig("arch-go.yml")
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Config: %+v\n", configuration)
		mainPackage, _ := packages.GetMainPackage()

		fmt.Printf("Module to analyze: %s\n", mainPackage)
		packages.GetBasicPackagesInfo()

	}
}