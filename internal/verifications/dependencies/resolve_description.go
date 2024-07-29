package dependencies

import (
	"fmt"
	"strings"

	"github.com/fdaines/arch-go/api/configuration"
)

func resolveDescription(rule configuration.DependenciesRule) string {
	var ruleDescriptions []string

	if rule.ShouldOnlyDependsOn != nil {
		ruleDescriptions = describeRules(rule.ShouldOnlyDependsOn, ruleDescriptions, true)
	}

	if rule.ShouldNotDependsOn != nil {
		ruleDescriptions = describeRules(rule.ShouldNotDependsOn, ruleDescriptions, false)
	}

	return fmt.Sprintf("Packages matching pattern '%s' should [%s]",
		rule.Package, strings.Join(ruleDescriptions, ","))
}

func describeRules(deps *configuration.Dependencies, ruleDescriptions []string, allowed bool) []string {
	kind := "only"

	if !allowed {
		kind = "not"
	}

	if len(deps.Internal) > 0 {
		ruleDescriptions = append(ruleDescriptions,
			fmt.Sprintf("'%s depend on internal packages that matches [%v]'", kind, deps.Internal))
	}

	if len(deps.External) > 0 {
		ruleDescriptions = append(ruleDescriptions,
			fmt.Sprintf("'%s depend on external packages that matches [%v]'", kind, deps.External))
	}

	if len(deps.Standard) > 0 {
		ruleDescriptions = append(ruleDescriptions,
			fmt.Sprintf("'%s depend on standard packages that matches [%v]'", kind, deps.Standard))
	}

	return ruleDescriptions
}
