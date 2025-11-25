package reports

import (
	"github.com/arch-go/arch-go/v2/api"
	"github.com/arch-go/arch-go/v2/api/configuration"
	"github.com/arch-go/arch-go/v2/internal/reports/model"
	"github.com/arch-go/arch-go/v2/internal/utils/values"
)

func resolveCompliance(result *api.Result, conf configuration.Config) *model.ThresholdSummary {
	var violations []string

	passesVerifications, totalVerifications := resolveTotals(result)
	rate := 0
	pass := true

	if totalVerifications > 0 {
		rate = (100 * passesVerifications) / totalVerifications
	}

	threshold := values.GetIntRef(0)
	if conf.Threshold != nil && conf.Threshold.Compliance != nil {
		threshold = conf.Threshold.Compliance
	}

	if threshold != nil && rate < *threshold {
		pass = false

		violations = append(violations, "")
	}

	return &model.ThresholdSummary{
		Rate:       rate,
		Threshold:  threshold,
		Pass:       pass,
		Violations: violations,
	}
}

func resolveTotals(result *api.Result) (int, int) {
	total := 0
	passes := 0

	countDependenciesRuleResults(result, &passes, &total)
	countFunctionsRuleResults(result, &passes, &total)
	countContentsRuleResults(result, &passes, &total)
	countNamingRuleResults(result, &passes, &total)

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
