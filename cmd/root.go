package cmd

import (
	"fmt"
	"os"

	"github.com/fdaines/arch-go/internal/common"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	commandToRun = func() bool { return true }
)

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
	success := commandToRun()
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
