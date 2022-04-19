package result

import (
	"github.com/fdaines/arch-go/internal/impl/model"
	"github.com/fdaines/arch-go/internal/utils/output"
)

func ResolveRulesSummary(verifications []model.RuleVerification) RulesSummary {
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

type RulesSummary struct {
	Total int32
	Succeeded int32
	Failed int32
	Details map[string]RulesSummaryDetail
}

type RulesSummaryDetail struct {
	Total int32
	Succeeded int32
	Failed int32
}

func (s RulesSummary) Print() {
	output.Print("--------------------------------------\n")
	output.Printf("Total Rules: \t%d\n", s.Total)
	output.Printf("Succeeded: \t%d\n", s.Succeeded)
	output.Printf("Failed: \t%d\n", s.Failed)
	output.Print("--------------------------------------\n")
}
