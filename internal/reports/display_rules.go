package reports

import (
	"io"

	"github.com/fatih/color"

	"github.com/fdaines/arch-go/internal/reports/model"
)

func displayRules(report *model.Report, output io.Writer) {
	color.Output = output
	if report.Details != nil {
		displayDetails(report.Details.ContentsVerificationDetails)
		displayDetails(report.Details.DependenciesVerificationDetails)
		displayDetails(report.Details.FunctionsVerificationDetails)
		displayDetails(report.Details.NamingVerificationDetails)
	}
}

func displayDetails(verification model.Verification) {
	for _, d := range verification.Details {
		printRuleStatus(d)
		printPackagesDetails(d)
	}
}

func printPackagesDetails(d model.VerificationDetails) {
	for _, p := range d.PackageDetails {
		if p.Status == "PASS" {
			color.Green("\tPackage '%s' passes\n", p.Package)
		} else {
			color.Red("\tPackage '%s' fails\n", p.Package)
			for _, str := range p.Details {
				color.Red("\t\t%s\n", str)
			}
		}
	}
}

func printRuleStatus(d model.VerificationDetails) {
	if d.Status == "PASS" {
		color.Green("[PASS] - %s\n", d.Rule)
	} else {
		color.Red("[FAIL] - %s\n", d.Rule)
	}
}
