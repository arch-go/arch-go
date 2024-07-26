package console

import (
	"fmt"
	"io"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/fdaines/arch-go/internal/reports/model"
)

func GenerateConsoleReport(report *model.Report, outputMirror io.Writer) {
	tw := table.NewWriter()
	tw.SetOutputMirror(outputMirror)

	appendSummary(tw, report)

	if report.Summary != nil && report.Summary.ComplianceThreshold != nil {
		appendFooter(tw, "Compliance Rate", report.Summary.ComplianceThreshold)
	}

	if report.Summary != nil && report.Summary.CoverageThreshold != nil {
		appendFooter(tw, "Coverage Rate", report.Summary.CoverageThreshold)
	}

	tw.Render()
}

func appendFooter(tw table.Writer, title string, threshold *model.ThresholdSummary) {
	rowConfig := table.RowConfig{
		AutoMerge:      true,
		AutoMergeAlign: text.AlignLeft,
	}
	complianceDetails := fmt.Sprintf("%3v%% [%s]", threshold.Rate, threshold.Status)

	tw.AppendFooter(table.Row{"", title, complianceDetails, complianceDetails, complianceDetails}, rowConfig)
}

func appendSummary(t table.Writer, report *model.Report) {
	if report.Details != nil {
		t.AppendHeader(table.Row{"#", "Rule Type", "Total", "Passed", "Failed"})
		appendSummaryRow(t, 1, "Dependencies Rules", report.Details.DependenciesVerificationDetails)
		appendSummaryRow(t, 2, "Functions Rules", report.Details.FunctionsVerificationDetails)
		appendSummaryRow(t, 3, "Contents Rules", report.Details.ContentsVerificationDetails)
		appendSummaryRow(t, 4, "Naming Rules", report.Details.NamingVerificationDetails)
	}
}

func appendSummaryRow(t table.Writer, idx int, title string, d model.Verification) {
	t.AppendRow(table.Row{idx, title, d.Total, d.Passed, d.Failed})
}
