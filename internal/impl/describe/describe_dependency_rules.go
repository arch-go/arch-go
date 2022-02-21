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
	for _,r := range rules {
		dependencyListPattern := "\t\t\t- '%s'\n"
		fmt.Fprintf(out, "\t* Packages that match pattern '%s',\n", r.Package)
		describeShouldOnlyDependsOn(r, out, dependencyListPattern)
		describeShouldNotDependsOn(r, out, dependencyListPattern)
		describeShouldOnlyDependsOnExternal(r, out, dependencyListPattern)
		describeShouldNotDependsOnExternal(r, out, dependencyListPattern)
	}
	fmt.Println()
}

func describeShouldNotDependsOnExternal(r *config.DependenciesRule, out io.Writer, dependencyListPattern string) {
	if r.ShouldNotDependsOnExternal != nil {
		fmt.Fprintf(out, "\t\t* Should not depends on external packages that matches\n")
		for _, p := range r.ShouldNotDependsOnExternal {
			fmt.Fprintf(out, dependencyListPattern, p)
		}
	}
}

func describeShouldOnlyDependsOnExternal(r *config.DependenciesRule, out io.Writer, dependencyListPattern string) {
	if r.ShouldOnlyDependsOnExternal != nil {
		fmt.Fprintf(out, "\t\t* Should only depends on external packages that matches\n")
		for _, p := range r.ShouldOnlyDependsOnExternal {
			fmt.Fprintf(out, dependencyListPattern, p)
		}
	}
}

func describeShouldNotDependsOn(r *config.DependenciesRule, out io.Writer, dependencyListPattern string) {
	if r.ShouldNotDependsOn != nil {
		fmt.Fprintf(out, "\t\t* Should not depends on packages that matches:\n")
		for _, p := range r.ShouldNotDependsOn {
			fmt.Fprintf(out, dependencyListPattern, p)
		}
	}
}

func describeShouldOnlyDependsOn(r *config.DependenciesRule, out io.Writer, dependencyListPattern string) {
	if r.ShouldOnlyDependsOn != nil {
		fmt.Fprintf(out, "\t\t* Should only depends on packages that matches:\n")
		for _, p := range r.ShouldOnlyDependsOn {
			fmt.Fprintf(out, dependencyListPattern, p)
		}
	}
}