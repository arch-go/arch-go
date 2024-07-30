package validators

import (
	"errors"

	"github.com/arch-go/arch-go/api/configuration"
)

func validateDependencyRules(rules []*configuration.DependenciesRule) error {
	for _, rule := range rules {
		if rule.Package == "" {
			return errors.New("dependencies rule - empty package")
		}

		if checkAtLeastOneCriteria(rule) {
			return errors.New(
				"dependencies rule - Should contain one of 'ShouldOnlyDependsOn' or 'ShouldNotDependsOn'")
		}

		if checkAtMostOneCriteria(rule) {
			return errors.New(
				"dependencies rule - Should contain only one of 'ShouldOnlyDependsOn' or 'ShouldNotDependsOn'")
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
			return errors.New("dependencies rule - ShouldOnlyDependsOn needs at least one of" +
				" 'External', 'Internal' or 'Standard'")
		}
	}

	return nil
}

func checkShouldNotDependsOn(rule *configuration.DependenciesRule) error {
	if rule.ShouldNotDependsOn != nil {
		if dependenciesSize(rule.ShouldNotDependsOn) == 0 {
			return errors.New(
				"dependencies rule - ShouldNotDependsOn needs at least one of 'External', 'Internal' or 'Standard'")
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
