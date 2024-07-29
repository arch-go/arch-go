package functions

import (
	"fmt"
	"strings"

	"github.com/fdaines/arch-go/api/configuration"
)

func resolveDescription(rule configuration.FunctionsRule) string {
	var ruleDescriptions []string

	if rule.MaxParameters != nil {
		ruleDescriptions = append(ruleDescriptions,
			fmt.Sprintf("'at most %d parameters'", *rule.MaxParameters))
	}

	if rule.MaxReturnValues != nil {
		ruleDescriptions = append(ruleDescriptions,
			fmt.Sprintf("'at most %d return values'", *rule.MaxReturnValues))
	}

	if rule.MaxLines != nil {
		ruleDescriptions = append(ruleDescriptions,
			fmt.Sprintf("'at most %d lines'", *rule.MaxLines))
	}

	if rule.MaxPublicFunctionPerFile != nil {
		ruleDescriptions = append(ruleDescriptions,
			fmt.Sprintf("'no more than %d public functions per file'", *rule.MaxPublicFunctionPerFile))
	}

	return fmt.Sprintf(
		"Functions in packages matching pattern '%s' should have [%s]",
		rule.Package,
		strings.Join(ruleDescriptions, ","),
	)
}
