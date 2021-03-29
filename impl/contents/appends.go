package contents

import (
	"fmt"
	"github.com/fdaines/arch-go/impl/model"
)

func appendError(results []*model.ContentsRuleResult, p string, s string) []*model.ContentsRuleResult {
	return append(results, &model.ContentsRuleResult{
		Description: fmt.Sprintf("Package '%s' %s", p, s),
		Passes:      false,
	})
}

func appendSuccess(results []*model.ContentsRuleResult, p string, s string) []*model.ContentsRuleResult {
	return append(results, &model.ContentsRuleResult{
		Description: fmt.Sprintf("Package '%s' %s", p, s),
		Passes:      true,
	})
}

