package result

import (
	"github.com/fdaines/arch-go/internal/config"
	"github.com/fdaines/arch-go/internal/impl/model"
	"github.com/fdaines/arch-go/internal/utils/output"
)

func ResolveRulesSummary(verifications []model.RuleVerification, configuration *config.Config) RulesSummary {
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

	if configuration.Threshold != nil && configuration.Threshold.Compliance != nil {
		summary.ComplianceThreshold = &ThresholdSummary{
			Rate:      int(100 * summary.Succeeded / summary.Total),
			Threshold: *configuration.Threshold.Compliance,
			Status:    "Fail",
		}
		if summary.ComplianceThreshold.Rate >= summary.ComplianceThreshold.Threshold {
			summary.ComplianceThreshold.Status = "Pass"
		}
	}

	summary.Status = summary.ComplianceThreshold.Status == "Pass"

	return summary
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
	Rate      int
	Threshold int
	Status    string
}

type RulesSummary struct {
	Total               int32
	Succeeded           int32
	Failed              int32
	Status              bool
	Details             map[string]RulesSummaryDetail
	ComplianceThreshold *ThresholdSummary
}

type RulesSummaryDetail struct {
	Total     int32
	Succeeded int32
	Failed    int32
}

func (s RulesSummary) Print() {
	output.Print("--------------------------------------\n")
	output.Printf("Total Rules: \t%d\n", s.Total)
	output.Printf("Succeeded: \t%d\n", s.Succeeded)
	output.Printf("Failed: \t%d\n", s.Failed)
	output.Print("--------------------------------------\n")
	output.Printf("Compliance: \t%d%% (%s)\n", s.ComplianceThreshold.Rate, s.ComplianceThreshold.Status)
}
