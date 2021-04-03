package functions

import (
	"fmt"
	"github.com/fdaines/arch-go/config"
	"github.com/fdaines/arch-go/impl/model"
	"github.com/fdaines/arch-go/utils/packages"
	"github.com/fdaines/arch-go/utils/text"
	"regexp"
)

func CheckRule(results []*model.FunctionsRuleResult, rule config.FunctionsRule, mainPackage string, pkgs []*packages.PackageInfo) []*model.FunctionsRuleResult {
	results = checkMaxParameters(results, rule, mainPackage, pkgs)
	results = checkMaxReturnValues(results, rule, mainPackage, pkgs)
	results = checkMaxPublicFunctions(results, rule, mainPackage, pkgs)

	return results
}

func checkMaxPublicFunctions(results []*model.FunctionsRuleResult, rule config.FunctionsRule, mainPackage string, pkgs []*packages.PackageInfo) []*model.FunctionsRuleResult {
	if rule.MaxPublicFunctionPerFile <= 0 {
		return results
	}
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(rule.Package))
	ruleResult := &model.FunctionsRuleResult{
		Description: fmt.Sprintf("Packages matching pattern '%s' should only contain %d public functions per file", rule.Package, rule.MaxPublicFunctionPerFile),
		Passes:      true,
	}
	for _, p := range pkgs {
		if packageRegExp.MatchString(p.Path) {
			functions, _ := retrieveFunctions(p, mainPackage)
			publicFunctions := map[string]int{}
			for _, fn := range functions {
				if fn.IsPublic {
					current, ok := publicFunctions[fn.FilePath]
					if !ok {
						current = 0
					}
					publicFunctions[fn.FilePath] = current + 1
				}
			}
			for key, value := range publicFunctions {
				if value > rule.MaxPublicFunctionPerFile {
					functionRuleViolation := &model.FunctionsRuleResultDetail{
						Package: p.Path,
						File:    key,
					}
					functionRuleViolation.Details = append(functionRuleViolation.Details,
						fmt.Sprintf("File %s has too many public functions (%d)", key, value))
					ruleResult.Failures = append(ruleResult.Failures, functionRuleViolation)
				}
			}
		}
	}
	ruleResult.Passes = len(ruleResult.Failures) == 0
	results = append(results, ruleResult)

	return results
}

func checkMaxReturnValues(results []*model.FunctionsRuleResult, rule config.FunctionsRule, mainPackage string, pkgs []*packages.PackageInfo) []*model.FunctionsRuleResult {
	if rule.MaxReturnValues <= 0 {
		return results
	}
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(rule.Package))
	ruleResult := &model.FunctionsRuleResult{
		Description: fmt.Sprintf("Functions in packages matching pattern '%s' should not return more than %d values", rule.Package, rule.MaxReturnValues),
		Passes:      true,
	}
	for _, p := range pkgs {
		if packageRegExp.MatchString(p.Path) {
			functions, _ := retrieveFunctions(p, mainPackage)
			for _, fn := range functions {
				functionRuleViolation := &model.FunctionsRuleResultDetail{
					Package: p.Path,
					Name:    fn.Name,
					File:    fn.File,
				}
				if fn.NumReturns > rule.MaxReturnValues {
					functionRuleViolation.Details = append(functionRuleViolation.Details,
						fmt.Sprintf("Function %s in file %s returns too many values (%d)",
							fn.Name, fn.FilePath, fn.NumReturns))
					ruleResult.Failures = append(ruleResult.Failures, functionRuleViolation)
				}
			}
		}
	}
	ruleResult.Passes = len(ruleResult.Failures) == 0
	results = append(results, ruleResult)

	return results
}

func checkMaxParameters(results []*model.FunctionsRuleResult, rule config.FunctionsRule, mainPackage string, pkgs []*packages.PackageInfo) []*model.FunctionsRuleResult {
	if rule.MaxParameters <= 0 {
		return results
	}
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(rule.Package))
	ruleResult := &model.FunctionsRuleResult{
		Description: fmt.Sprintf("Functions in packages matching pattern '%s' should not receive more than %d parameters", rule.Package, rule.MaxParameters),
		Passes:      true,
	}
	for _, p := range pkgs {
		if packageRegExp.MatchString(p.Path) {
			functions, _ := retrieveFunctions(p, mainPackage)
			for _, fn := range functions {
				functionRuleViolation := &model.FunctionsRuleResultDetail{
					Package: p.Path,
					Name:    fn.Name,
					File:    fn.File,
				}
				if fn.NumParams > rule.MaxParameters {
					functionRuleViolation.Details = append(functionRuleViolation.Details,
						fmt.Sprintf("Function %s in file %s receive too many parameters (%d)",
							fn.Name, fn.FilePath, fn.NumParams))
					ruleResult.Failures = append(ruleResult.Failures, functionRuleViolation)
				}
			}
		}
	}
	ruleResult.Passes = len(ruleResult.Failures) == 0
	results = append(results, ruleResult)

	return results
}
