package cycles

import (
	"fmt"
	"github.com/fdaines/arch-go/config"
	"github.com/fdaines/arch-go/impl/model"
	"github.com/fdaines/arch-go/utils/arrays"
	"github.com/fdaines/arch-go/utils/packages"
	"github.com/fdaines/arch-go/utils/text"
	"regexp"
	"strings"
)

func CheckRule(results []*model.CyclesRuleResult, rule config.CyclesRule, mainPackage string, pkgs []*packages.PackageInfo) []*model.CyclesRuleResult {
	result := &model.CyclesRuleResult{
		Description: fmt.Sprintf("Package '%s' should not have cycles", rule.Package),
		Passes:      true,
	}
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(rule.Package))
	pkgsMap := makePackageInfoMap(pkgs)
	for _, p := range pkgs {
		if packageRegExp.MatchString(p.Path) {
			if rule.ShouldNotContainCycles {
				hasCycles, cycle := searchForCycles(p, mainPackage, pkgsMap)
				if hasCycles {
					failure := &model.CyclesRuleResultDetail{
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

func makePackageInfoMap(pkgs []*packages.PackageInfo) map[string]*packages.PackageInfo {
	pkgsMap := make(map[string]*packages.PackageInfo)
	for _, p := range pkgs {
		pkgsMap[p.Path] = p
	}
	return pkgsMap
}

func searchForCycles(p *packages.PackageInfo, mainPackage string, pkgsMap map[string]*packages.PackageInfo) (bool, []string) {
	return checkDependencies([]string{}, p, mainPackage, pkgsMap)
}

func checkDependencies(imports []string, p *packages.PackageInfo, mainPackage string, pkgsMap map[string]*packages.PackageInfo) (bool, []string) {
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
