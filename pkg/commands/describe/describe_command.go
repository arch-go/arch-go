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
		configuration, err := config.LoadConfig(viper.ConfigFileUsed())
		if err != nil {
			fmt.Fprintf(dc.Output, "Error: %+v\n", err)
			exitCode = 1
			return
		}
		err = validators.ValidateConfiguration(configuration)
		if err != nil {
			fmt.Fprintf(dc.Output, "Invalid Configuration: %+v\n", err)
			exitCode = 1
			return
		} else {
			describeDependencyRules(configuration.DependenciesRules, dc.Output)
			describeFunctionRules(configuration.FunctionsRules, dc.Output)
			describeContentRules(configuration.ContentRules, dc.Output)
			describeNamingRules(configuration.NamingRules, dc.Output)
			describeThresholdRules(configuration.Threshold, dc.Output)
		}
	})
	os.Exit(exitCode)
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
