package reports

import (
	"github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/pkg/config"
	"github.com/fdaines/arch-go/pkg/verifications"
)

func GenerateReport(result *verifications.Result, moduleInfo model.ModuleInfo, config config.Config) *Report {
	compliance := resolveCompliance(result, config)
	coverage := resolveCoverage(result, moduleInfo, config)
	details := resolveReportDetails(result)
	total, passed, failed := retrieveTotals(details)
	return &Report{
		Summary: &ReportSummary{
			Status:              resolveGlobalStatus(compliance, coverage),
			Total:               total,
			Passed:              passed,
			Failed:              failed,
			Time:                result.Time,
			Duration:            result.Duration,
			ComplianceThreshold: compliance,
			CoverageThreshold:   coverage,
		},
		Details: details,
	}
}

func retrieveTotals(details *ReportDetails) (int, int, int) {
	total := details.DependenciesVerificationDetails.Total + details.FunctionsVerificationDetails.Total + details.ContentsVerificationDetails.Total + details.NamingVerificationDetails.Total
	passed := details.DependenciesVerificationDetails.Passed + details.FunctionsVerificationDetails.Passed + details.ContentsVerificationDetails.Passed + details.NamingVerificationDetails.Passed
	failed := details.DependenciesVerificationDetails.Failed + details.FunctionsVerificationDetails.Failed + details.ContentsVerificationDetails.Failed + details.NamingVerificationDetails.Failed

	return total, passed, failed
}
