package result

import "github.com/fdaines/arch-go/internal/utils/output"

type ResultSummary struct {
	Rules     int
	Succeeded int
	Failed    int
}

func (r *ResultSummary) Status() bool {
	return r.Failed == 0
}

func (r *ResultSummary) Print() {
	output.Print("--------------------------------------")
	output.Printf("Total Rules: \t%d\n", r.Rules)
	output.Printf("Succeeded: \t%d\n", r.Succeeded)
	output.Printf("Failed: \t%d\n", r.Failed)
	output.Print("--------------------------------------")
}
