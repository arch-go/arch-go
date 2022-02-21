package impl

import (
	"fmt"
	"github.com/fdaines/arch-go/internal/common"
	"github.com/fdaines/arch-go/internal/config"
	"github.com/fdaines/arch-go/internal/utils"
	"io"
	"os"
	"strings"
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

func describeDependencyRules(rules []*config.DependenciesRule, out io.Writer) {
	fmt.Fprintf(out, "Dependency Rules\n")
	if len(rules) == 0 {
		fmt.Fprintf(out, common.NoRulesDefined)
		return
	}
	for _,r := range rules {
		dependencyListPattern := "\t\t\t- '%s'\n"
		fmt.Fprintf(out, "\t* Packages that match pattern '%s',\n", r.Package)
		if r.ShouldOnlyDependsOn != nil {
			fmt.Fprintf(out, "\t\t* Should only depends on packages that matches:\n")
			for _,p := range r.ShouldOnlyDependsOn {
				fmt.Fprintf(out, dependencyListPattern, p)
			}
		}
		if r.ShouldNotDependsOn != nil {
			fmt.Fprintf(out, "\t\t* Should not depends on packages that matches:\n")
			for _,p := range r.ShouldNotDependsOn {
				fmt.Fprintf(out, dependencyListPattern, p)
			}
		}
		if r.ShouldOnlyDependsOnExternal != nil {
			fmt.Fprintf(out, "\t\t* Should only depends on external packages that matches\n")
			for _,p := range r.ShouldOnlyDependsOnExternal {
				fmt.Fprintf(out, dependencyListPattern, p)
			}
		}
		if r.ShouldNotDependsOnExternal != nil {
			fmt.Fprintf(out, "\t\t* Should not depends on external packages that matches\n")
			for _,p := range r.ShouldNotDependsOnExternal {
				fmt.Fprintf(out, dependencyListPattern, p)
			}
		}
	}
	fmt.Println()
}

func describeFunctionRules(rules []*config.FunctionsRule, out io.Writer) {
	fmt.Fprintf(out, "Function Rules\n")
	if len(rules) == 0 {
		fmt.Fprintf(out, common.NoRulesDefined)
		return
	}
	for _,r := range rules {
		fmt.Fprintf(out, "\t* Packages that match pattern '%s' should comply with the following rules:\n", r.Package)
		if r.MaxLines > 0 {
			fmt.Fprintf(out, "\t\t* Functions should not have more than %d lines\n", r.MaxLines)
		}
		if r.MaxParameters > 0 {
			fmt.Fprintf(out, "\t\t* Functions should not have more than %d parameters\n", r.MaxParameters)
		}
		if r.MaxReturnValues > 0 {
			fmt.Fprintf(out, "\t\t* Functions should not have more than %d return values\n", r.MaxReturnValues)
		}
		if r.MaxPublicFunctionPerFile > 0 {
			fmt.Fprintf(out, "\t\t* Files should not have more than %d public functions\n", r.MaxPublicFunctionPerFile)
		}
	}
	fmt.Println()
}

func describeContentRules(rules []*config.ContentsRule, out io.Writer) {
	fmt.Fprintf(out, "Content Rules\n")
	if len(rules) == 0 {
		fmt.Fprintf(out, common.NoRulesDefined)
		return
	}
	for _,r := range rules {
		fmt.Fprintf(out, "\t* Packages that match pattern '%s' %s\n", r.Package, resolveContentRule(r))
	}
	fmt.Fprintln(out)
}

func resolveContentRule(r *config.ContentsRule) string {
	var shouldNotContain []string
	if r.ShouldOnlyContainStructs {
		return "should only contain structs"
	}
	if r.ShouldOnlyContainInterfaces {
		return "should only contain interfaces"
	}
	if r.ShouldOnlyContainFunctions {
		return "should only contain functions"
	}
	if r.ShouldOnlyContainMethods {
		return "should only contain methods"
	}
	if r.ShouldNotContainStructs {
		shouldNotContain = append(shouldNotContain, "structs")
	}
	if r.ShouldNotContainInterfaces {
		shouldNotContain = append(shouldNotContain, "interfaces")
	}
	if r.ShouldNotContainFunctions {
		shouldNotContain = append(shouldNotContain, "functions")
	}
	if r.ShouldNotContainMethods {
		shouldNotContain = append(shouldNotContain, "methods")
	}
	return fmt.Sprintf("should not contain %s", strings.Join(shouldNotContain, " or "))
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

func describeNamingRules(rules []*config.NamingRule, out io.Writer) {
	fmt.Fprintf(out, "Naming Rules\n")
	if len(rules) == 0 {
		fmt.Fprintf(out, common.NoRulesDefined)
		return
	}
	for _,r := range rules {
		fmt.Fprintf(out, "\t* Packages that match pattern '%s' should comply with:\n", r.Package)
		if r.InterfaceImplementationNamingRule != nil {
			namingRule := ""
			if r.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith != "" {
				namingRule = fmt.Sprintf("should have simple name ending with '%s'",
					r.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith)
			}
			if r.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith != "" {
				namingRule = fmt.Sprintf("should have simple name starting with '%s'",
					r.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith)
			}
			fmt.Fprintf(out, "\t\t* Structs that implement interfaces matching name '%s' %s\n",
				r.InterfaceImplementationNamingRule.StructsThatImplement, namingRule)

		}
	}
	fmt.Fprintln(out)
}
