package describe

import (
	"fmt"
	"github.com/fdaines/arch-go/internal/utils/timer"
	"github.com/fdaines/arch-go/old/config"
	"io"
	"os"
)

func DescribeArchitectureGuidelines(out io.Writer) {
	timer.ExecuteWithTimer(func() {
		configuration, err := config.LoadConfig("arch-go.yml")
		if err != nil {
			fmt.Fprintf(out, "Error: %+v\n", err)
			os.Exit(1)
		} else {
			describeDependencyRules(configuration.DependenciesRules, out)
			describeFunctionRules(configuration.FunctionsRules, out)
			describeContentRules(configuration.ContentRules, out)
			describeNamingRules(configuration.NamingRules, out)
			describeThresholdRules(configuration.Threshold, out)
		}
	})
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
}
