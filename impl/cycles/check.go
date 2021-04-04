package cycles

import (
	"fmt"
	"github.com/fdaines/arch-go/config"
	model2 "github.com/fdaines/arch-go/model"
	result2 "github.com/fdaines/arch-go/model/result"
	"github.com/fdaines/arch-go/utils/arrays"
	"github.com/fdaines/arch-go/utils/text"
	"regexp"
	"strings"
)

func CheckRule(results []*result2.CyclesRuleResult, rule config.CyclesRule, module *model2.ModuleInfo) []*result2.CyclesRuleResult {
	result := &result2.CyclesRuleResult{
		Description: fmt.Sprintf("Packages matching pattern '%s' should not have cycles", rule.Package),
		Passes:      true,
	}
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(rule.Package))
	pkgsMap := makePackageInfoMap(module.Packages)
	for _, p := range module.Packages {
		if packageRegExp.MatchString(p.Path) {
			if rule.ShouldNotContainCycles {
				hasCycles, cycle := searchForCycles(p, module.MainPackage, pkgsMap)
				if hasCycles {
					failure := &result2.CyclesRuleResultDetail{
						Package: p.Path,
						Details: cycle,
					}
					result.Failures = append(result.Failures, failure)
				}
			}
		}
	}
	result.Passes = len(result.Failures) == 0
	results = append(results, result)

	return results
}

func makePackageInfoMap(pkgs []*model2.PackageInfo) map[string]*model2.PackageInfo {
	pkgsMap := make(map[string]*model2.PackageInfo)
	for _, p := range pkgs {
		pkgsMap[p.Path] = p
	}
	return pkgsMap
}

func searchForCycles(p *model2.PackageInfo, mainPackage string, pkgsMap map[string]*model2.PackageInfo) (bool, []string) {
	return checkDependencies([]string{}, p, mainPackage, pkgsMap)
}

func checkDependencies(imports []string, p *model2.PackageInfo, mainPackage string, pkgsMap map[string]*model2.PackageInfo) (bool, []string) {
	for _, pkg := range p.PackageData.Imports {
		if strings.HasPrefix(pkg, mainPackage) {
			if arrays.Contains(imports, pkg) {
				return true, append(imports, pkg)
			} else {
				hasCycles, cyclicDependencyPath := checkDependencies(append(imports, pkg), pkgsMap[pkg], mainPackage, pkgsMap)
				if hasCycles {
					return true, cyclicDependencyPath
				}
			}
		}
	}

	return false, imports
}
