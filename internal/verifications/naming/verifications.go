package naming

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fdaines/arch-go/api/configuration"

	"github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/internal/utils/text"
)

func CheckRules(moduleInfo model.ModuleInfo, rules []*configuration.NamingRule) *RulesResult {
	result := &RulesResult{
		Passes: true,
	}

	for _, it := range rules {
		result.Results = append(result.Results, CheckRule(moduleInfo, *it))
	}

	// Update result.Passes based on each rule result
	for _, r := range result.Results {
		result.Passes = result.Passes && r.Passes
	}

	return result
}

func CheckRule(moduleInfo model.ModuleInfo, rule configuration.NamingRule) *RuleResult {
	result := &RuleResult{
		Rule:        rule,
		Description: resolveDescription(rule),
		Passes:      true,
	}

	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(rule.Package))
	for _, it := range moduleInfo.Packages {
		if it != nil && packageRegExp.MatchString(it.Path) {
			pass, details := checkNamingRule(it, rule, moduleInfo)
			result.Passes = result.Passes && pass
			result.Verifications = append(
				result.Verifications,
				Verification{
					Package: it.Path,
					Passes:  pass,
					Details: details,
				},
			)
		}
	}

	return result
}

func checkNamingRule(pkg *model.PackageInfo, rule configuration.NamingRule, module model.ModuleInfo) (bool, []string) {
	if rule.InterfaceImplementationNamingRule != nil {
		interfaces, _ := getInterfacesMatching(pkg, rule.InterfaceImplementationNamingRule.StructsThatImplement)
		return checkInterfaceImplementationNamingRule(interfaces, rule, module.Packages)
	}

	return true, []string{}
}

func checkInterfaceImplementationNamingRule(interfaces []InterfaceDescription, rule configuration.NamingRule, pkgs []*model.PackageInfo) (bool, []string) {
	var details []string
	var passes bool
	ruleResult := true

	for _, pkg := range pkgs {
		if packageMustBeAnalyzed(pkg, rule.Package) {
			passes, details = analyzePackage(interfaces, pkg, details, rule)
			ruleResult = ruleResult && passes
		}
	}

	return ruleResult, details
}

func analyzePackage(interfaces []InterfaceDescription, pkg *model.PackageInfo, details []string, rule configuration.NamingRule) (bool, []string) {
	structs, _ := getStructsWithMethods(pkg)
	passes, details := analyzeStructs(interfaces, pkg, details, rule, structs)

	return passes, details
}

func analyzeStructs(interfaces []InterfaceDescription, pkg *model.PackageInfo, details []string, rule configuration.NamingRule, structs []StructDescription) (bool, []string) {
	passes := true
	if len(structs) > 0 {
		for _, s := range structs {
			for _, i := range interfaces {
				pass := checkStruct(s, i, rule.InterfaceImplementationNamingRule)
				if !pass {
					passes = false
					details = append(details, fmt.Sprintf("Struct [%s] in Package [%s] does not match Naming Rule", s.Name, pkg.Path))
				}
			}
		}
	}
	return passes, details
}

func checkStruct(s StructDescription, i InterfaceDescription, rule *configuration.InterfaceImplementationRule) bool {
	if implementsInterface(s, i) {
		return checkStructName(s.Name, rule)
	}
	return true
}

func checkStructName(name string, rule *configuration.InterfaceImplementationRule) bool {
	if rule.ShouldHaveSimpleNameEndingWith != nil {
		return strings.HasSuffix(name, *rule.ShouldHaveSimpleNameEndingWith)
	}
	if rule.ShouldHaveSimpleNameStartingWith != nil {
		return strings.HasPrefix(name, *rule.ShouldHaveSimpleNameStartingWith)
	}
	return false
}
