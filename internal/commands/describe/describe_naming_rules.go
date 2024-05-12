package describe

import (
	"fmt"
	"io"

	"github.com/fdaines/arch-go/api/configuration"

	"github.com/fdaines/arch-go/internal/common"
)

func describeNamingRules(rules []*configuration.NamingRule, out io.Writer) {
	fmt.Fprintf(out, "Naming Rules\n")
	if len(rules) == 0 {
		fmt.Fprintf(out, common.NoRulesDefined)
		return
	}
	for _, r := range rules {
		fmt.Fprintf(out, "\t* Packages that match pattern '%s' should comply with:\n", r.Package)
		describeInterfaceImplementationNamingRule(r, out)
	}
}

func describeInterfaceImplementationNamingRule(r *configuration.NamingRule, out io.Writer) {
	if r.InterfaceImplementationNamingRule != nil {
		namingRule := ""
		if r.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith != nil {
			namingRule = fmt.Sprintf("should have simple name ending with '%s'",
				*r.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith)
		}
		if r.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith != nil {
			namingRule = fmt.Sprintf("should have simple name starting with '%s'",
				*r.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith)
		}
		fmt.Fprintf(out, "\t\t* Structs that implement interfaces matching name '%s' %s\n",
			r.InterfaceImplementationNamingRule.StructsThatImplement, namingRule)

	}
}
