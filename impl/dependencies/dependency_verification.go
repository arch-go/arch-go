package dependencies

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/fdaines/arch-go/config"
	"github.com/fdaines/arch-go/impl/model"
	"github.com/fdaines/arch-go/utils/output"
	"github.com/fdaines/arch-go/utils/packages"
	"github.com/fdaines/arch-go/utils/text"
	"regexp"
	"strings"
)

type DependencyRuleVerification struct {
	Module         string
	Description    string
	Rule           *config.DependenciesRule
	PackageDetails []model.PackageVerification
	Passes         bool
}

func NewDependencyRuleVerification(module string, rule *config.DependenciesRule) *DependencyRuleVerification {
	var ruleDescriptions []string
	if len(rule.ShouldOnlyDependsOn) > 0 {
		ruleDescriptions = append(ruleDescriptions, fmt.Sprintf("'only depends on [%v]'", rule.ShouldOnlyDependsOn))
	}
	if len(rule.ShouldNotDependsOn) > 0 {
		ruleDescriptions = append(ruleDescriptions, fmt.Sprintf("'not depends on [%v]'", rule.ShouldNotDependsOn))
	}
	if len(rule.ShouldOnlyDependsOnExternal) > 0 {
		ruleDescriptions = append(ruleDescriptions, fmt.Sprintf("'only depends on the following external dependencies [%v]'", rule.ShouldOnlyDependsOnExternal))
	}
	description := fmt.Sprintf("Packages matching pattern '%s' should [%s]", rule.Package, strings.Join(ruleDescriptions, ","))

	return &DependencyRuleVerification{
		Module:      module,
		Rule:        rule,
		Description: description,
		Passes:      true,
	}
}

func (d *DependencyRuleVerification) Verify() {
	result := true
	for index, pd := range d.PackageDetails {
		d.PackageDetails[index].Passes = true
		output.PrintVerbose("Checking dependency rules for package: %s\n", pd.Package.Path)
		if len(d.Rule.ShouldOnlyDependsOn) > 0 {
			for _, pkg := range pd.Package.PackageData.Imports {
				if strings.HasPrefix(pkg, d.Module) {
					success := false
					output.PrintVerbose("Check if imported package '%s' complies with allowed imports\n", pkg)
					for _, allowedImport := range d.Rule.ShouldOnlyDependsOn {
						allowedImportRegexp, _ := regexp.Compile(text.PreparePackageRegexp(allowedImport))
						success = success || allowedImportRegexp.MatchString(pkg)
					}
					if !success {
						d.PackageDetails[index].Details = append(d.PackageDetails[index].Details, fmt.Sprintf("ShouldOnlyDependsOn rule doesn't contains imported package '%s'", pkg))
						d.PackageDetails[index].Passes = false
					}
					result = result && success
				}
			}
		}
		if len(d.Rule.ShouldNotDependsOn) > 0 {
			for index, pkg := range pd.Package.PackageData.Imports {
				if strings.HasPrefix(pkg, d.Module) {
					fails := false
					output.PrintVerbose("Check if imported package '%s' is one of the restricted packages\n", pkg)
					for _, notAllowedImport := range d.Rule.ShouldNotDependsOn {
						notAllowedImportRegexp, _ := regexp.Compile(text.PreparePackageRegexp(notAllowedImport))
						fails = fails || notAllowedImportRegexp.MatchString(pkg)
					}
					if fails {
						d.PackageDetails[index].Details = append(d.PackageDetails[index].Details, fmt.Sprintf("ShouldNotDependsOn rule contains imported package '%s'", pkg))
						d.PackageDetails[index].Passes = false
					}
					result = result && !fails
				}
			}
		}
		if len(d.Rule.ShouldOnlyDependsOnExternal) > 0 {
			for _, pkg := range pd.Package.PackageData.Imports {
				if !strings.HasPrefix(pkg, d.Module) && packages.IsExternalPackage(pkg) {
					success := false
					output.PrintVerbose("Check if imported package '%s' complies with allowed imports\n", pkg)
					for _, allowedImport := range d.Rule.ShouldOnlyDependsOnExternal {
						allowedImportRegexp, _ := regexp.Compile(text.PreparePackageRegexp(allowedImport))
						success = success || allowedImportRegexp.MatchString(pkg)
					}
					if !success {
						d.PackageDetails[index].Details = append(d.PackageDetails[index].Details, fmt.Sprintf("ShouldOnlyDependsOnExternal rule doesn't contains imported package '%s'", pkg))
						d.PackageDetails[index].Passes = false
					}
					result = result && success
				}
			}
		}

		d.Passes = result
	}
}

func (d *DependencyRuleVerification) PrintResults() {
	if d.Passes {
		color.Green("[PASS] - %s\n", d.Description)
	} else {
		color.Red("[FAIL] - %s\n", d.Description)
	}
	for _, p := range d.PackageDetails {
		if p.Passes {
			color.Green("\tPackage '%s' passes\n", p.Package.Path)
		} else {
			color.Red("\tPackage '%s' fails\n", p.Package.Path)
			for _, str := range p.Details {
				color.Red("\t\t%s\n", str)
			}
		}
	}
}
