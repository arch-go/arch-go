package describe

import (
	"fmt"
	"github.com/fdaines/arch-go/internal/common"
	"github.com/fdaines/arch-go/old/config"
	"io"
)

func describeNamingRules(rules []*config.NamingRule, out io.Writer) {
	fmt.Fprintf(out, "Naming Rules\n")
	if len(rules) == 0 {
		fmt.Fprintf(out, common.NoRulesDefined)
		return
	}
	for _, r := range rules {
		fmt.Fprintf(out, "\t* Packages that match pattern '%s' should comply with:\n", r.Package)
		describeInterfaceImplementationNamingRule(r, out)
	}
	fmt.Fprintln(out)
}

func describeInterfaceImplementationNamingRule(r *config.NamingRule, out io.Writer) {
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
