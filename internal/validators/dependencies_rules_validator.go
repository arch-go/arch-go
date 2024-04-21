package validators

import (
	"fmt"

	"github.com/fdaines/arch-go/pkg/config"
)

func validateDependencyRules(rules []*config.DependenciesRule) error {
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

func checkShouldOnlyDependsOn(rule *config.DependenciesRule) error {
	if rule.ShouldOnlyDependsOn != nil {
		if len(rule.ShouldOnlyDependsOn.External)+len(rule.ShouldOnlyDependsOn.Internal)+len(rule.ShouldOnlyDependsOn.Standard) == 0 {
			return fmt.Errorf("dependencies rule - ShouldOnlyDependsOn needs at least one of 'External', 'Internal' or 'Standard'")
		}
	}
	return nil
}

func checkShouldNotDependsOn(rule *config.DependenciesRule) error {
	if rule.ShouldNotDependsOn != nil {
		if len(rule.ShouldNotDependsOn.External)+len(rule.ShouldNotDependsOn.Internal)+len(rule.ShouldNotDependsOn.Standard) == 0 {
			return fmt.Errorf("dependencies rule - ShouldNotDependsOn needs at least one of 'External', 'Internal' or 'Standard'")
		}
	}
	return nil
}

func checkAtMostOneCriteria(rule *config.DependenciesRule) bool {
	return rule.ShouldNotDependsOn != nil && rule.ShouldOnlyDependsOn != nil
}

func checkAtLeastOneCriteria(rule *config.DependenciesRule) bool {
	return rule.ShouldNotDependsOn == nil && rule.ShouldOnlyDependsOn == nil
}
