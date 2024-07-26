package reports

import (
	"fmt"
	"io"

	"github.com/fatih/color"

	"github.com/fdaines/arch-go/internal/common"
	"github.com/fdaines/arch-go/internal/reports/console"
	"github.com/fdaines/arch-go/internal/reports/html"
	"github.com/fdaines/arch-go/internal/reports/model"
)

func DisplayResult(report *model.Report, output io.Writer) {
	displayRules(report, output)

	if common.HTML {
		html.GenerateHTMLReport(report, output)
	} else {
		console.GenerateConsoleReport(report, output)
	}

	displaySummary(report.Summary, output)
}

func displaySummary(summary *model.ReportSummary, output io.Writer) {
	const lineSeparator = "--------------------------------------\n"

	color.Output = output

	fmt.Fprintf(output, lineSeparator)
	fmt.Fprintf(output, "\tExecution Summary\n")
	fmt.Fprintf(output, lineSeparator)
	fmt.Fprintf(output, "Total Rules: \t%d\n", summary.Total)
	fmt.Fprintf(output, "Succeeded: \t%d\n", summary.Passed)
	fmt.Fprintf(output, "Failed: \t%d\n", summary.Failed)
	fmt.Fprintf(output, lineSeparator)

	if summary.ComplianceThreshold != nil {
		complianceSummary := fmt.Sprintf("Compliance: %8d%% (%s)\n",
			summary.ComplianceThreshold.Rate,
			summary.ComplianceThreshold.Status)
		if summary.ComplianceThreshold.Status == passStatus {
			color.Green(complianceSummary)
		} else {
			color.Red(complianceSummary)
		}
	}
	if summary.CoverageThreshold != nil {
		complianceSummary := fmt.Sprintf("Coverage: %10d%% (%s)\n",
			summary.CoverageThreshold.Rate,
			summary.CoverageThreshold.Status)
		if summary.CoverageThreshold.Status == passStatus {
			color.Green(complianceSummary)
		} else {
			color.Red(complianceSummary)
		}
	}
}
