package contents

import (
	"fmt"
	"strings"

	"github.com/fdaines/arch-go/api/configuration"
)

func resolveDescription(rule configuration.ContentsRule) string {
	var ruleDescriptions []string
	if rule.ShouldOnlyContainStructs {
		ruleDescriptions = append(ruleDescriptions, "'should only contain structs'")
	}
	if rule.ShouldOnlyContainInterfaces {
		ruleDescriptions = append(ruleDescriptions, "'should only contain interfaces'")
	}
	if rule.ShouldOnlyContainFunctions {
		ruleDescriptions = append(ruleDescriptions, "'should only contain functions'")
	}
	if rule.ShouldOnlyContainMethods {
		ruleDescriptions = append(ruleDescriptions, "'should only contain methods'")
	}
	if rule.ShouldNotContainStructs {
		ruleDescriptions = append(ruleDescriptions, "'should not contain structs'")
	}
	if rule.ShouldNotContainInterfaces {
		ruleDescriptions = append(ruleDescriptions, "'should not contain interfaces'")
	}
	if rule.ShouldNotContainFunctions {
		ruleDescriptions = append(ruleDescriptions, "'should not contain functions'")
	}
	if rule.ShouldNotContainMethods {
		ruleDescriptions = append(ruleDescriptions, "'should not contain methods'")
	}
	return fmt.Sprintf("Packages matching pattern '%s' should complies with [%s]", rule.Package, strings.Join(ruleDescriptions, ","))
}
