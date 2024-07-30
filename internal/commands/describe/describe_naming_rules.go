package describe

import (
	"fmt"
	"io"

	"github.com/arch-go/arch-go/api/configuration"
	"github.com/arch-go/arch-go/internal/common"
)

func describeNamingRules(rules []*configuration.NamingRule, out io.Writer) {
	fmt.Fprint(out, "Naming Rules\n")

	if len(rules) == 0 {
		fmt.Fprint(out, common.NoRulesDefined)

		return
	}

	for _, r := range rules {
		fmt.Fprintf(out, "\t* Packages that match pattern '%s' should comply with:\n", r.Package)
		describeInterfaceImplementationNamingRule(r, out)
	}
}

func describeInterfaceImplementationNamingRule(rule *configuration.NamingRule, out io.Writer) {
	if rule.InterfaceImplementationNamingRule != nil {
		namingRule := ""

		if rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith != nil {
			namingRule = fmt.Sprintf("should have simple name ending with '%s'",
				*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith)
		}

		if rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith != nil {
			namingRule = fmt.Sprintf("should have simple name starting with '%s'",
				*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith)
		}

		fmt.Fprintf(out, "\t\t* Structs that implement interfaces matching name '%s' %s\n",
			rule.InterfaceImplementationNamingRule.StructsThatImplement, namingRule)
	}
}
