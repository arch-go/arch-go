package reports

import (
	"io"

	"github.com/fatih/color"

	"github.com/arch-go/arch-go/internal/reports/model"
)

func displayRules(report *model.Report, output io.Writer) {
	color.Output = output

	if report.Compliance.Details != nil {
		displayDetails(report.Compliance.Details.ContentsVerificationDetails)
		displayDetails(report.Compliance.Details.DependenciesVerificationDetails)
		displayDetails(report.Compliance.Details.FunctionsVerificationDetails)
		displayDetails(report.Compliance.Details.NamingVerificationDetails)
	}
}

func displayDetails(verification model.Verification) {
	for _, d := range verification.Details {
		printRuleStatus(d)
		printPackagesDetails(d)
	}
}

func printPackagesDetails(d model.VerificationDetails) {
	for _, pd := range d.PackageDetails {
		if pd.Pass {
			color.Green("\tPackage '%s' passes\n", pd.Package)
		} else {
			color.Red("\tPackage '%s' fails\n", pd.Package)

			for _, str := range pd.Details {
				color.Red("\t\t%s\n", str)
			}
		}
	}
}

func printRuleStatus(d model.VerificationDetails) {
	if d.Pass {
		color.Green("[PASS] - %s\n", d.Rule)
	} else {
		color.Red("[FAIL] - %s\n", d.Rule)
	}
}
