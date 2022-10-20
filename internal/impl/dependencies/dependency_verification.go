package dependencies

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/fdaines/arch-go/internal/config"
	"github.com/fdaines/arch-go/internal/model"
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
	if rule.ShouldOnlyDependsOn != nil {
		ruleDescriptions = describeRules(rule.ShouldOnlyDependsOn, ruleDescriptions, true)
	}
	if rule.ShouldNotDependsOn != nil {
		ruleDescriptions = describeRules(rule.ShouldNotDependsOn, ruleDescriptions, false)
	}
	description := fmt.Sprintf("Packages matching pattern '%s' should [%s]", rule.Package, strings.Join(ruleDescriptions, ","))

	return &DependencyRuleVerification{
		Module:      module,
		Rule:        rule,
		Description: description,
		Passes:      true,
	}
}

func describeRules(d *config.Dependencies, ruleDescriptions []string, allowed bool) []string {
	kind := "only"
	if !allowed {
		kind = "not"
	}
	if len(d.Internal) > 0 {
		ruleDescriptions = append(ruleDescriptions, fmt.Sprintf("'%s depend on internal packages that matches [%v]'", kind, d.Internal))
	}
	if len(d.External) > 0 {
		ruleDescriptions = append(ruleDescriptions, fmt.Sprintf("'%s depend on external packages that matches [%v]'", kind, d.External))
	}
	if len(d.Standard) > 0 {
		ruleDescriptions = append(ruleDescriptions, fmt.Sprintf("'%s depend on standard packages that matches [%v]'", kind, d.Standard))
	}
	return ruleDescriptions
}

func (d *DependencyRuleVerification) Verify() bool {
	result := true
	for index, pd := range d.PackageDetails {
		d.PackageDetails[index].Passes = true
		output.PrintVerbose("Checking dependency rules for package: %s\n", pd.Package.Path)
		result = d.checksShouldOnlyDependsOn(pd, result, index)
		result = d.checksShouldNotDependsOn(pd, result, index)

		d.Passes = result
	}
	return d.Passes
}

func (d *DependencyRuleVerification) checksShouldNotDependsOn(pd model.PackageVerification, result bool, index int) bool {
	if d.Rule.ShouldNotDependsOn != nil {
		for _, pkg := range pd.Package.PackageData.Imports {
			result = d.checkComplianceWithRestrictedInternalImports(pkg, index, result)
			result = d.checkComplianceWithRestrictedExternalImports(pkg, index, result)
			result = d.checkComplianceWithRestrictedStandardImports(pkg, index, result)
		}
	}
	return result
}

func (d *DependencyRuleVerification) checksShouldOnlyDependsOn(pd model.PackageVerification, result bool, index int) bool {
	if d.Rule.ShouldOnlyDependsOn != nil {
		for _, pkg := range pd.Package.PackageData.Imports {
			result = d.checkComplianceWithAllowedInternalImports(pkg, index, result)
			result = d.checkComplianceWithAllowedExternalImports(pkg, index, result)
			result = d.checkComplianceWithAllowedStandardImports(pkg, index, result)
		}
	}
	return result
}

func (d *DependencyRuleVerification) checkComplianceWithAllowedStandardImports(pkg string, index int, result bool) bool {
	if len(d.Rule.ShouldOnlyDependsOn.Standard) == 0 {
		return result
	}
	if !strings.HasPrefix(pkg, d.Module) && packages.IsStandardPackage(pkg) {
		success := false
		output.PrintVerbose("Check if imported package '%s' complies with allowed standard imports\n", pkg)
		for _, allowedImport := range d.Rule.ShouldOnlyDependsOn.Standard {
			allowedImportRegexp, _ := regexp.Compile(text.PreparePackageRegexp(allowedImport))
			success = success || allowedImportRegexp.MatchString(pkg)
		}
		if !success {
			d.PackageDetails[index].Details = append(d.PackageDetails[index].Details, fmt.Sprintf("ShouldOnlyDependsOn.Standard rule doesn't contains imported package '%s'", pkg))
			d.PackageDetails[index].Passes = false
		}
		result = result && success
	}
	return result
}

func (d *DependencyRuleVerification) checkComplianceWithRestrictedStandardImports(pkg string, index int, result bool) bool {
	if len(d.Rule.ShouldNotDependsOn.Standard) == 0 {
		return result
	}
	if !strings.HasPrefix(pkg, d.Module) && packages.IsStandardPackage(pkg) {
		fails := false
		output.PrintVerbose("Check if imported package '%s' complies with restricted standard imports\n", pkg)
		for _, restrictedImport := range d.Rule.ShouldNotDependsOn.Standard {
			restrictedImportRegexp, err := regexp.Compile(text.PreparePackageRegexp(restrictedImport))
			if err != nil {
				output.Printf("Error compiling restricted imports expresion: %+v\n", err)
			}
			fails = fails || restrictedImportRegexp.MatchString(pkg)
		}
		if fails {
			d.PackageDetails[index].Details = append(d.PackageDetails[index].Details, fmt.Sprintf("ShouldNotDependsOn.Standard rule contains imported package '%s'", pkg))
			d.PackageDetails[index].Passes = false
		}
		result = result && !fails
	}
	return result
}

