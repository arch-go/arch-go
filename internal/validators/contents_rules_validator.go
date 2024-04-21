package validators

import (
	"fmt"

	"github.com/fdaines/arch-go/pkg/config"
)

func validateContentRules(rules []*config.ContentsRule) error {
	for _, rule := range rules {
		if rule.Package == "" {
			return fmt.Errorf("content rule - empty package")
		}
		if countTrueValues(rule) == 0 {
			return fmt.Errorf("content rule - At least one criteria should be set")
		}
		if checkShouldOnlyRule(rule.ShouldOnlyContainFunctions, rule) {
			return fmt.Errorf("content rule - if ShouldOnlyContainFunctions is set, then it should be the only parameter")
		}
		if checkShouldOnlyRule(rule.ShouldOnlyContainStructs, rule) {
			return fmt.Errorf("content rule - if ShouldOnlyContainStructs is set, then it should be the only parameter")
		}
		if checkShouldOnlyRule(rule.ShouldOnlyContainMethods, rule) {
			return fmt.Errorf("content rule - if ShouldOnlyContainMethods is set, then it should be the only parameter")
		}
		if checkShouldOnlyRule(rule.ShouldOnlyContainInterfaces, rule) {
			return fmt.Errorf("content rule - if ShouldOnlyContainInterfaces is set, then it should be the only parameter")
		}
	}
	return nil
}

func checkShouldOnlyRule(shouldOnlyRule bool, rule *config.ContentsRule) bool {
	return shouldOnlyRule && countTrueValues(rule) > 1
}

func countTrueValues(rule *config.ContentsRule) int32 {
	return trueValues(
		rule.ShouldOnlyContainFunctions,
		rule.ShouldOnlyContainInterfaces,
		rule.ShouldOnlyContainMethods,
		rule.ShouldOnlyContainStructs,
		rule.ShouldNotContainFunctions,
		rule.ShouldNotContainInterfaces,
		rule.ShouldNotContainMethods,
		rule.ShouldNotContainStructs,
	)
}
