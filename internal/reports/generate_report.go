package reports

import (
	"github.com/fdaines/arch-go/api"
	"github.com/fdaines/arch-go/api/configuration"
	"github.com/fdaines/arch-go/internal/common"
	"github.com/fdaines/arch-go/internal/model"
	reportModel "github.com/fdaines/arch-go/internal/reports/model"
)

func GenerateReport(result *api.Result, moduleInfo model.ModuleInfo, config configuration.Config) *reportModel.Report {
	compliance := resolveCompliance(result, config)
	coverage := resolveCoverage(result, moduleInfo, config)
	details := resolveReportDetails(result)
	total, passed, failed := retrieveTotals(details)

	return &reportModel.Report{
		ArchGoVersion: common.Version,
		Summary: &reportModel.ReportSummary{
			Status:              resolveGlobalStatus(compliance, coverage),
			Total:               total,
			Passed:              passed,
			Failed:              failed,
			Time:                result.Time,
			Duration:            result.Duration,
			ComplianceThreshold: compliance,
			CoverageThreshold:   coverage,
		},
		Details:      details,
		CoverageInfo: generateCoverageInfo(moduleInfo, result),
	}
}

func generateCoverageInfo(moduleInfo model.ModuleInfo, result *api.Result) []reportModel.CoverageInfo {
	var coverageInfo []reportModel.CoverageInfo

	for _, pkg := range moduleInfo.Packages {
		cr := countContentsRulesVerifications(pkg.Path, result)
		dr := countDependenciesRulesVerifications(pkg.Path, result)
		fr := countFunctionsRulesVerifications(pkg.Path, result)
		nr := countNamingRulesVerifications(pkg.Path, result)
		status := "NO"

		if cr+dr+fr+nr > 0 {
			status = "YES"
		}

		coverageInfo = append(coverageInfo, reportModel.CoverageInfo{
			Package:           pkg.Path,
			ContensRules:      cr,
			DependenciesRules: dr,
			FunctionsRules:    fr,
			NamingRules:       nr,
			Status:            status,
		})
	}

	return coverageInfo
}

func retrieveTotals(details *reportModel.ReportDetails) (int, int, int) {
	total :=
		details.DependenciesVerificationDetails.Total +
			details.FunctionsVerificationDetails.Total +
			details.ContentsVerificationDetails.Total +
			details.NamingVerificationDetails.Total
	passed :=
		details.DependenciesVerificationDetails.Passed +
			details.FunctionsVerificationDetails.Passed +
			details.ContentsVerificationDetails.Passed +
			details.NamingVerificationDetails.Passed
	failed :=
		details.DependenciesVerificationDetails.Failed +
			details.FunctionsVerificationDetails.Failed +
			details.ContentsVerificationDetails.Failed +
			details.NamingVerificationDetails.Failed

	return total, passed, failed
}

func countContentsRulesVerifications(pkg string, result *api.Result) int {
	var total int

	if result.ContentsRuleResult != nil {
		for _, r := range result.ContentsRuleResult.Results {
			for _, v := range r.Verifications {
				if v.Package == pkg {
					total++
				}
			}
		}
	}

	return total
}

func countDependenciesRulesVerifications(pkg string, result *api.Result) int {
	var total int

	if result.DependenciesRuleResult != nil {
		for _, r := range result.DependenciesRuleResult.Results {
			for _, v := range r.Verifications {
				if v.Package == pkg {
					total++
				}
			}
		}
	}

	return total
}

func countFunctionsRulesVerifications(pkg string, result *api.Result) int {
	var total int

	if result.FunctionsRuleResult != nil {
		for _, r := range result.FunctionsRuleResult.Results {
			for _, v := range r.Verifications {
				if v.Package == pkg {
					total++
				}
			}
		}
	}

	return total
}

func countNamingRulesVerifications(pkg string, result *api.Result) int {
	var total int

	if result.NamingRuleResult != nil {
		for _, r := range result.NamingRuleResult.Results {
			for _, v := range r.Verifications {
				if v.Package == pkg {
					total++
				}
			}
		}
	}

	return total
}
