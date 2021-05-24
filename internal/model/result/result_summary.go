package result

import "github.com/fdaines/arch-go/internal/utils/output"

type ResultSummary struct {
	rules   int
	success int
	failed  int
}

func (r *ResultSummary) Status() bool {
	return r.failed == 0
}

func (r *ResultSummary) Print() {
	output.Print("--------------------------------------")
	output.Printf("Total Rules: \t%d\n", r.rules)
	output.Printf("Succeeded: \t%d\n", r.success)
	output.Printf("Failed: \t%d\n", r.failed)
	output.Print("--------------------------------------")
}
