package console

import (
	"fmt"
	"github.com/fdaines/arch-go/internal/model/result"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"sort"
)

func GenerateConsoleReport(summary result.RulesSummary) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	header := table.Row{"#", "Rule Type", "Total", "Succeeded", "Failed"}
	t.AppendHeader(header)
	idx := 1
	keys := make([]string, 0, len(summary.Details))
	for k := range summary.Details {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		row := table.Row{idx, k, summary.Details[k].Total, summary.Details[k].Succeeded, summary.Details[k].Failed}
		t.AppendRow(row)
		idx++
	}

	if summary.ComplianceThreshold != nil {
		rowConfig := table.RowConfig{
			AutoMerge:      true,
			AutoMergeAlign: text.AlignLeft,
		}
		complianceDetails := fmt.Sprintf("%v%% [%s]",
			summary.ComplianceThreshold.Rate,
			summary.ComplianceThreshold.Status)
		t.AppendFooter(table.Row{"", "Compliance Rate", complianceDetails, complianceDetails, complianceDetails}, rowConfig)
	}

	t.Render()
}
