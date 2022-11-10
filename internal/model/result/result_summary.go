package result

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/fdaines/arch-go/internal/config"
	"github.com/fdaines/arch-go/internal/impl/model"
	baseModel "github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/internal/utils/output"
	"github.com/fdaines/arch-go/internal/utils/packages"
)

func ResolveRulesSummary(pkgs []*baseModel.PackageInfo, verifications []model.RuleVerification, configuration *config.Config) RulesSummary {
	summary := NewRulesSummary()
	for _, v := range verifications {
		current := summary.Details[v.Type()]
		if v.Status() {
			current.Succeeded++
			summary.Succeeded++
		} else {
			current.Failed++
			summary.Failed++
		}
		current.Total++
		summary.Total++

		summary.Details[v.Type()] = current
	}

	resolveComplianceThreshold(&summary, configuration)
	resolveCoverageThreshold(&summary, configuration, verifications, pkgs)

	summary.Status = summary.ComplianceThreshold.Status == "Pass" &&
		summary.CoverageThreshold.Status == "Pass"

	return summary
}

func resolveComplianceThreshold(summary *RulesSummary, configuration *config.Config) {
	if configuration.Threshold == nil || configuration.Threshold.Compliance == nil {
		return
	}

	summary.ComplianceThreshold = &ThresholdSummary{
		Rate:      int(100 * summary.Succeeded / summary.Total),
		Threshold: *configuration.Threshold.Compliance,
		Status:    "Fail",
	}
	if summary.ComplianceThreshold.Rate >= summary.ComplianceThreshold.Threshold {
		summary.ComplianceThreshold.Status = "Pass"
	}
}

func resolveCoverageThreshold(summary *RulesSummary, configuration *config.Config, verifications []model.RuleVerification, pkgs []*baseModel.PackageInfo) {
	if configuration.Threshold == nil || configuration.Threshold.Coverage == nil {
		return
	}

	coveredPackages := packages.ResolveCoveredPackages(verifications)
	modulePackages := packages.ResolveTotalPackages(pkgs)
	modulePackagesQuantity := len(modulePackages)
	uncoveredPackages := packages.ResolveUncoveredPackages(modulePackages, coveredPackages)

	summary.CoverageThreshold = &ThresholdSummary{
		Rate:       100 * len(coveredPackages) / modulePackagesQuantity,
		Threshold:  *configuration.Threshold.Coverage,
		Status:     "Fail",
		Violations: uncoveredPackages,
	}
	if summary.CoverageThreshold.Rate >= summary.CoverageThreshold.Threshold {
		summary.CoverageThreshold.Status = "Pass"
	}
}

func NewRulesSummary() RulesSummary {
	summary := RulesSummary{}
	summary.Details = make(map[string]RulesSummaryDetail)
	summary.Details["DependenciesRule"] = RulesSummaryDetail{}
	summary.Details["FunctionsRule"] = RulesSummaryDetail{}
	summary.Details["ContentRule"] = RulesSummaryDetail{}
	summary.Details["CycleRule"] = RulesSummaryDetail{}
	summary.Details["NamingRule"] = RulesSummaryDetail{}

	return summary
}

type ThresholdSummary struct {
	Rate       int
	Threshold  int
	Status     string
	Violations []string
}

type RulesSummary struct {
	Total               int32
	Succeeded           int32
	Failed              int32
	Status              bool
	Details             map[string]RulesSummaryDetail
	ComplianceThreshold *ThresholdSummary
	CoverageThreshold   *ThresholdSummary
}

type RulesSummaryDetail struct {
	Total     int32
	Succeeded int32
	Failed    int32
}

func (s RulesSummary) Print() {
	const lineSeparator = "--------------------------------------\n"
	output.Print(lineSeparator)
	output.Print("\tExecution Summary\n")
	output.Print(lineSeparator)
	output.Printf("Total Rules: \t%d\n", s.Total)
	output.Printf("Succeeded: \t%d\n", s.Succeeded)
	output.Printf("Failed: \t%d\n", s.Failed)
	output.Print(lineSeparator)
	if s.ComplianceThreshold != nil {
		complianceSummary := fmt.Sprintf("Compliance: %8d%% (%s)\n", s.ComplianceThreshold.Rate, s.ComplianceThreshold.Status)
		if s.ComplianceThreshold.Status == "Pass" {
			color.Green(complianceSummary)
		} else {
			color.Red(complianceSummary)
		}
	}
	if s.CoverageThreshold != nil {
		complianceSummary := fmt.Sprintf("Coverage: %10d%% (%s)\n", s.CoverageThreshold.Rate, s.CoverageThreshold.Status)
		if s.CoverageThreshold.Status == "Pass" {
			color.Green(complianceSummary)
		} else {
			color.Red(complianceSummary)
		}
	}
}
