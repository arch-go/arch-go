package naming

import (
	"fmt"

	"github.com/arch-go/arch-go/v2/api/configuration"
)

func resolveDescription(rule configuration.NamingRule) string {
	var description string

	if rule.InterfaceImplementationNamingRule != nil {
		if rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith != nil {
			if rule.InterfaceImplementationNamingRule.StructsThatImplement.Internal != nil {
				description = internalStartingWith(rule)
			}

			if rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard != nil {
				description = standardStartingWith(rule)
			}

			if rule.InterfaceImplementationNamingRule.StructsThatImplement.External != nil {
				description = externalStartingWith(rule)
			}
		}

		if rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith != nil {
			if rule.InterfaceImplementationNamingRule.StructsThatImplement.Internal != nil {
				description = internalEndingWith(rule)
			}

			if rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard != nil {
				description = standardEndingWith(rule)
			}

			if rule.InterfaceImplementationNamingRule.StructsThatImplement.External != nil {
				description = externalEndingWith(rule)
			}
		}
	}

	ruleDescription := fmt.Sprintf("Packages matching pattern '%s' should comply with [%s]",
		rule.Package, description)

	return ruleDescription
}

func internalStartingWith(rule configuration.NamingRule) string {
	return fmt.Sprintf(
		"structs that implement '%s' should have simple name starting with '%s'",
		*rule.InterfaceImplementationNamingRule.StructsThatImplement.Internal,
		*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith,
	)
}

func standardStartingWith(rule configuration.NamingRule) string {
	return fmt.Sprintf(
		"structs that implement '%s' from standard package '%s' should have simple name starting with '%s'",
		rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard.Interface,
		rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard.Package,
		*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith,
	)
}

func externalStartingWith(rule configuration.NamingRule) string {
	return fmt.Sprintf(
		"structs that implement '%s' from external package '%s' should have simple name starting with '%s'",
		rule.InterfaceImplementationNamingRule.StructsThatImplement.External.Interface,
		rule.InterfaceImplementationNamingRule.StructsThatImplement.External.Package,
		*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith,
	)
}

func internalEndingWith(rule configuration.NamingRule) string {
	return fmt.Sprintf(
		"structs that implement '%s' should have simple name ending with '%s'",
		*rule.InterfaceImplementationNamingRule.StructsThatImplement.Internal,
		*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith,
	)
}

func standardEndingWith(rule configuration.NamingRule) string {
	return fmt.Sprintf(
		"structs that implement '%s' from standard package '%s' should have simple name ending with '%s'",
		rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard.Interface,
		rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard.Package,
		*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith,
	)
}

func externalEndingWith(rule configuration.NamingRule) string {
	return fmt.Sprintf(
		"structs that implement '%s' from external package '%s' should have simple name ending with '%s'",
		rule.InterfaceImplementationNamingRule.StructsThatImplement.External.Interface,
		rule.InterfaceImplementationNamingRule.StructsThatImplement.External.Package,
		*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith,
	)
}
