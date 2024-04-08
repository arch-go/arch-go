package naming

import (
	"fmt"
	"github.com/fdaines/arch-go/old/config"
)

func resolveRuleDescription(rule *config.NamingRule) string {
	var description string

	if rule.InterfaceImplementationNamingRule != nil {
		description = resolveInterfaceImplementationDescription(rule.InterfaceImplementationNamingRule)
	}
	ruleDescription := fmt.Sprintf("Packages matching pattern '%s' should comply with [%s]", rule.Package, description)

	return ruleDescription
}

func resolveInterfaceImplementationDescription(rule *config.InterfaceImplementationRule) string {
	description := ""
	if rule.ShouldHaveSimpleNameStartingWith != "" {
		description = fmt.Sprintf("structs that implement '%s' should have simple name starting with '%s'", rule.StructsThatImplement, rule.ShouldHaveSimpleNameStartingWith)
	}
	if rule.ShouldHaveSimpleNameEndingWith != "" {
		description = fmt.Sprintf("structs that implement '%s' should have simple name ending with '%s'", rule.StructsThatImplement, rule.ShouldHaveSimpleNameEndingWith)
	}

	return description
}
