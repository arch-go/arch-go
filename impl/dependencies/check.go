package dependencies

import (
	"fmt"
	"github.com/fdaines/arch-go/config"
	"github.com/fdaines/arch-go/model"
	"github.com/fdaines/arch-go/utils/output"
	"github.com/fdaines/arch-go/utils/packages"
	"github.com/fdaines/arch-go/utils/text"
	"regexp"
	"strings"
)

func CheckDependenciesRule(results []*model.DependenciesRuleResult, r config.DependenciesRule, mainPackage string, pkgs []*packages.PackageInfo) []*model.DependenciesRuleResult {
	if len(r.ShouldOnlyDependsOn)>0 {
		shouldOnlyImportResult := checkShouldOnlyImportRule(r, mainPackage, pkgs)
		results = append(results, shouldOnlyImportResult)
	}
	if len(r.ShouldNotDependsOn)>0 {
		shouldNotImportResult := checkShouldNotImportRule(r, mainPackage, pkgs)
		results = append(results, shouldNotImportResult)
	}
	return results
}

func checkShouldOnlyImportRule(rule config.DependenciesRule, mainPackage string, pkgs []*packages.PackageInfo) *model.DependenciesRuleResult {
	ruleResult := &model.DependenciesRuleResult{
		Description: fmt.Sprintf("Package '%s' should only depends on: %v", rule.Package, rule.ShouldOnlyDependsOn),
		Passes: true,
	}
	output.PrintVerbose("Check rule: package '%s' should only depends on: %v\n", rule.Package, rule.ShouldOnlyDependsOn)
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(rule.Package))
	for _, p := range pkgs {
		if packageRegExp.MatchString(p.Path) {
			failureDetails := &model.DependenciesRuleFailureDetail{Package: p.Path}
			output.PrintVerbose("Checking rule for package: %s\n", p.Path)
			result := true
			for _, pkg := range p.PackageData.Imports {
				if strings.HasPrefix(pkg, mainPackage) {
					success := false
					output.PrintVerbose("Check if imported package '%s' complies with allowed imports\n", pkg)
					for _, allowedImport := range rule.ShouldOnlyDependsOn {
						allowedImportRegexp, _ := regexp.Compile(text.PreparePackageRegexp(allowedImport))
						success = success || allowedImportRegexp.MatchString(pkg)
					}
					if !success {
						failureDetails.Details = append(failureDetails.Details, fmt.Sprintf("imports package '%s'", pkg))
					}
					result = result && success
				}
			}
			if !result {
				ruleResult.Passes = false
				ruleResult.Failures = append(ruleResult.Failures, failureDetails)
			}
		}
	}

	return ruleResult
}

func checkShouldNotImportRule(rule config.DependenciesRule, mainPackage string, pkgs []*packages.PackageInfo) *model.DependenciesRuleResult {
	ruleResult := &model.DependenciesRuleResult{
		Description: fmt.Sprintf("Package '%s' should not depends on: %v", rule.Package, rule.ShouldNotDependsOn),
		Passes: true,
	}
	output.PrintVerbose("Check rule: package '%s' should not depends on: %v\n", rule.Package, rule.ShouldNotDependsOn)
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(rule.Package))
	for _, p := range pkgs {
		if packageRegExp.MatchString(p.Path) {
			failureDetails := &model.DependenciesRuleFailureDetail{Package: p.Path}
			output.PrintVerbose("Checking rule for package: %s\n", p.Path)
			result := true
			for _, pkg := range p.PackageData.Imports {
				if strings.HasPrefix(pkg, mainPackage) {
					fails := false
					output.PrintVerbose("Check if imported package '%s' is one of the restricted packages\n", pkg)
					for _, notAllowedImport := range rule.ShouldNotDependsOn {
						notAllowedImportRegexp, _ := regexp.Compile(text.PreparePackageRegexp(notAllowedImport))
						fails = fails || notAllowedImportRegexp.MatchString(pkg)
					}
					if fails {
						failureDetails.Details = append(failureDetails.Details, fmt.Sprintf("imports package '%s'", pkg))
					}
					result = result && !fails
				}
			}
			if !result {
				ruleResult.Passes = false
				ruleResult.Failures = append(ruleResult.Failures, failureDetails)
			}
		}
	}

	return ruleResult
}