package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/fdaines/arch-go/api"
	"github.com/fdaines/arch-go/api/configuration"
	"github.com/fdaines/arch-go/internal/common"
	"github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/internal/reports"
	"github.com/fdaines/arch-go/internal/utils/packages"
)

//nolint:gochecknoglobals
var (
	commandToRun = runRootCommand
	rootCmd      = &cobra.Command{
		Use:     "arch-go",
		Version: common.Version,
		Short:   "Architecture checks for Go",
		Long: `Architecture checks for Go:
* Dependencies
* Package contents
* Function rules
* Naming rules`,
		Run: runCommand,
	}
)

//nolint:gochecknoinits
func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&common.Verbose, "verbose", "v", false, "Verbose Output")
	rootCmd.PersistentFlags().BoolVar(&common.Html, "html", false, "Generate HTML Report")
	rootCmd.PersistentFlags().StringVar(&common.Color, "color", "auto", "Print colors (auto, yes, no)")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func runCommand(cmd *cobra.Command, _ []string) {
	configureColor()
	fmt.Fprintf(cmd.OutOrStdout(), "Running arch-go command\n")
	fmt.Fprintf(cmd.OutOrStdout(), "Using configuration file: %s\n", viper.ConfigFileUsed())

	success := commandToRun(cmd.OutOrStdout())
	if !success {
		os.Exit(1)
	}
}

func configureColor() {
	// Skip when set to auto or anything else
	if strings.ToLower(common.Color) == "yes" {
		color.NoColor = false
	}

	if strings.ToLower(common.Color) == "no" {
		color.NoColor = true
	}
}

// initConfig reads in configuration file and ENV variables if set.
func initConfig() {
	// Find current directory.
	pwd, err := os.Getwd()
	cobra.CheckErr(err)

	// Search configuration in running directory with name "arch-go.yml".
	viper.AddConfigPath(pwd)
	viper.SetConfigType("yaml")
	viper.SetConfigName("arch-go")

	err = viper.ReadInConfig()
	cobra.CheckErr(err)
}

func runRootCommand(out io.Writer) bool {
	conf, err := configuration.LoadConfig(viper.ConfigFileUsed())
	if err != nil {
		fmt.Fprintf(out, "Error: %+v\n", err)
		os.Exit(1)

		return false
	}

	mainPackage, _ := packages.GetMainPackage()
	packages, _ := packages.GetBasicPackagesInfo(mainPackage, out, common.Verbose)
	moduleInfo := model.ModuleInfo{
		MainPackage: mainPackage,
		Packages:    packages,
	}

	result := api.CheckArchitecture(moduleInfo, *conf)
	report := reports.GenerateReport(result, moduleInfo, *conf)
	reports.DisplayResult(report, out)

	return result.Passes
}
