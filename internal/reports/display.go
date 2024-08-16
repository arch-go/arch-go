package reports

import (
	"fmt"
	"io"

	"github.com/fatih/color"

	"github.com/arch-go/arch-go/internal/common"
	"github.com/arch-go/arch-go/internal/reports/console"
	"github.com/arch-go/arch-go/internal/reports/html"
	"github.com/arch-go/arch-go/internal/reports/json"
	"github.com/arch-go/arch-go/internal/reports/model"
	"github.com/arch-go/arch-go/internal/reports/utils"
)

func DisplayResult(report *model.Report, output io.Writer) {
	displayRules(report, output)

	generateHTMLReport := common.HTML
	if generateHTMLReport {
		html.GenerateHTMLReport(report, output)
	}

	generateJSONReport := common.HTML
	if generateJSONReport {
		json.GenerateReport(report, output)
	}

	console.GenerateConsoleReport(report, output)

	displaySummary(report, output)
}

func displaySummary(report *model.Report, output io.Writer) {
	const lineSeparator = "--------------------------------------\n"

	color.Output = output

	fmt.Fprint(output, lineSeparator)
	fmt.Fprint(output, "\tExecution Summary\n")
	fmt.Fprint(output, lineSeparator)
	fmt.Fprintf(output, "Total Rules: \t%d\n", report.Compliance.Total)
	fmt.Fprintf(output, "Succeeded: \t%d\n", report.Compliance.Passed)
	fmt.Fprintf(output, "Failed: \t%d\n", report.Compliance.Failed)
	fmt.Fprint(output, lineSeparator)

	if report.Compliance.Threshold != nil {
		complianceSummary := fmt.Sprintf("Compliance: %8d%% (%s)\n",
			report.Compliance.Rate,
			utils.ResolveStatus(report.Compliance.Pass))
		if report.Compliance.Pass {
			color.Green(complianceSummary)
		} else {
			color.Red(complianceSummary)
		}
	}

	if report.Coverage.Threshold != nil {
		coverageSummary := fmt.Sprintf("Coverage: %10d%% (%s)\n",
			report.Coverage.Rate,
			utils.ResolveStatus(report.Coverage.Pass))
		if report.Coverage.Pass {
			color.Green(coverageSummary)
		} else {
			color.Red(coverageSummary)
		}
	}
}
