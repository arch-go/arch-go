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
	return fmt.Sprintf("Packages matching pattern '%s' should [%s]", rule.Package, strings.Join(ruleDescriptions, ","))
}

func describeRules(d *configuration.Dependencies, ruleDescriptions []string, allowed bool) []string {
	kind := "only"
	if !allowed {
		kind = "not"
	}
	if len(d.Internal) > 0 {
		ruleDescriptions = append(ruleDescriptions, fmt.Sprintf("'%s depend on internal packages that matches [%v]'", kind, d.Internal))
	}
	if len(d.External) > 0 {
		ruleDescriptions = append(ruleDescriptions, fmt.Sprintf("'%s depend on external packages that matches [%v]'", kind, d.External))
	}
	if len(d.Standard) > 0 {
		ruleDescriptions = append(ruleDescriptions, fmt.Sprintf("'%s depend on standard packages that matches [%v]'", kind, d.Standard))
	}
	return ruleDescriptions
}
