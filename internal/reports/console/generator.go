package console

import (
	"fmt"
	"os"
	"sort"

	"github.com/fdaines/arch-go/old/model/result"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func GenerateConsoleReport(report result.Report) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	header := table.Row{"#", "Rule Type", "Total", "Succeeded", "Failed"}
	t.AppendHeader(header)
	idx := 1
	keys := make([]string, 0, len(report.Summary.Details))
	for k := range report.Summary.Details {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		row := table.Row{idx, k, report.Summary.Details[k].Total, report.Summary.Details[k].Succeeded, report.Summary.Details[k].Failed}
		t.AppendRow(row)
		idx++
	}

	if report.Summary.ComplianceThreshold != nil {
		rowConfig := table.RowConfig{
			AutoMerge:      true,
			AutoMergeAlign: text.AlignLeft,
		}
		complianceDetails := fmt.Sprintf("%3v%% [%s]",
			report.Summary.ComplianceThreshold.Rate,
			report.Summary.ComplianceThreshold.Status)
		t.AppendFooter(table.Row{"", "Compliance Rate", complianceDetails, complianceDetails, complianceDetails}, rowConfig)
		coverageDetails := fmt.Sprintf("%3v%% [%s]",
			report.Summary.CoverageThreshold.Rate,
			report.Summary.CoverageThreshold.Status)
		t.AppendFooter(table.Row{"", "Coverage Rate", coverageDetails, coverageDetails, coverageDetails}, rowConfig)
	}

	t.Render()
}
