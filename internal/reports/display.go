package reports

import (
	"fmt"
	"io"

	"github.com/fatih/color"

	"github.com/arch-go/arch-go/internal/common"
	"github.com/arch-go/arch-go/internal/reports/console"
	"github.com/arch-go/arch-go/internal/reports/html"
	"github.com/arch-go/arch-go/internal/reports/model"
	"github.com/arch-go/arch-go/internal/reports/utils"
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

	fmt.Fprint(output, lineSeparator)
	fmt.Fprint(output, "\tExecution Summary\n")
	fmt.Fprint(output, lineSeparator)
	fmt.Fprintf(output, "Total Rules: \t%d\n", summary.Total)
	fmt.Fprintf(output, "Succeeded: \t%d\n", summary.Passed)
	fmt.Fprintf(output, "Failed: \t%d\n", summary.Failed)
	fmt.Fprint(output, lineSeparator)

	if summary.ComplianceThreshold != nil {
		complianceSummary := fmt.Sprintf("Compliance: %8d%% (%s)\n",
			summary.ComplianceThreshold.Rate,
			utils.ResolveStatus(summary.ComplianceThreshold.Pass))
		if summary.ComplianceThreshold.Pass {
			color.Green(complianceSummary)
		} else {
			color.Red(complianceSummary)
		}
	}

	if summary.CoverageThreshold != nil {
		complianceSummary := fmt.Sprintf("Coverage: %10d%% (%s)\n",
			summary.CoverageThreshold.Rate,
			utils.ResolveStatus(summary.CoverageThreshold.Pass))
		if summary.CoverageThreshold.Pass {
			color.Green(complianceSummary)
		} else {
			color.Red(complianceSummary)
		}
	}
}
