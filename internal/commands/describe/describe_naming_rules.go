package describe

import (
	"fmt"
	"io"

	"github.com/arch-go/arch-go/v2/api/configuration"
	"github.com/arch-go/arch-go/v2/internal/common"
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
			if rule.InterfaceImplementationNamingRule.StructsThatImplement.Internal != nil {
				namingRule = internalEndingWith(rule)
			}

			if rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard != nil {
				namingRule = standardEndingWith(rule)
			}

			if rule.InterfaceImplementationNamingRule.StructsThatImplement.External != nil {
				namingRule = externalEndingWith(rule)
			}
		}

		if rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith != nil {
			if rule.InterfaceImplementationNamingRule.StructsThatImplement.Internal != nil {
				namingRule = internalStartingWith(rule)
			}

			if rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard != nil {
				namingRule = standardStartingWith(rule)
			}

			if rule.InterfaceImplementationNamingRule.StructsThatImplement.External != nil {
				namingRule = externalStartingWith(rule)
			}
		}

		fmt.Fprintf(out, "\t\t* Structs that implement interfaces matching name %s\n", namingRule)
	}
}

func internalEndingWith(rule *configuration.NamingRule) string {
	return fmt.Sprintf(
		"'%s' should have simple name ending with '%s'",
		*rule.InterfaceImplementationNamingRule.StructsThatImplement.Internal,
		*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith,
	)
}

func standardEndingWith(rule *configuration.NamingRule) string {
	return fmt.Sprintf(
		"'%s' from standard package '%s' should have simple name ending with '%s'",
		rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard.Interface,
		rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard.Package,
		*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith,
	)
}

func externalEndingWith(rule *configuration.NamingRule) string {
	return fmt.Sprintf(
		"'%s' from external package '%s' should have simple name ending with '%s'",
		rule.InterfaceImplementationNamingRule.StructsThatImplement.External.Interface,
		rule.InterfaceImplementationNamingRule.StructsThatImplement.External.Package,
		*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith,
	)
}

func internalStartingWith(rule *configuration.NamingRule) string {
	return fmt.Sprintf(
		"'%s' should have simple name starting with '%s'",
		*rule.InterfaceImplementationNamingRule.StructsThatImplement.Internal,
		*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith,
	)
}

func standardStartingWith(rule *configuration.NamingRule) string {
	return fmt.Sprintf(
		"'%s' from standard package '%s' should have simple name starting with '%s'",
		rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard.Interface,
		rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard.Package,
		*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith,
	)
}

func externalStartingWith(rule *configuration.NamingRule) string {
	return fmt.Sprintf(
		"'%s' from external package '%s' should have simple name starting with '%s'",
		rule.InterfaceImplementationNamingRule.StructsThatImplement.External.Interface,
		rule.InterfaceImplementationNamingRule.StructsThatImplement.External.Package,
		*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith,
	)
}
