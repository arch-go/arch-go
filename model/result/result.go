package result

import (
	"github.com/fatih/color"
	"github.com/fdaines/arch-go/utils/output"
	"strings"
)

type Result struct {
	DependenciesRulesResults []*DependenciesRuleResult
	ContentsRuleResults      []*ContentsRuleResult
	CyclesRuleResults        []*CyclesRuleResult
	FunctionsRulesResults    []*FunctionsRuleResult
}

func (r *Result) Print() *ResultSummary {
	summary := &ResultSummary{}
	output.Print("--------------------------------------")
	r.printDependenciesResults(summary)
	r.printContentResults(summary)
	r.printCyclerResults(summary)
	r.printFunctionResults(summary)

	return summary
}

func (r *Result) printFunctionResults(summary *ResultSummary) {
	for _, fr := range r.FunctionsRulesResults {
		summary.rules++
		if fr.Passes {
			summary.success++
			color.Green("[PASS] - %s\n", fr.Description)
		} else {
			summary.failed++
			color.Red("[FAIL] - %s\n", fr.Description)
			for _, fd := range fr.Failures {
				for _, str := range fd.Details {
					color.Red("\t%s\n", str)
				}
			}
		}
	}
}

func (r *Result) printCyclerResults(summary *ResultSummary) {
	for _, cr := range r.CyclesRuleResults {
		summary.rules++
		if cr.Passes {
			summary.success++
			color.Green("[PASS] - %s\n", cr.Description)
		} else {
			summary.failed++
			color.Red("[FAIL] - %s\n", cr.Description)
			for _, fd := range cr.Failures {
				color.Red("\tPackage '%s' fails\n", fd.Package)
				for idx, str := range fd.Details {
					spaces := strings.Repeat(" ", idx+1)
					color.Red("\t%s + imports %s\n", spaces, str)
				}
			}
		}
	}

}

func (r *Result) printContentResults(summary *ResultSummary) {
	for _, cr := range r.ContentsRuleResults {
		summary.rules++
		if cr.Passes {
			summary.success++
			color.Green("[PASS] - %s\n", cr.Description)
		} else {
			summary.failed++
			color.Red("[FAIL] - %s\n", cr.Description)
		}
	}
}

func (r *Result) printDependenciesResults(summary *ResultSummary) {
	for _, dr := range r.DependenciesRulesResults {
		summary.rules++
		if dr.Passes {
			summary.success++
			color.Green("[PASS] - %s\n", dr.Description)
		} else {
			summary.failed++
			color.Red("[FAIL] - %s\n", dr.Description)
			for _, fd := range dr.Failures {
				color.Red("\tPackage '%s' fails\n", fd.Package)
				for _, str := range fd.Details {
					color.Red("\t\t%s\n", str)
				}
			}
		}
	}
}