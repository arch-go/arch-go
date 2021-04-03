package dependencies

import (
	"fmt"
	"github.com/fdaines/arch-go/config"
	model2 "github.com/fdaines/arch-go/model"
	result2 "github.com/fdaines/arch-go/model/result"
	"github.com/fdaines/arch-go/utils/output"
	"github.com/fdaines/arch-go/utils/text"
	"regexp"
	"strings"
)

func CheckDependenciesRule(results []*result2.DependenciesRuleResult, r config.DependenciesRule, module *model2.ModuleInfo) []*result2.DependenciesRuleResult {
	if len(r.ShouldOnlyDependsOn) > 0 {
		shouldOnlyImportResult := checkShouldOnlyImportRule(r, module)
		results = append(results, shouldOnlyImportResult)
	}
	if len(r.ShouldNotDependsOn) > 0 {
		shouldNotImportResult := checkShouldNotImportRule(r, module)
		results = append(results, shouldNotImportResult)
	}
	return results
}

func checkShouldOnlyImportRule(rule config.DependenciesRule, module *model2.ModuleInfo) *result2.DependenciesRuleResult {
	ruleResult := &result2.DependenciesRuleResult{
		Description: fmt.Sprintf("Package '%s' should only depends on: %v", rule.Package, rule.ShouldOnlyDependsOn),
		Passes:      true,
	}
	output.PrintVerbose("Check rule: package '%s' should only depends on: %v\n", rule.Package, rule.ShouldOnlyDependsOn)
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(rule.Package))
	for _, p := range module.Packages {
		if packageRegExp.MatchString(p.Path) {
			failureDetails := &result2.DependenciesRuleFailureDetail{Package: p.Path}
			output.PrintVerbose("Checking rule for package: %s\n", p.Path)
			result := true
			for _, pkg := range p.PackageData.Imports {
				if strings.HasPrefix(pkg, module.MainPackage) {
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

func checkShouldNotImportRule(rule config.DependenciesRule, module *model2.ModuleInfo) *result2.DependenciesRuleResult {
	ruleResult := &result2.DependenciesRuleResult{
		Description: fmt.Sprintf("Package '%s' should not depends on: %v", rule.Package, rule.ShouldNotDependsOn),
		Passes:      true,
	}
	output.PrintVerbose("Check rule: package '%s' should not depends on: %v\n", rule.Package, rule.ShouldNotDependsOn)
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(rule.Package))
	for _, p := range module.Packages {
		if packageRegExp.MatchString(p.Path) {
			failureDetails := &result2.DependenciesRuleFailureDetail{Package: p.Path}
			output.PrintVerbose("Checking rule for package: %s\n", p.Path)
			result := true
			for _, pkg := range p.PackageData.Imports {
				if strings.HasPrefix(pkg, module.MainPackage) {
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
