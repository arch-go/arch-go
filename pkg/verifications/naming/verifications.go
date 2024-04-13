package naming

import (
	"fmt"
	"regexp"

	"github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/internal/utils/text"
	"github.com/fdaines/arch-go/pkg/config"
)

func CheckRules(moduleInfo model.ModuleInfo, functionRules []*config.NamingRule) *RulesResult {
	result := &RulesResult{
		Passes: true,
	}

	for _, it := range functionRules {
		result.Results = append(result.Results, CheckRule(moduleInfo, *it))
	}

	// Update result.Passes based on each rule result
	for _, r := range result.Results {
		result.Passes = result.Passes && r.Passes
	}

	return result
}

func CheckRule(moduleInfo model.ModuleInfo, rule config.NamingRule) *RuleResult {
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

func checkNamingRule(pkg *model.PackageInfo, rule config.NamingRule, module model.ModuleInfo) (bool, []string) {
	if rule.InterfaceImplementationNamingRule != nil {
		interfaces, _ := getInterfacesMatching(pkg, rule.InterfaceImplementationNamingRule.StructsThatImplement)
		return checkInterfaceImplementationNamingRule(interfaces, rule, module.Packages)
	}

	return true, []string{}
}

func checkInterfaceImplementationNamingRule(interfaces []InterfaceDescription, rule config.NamingRule, pkgs []*model.PackageInfo) (bool, []string) {
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(rule.Package))

	var details []string
	var passes bool

	for _, pkg := range pkgs {
		if pkg != nil && packageRegExp.MatchString(pkg.Path) {
			structs, _ := getStructsWithMethods(pkg)
			if len(structs) > 0 {
				passes = true
				for _, s := range structs {
					for _, i := range interfaces {
						pass := checkStruct(s, i)
						if !pass {
							passes = false
							details = append(details, fmt.Sprintf("Struct [%s] in Package [%s] does not match Naming Rule", s.Name, pkg.Path))
						}
					}
				}
			}

		}
	}

	return passes, details
}

func checkStruct(s StructDescription, i InterfaceDescription) bool {
	return false
}
