package naming

import (
	"fmt"

	"github.com/fdaines/arch-go/pkg/config"
)

func resolveDescription(rule config.NamingRule) string {
	var description string

	if rule.InterfaceImplementationNamingRule != nil {
		if rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith != nil {
			description = fmt.Sprintf(
				"structs that implement '%s' should have simple name starting with '%s'",
				rule.InterfaceImplementationNamingRule.StructsThatImplement,
				*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith,
			)
		}
		if rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith != nil {
			description = fmt.Sprintf(
				"structs that implement '%s' should have simple name ending with '%s'",
				rule.InterfaceImplementationNamingRule.StructsThatImplement,
				*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith,
			)
		}
	}
	ruleDescription := fmt.Sprintf("Packages matching pattern '%s' should comply with [%s]", rule.Package, description)

	return ruleDescription
}
