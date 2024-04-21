package validators

import (
	"fmt"

	"github.com/fdaines/arch-go/pkg/config"
)

func ValidateConfiguration(configuration *config.Config) error {
	if configuration == nil {
		return fmt.Errorf("configuration file not found")
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

func checkRulesQuantity(c *config.Config) error {
	fmt.Printf("Config: %+v\n", c)
	total := len(c.ContentRules) + len(c.DependenciesRules) + len(c.FunctionsRules) + len(c.NamingRules)
	if total == 0 {
		return fmt.Errorf("configuration file should have at least one rule")
	}
	return nil
}
