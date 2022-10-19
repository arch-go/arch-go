package describe

import (
	"fmt"
	"io"

	"github.com/fdaines/arch-go/internal/common"
	"github.com/fdaines/arch-go/internal/config"
)

func describeDependencyRules(rules []*config.DependenciesRule, out io.Writer) {
	fmt.Fprintf(out, "Dependency Rules\n")
	if len(rules) == 0 {
		fmt.Fprintf(out, common.NoRulesDefined)
		return
	}
	for _, r := range rules {
		dependencyListPattern := "\t\t\t\t- '%s'\n"
		fmt.Fprintf(out, "\t* Packages that match pattern '%s',\n", r.Package)
		describeShouldOnlyDependsOn(r, out, dependencyListPattern)
		describeShouldNotDependsOn(r, out, dependencyListPattern)
	}
	fmt.Println()
}

func describeShouldNotDependsOn(r *config.DependenciesRule, out io.Writer, dependencyListPattern string) {
	if r.ShouldNotDependsOn != nil {
		fmt.Fprintf(out, "\t\t* Should not depends on:\n")
		describeDependencies(r.ShouldOnlyDependsOn, out, dependencyListPattern)
	}
}

func describeShouldOnlyDependsOn(r *config.DependenciesRule, out io.Writer, dependencyListPattern string) {
	if r.ShouldOnlyDependsOn != nil {
		fmt.Fprintf(out, "\t\t* Should only depends on:\n")
		describeDependencies(r.ShouldOnlyDependsOn, out, dependencyListPattern)
	}
}

func describeDependencies(d *config.Dependencies, out io.Writer, dependencyListPattern string) {
	if len(d.Internal) > 0 {
		fmt.Fprintf(out, "\t\t\t* Internal Packages that matches:\n")
		for _, p := range d.Internal {
			fmt.Fprintf(out, dependencyListPattern, p)
		}
	}
	if len(d.External) > 0 {
		fmt.Fprintf(out, "\t\t\t* External Packages that matches:\n")
		for _, p := range d.External {
			fmt.Fprintf(out, dependencyListPattern, p)
		}
	}
	if len(d.Standard) > 0 {
		fmt.Fprintf(out, "\t\t\t* StandardLib Packages that matches:\n")
		for _, p := range d.Standard {
			fmt.Fprintf(out, dependencyListPattern, p)
		}
	}
}
