package naming

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/arch-go/arch-go/v2/internal/utils/output"
	"github.com/arch-go/arch-go/v2/internal/utils/packages"
	"github.com/arch-go/arch-go/v2/internal/utils/text"

	"github.com/arch-go/arch-go/v2/api/configuration"
	"github.com/arch-go/arch-go/v2/internal/model"
)

func CheckRules(moduleInfo model.ModuleInfo, rules []*configuration.NamingRule) *RulesResult {
	result := &RulesResult{
		Passes: true,
	}

	for _, it := range rules {
		result.Results = append(result.Results, CheckRule(moduleInfo, *it))
	}

	// Update result.Pass based on each rule result
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
	result := true
	var details []string

	if rule.InterfaceImplementationNamingRule != nil {
		pass1, details1 := checkInternalInterfaces(pkg, rule, module)
		pass2, details2 := checkExternalInterfaces(pkg, rule, module)
		pass3, details3 := checkStandardInterfaces(pkg, rule, module)

		result = result && pass1 && pass2 && pass3
		details = append(details, details1...)
		details = append(details, details2...)
		details = append(details, details3...)
	}

	return result, details
}

func checkInternalInterfaces(pkg *model.PackageInfo, rule configuration.NamingRule, module model.ModuleInfo) (bool, []string) {
	if rule.InterfaceImplementationNamingRule.StructsThatImplement.Internal == nil {
		return true, nil
	}

	interfaces, err := getInterfacesMatching(pkg, *rule.InterfaceImplementationNamingRule.StructsThatImplement.Internal)
	if err != nil {
		return false, []string{err.Error()}
	}

	return checkInterfaceImplementationNamingRule(interfaces, rule, module.Packages)
}

func checkExternalInterfaces(_ *model.PackageInfo, rule configuration.NamingRule, module model.ModuleInfo) (bool, []string) {
	if rule.InterfaceImplementationNamingRule.StructsThatImplement.External == nil {
		return true, nil
	}

	pkgs, err := packages.GetBasicPackagesInfo(rule.InterfaceImplementationNamingRule.StructsThatImplement.External.Package, output.CreateNilWriter(), false)
	if err != nil {
		return false, []string{err.Error()}
	}

	var interfaces []InterfaceDescription
	for _, pkg := range pkgs {
		pkgInterfaces, err := getInterfacesMatching(pkg, rule.InterfaceImplementationNamingRule.StructsThatImplement.External.Interface)
		if err != nil {
			return false, []string{err.Error()}
		}
		interfaces = append(interfaces, pkgInterfaces...)
	}

	return checkInterfaceImplementationNamingRule(interfaces, rule, module.Packages)
}

func checkStandardInterfaces(_ *model.PackageInfo, rule configuration.NamingRule, module model.ModuleInfo) (bool, []string) {
	if rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard == nil {
		return true, nil
	}

	pkgs, err := packages.GetBasicPackagesInfo(rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard.Package, output.CreateNilWriter(), false)
	if err != nil {
		return false, []string{err.Error()}
	}

	var interfaces []InterfaceDescription
	for _, pkg := range pkgs {
		pkgInterfaces, err := getInterfacesMatching(pkg, rule.InterfaceImplementationNamingRule.StructsThatImplement.Standard.Interface)
		if err != nil {
			return false, []string{err.Error()}
		}
		interfaces = append(interfaces, pkgInterfaces...)
	}

	return checkInterfaceImplementationNamingRule(interfaces, rule, module.Packages)
}

func checkInterfaceImplementationNamingRule(
	interfaces []InterfaceDescription, rule configuration.NamingRule, pkgs []*model.PackageInfo,
) (bool, []string) {
	var (
		details []string
		passes  bool
	)

	ruleResult := true

	for _, pkg := range pkgs {
		if packageMustBeAnalyzed(pkg, rule.Package) {
			passes, details = analyzePackage(interfaces, pkg, details, rule)
			ruleResult = ruleResult && passes
		}
	}

	return ruleResult, details
}

func analyzePackage(
	interfaces []InterfaceDescription,
	pkg *model.PackageInfo,
	details []string,
	rule configuration.NamingRule,
) (bool, []string) {
	structs, _ := getStructsWithMethods(pkg)
	passes, details := analyzeStructs(interfaces, pkg, details, rule, structs)

	return passes, details
}

func analyzeStructs(
	interfaces []InterfaceDescription,
	pkg *model.PackageInfo,
	details []string,
	rule configuration.NamingRule,
	structs []StructDescription,
) (bool, []string) {
	passes := true

	if len(structs) > 0 {
		for _, strkt := range structs {
			for _, iface := range interfaces {
				pass := checkStruct(strkt, iface, rule.InterfaceImplementationNamingRule)
				if !pass {
					passes = false

					details = append(details,
						fmt.Sprintf("Struct [%s] in Package [%s] does not match Naming Rule", strkt.Name, pkg.Path))
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
