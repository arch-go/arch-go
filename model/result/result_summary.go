package result

import "github.com/fdaines/arch-go/utils/output"

type ResultSummary struct {
	rules   int
	success int
	failed  int
	Result  bool
}

func newResultSummary(rules int, success int, failed int) *ResultSummary {
	return &ResultSummary{
		rules:   rules,
		success: success,
		failed:  failed,
		Result:  failed == 0,
	}
}

func (r *ResultSummary) Print() {
	output.Print("--------------------------------------")
	output.Printf("Total Rules: \t%d\n", r.rules)
	output.Printf("Succeeded: \t%d\n", r.success)
	output.Printf("Failed: \t%d\n", r.failed)
	output.Print("--------------------------------------")
}
