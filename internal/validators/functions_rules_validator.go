package validators

import (
	"fmt"

	"github.com/fdaines/arch-go/api/configuration"
	"github.com/fdaines/arch-go/internal/utils/values"
)

func validateFunctionRules(rules []*configuration.FunctionsRule) error {
	for _, rule := range rules {
		if rule.Package == "" {
			return fmt.Errorf("function rule - empty package")
		}
		if countNotNil(rule.MaxParameters, rule.MaxLines, rule.MaxReturnValues, rule.MaxPublicFunctionPerFile) == 0 {
			return fmt.Errorf("function rule - At least one criteria should be set")
		}

		if values.IsLessThanZero(rule.MaxParameters) {
			return fmt.Errorf("function rule - MaxParameters is less than zero")
		}
		if values.IsLessThanZero(rule.MaxLines) {
			return fmt.Errorf("function rule - MaxLines is less than zero")
		}
		if values.IsLessThanZero(rule.MaxReturnValues) {
			return fmt.Errorf("function rule - MaxReturnValues is less than zero")
		}
		if values.IsLessThanZero(rule.MaxPublicFunctionPerFile) {
			return fmt.Errorf("function rule - MaxPublicFunctionPerFile is less than zero")
		}
	}
	return nil
}
