package console

import (
	"github.com/fdaines/arch-go/internal/model/result"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

func GenerateConsoleReport(summary result.RulesSummary) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	header := table.Row{"#", "Rule Type", "Total", "Succeeded", "Failed"}
	t.AppendHeader(header)
	idx := 1
	for k, v := range summary.Details {
		row := table.Row{idx, k, v.Total, v.Succeeded, v.Failed}
		t.AppendRow(row)
		idx++
	}
	t.Render()
}