package describe

import (
	"fmt"
	"io"

	"github.com/fdaines/arch-go/api/configuration"
	"github.com/fdaines/arch-go/internal/common"
)

func describeDependencyRules(rules []*configuration.DependenciesRule, out io.Writer) {
	fmt.Fprint(out, "Dependency Rules\n")

	if len(rules) == 0 {
		fmt.Fprint(out, common.NoRulesDefined)

		return
	}

	for _, r := range rules {
		dependencyListPattern := "\t\t\t\t- '%s'\n"

		fmt.Fprintf(out, "\t* Packages that match pattern '%s',\n", r.Package)
		describeShouldOnlyDependsOn(r, out, dependencyListPattern)
		describeShouldNotDependsOn(r, out, dependencyListPattern)
	}
}

func describeShouldNotDependsOn(rule *configuration.DependenciesRule, out io.Writer, dependencyListPattern string) {
	if rule.ShouldNotDependsOn != nil {
		fmt.Fprint(out, "\t\t* Should not depends on:\n")
		describeDependencies(rule.ShouldNotDependsOn, out, dependencyListPattern)
	}
}

func describeShouldOnlyDependsOn(rule *configuration.DependenciesRule, out io.Writer, dependencyListPattern string) {
	if rule.ShouldOnlyDependsOn != nil {
		fmt.Fprint(out, "\t\t* Should only depends on:\n")
		describeDependencies(rule.ShouldOnlyDependsOn, out, dependencyListPattern)
	}
}

func describeDependencies(deps *configuration.Dependencies, out io.Writer, dependencyListPattern string) {
	if len(deps.Internal) > 0 {
		fmt.Fprint(out, "\t\t\t* Internal Packages that matches:\n")

		for _, p := range deps.Internal {
			fmt.Fprintf(out, dependencyListPattern, p)
		}
	}

	if len(deps.External) > 0 {
		fmt.Fprint(out, "\t\t\t* External Packages that matches:\n")

		for _, p := range deps.External {
			fmt.Fprintf(out, dependencyListPattern, p)
		}
	}

	if len(deps.Standard) > 0 {
		fmt.Fprint(out, "\t\t\t* StandardLib Packages that matches:\n")

		for _, p := range deps.Standard {
			fmt.Fprintf(out, dependencyListPattern, p)
		}
	}
}
