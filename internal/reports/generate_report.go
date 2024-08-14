package reports

import (
	"github.com/arch-go/arch-go/api"
	"github.com/arch-go/arch-go/api/configuration"
	"github.com/arch-go/arch-go/internal/common"
	"github.com/arch-go/arch-go/internal/model"
	reportModel "github.com/arch-go/arch-go/internal/reports/model"
	"github.com/arch-go/arch-go/internal/reports/utils"
)

func GenerateReport(result *api.Result, moduleInfo model.ModuleInfo, config configuration.Config) *reportModel.Report {
	compliance := resolveCompliance(result, config)
	coverage := resolveCoverage(result, moduleInfo, config)
	details := resolveReportDetails(result)
	total, passed, failed := retrieveTotals(details)

	return &reportModel.Report{
		ArchGoVersion: common.Version,
		Summary: &reportModel.ReportSummary{
			Pass:                utils.ResolveGlobalStatus(compliance, coverage),
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

	if len(moduleInfo.Packages) != 0 {
		coverageInfo = make([]reportModel.CoverageInfo, len(moduleInfo.Packages))
	}

	for i, pkg := range moduleInfo.Packages {
		cr := countContentsRulesVerifications(pkg.Path, result)
		dr := countDependenciesRulesVerifications(pkg.Path, result)
		fr := countFunctionsRulesVerifications(pkg.Path, result)
		nr := countNamingRulesVerifications(pkg.Path, result)

		coverageInfo[i] = reportModel.CoverageInfo{
			Package:           pkg.Path,
			ContentsRules:     cr,
			DependenciesRules: dr,
			FunctionsRules:    fr,
			NamingRules:       nr,
			Covered:           cr+dr+fr+nr > 0,
		}
	}

	return coverageInfo
}

func retrieveTotals(details *reportModel.ReportDetails) (int, int, int) {
	total := details.DependenciesVerificationDetails.Total +
		details.FunctionsVerificationDetails.Total +
		details.ContentsVerificationDetails.Total +
		details.NamingVerificationDetails.Total
	passed := details.DependenciesVerificationDetails.Passed +
		details.FunctionsVerificationDetails.Passed +
		details.ContentsVerificationDetails.Passed +
		details.NamingVerificationDetails.Passed
	failed := details.DependenciesVerificationDetails.Failed +
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
