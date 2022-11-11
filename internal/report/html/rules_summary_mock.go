package html

import "github.com/fdaines/arch-go/internal/model/result"

func createRulesSummaryMock(withCompliance, withCoverage bool) result.RulesSummary {
	summary := result.RulesSummary{
		Total:               200,
		Succeeded:           180,
		Failed:              20,
		Status:              true,
		Details:             map[string]result.RulesSummaryDetail{},
		ComplianceThreshold: nil,
		CoverageThreshold:   nil,
	}
	if withCompliance {
		summary.ComplianceThreshold = &result.ThresholdSummary{
			Rate:      90,
			Threshold: 90,
			Status:    "Pass",
		}
	}
	if withCoverage {
		summary.CoverageThreshold = &result.ThresholdSummary{
			Rate:      85,
			Threshold: 90,
			Status:    "Fail",
			Violations: []string{
				"foobar",
				"barfoo",
			},
		}
	}

	return summary
}
