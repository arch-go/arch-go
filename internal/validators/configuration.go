package validators

import (
	"errors"

	"github.com/fdaines/arch-go/api/configuration"
)

func ValidateConfiguration(configuration *configuration.Config) error {
	if configuration == nil {
		return errors.New("configuration file not found")
	}

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

func checkRulesQuantity(c *configuration.Config) error {
	if countRules(c) == 0 {
		return errors.New("configuration file should have at least one rule")
	}

	return nil
}

func countRules(c *configuration.Config) int {
	return len(c.ContentRules) + len(c.DependenciesRules) + len(c.FunctionsRules) + len(c.NamingRules)
}