func (d *DependencyRuleVerification) checkComplianceWithAllowedExternalImports(pkg string, index int, result bool) bool {
	if len(d.Rule.ShouldOnlyDependsOn.External) == 0 {
		return result
	}
	if !strings.HasPrefix(pkg, d.Module) && packages.IsExternalPackage(pkg) {
		success := false
		output.PrintVerbose("Check if imported package '%s' complies with allowed external imports\n", pkg)
		for _, allowedImport := range d.Rule.ShouldOnlyDependsOn.External {
			allowedImportRegexp, _ := regexp.Compile(text.PreparePackageRegexp(allowedImport))
			success = success || allowedImportRegexp.MatchString(pkg)
		}
		if !success {
			d.PackageDetails[index].Details = append(d.PackageDetails[index].Details, fmt.Sprintf("ShouldOnlyDependsOn.External rule doesn't contains imported package '%s'", pkg))
			d.PackageDetails[index].Passes = false
		}
		result = result && success
	}
	return result
}

func (d *DependencyRuleVerification) checkComplianceWithRestrictedExternalImports(pkg string, index int, result bool) bool {
	if len(d.Rule.ShouldNotDependsOn.External) == 0 {
		return result
	}
	if !strings.HasPrefix(pkg, d.Module) && packages.IsExternalPackage(pkg) {
		fails := false
		output.PrintVerbose("Check if imported package '%s' complies with restricted external imports\n", pkg)
		for _, restrictedImport := range d.Rule.ShouldNotDependsOn.External {
			restrictedImportRegexp, err := regexp.Compile(text.PreparePackageRegexp(restrictedImport))
			if err != nil {
				output.Printf("Error compiling restricted imports expresion: %+v\n", err)
			}
			fails = fails || restrictedImportRegexp.MatchString(pkg)
		}
		if fails {
			d.PackageDetails[index].Details = append(d.PackageDetails[index].Details, fmt.Sprintf("ShouldNotDependsOn.External rule contains imported package '%s'", pkg))
			d.PackageDetails[index].Passes = false
		}
		result = result && !fails
	}
	return result
}

func (d *DependencyRuleVerification) checkComplianceWithRestrictedInternalImports(pkg string, index int, result bool) bool {
	if len(d.Rule.ShouldNotDependsOn.Internal) == 0 {
		return result
	}
	if strings.HasPrefix(pkg, d.Module) {
		fails := false
		output.PrintVerbose("Check if imported package '%s' is one of the restricted internal packages\n", pkg)
		for _, notAllowedImport := range d.Rule.ShouldNotDependsOn.Internal {
			notAllowedImportRegexp, err := regexp.Compile(text.PreparePackageRegexp(notAllowedImport))
			if err != nil {
				output.Printf("Error compiling allowed imports expresion: %+v\n", err)
			}
			fails = fails || notAllowedImportRegexp.MatchString(pkg)
		}
		if fails {
			d.PackageDetails[index].Details = append(d.PackageDetails[index].Details, fmt.Sprintf("ShouldNotDependsOn.Internal rule contains imported package '%s'", pkg))
			d.PackageDetails[index].Passes = false
		}
		result = result && !fails
	}
	return result
}

func (d *DependencyRuleVerification) checkComplianceWithAllowedInternalImports(pkg string, index int, result bool) bool {
	if len(d.Rule.ShouldOnlyDependsOn.Internal) == 0 {
		return result
	}
	if strings.HasPrefix(pkg, d.Module) {
		success := false
		output.PrintVerbose("Check if imported package '%s' complies with allowed internal imports\n", pkg)
		for _, allowedImport := range d.Rule.ShouldOnlyDependsOn.Internal {
			allowedImportRegexp, _ := regexp.Compile(text.PreparePackageRegexp(allowedImport))
			success = success || allowedImportRegexp.MatchString(pkg)
		}
		if !success {
			d.PackageDetails[index].Details = append(d.PackageDetails[index].Details, fmt.Sprintf("ShouldOnlyDependsOn.Internal rule doesn't contains imported package '%s'", pkg))
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
	isValid = d.validateDependencies(d.Rule.ShouldOnlyDependsOn, "ShouldOnlyDependsOn", isValid)
	isValid = d.validateDependencies(d.Rule.ShouldNotDependsOn, "ShouldNotDependsOn", isValid)

	return isValid
}

func (d *DependencyRuleVerification) validateDependencies(dependencies *config.Dependencies, baseRule string, isValid bool) bool {
	if dependencies == nil {
		return isValid
	}
	for _, sodo := range dependencies.Internal {
		_, err := regexp.Compile(text.PreparePackageRegexp(sodo))
		if err != nil {
			color.Red("[%s] - Invalid pattern in %s.Internal: %s\n", d.Description, baseRule, sodo)
			isValid = false
		}
	}
	for _, sodoe := range dependencies.External {
		_, err := regexp.Compile(text.PreparePackageRegexp(sodoe))
		if err != nil {
			color.Red("[%s] - Invalid pattern in %s.External: %s\n", d.Description, baseRule, sodoe)
			isValid = false
		}
	}
	for _, sodos := range dependencies.Standard {
		_, err := regexp.Compile(text.PreparePackageRegexp(sodos))
		if err != nil {
			color.Red("[%s] - Invalid pattern in %s.Standard: %s\n", d.Description, baseRule, sodos)
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
