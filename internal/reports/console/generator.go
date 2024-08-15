package console

import (
	"fmt"
	"io"

	"github.com/arch-go/arch-go/internal/reports/utils"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/arch-go/arch-go/internal/reports/model"
)

func GenerateConsoleReport(report *model.Report, outputMirror io.Writer) {
	tw := table.NewWriter()
	tw.SetOutputMirror(outputMirror)

	appendSummary(tw, report)

	if report.Compliance.Threshold != nil {
		appendFooter(tw, "Compliance Rate", report.Compliance.Rate, report.Compliance.Pass)
	}

	if report.Coverage.Threshold != nil {
		appendFooter(tw, "Coverage Rate", report.Coverage.Rate, report.Coverage.Pass)
	}

	tw.Render()
}

func appendFooter(tw table.Writer, title string, rate int, pass bool) {
	rowConfig := table.RowConfig{
		AutoMerge:      true,
		AutoMergeAlign: text.AlignLeft,
	}
	complianceDetails := fmt.Sprintf("%3v%% [%s]", rate, utils.ResolveStatus(pass))

	tw.AppendFooter(table.Row{"", title, complianceDetails, complianceDetails, complianceDetails}, rowConfig)
}

func appendSummary(tw table.Writer, report *model.Report) {
	if report.Details != nil {
		tw.AppendHeader(table.Row{"#", "Rule Type", "Total", "Passed", "Failed"})
		appendSummaryRow(tw, 1, "Dependencies Rules", report.Details.DependenciesVerificationDetails)
		appendSummaryRow(tw, 2, "Functions Rules", report.Details.FunctionsVerificationDetails)
		appendSummaryRow(tw, 3, "Contents Rules", report.Details.ContentsVerificationDetails)
		appendSummaryRow(tw, 4, "Naming Rules", report.Details.NamingVerificationDetails)
	}
}

func appendSummaryRow(tw table.Writer, idx int, title string, d model.Verification) {
	tw.AppendRow(table.Row{idx, title, d.Total, d.Passed, d.Failed})
}
