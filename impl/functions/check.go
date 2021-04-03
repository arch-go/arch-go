package functions

import (
	"fmt"
	"github.com/fdaines/arch-go/config"
	model2 "github.com/fdaines/arch-go/model"
	"github.com/fdaines/arch-go/model/result"
	"github.com/fdaines/arch-go/utils/text"
	"regexp"
)

func CheckRule(results []*result.FunctionsRuleResult, rule config.FunctionsRule, module *model2.ModuleInfo) []*result.FunctionsRuleResult {
	results = checkMaxParameters(results, rule, module)
	results = checkMaxReturnValues(results, rule, module)
	results = checkMaxPublicFunctions(results, rule, module)

	return results
}

func checkMaxPublicFunctions(results []*result.FunctionsRuleResult, rule config.FunctionsRule, module *model2.ModuleInfo) []*result.FunctionsRuleResult {
	if rule.MaxPublicFunctionPerFile <= 0 {
		return results
	}
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(rule.Package))
	ruleResult := &result.FunctionsRuleResult{
		Description: fmt.Sprintf("Packages matching pattern '%s' should only contain %d public functions per file", rule.Package, rule.MaxPublicFunctionPerFile),
		Passes:      true,
	}
	for _, p := range module.Packages {
		if packageRegExp.MatchString(p.Path) {
			functions, _ := retrieveFunctions(p, module.MainPackage)
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
					functionRuleViolation := &result.FunctionsRuleResultDetail{
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

func checkMaxReturnValues(results []*result.FunctionsRuleResult, rule config.FunctionsRule, module *model2.ModuleInfo) []*result.FunctionsRuleResult {
	if rule.MaxReturnValues <= 0 {
		return results
	}
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(rule.Package))
	ruleResult := &result.FunctionsRuleResult{
		Description: fmt.Sprintf("Functions in packages matching pattern '%s' should not return more than %d values", rule.Package, rule.MaxReturnValues),
		Passes:      true,
	}
	for _, p := range module.Packages {
		if packageRegExp.MatchString(p.Path) {
			functions, _ := retrieveFunctions(p, module.MainPackage)
			for _, fn := range functions {
				functionRuleViolation := &result.FunctionsRuleResultDetail{
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

func checkMaxParameters(results []*result.FunctionsRuleResult, rule config.FunctionsRule, module *model2.ModuleInfo) []*result.FunctionsRuleResult {
	if rule.MaxParameters <= 0 {
		return results
	}
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(rule.Package))
	ruleResult := &result.FunctionsRuleResult{
		Description: fmt.Sprintf("Functions in packages matching pattern '%s' should not receive more than %d parameters", rule.Package, rule.MaxParameters),
		Passes:      true,
	}
	for _, p := range module.Packages {
		if packageRegExp.MatchString(p.Path) {
			functions, _ := retrieveFunctions(p, module.MainPackage)
			for _, fn := range functions {
				functionRuleViolation := &result.FunctionsRuleResultDetail{
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
