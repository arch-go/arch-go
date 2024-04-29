package reports

import (
	"github.com/fdaines/arch-go/api"
	"github.com/fdaines/arch-go/api/configuration"
	"github.com/fdaines/arch-go/internal/reports/model"
)

func resolveCompliance(r *api.Result, c configuration.Config) *model.ThresholdSummary {
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

func resolveTotals(r *api.Result) (int, int) {
	total := 0
	passes := 0
	countDependenciesRuleResults(r, &passes, &total)
	countFunctionsRuleResults(r, &passes, &total)
	countContentsRuleResults(r, &passes, &total)
	countNamingRuleResults(r, &passes, &total)
	return passes, total
}

func countDependenciesRuleResults(r *api.Result, passes *int, total *int) {
	if r.DependenciesRuleResult != nil {
		for _, dr := range r.DependenciesRuleResult.Results {
			if dr.Passes {
				*passes++
			}
			*total++
		}
	}
}

func countFunctionsRuleResults(r *api.Result, passes *int, total *int) {
	if r.FunctionsRuleResult != nil {
		for _, dr := range r.FunctionsRuleResult.Results {
			if dr.Passes {
				*passes++
			}
			*total++
		}
	}
}

func countContentsRuleResults(r *api.Result, passes *int, total *int) {
	if r.ContentsRuleResult != nil {
		for _, dr := range r.ContentsRuleResult.Results {
			if dr.Passes {
				*passes++
			}
			*total++
		}
	}
}

func countNamingRuleResults(r *api.Result, passes *int, total *int) {
	if r.NamingRuleResult != nil {
		for _, dr := range r.NamingRuleResult.Results {
			if dr.Passes {
				*passes++
			}
			*total++
		}
	}
}
