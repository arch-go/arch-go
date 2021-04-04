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

type DependencyRule struct {
	results []*result2.DependenciesRuleResult
	rule *config.DependenciesRule
	module *model2.ModuleInfo
}

func NewDependencyRule(results []*result2.DependenciesRuleResult, rule *config.DependenciesRule, module *model2.ModuleInfo) *DependencyRule {
	return &DependencyRule{
		rule: rule,
		results: results,
		module: module,
	}
}

func (dr *DependencyRule) CheckRule() []*result2.DependenciesRuleResult {
	if len(dr.rule.ShouldOnlyDependsOn) > 0 {
		shouldOnlyImportResult := dr.checkShouldOnlyImportRule()
		dr.results = append(dr.results, shouldOnlyImportResult)
	}
	if len(dr.rule.ShouldNotDependsOn) > 0 {
		shouldNotImportResult := dr.checkShouldNotImportRule()
		dr.results = append(dr.results, shouldNotImportResult)
	}
	return dr.results
}

func (dr *DependencyRule) checkShouldOnlyImportRule() *result2.DependenciesRuleResult {
	ruleResult := &result2.DependenciesRuleResult{
		Description: fmt.Sprintf("Packages matching pattern '%s' should only depends on: %v", dr.rule.Package, dr.rule.ShouldOnlyDependsOn),
		Passes:      true,
	}
	output.PrintVerbose("Check rule: package '%s' should only depends on: %v\n", dr.rule.Package, dr.rule.ShouldOnlyDependsOn)
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(dr.rule.Package))
	for _, p := range dr.module.Packages {
		if packageRegExp.MatchString(p.Path) {
			failureDetails := &result2.DependenciesRuleFailureDetail{Package: p.Path}
			output.PrintVerbose("Checking rule for package: %s\n", p.Path)
			result := true
			for _, pkg := range p.PackageData.Imports {
				if strings.HasPrefix(pkg, dr.module.MainPackage) {
					success := false
					output.PrintVerbose("Check if imported package '%s' complies with allowed imports\n", pkg)
					for _, allowedImport := range dr.rule.ShouldOnlyDependsOn {
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

func (dr *DependencyRule) checkShouldNotImportRule() *result2.DependenciesRuleResult {
	ruleResult := &result2.DependenciesRuleResult{
		Description: fmt.Sprintf("Packages matching pattern '%s' should not depends on: %v", dr.rule.Package, dr.rule.ShouldNotDependsOn),
		Passes:      true,
	}
	output.PrintVerbose("Check rule: package '%s' should not depends on: %v\n", dr.rule.Package, dr.rule.ShouldNotDependsOn)
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(dr.rule.Package))
	for _, p := range dr.module.Packages {
		if packageRegExp.MatchString(p.Path) {
			failureDetails := &result2.DependenciesRuleFailureDetail{Package: p.Path}
			output.PrintVerbose("Checking rule for package: %s\n", p.Path)
			result := true
			for _, pkg := range p.PackageData.Imports {
				if strings.HasPrefix(pkg, dr.module.MainPackage) {
					fails := false
					output.PrintVerbose("Check if imported package '%s' is one of the restricted packages\n", pkg)
					for _, notAllowedImport := range dr.rule.ShouldNotDependsOn {
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
