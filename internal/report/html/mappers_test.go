package html

import (
	"github.com/fdaines/arch-go/internal/model/result"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMappers(t *testing.T) {
	t.Run("Calls resolveRulesSummary function", func(t *testing.T) {
		report := sampleReport()
		htmlReport := &HtmlReport{}

		resolveRulesSummary(report, htmlReport)

		assert.Equal(t, 4, len(htmlReport.RulesSummary))
		assert.Equal(t, "Content Rule", htmlReport.RulesSummary[0].Type)
		assert.Equal(t, "Dependency Rule", htmlReport.RulesSummary[1].Type)
		assert.Equal(t, "Function Rule", htmlReport.RulesSummary[2].Type)
		assert.Equal(t, "Naming Rule", htmlReport.RulesSummary[3].Type)
		assert.Equal(t, 2, htmlReport.RulesSummary[1].Succeeded)
		assert.Equal(t, 1, htmlReport.RulesSummary[1].Failed)
		assert.Equal(t, 3, htmlReport.RulesSummary[1].Total)

	})

	t.Run("Calls resolveCoverageResults function", func(t *testing.T) {
		report := sampleReport()
		htmlReport := &HtmlReport{}

		resolveCoverageResults(report, htmlReport)

		assert.Equal(t, 2, len(htmlReport.UncoveredPackages))
		assert.Equal(t, 34, htmlReport.CoverageResult.Rate)
		assert.Equal(t, 98, htmlReport.CoverageResult.Threshold)
		assert.Equal(t, 25, htmlReport.CoverageResult.Total)
		assert.Equal(t, 2, htmlReport.CoverageResult.Uncovered)
		assert.Equal(t, "red", htmlReport.CoverageResult.Color)
	})

	t.Run("Calls resolveComplianceResults function", func(t *testing.T) {
		report := sampleReport()
		htmlReport := &HtmlReport{}

		resolveComplianceResults(report, htmlReport)

		assert.Equal(t, 89, htmlReport.ComplianceResult.Rate)
		assert.Equal(t, 80, htmlReport.ComplianceResult.Threshold)
		assert.Equal(t, 10, htmlReport.ComplianceResult.Total)
		assert.Equal(t, "green", htmlReport.ComplianceResult.Color)
	})
}

func sampleReport() result.Report {
	return result.Report{
		TotalPackages: 25,
		Summary: result.RulesSummary{
			Total:     10,
			Succeeded: 8,
			Failed:    2,
			Status:    true,
			Details: map[string]result.RulesSummaryDetail{
				"Dependency Rule": result.RulesSummaryDetail{
					Total:     3,
					Succeeded: 2,
					Failed:    1,
				},
			},
			ComplianceThreshold: &result.ThresholdSummary{
				Rate:      89,
				Threshold: 80,
				Status:    "Pass",
			},
			CoverageThreshold: &result.ThresholdSummary{
				Rate:       34,
				Threshold:  98,
				Status:     "Fail",
				Violations: []string{"foo", "bar"},
			},
		},
	}
}
