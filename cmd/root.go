package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/internal/reports"
	"github.com/fdaines/arch-go/internal/utils/packages"
	"github.com/fdaines/arch-go/pkg/config"
	"github.com/fdaines/arch-go/pkg/verifications"

	"github.com/fdaines/arch-go/internal/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var commandToRun = runRootCommand

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
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

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolVarP(&common.Verbose, "verbose", "v", false, "Verbose Output")
	rootCmd.Flags().BoolVar(&common.Html, "html", false, "Generate HTML Report")
}

func runCommand(cmd *cobra.Command, args []string) {
	fmt.Fprintf(cmd.OutOrStdout(), "Running arch-go command\n")
	fmt.Fprintf(cmd.OutOrStdout(), "Using config file: %s\n", viper.ConfigFileUsed())
	success := commandToRun(cmd.OutOrStdout())
	if !success {
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find current directory.
	pwd, err := os.Getwd()
	cobra.CheckErr(err)

	// Search config in running directory with name "arch-go.yml".
	viper.AddConfigPath(pwd)
	viper.SetConfigType("yaml")
	viper.SetConfigName("arch-go")

	err = viper.ReadInConfig()
	cobra.CheckErr(err)
}

func runRootCommand(out io.Writer) bool {
	configuration, err := config.LoadConfig(viper.ConfigFileUsed())
	if err != nil {
		fmt.Fprintf(out, "Error: %+v\n", err)
		os.Exit(1)
		return false
	}
	mainPackage, _ := packages.GetMainPackage()
	packages, _ := packages.GetBasicPackagesInfo(mainPackage, out, false)
	moduleInfo := model.ModuleInfo{
		MainPackage: mainPackage,
		Packages:    packages,
	}

	result := verifications.CheckArchitecture(moduleInfo, *configuration)
	report := reports.GenerateReport(result, moduleInfo, *configuration)
	reports.DisplayResult(report, out)

	return true
}
