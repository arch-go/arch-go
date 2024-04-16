package reports

import (
	"github.com/fdaines/arch-go/internal/reports/model"
	"github.com/fdaines/arch-go/pkg/config"
	"github.com/fdaines/arch-go/pkg/verifications"
)

func resolveCompliance(r *verifications.Result, c config.Config) *model.ThresholdSummary {
	passesVerifications, totalVerifications := resolveTotals(r)

	rate := 0
	if totalVerifications > 0 {
		rate = (100 * passesVerifications) / totalVerifications
	}

	threshold := 0
	if c.Threshold != nil && c.Threshold.Compliance != nil {
		threshold = *c.Threshold.Compliance
	}

	status := passStatus
	var violations []string
	if rate < threshold {
		status = failStatus
		violations = append(violations, "")
	}

	return &model.ThresholdSummary{
		Rate:       rate,
		Threshold:  threshold,
		Status:     status,
		Violations: violations,
	}
}

func resolveTotals(r *verifications.Result) (int, int) {
	total := 0
	passes := 0
	if r.DependenciesRuleResult != nil {
		for _, dr := range r.DependenciesRuleResult.Results {
			if dr.Passes {
				passes++
			}
			total++
		}
	}
	if r.FunctionsRuleResult != nil {
		for _, dr := range r.FunctionsRuleResult.Results {
			if dr.Passes {
				passes++
			}
			total++
		}
	}
	if r.ContentsRuleResult != nil {
		for _, dr := range r.ContentsRuleResult.Results {
			if dr.Passes {
				passes++
			}
			total++
		}
	}
	if r.NamingRuleResult != nil {
		for _, dr := range r.NamingRuleResult.Results {
			if dr.Passes {
				passes++
			}
			total++
		}
	}
	return passes, total
}
