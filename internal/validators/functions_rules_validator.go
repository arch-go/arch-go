package validators

import (
	"errors"

	"github.com/arch-go/arch-go/api/configuration"
	"github.com/arch-go/arch-go/internal/utils/values"
)

func validateFunctionRules(rules []*configuration.FunctionsRule) error {
	for _, rule := range rules {
		if rule.Package == "" {
			return errors.New("function rule - empty package")
		}

		if countNotNil(rule.MaxParameters, rule.MaxLines, rule.MaxReturnValues, rule.MaxPublicFunctionPerFile) == 0 {
			return errors.New("function rule - At least one criteria should be set")
		}

		if values.IsLessThanZero(rule.MaxParameters) {
			return errors.New("function rule - MaxParameters is less than zero")
		}

		if values.IsLessThanZero(rule.MaxLines) {
			return errors.New("function rule - MaxLines is less than zero")
		}

		if values.IsLessThanZero(rule.MaxReturnValues) {
			return errors.New("function rule - MaxReturnValues is less than zero")
		}

		if values.IsLessThanZero(rule.MaxPublicFunctionPerFile) {
			return errors.New("function rule - MaxPublicFunctionPerFile is less than zero")
		}
	}

	return nil
}
