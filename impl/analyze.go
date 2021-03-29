package impl

import (
	"github.com/fdaines/arch-go/config"
	"github.com/fdaines/arch-go/impl/contents"
	"github.com/fdaines/arch-go/impl/dependencies"
	"github.com/fdaines/arch-go/impl/model"
	"github.com/fdaines/arch-go/utils/output"
	"github.com/fdaines/arch-go/utils/packages"
)

func CheckArchitecture(config *config.Config, mainPackage string, pkgs []*packages.PackageInfo) *model.Result {
	result := &model.Result{}

	output.Printf("Module: %s\n", mainPackage)
	result.DependenciesRulesResults = checkDependencies(config.DependenciesRules, mainPackage, pkgs)
	result.ContentsRuleResults = checkContents(config.ContentRules, mainPackage, pkgs)

	return result
}

func checkDependencies(rules []config.DependenciesRule, mainPackage string, pkgs []*packages.PackageInfo) []*model.DependenciesRuleResult {
	results := []*model.DependenciesRuleResult{}
	for _, r := range rules {
		results = dependencies.CheckDependenciesRule(results, r, mainPackage, pkgs)
	}
	return results
}

func checkContents(rules []config.ContentsRule, mainPackage string, pkgs []*packages.PackageInfo) []*model.ContentsRuleResult {
	results := []*model.ContentsRuleResult{}
	for _, rule := range rules {
		output.Printf("%+v\n", rule.Package)
		results = contents.CheckRule(results, rule, mainPackage, pkgs)
	}
	output.Printf("Resultados: %+v\n", results)
	return results
}
