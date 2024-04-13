package functions

import (
	"fmt"
	"strings"

	"github.com/fdaines/arch-go/pkg/config"
)

func resolveDescription(r config.FunctionsRule) string {
	var ruleDescriptions []string
	if r.MaxParameters != nil {
		ruleDescriptions = append(ruleDescriptions, fmt.Sprintf("'at most %d parameters'", *r.MaxParameters))
	}
	if r.MaxReturnValues != nil {
		ruleDescriptions = append(ruleDescriptions, fmt.Sprintf("'at most %d return values'", *r.MaxReturnValues))
	}
	if r.MaxLines != nil {
		ruleDescriptions = append(ruleDescriptions, fmt.Sprintf("'at most %d lines'", *r.MaxLines))
	}
	if r.MaxPublicFunctionPerFile != nil {
		ruleDescriptions = append(ruleDescriptions, fmt.Sprintf("'no more than %d public functions per file'", *r.MaxPublicFunctionPerFile))
	}
	return fmt.Sprintf(
		"Functions in packages matching pattern '%s' should have [%s]",
		r.Package,
		strings.Join(ruleDescriptions, ","),
	)
}
