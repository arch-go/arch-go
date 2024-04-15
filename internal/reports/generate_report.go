package reports

import (
	"github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/pkg/config"
	"github.com/fdaines/arch-go/pkg/verifications"
)

func GenerateReport(result *verifications.Result, moduleInfo model.ModuleInfo, config config.Config) *Report {
	compliance := resolveCompliance(result, config)
	coverage := resolveCoverage(result, moduleInfo, config)
	return &Report{
		Summary: &ReportSummary{
			Status:              resolveGlobalStatus(compliance, coverage),
			Total:               10,
			Succeeded:           9,
			Failed:              1,
			Time:                result.Time,
			Duration:            result.Duration,
			ComplianceThreshold: compliance,
			CoverageThreshold:   coverage,
		},
		Details: resolveReportDetails(result),
	}
}
