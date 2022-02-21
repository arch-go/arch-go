package describe

import (
	"fmt"
	"github.com/fdaines/arch-go/internal/common"
	"github.com/fdaines/arch-go/internal/config"
	"github.com/fdaines/arch-go/internal/utils"
	"io"
	"os"
)

func DescribeArchitectureGuidelines(out io.Writer) {
	utils.ExecuteWithTimer(func() {
		configuration, err := config.LoadConfig("arch-go.yml")
		if err != nil {
			fmt.Fprintf(out, "Error: %+v\n", err)
			os.Exit(1)
		} else {
			describeDependencyRules(configuration.DependenciesRules, out)
			describeFunctionRules(configuration.FunctionsRules, out)
			describeContentRules(configuration.ContentRules, out)
			describeCyclesRules(configuration.CyclesRules, out)
			describeNamingRules(configuration.NamingRules, out)
		}
	})
}

func describeCyclesRules(rules []*config.CyclesRule, out io.Writer) {
	fmt.Fprintf(out, "Cycles Rules\n")
	if len(rules) == 0 {
		fmt.Fprintf(out, common.NoRulesDefined)
		return
	}
	for _,r := range rules {
		if r.ShouldNotContainCycles {
			fmt.Fprintf(out, "\t* Packages that match pattern '%s' should not contain cycles\n", r.Package)
		}
	}
	fmt.Fprintln(out)
}
