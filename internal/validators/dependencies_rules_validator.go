package validators

import (
	"fmt"

	"github.com/fdaines/arch-go/api/configuration"
)

func validateDependencyRules(rules []*configuration.DependenciesRule) error {
	for _, rule := range rules {
		if rule.Package == "" {
			return fmt.Errorf("dependencies rule - empty package")
		}

		if checkAtLeastOneCriteria(rule) {
			return fmt.Errorf("dependencies rule - Should contain one of 'ShouldOnlyDependsOn' or 'ShouldNotDependsOn'")
		}

		if checkAtMostOneCriteria(rule) {
			return fmt.Errorf("dependencies rule - Should contain only one of 'ShouldOnlyDependsOn' or 'ShouldNotDependsOn'")
		}

		err := checkShouldNotDependsOn(rule)
		if err != nil {
			return err
		}

		err = checkShouldOnlyDependsOn(rule)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkShouldOnlyDependsOn(rule *configuration.DependenciesRule) error {
	if rule.ShouldOnlyDependsOn != nil {
		if dependenciesSize(rule.ShouldOnlyDependsOn) == 0 {
			return fmt.Errorf("dependencies rule - ShouldOnlyDependsOn needs at least one of 'External', 'Internal' or 'Standard'")
		}
	}

	return nil
}

func checkShouldNotDependsOn(rule *configuration.DependenciesRule) error {
	if rule.ShouldNotDependsOn != nil {
		if dependenciesSize(rule.ShouldNotDependsOn) == 0 {
			return fmt.Errorf("dependencies rule - ShouldNotDependsOn needs at least one of 'External', 'Internal' or 'Standard'")
		}
	}
	return nil
}

func dependenciesSize(d *configuration.Dependencies) int {
	return len(d.External) + len(d.Internal) + len(d.Standard)
}

func checkAtMostOneCriteria(rule *configuration.DependenciesRule) bool {
	return rule.ShouldNotDependsOn != nil && rule.ShouldOnlyDependsOn != nil
}

func checkAtLeastOneCriteria(rule *configuration.DependenciesRule) bool {
	return rule.ShouldNotDependsOn == nil && rule.ShouldOnlyDependsOn == nil
}
