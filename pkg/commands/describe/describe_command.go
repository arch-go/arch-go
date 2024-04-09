package describe

import (
	"fmt"
	"github.com/fdaines/arch-go/internal/validators"
	"io"
	"os"

	"github.com/fdaines/arch-go/old/config"
	"github.com/fdaines/arch-go/old/utils"
	"github.com/fdaines/arch-go/pkg/commands"
	"github.com/spf13/viper"
)

type command struct {
	commands.BaseCommand
}

func NewCommand(output io.Writer) command {
	return command{
		commands.BaseCommand{Output: output},
	}
}

func (dc command) Run() {
	var exitCode int
	utils.ExecuteWithTimer(func() {
		exitCode = runDescribeCommand(dc)
	})
	os.Exit(exitCode)
}

func runDescribeCommand(dc command) int {
	configuration, err := config.LoadConfig(viper.ConfigFileUsed())
	if err != nil {
		fmt.Fprintf(dc.Output, "Error: %+v\n", err)
		return 1
	}
	err = validators.ValidateConfiguration(configuration)
	if err != nil {
		fmt.Fprintf(dc.Output, "Invalid Configuration: %+v\n", err)
		return 1
	}
	describeDependencyRules(configuration.DependenciesRules, dc.Output)
	describeFunctionRules(configuration.FunctionsRules, dc.Output)
	describeContentRules(configuration.ContentRules, dc.Output)
	describeNamingRules(configuration.NamingRules, dc.Output)
	describeThresholdRules(configuration.Threshold, dc.Output)

	return 0
}

func describeThresholdRules(threshold *config.Threshold, out io.Writer) {
	if threshold == nil {
		return
	}

	fmt.Fprintf(out, "Threshold Rules\n")
	if threshold.Compliance != nil {
		fmt.Fprintf(out,
			"\t* The module must comply with at least %d%% of the rules described above.\n",
			*threshold.Compliance,
		)
	}
	if threshold.Coverage != nil {
		fmt.Fprintf(out,
			"\t* The rules described above must cover at least %d%% of the packages in this module.\n",
			*threshold.Coverage,
		)
	}
	fmt.Fprintln(out)
}
