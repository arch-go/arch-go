package impl

import (
	"github.com/fdaines/arch-go/config"
	"github.com/fdaines/arch-go/impl/contents"
	"github.com/fdaines/arch-go/impl/cycles"
	"github.com/fdaines/arch-go/impl/dependencies"
	"github.com/fdaines/arch-go/impl/functions"
	model2 "github.com/fdaines/arch-go/model"
	"github.com/fdaines/arch-go/model/result"
	"github.com/fdaines/arch-go/utils/output"
)

func CheckArchitecture(config *config.Config, module *model2.ModuleInfo) *result.Result {
	result := &result.Result{}

	output.Printf("Analyze Module: %s\n", module.MainPackage)
	result.DependenciesRulesResults = checkDependencies(config.DependenciesRules, module)
	result.ContentsRuleResults = checkContents(config.ContentRules, module)
	result.CyclesRuleResults = checkCycles(config.CyclesRules, module)
	result.FunctionsRulesResults = checkFunctions(config.FunctionsRules, module)

	return result
}

func checkCycles(rules []config.CyclesRule, module *model2.ModuleInfo) []*result.CyclesRuleResult {
	results := []*result.CyclesRuleResult{}
	for _, rule := range rules {
		results = cycles.CheckRule(results, rule, module)
	}
	return results
}

func checkFunctions(rules []config.FunctionsRule, module *model2.ModuleInfo) []*result.FunctionsRuleResult {
	results := []*result.FunctionsRuleResult{}
	for _, rule := range rules {
		results = functions.CheckRule(results, rule, module)
	}
	return results
}

func checkDependencies(rules []config.DependenciesRule, module *model2.ModuleInfo) []*result.DependenciesRuleResult {
	results := []*result.DependenciesRuleResult{}
	for _, rule := range rules {
		results = dependencies.CheckDependenciesRule(results, rule, module)
	}
	return results
}

func checkContents(rules []config.ContentsRule, module *model2.ModuleInfo) []*result.ContentsRuleResult {
	results := []*result.ContentsRuleResult{}
	for _, rule := range rules {
		results = contents.CheckRule(results, rule, module)
	}
	return results
}
