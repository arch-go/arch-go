package functions

import (
	"fmt"
	"github.com/fdaines/arch-go/model/result"
)

func appendError(results []*result.ContentsRuleResult, p string, s string) []*result.ContentsRuleResult {
	return append(results, &result.ContentsRuleResult{
		Description: fmt.Sprintf("Package '%s' %s", p, s),
		Passes:      false,
	})
}

func appendSuccess(results []*result.ContentsRuleResult, p string, s string) []*result.ContentsRuleResult {
	return append(results, &result.ContentsRuleResult{
		Description: fmt.Sprintf("Package '%s' %s", p, s),
		Passes:      true,
	})
}
