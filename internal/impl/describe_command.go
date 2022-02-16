package impl

import (
	"fmt"
	"github.com/fdaines/arch-go/internal/config"
	"github.com/fdaines/arch-go/internal/utils"
	"os"
	"strings"
)

func DescribeArchitectureGuidelines() {
	utils.ExecuteWithTimer(func() {
		configuration, err := config.LoadConfig("arch-go.yml")
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
			os.Exit(1)
		} else {
			describeDependencyRules(configuration.DependenciesRules)
			describeFunctionRules(configuration.FunctionsRules)
			describeContentRules(configuration.ContentRules)
			describeCyclesRules(configuration.CyclesRules)
			describeNamingRules(configuration.NamingRules)
		}
	})
}

func describeDependencyRules(rules []*config.DependenciesRule) {
	fmt.Printf("Dependency Rules\n")
	for _,r := range rules {
		fmt.Printf("\t* Packages that match pattern '%s',\n", r.Package)
		if r.ShouldOnlyDependsOn != nil {
			fmt.Printf("\t\t* Should only depends on packages that matches:\n")
			for _,p := range r.ShouldOnlyDependsOn {
				fmt.Printf("\t\t\t- '%s'\n", p)
			}
		}
		if r.ShouldNotDependsOn != nil {
			fmt.Printf("\t\t* Should not depends on packages that matches:\n")
			for _,p := range r.ShouldNotDependsOn {
				fmt.Printf("\t\t\t- '%s'\n", p)
			}
		}
		if r.ShouldOnlyDependsOnExternal != nil {
			fmt.Printf("\t\t* Should only depends on external packages that matches\n")
			for _,p := range r.ShouldOnlyDependsOnExternal {
				fmt.Printf("\t\t\t- '%s'\n", p)
			}
		}
		if r.ShouldNotDependsOnExternal != nil {
			fmt.Printf("\t\t* Should not depends on external packages that matches\n")
			for _,p := range r.ShouldNotDependsOnExternal {
				fmt.Printf("\t\t\t- '%s'\n", p)
			}
		}
	}
	fmt.Println()
}

func describeFunctionRules(rules []*config.FunctionsRule) {
	fmt.Printf("Function Rules\n")
	for _,r := range rules {
		fmt.Printf("\t* Packages that match pattern '%s' should comply with the following rules:\n", r.Package)
		if r.MaxLines > 0 {
			fmt.Printf("\t\t* Functions should not have more than %d lines\n", r.MaxLines)
		}
		if r.MaxParameters > 0 {
			fmt.Printf("\t\t* Functions should not have more than %d parameters\n", r.MaxParameters)
		}
		if r.MaxReturnValues > 0 {
			fmt.Printf("\t\t* Functions should not have more than %d return values\n", r.MaxReturnValues)
		}
		if r.MaxPublicFunctionPerFile > 0 {
			fmt.Printf("\t\t* Files should not have more than %d public functions\n", r.MaxPublicFunctionPerFile)
		}
	}
	fmt.Println()
}

func describeContentRules(rules []*config.ContentsRule) {
	fmt.Printf("Content Rules\n")
	for _,r := range rules {
		fmt.Printf("\t* Packages that match pattern '%s' %s\n", r.Package, resolveContentRule(r))
	}
	fmt.Println()
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

func describeCyclesRules(rules []*config.CyclesRule) {
	fmt.Printf("Cycles Rules\n")
	if len(rules) == 0 {
		fmt.Printf("\t* No rules defined\n")
	}
	for _,r := range rules {
		if r.ShouldNotContainCycles {
			fmt.Printf("\t* Packages that match pattern '%s' should not contain cycles\n", r.Package)
		}
	}
	fmt.Println()
}

func describeNamingRules(rules []*config.NamingRule) {
	fmt.Printf("Naming Rules\n")
	for _,r := range rules {
		fmt.Printf("\t* Packages that match pattern '%s' should comply with:\n", r.Package)
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
			fmt.Printf("\t\t* Structs that implement interfaces matching name '%s' %s\n",
				r.InterfaceImplementationNamingRule.StructsThatImplement, namingRule)

		}
	}
	fmt.Println()
}
