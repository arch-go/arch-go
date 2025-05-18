package naming

import (
	"fmt"

	"github.com/arch-go/arch-go/api/configuration"
)

func resolveDescription(rule configuration.NamingRule) string {
	var description string

	if rule.InterfaceImplementationNamingRule != nil {
		if rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith != nil {
			if rule.InterfaceImplementationNamingRule.StructsThatImplement.Internal != nil {
				description = fmt.Sprintf(
					"structs that implement '%s' should have simple name starting with '%s'",
					*rule.InterfaceImplementationNamingRule.StructsThatImplement.Internal,
					*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith,
				)
			}

			if rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard != nil {
				description = fmt.Sprintf(
					"structs that implement '%s' from standard package '%s' should have simple name starting with '%s'",
					rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard.Interface,
					rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard.Package,
					*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith,
				)
			}

			if rule.InterfaceImplementationNamingRule.StructsThatImplement.External != nil {
				description = fmt.Sprintf(
					"structs that implement '%s' from external package '%s' should have simple name starting with '%s'",
					rule.InterfaceImplementationNamingRule.StructsThatImplement.External.Interface,
					rule.InterfaceImplementationNamingRule.StructsThatImplement.External.Package,
					*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith,
				)
			}
		}

		if rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith != nil {
			if rule.InterfaceImplementationNamingRule.StructsThatImplement.Internal != nil {
				description = fmt.Sprintf(
					"structs that implement '%s' should have simple name ending with '%s'",
					*rule.InterfaceImplementationNamingRule.StructsThatImplement.Internal,
					*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith,
				)
			}

			if rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard != nil {
				description = fmt.Sprintf(
					"structs that implement '%s' from standard package '%s' should have simple name ending with '%s'",
					rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard.Interface,
					rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard.Package,
					*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith,
				)
			}

			if rule.InterfaceImplementationNamingRule.StructsThatImplement.External != nil {
				description = fmt.Sprintf(
					"structs that implement '%s' from external package '%s' should have simple name ending with '%s'",
					rule.InterfaceImplementationNamingRule.StructsThatImplement.External.Interface,
					rule.InterfaceImplementationNamingRule.StructsThatImplement.External.Package,
					*rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith,
				)
			}
		}
	}

	ruleDescription := fmt.Sprintf("Packages matching pattern '%s' should comply with [%s]",
		rule.Package, description)

	return ruleDescription
}
