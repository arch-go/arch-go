package validators

import (
	"fmt"
	"github.com/fdaines/arch-go/old/config"
)

func ValidateConfiguration(configuration *config.Config) error {
	err := checkRulesQuantity(configuration)
	if err != nil {
		return err
	}
	err = validateDependencyRules(configuration.DependenciesRules)
	if err != nil {
		return err
	}
	err = validateFunctionRules(configuration.FunctionsRules)
	if err != nil {
		return err
	}
	err = validateContentRules(configuration.ContentRules)
	if err != nil {
		return err
	}
	return nil
}

func checkRulesQuantity(c *config.Config) error {
	total := len(c.ContentRules) + len(c.DependenciesRules) + len(c.FunctionsRules) + len(c.NamingRules)
	if total == 0 {
		return fmt.Errorf("configuration file should have at least one rule")
	}
	return nil
}

func validateDependencyRules(rules []*config.DependenciesRule) error {
	for _, rule := range rules {
		if rule.Package == "" {
			return fmt.Errorf("dependencies rule - empty package")
		}
		if rule.ShouldNotDependsOn == nil && rule.ShouldOnlyDependsOn == nil {
			return fmt.Errorf("dependencies rule - Should contain one of 'ShouldOnlyDependsOn' or 'ShouldNotDependsOn'")
		}
		if rule.ShouldNotDependsOn != nil && rule.ShouldOnlyDependsOn != nil {
			return fmt.Errorf("dependencies rule - Should contain only one of 'ShouldOnlyDependsOn' or 'ShouldNotDependsOn'")
		}
		if rule.ShouldNotDependsOn != nil {
			if len(rule.ShouldNotDependsOn.External)+len(rule.ShouldNotDependsOn.Internal)+len(rule.ShouldNotDependsOn.Standard) == 0 {
				return fmt.Errorf("dependencies rule - ShouldNotDependsOn needs at least one of 'External', 'Internal' or 'Standard'")
			}
		}
		if rule.ShouldOnlyDependsOn != nil {
			if len(rule.ShouldOnlyDependsOn.External)+len(rule.ShouldOnlyDependsOn.Internal)+len(rule.ShouldOnlyDependsOn.Standard) == 0 {
				return fmt.Errorf("dependencies rule - ShouldOnlyDependsOn needs at least one of 'External', 'Internal' or 'Standard'")
			}
		}
	}
	return nil
}

func validateFunctionRules(rules []*config.FunctionsRule) error {
	for _, rule := range rules {
		if rule.Package == "" {
			return fmt.Errorf("function rule - empty package")
		}
		if rule.MaxParameters+rule.MaxLines+rule.MaxReturnValues+rule.MaxPublicFunctionPerFile == 0 {
			return fmt.Errorf("function rule - At least one criteria should be set")
		}

		if rule.MaxParameters < 0 {
			return fmt.Errorf("function rule - MaxParameters is less than zero")
		}
		if rule.MaxLines < 0 {
			return fmt.Errorf("function rule - MaxLines is less than zero")
		}
		if rule.MaxReturnValues < 0 {
			return fmt.Errorf("function rule - MaxReturnValues is less than zero")
		}
		if rule.MaxPublicFunctionPerFile < 0 {
			return fmt.Errorf("function rule - MaxPublicFunctionPerFile is less than zero")
		}
	}
	return nil
}

func validateContentRules(rules []*config.ContentsRule) error {
	for _, rule := range rules {
		if rule.Package == "" {
			return fmt.Errorf("content rule - empty package")
		}
		if countTrueValues(rule) == 0 {
			return fmt.Errorf("content rule - At least one criteria should be set")
		}
		if rule.ShouldOnlyContainFunctions && countTrueValues(rule) > 1 {
			return fmt.Errorf("content rule - if ShouldOnlyContainFunctions is set, then it should be the only parameter")
		}
		if rule.ShouldOnlyContainStructs && countTrueValues(rule) > 1 {
			return fmt.Errorf("content rule - if ShouldOnlyContainStructs is set, then it should be the only parameter")
		}
		if rule.ShouldOnlyContainMethods && countTrueValues(rule) > 1 {
			return fmt.Errorf("content rule - if ShouldOnlyContainMethods is set, then it should be the only parameter")
		}
		if rule.ShouldOnlyContainInterfaces && countTrueValues(rule) > 1 {
			return fmt.Errorf("content rule - if ShouldOnlyContainInterfaces is set, then it should be the only parameter")
		}
	}
	return nil
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

func trueValues(v ...bool) int32 {
	var counter int32
	for _, it := range v {
		if it {
			counter++
		}
	}
	return counter
}
