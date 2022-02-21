package dependencies

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/fdaines/arch-go/internal/config"
	"github.com/fdaines/arch-go/internal/impl/model"
	"github.com/fdaines/arch-go/internal/utils/output"
	"github.com/fdaines/arch-go/internal/utils/packages"
	"github.com/fdaines/arch-go/internal/utils/text"
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
	if len(rule.ShouldNotDependsOnExternal) > 0 {
		ruleDescriptions = append(ruleDescriptions, fmt.Sprintf("'not depends on the following external dependencies [%v]'", rule.ShouldNotDependsOnExternal))
	}
	description := fmt.Sprintf("Packages matching pattern '%s' should [%s]", rule.Package, strings.Join(ruleDescriptions, ","))

	return &DependencyRuleVerification{
		Module:      module,
		Rule:        rule,
		Description: description,
		Passes:      true,
	}
}

func (d *DependencyRuleVerification) Verify() bool {
	result := true
	for index, pd := range d.PackageDetails {
		d.PackageDetails[index].Passes = true
		output.PrintVerbose("Checking dependency rules for package: %s\n", pd.Package.Path)
		result = d.checksShouldOnlyDependsOn(pd, result, index)
		result = d.checksShouldNotDependsOn(pd, result, index)
		result = d.checksShouldOnlyDependsOnExternal(pd, result, index)
		result = d.checksShouldNotDependsOnExternal(pd, result, index)

		d.Passes = result
	}
	return d.Passes
}

func (d *DependencyRuleVerification) checksShouldNotDependsOnExternal(pd model.PackageVerification, result bool, index int) bool {
	if len(d.Rule.ShouldNotDependsOnExternal) > 0 {
		for _, pkg := range pd.Package.PackageData.Imports {
			result = d.checkComplianceWithRestrictedExternalImports(pkg, index, result)
		}
	}
	return result
}

func (d *DependencyRuleVerification) checksShouldOnlyDependsOnExternal(pd model.PackageVerification, result bool, index int) bool {
	if len(d.Rule.ShouldOnlyDependsOnExternal) > 0 {
		for _, pkg := range pd.Package.PackageData.Imports {
			result = d.checkComplianceWithAllowedExternalImports(pkg, index, result)
		}
	}
	return result
}

func (d *DependencyRuleVerification) checksShouldNotDependsOn(pd model.PackageVerification, result bool, index int) bool {
	if len(d.Rule.ShouldNotDependsOn) > 0 {
		for _, pkg := range pd.Package.PackageData.Imports {
			result = d.checkComplianceWithRestrictedInternalImports(pkg, index, result)
		}
	}
	return result
}

func (d *DependencyRuleVerification) checksShouldOnlyDependsOn(pd model.PackageVerification, result bool, index int) bool {
	if len(d.Rule.ShouldOnlyDependsOn) > 0 {
		for _, pkg := range pd.Package.PackageData.Imports {
			result = d.checkComplianceWithAllowedInternalImports(pkg, index, result)
		}
	}
	return result
}

func (d *DependencyRuleVerification) checkComplianceWithAllowedExternalImports(pkg string, index int, result bool) bool {
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
	return result
}

func (d *DependencyRuleVerification) checkComplianceWithRestrictedExternalImports(pkg string, index int, result bool) bool {
	if !strings.HasPrefix(pkg, d.Module) && packages.IsExternalPackage(pkg) {
		fails := false
		output.PrintVerbose("Check if imported package '%s' complies with restricted imports\n", pkg)
		for _, restrictedImport := range d.Rule.ShouldNotDependsOnExternal {
			restrictedImportRegexp, err := regexp.Compile(text.PreparePackageRegexp(restrictedImport))
			if err != nil {
				output.Printf("Error compiling restricted imports expresion: %+v\n", err)
			}
			fails = fails || restrictedImportRegexp.MatchString(pkg)
		}
		if fails {
			d.PackageDetails[index].Details = append(d.PackageDetails[index].Details, fmt.Sprintf("ShouldNotDependsOnExternal rule contains imported package '%s'", pkg))
			d.PackageDetails[index].Passes = false
		}
		result = result && !fails
	}
	return result
}

func (d *DependencyRuleVerification) checkComplianceWithRestrictedInternalImports(pkg string, index int, result bool) bool {
	if strings.HasPrefix(pkg, d.Module) {
		fails := false
		output.PrintVerbose("Check if imported package '%s' is one of the restricted packages\n", pkg)
		for _, notAllowedImport := range d.Rule.ShouldNotDependsOn {
			notAllowedImportRegexp, err := regexp.Compile(text.PreparePackageRegexp(notAllowedImport))
			if err != nil {
				output.Printf("Error compiling allowed imports expresion: %+v\n", err)
			}
			fails = fails || notAllowedImportRegexp.MatchString(pkg)
		}
		if fails {
			d.PackageDetails[index].Details = append(d.PackageDetails[index].Details, fmt.Sprintf("ShouldNotDependsOn rule contains imported package '%s'", pkg))
			d.PackageDetails[index].Passes = false
		}
		result = result && !fails
	}
	return result
}

func (d *DependencyRuleVerification) checkComplianceWithAllowedInternalImports(pkg string, index int, result bool) bool {
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
	return result
}

func (d *DependencyRuleVerification) Type() string {
	return "DependenciesRule"
}

func (d *DependencyRuleVerification) Name() string {
	return d.Description
}

func (d *DependencyRuleVerification) Status() bool {
	return d.Passes
}

func (d *DependencyRuleVerification) ValidatePatterns() bool {
	isValid := true
	_, err := regexp.Compile(text.PreparePackageRegexp(d.Rule.Package))
	if err != nil {
		color.Red("[%s] - Invalid Package Pattern: %s\n", d.Description, d.Rule.Package)
		isValid = false
	}
	for _, sodo := range d.Rule.ShouldOnlyDependsOn {
		_, err := regexp.Compile(text.PreparePackageRegexp(sodo))
		if err != nil {
			color.Red("[%s] - Invalid pattern in ShouldOnlyDependsOn: %s\n", d.Description, sodo)
			isValid = false
		}
	}
	for _, sndo := range d.Rule.ShouldNotDependsOn {
		_, err := regexp.Compile(text.PreparePackageRegexp(sndo))
		if err != nil {
			color.Red("[%s] - Invalid pattern in ShouldNotDependsOn: %s\n", d.Description, sndo)
			isValid = false
		}
	}
	for _, sodoe := range d.Rule.ShouldOnlyDependsOnExternal {
		_, err := regexp.Compile(text.PreparePackageRegexp(sodoe))
		if err != nil {
			color.Red("[%s] - Invalid pattern in ShouldOnlyDependsOnExternal: %s\n", d.Description, sodoe)
			isValid = false
		}
	}
	for _, sndoe := range d.Rule.ShouldNotDependsOnExternal {
		_, err := regexp.Compile(text.PreparePackageRegexp(sndoe))
		if err != nil {
			color.Red("[%s] - Invalid pattern in ShouldNotDependsOnExternal: %s\n", d.Description, sndoe)
			isValid = false
		}
	}
	return isValid
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

func (d *DependencyRuleVerification) GetVerifications() []model.PackageVerification {
	return d.PackageDetails
}