package functions

import (
	"fmt"
	"github.com/fdaines/arch-go/config"
	"github.com/fdaines/arch-go/model"
	"github.com/fdaines/arch-go/model/result"
	"github.com/fdaines/arch-go/utils/text"
	"regexp"
)

type FunctionsRule struct {
	results []*result.FunctionsRuleResult
	rule    *config.FunctionsRule
	module  *model.ModuleInfo
}

func NewFunctionRule(results []*result.FunctionsRuleResult, rule *config.FunctionsRule, module *model.ModuleInfo) *FunctionsRule {
	return &FunctionsRule{
		rule:    rule,
		results: results,
		module:  module,
	}
}

func (f *FunctionsRule) CheckRule() []*result.FunctionsRuleResult {
	f.checkMaxParameters()
	f.checkMaxReturnValues()
	f.checkMaxPublicFunctions()

	return f.results
}

func (f *FunctionsRule) checkMaxPublicFunctions() {
	if f.rule.MaxPublicFunctionPerFile <= 0 {
		return
	}
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(f.rule.Package))
	ruleResult := &result.FunctionsRuleResult{
		Description: fmt.Sprintf("Packages matching pattern '%s' should only contain %d public functions per file",
			f.rule.Package, f.rule.MaxPublicFunctionPerFile),
		Passes: true,
	}
	for _, p := range f.module.Packages {
		if packageRegExp.MatchString(p.Path) {
			functions, _ := retrieveFunctions(p, f.module.MainPackage)
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
				if value > f.rule.MaxPublicFunctionPerFile {
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
	f.results = append(f.results, ruleResult)
}

func (f *FunctionsRule) checkMaxReturnValues() {
	if f.rule.MaxReturnValues <= 0 {
		return
	}
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(f.rule.Package))
	ruleResult := &result.FunctionsRuleResult{
		Description: fmt.Sprintf("Functions in packages matching pattern '%s' should not return more than %d values",
			f.rule.Package, f.rule.MaxReturnValues),
		Passes: true,
	}
	for _, p := range f.module.Packages {
		if packageRegExp.MatchString(p.Path) {
			functions, _ := retrieveFunctions(p, f.module.MainPackage)
			for _, fn := range functions {
				functionRuleViolation := &result.FunctionsRuleResultDetail{
					Package: p.Path,
					Name:    fn.Name,
					File:    fn.File,
				}
				if fn.NumReturns > f.rule.MaxReturnValues {
					functionRuleViolation.Details = append(functionRuleViolation.Details,
						fmt.Sprintf("Function %s in file %s returns too many values (%d)",
							fn.Name, fn.FilePath, fn.NumReturns))
					ruleResult.Failures = append(ruleResult.Failures, functionRuleViolation)
				}
			}
		}
	}
	ruleResult.Passes = len(ruleResult.Failures) == 0
	f.results = append(f.results, ruleResult)
}

func (f *FunctionsRule) checkMaxParameters() {
	if f.rule.MaxParameters <= 0 {
		return
	}
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(f.rule.Package))
	ruleResult := &result.FunctionsRuleResult{
		Description: fmt.Sprintf("Functions in packages matching pattern '%s' should not receive more than %d parameters",
			f.rule.Package, f.rule.MaxParameters),
		Passes: true,
	}
	for _, p := range f.module.Packages {
		if packageRegExp.MatchString(p.Path) {
			functions, _ := retrieveFunctions(p, f.module.MainPackage)
			for _, fn := range functions {
				functionRuleViolation := &result.FunctionsRuleResultDetail{
					Package: p.Path,
					Name:    fn.Name,
					File:    fn.File,
				}
				if fn.NumParams > f.rule.MaxParameters {
					functionRuleViolation.Details = append(functionRuleViolation.Details,
						fmt.Sprintf("Function %s in file %s receive too many parameters (%d)",
							fn.Name, fn.FilePath, fn.NumParams))
					ruleResult.Failures = append(ruleResult.Failures, functionRuleViolation)
				}
			}
		}
	}
	ruleResult.Passes = len(ruleResult.Failures) == 0
	f.results = append(f.results, ruleResult)
}
