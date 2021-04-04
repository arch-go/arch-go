package cycles

import (
	"fmt"
	"github.com/fdaines/arch-go/config"
	"github.com/fdaines/arch-go/model"
	"github.com/fdaines/arch-go/model/result"
	"github.com/fdaines/arch-go/utils/arrays"
	"github.com/fdaines/arch-go/utils/text"
	"regexp"
	"strings"
)

type CycleRule struct {
	results []*result.CyclesRuleResult
	rule *config.CyclesRule
	module *model.ModuleInfo
}

func NewCycleRule(results []*result.CyclesRuleResult, rule *config.CyclesRule, module *model.ModuleInfo) *CycleRule {
	return &CycleRule{
		rule: rule,
		results: results,
		module: module,
	}
}

func (c CycleRule) CheckRule() []*result.CyclesRuleResult {
	resultInstance := &result.CyclesRuleResult{
		Description: fmt.Sprintf("Packages matching pattern '%s' should not have cycles", c.rule.Package),
		Passes:      true,
	}
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(c.rule.Package))
	pkgsMap := makePackageInfoMap(c.module.Packages)
	for _, p := range c.module.Packages {
		if packageRegExp.MatchString(p.Path) {
			if c.rule.ShouldNotContainCycles {
				hasCycles, cycle := searchForCycles(p, c.module.MainPackage, pkgsMap)
				if hasCycles {
					failure := &result.CyclesRuleResultDetail{
						Package: p.Path,
						Details: cycle,
					}
					resultInstance.Failures = append(resultInstance.Failures, failure)
				}
			}
		}
	}
	resultInstance.Passes = len(resultInstance.Failures) == 0
	c.results = append(c.results, resultInstance)

	return c.results
}

func makePackageInfoMap(pkgs []*model.PackageInfo) map[string]*model.PackageInfo {
	pkgsMap := make(map[string]*model.PackageInfo)
	for _, p := range pkgs {
		pkgsMap[p.Path] = p
	}
	return pkgsMap
}

func searchForCycles(p *model.PackageInfo, mainPackage string, pkgsMap map[string]*model.PackageInfo) (bool, []string) {
	return checkDependencies([]string{}, p, mainPackage, pkgsMap)
}

func checkDependencies(imports []string, p *model.PackageInfo, mainPackage string, pkgsMap map[string]*model.PackageInfo) (bool, []string) {
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
