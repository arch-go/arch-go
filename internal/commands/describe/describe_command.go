package describe

import (
	"fmt"
	"io"
	"os"

	"github.com/fdaines/arch-go/internal/commands"
	"github.com/fdaines/arch-go/pkg/archgo/configuration"

	"github.com/fdaines/arch-go/internal/utils/timer"

	"github.com/fdaines/arch-go/internal/validators"
)

type describeCommand struct {
	commands.BaseCommand
	configuration *configuration.Config
}

func NewCommand(configuration *configuration.Config, output io.Writer) describeCommand {
	return describeCommand{
		BaseCommand:   commands.BaseCommand{Output: output},
		configuration: configuration,
	}
}

func (dc describeCommand) Run() {
	var exitCode int
	timer.ExecuteWithTimer(func() {
		exitCode = runDescribeCommand(dc)
	})
	os.Exit(exitCode)
}

func runDescribeCommand(dc describeCommand) int {
	err := validators.ValidateConfiguration(dc.configuration)
	if err != nil {
		fmt.Fprintf(dc.Output, "Invalid Configuration: %+v\n", err)
		return 1
	}
	describeDependencyRules(dc.configuration.DependenciesRules, dc.Output)
	describeFunctionRules(dc.configuration.FunctionsRules, dc.Output)
	describeContentRules(dc.configuration.ContentRules, dc.Output)
	describeNamingRules(dc.configuration.NamingRules, dc.Output)
	describeThresholdRules(dc.configuration.Threshold, dc.Output)

	return 0
}

func describeThresholdRules(threshold *configuration.Threshold, out io.Writer) {
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
