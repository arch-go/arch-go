package reports

import (
	"github.com/fdaines/arch-go/internal/model"
	model2 "github.com/fdaines/arch-go/internal/reports/model"
	"github.com/fdaines/arch-go/pkg/archgo"
	"github.com/fdaines/arch-go/pkg/archgo/configuration"
)

func resolveCoverage(r *archgo.Result, m model.ModuleInfo, c configuration.Config) *model2.ThresholdSummary {
	moduleContents := checkPackagesCoverage(r, m)
	var uncoveredPackages []string
	for pkg, verified := range moduleContents {
		if !verified {
			uncoveredPackages = append(uncoveredPackages, pkg)
		}
	}

	totalPackages := len(moduleContents)
	coveredPackages := totalPackages - len(uncoveredPackages)
	rate := 0
	if totalPackages > 0 {
		rate = (100 * coveredPackages) / totalPackages
	}

	threshold := 0
	if c.Threshold != nil && c.Threshold.Coverage != nil {
		threshold = *c.Threshold.Coverage
	}

	status := passStatus
	if rate < threshold {
		status = failStatus
	}

	return &model2.ThresholdSummary{
		Rate:       rate,
		Threshold:  threshold,
		Status:     status,
		Violations: uncoveredPackages,
	}
}

func checkPackagesCoverage(r *archgo.Result, m model.ModuleInfo) map[string]bool {
	moduleContents := make(map[string]bool)
	for _, pkg := range m.Packages {
		moduleContents[pkg.Path] = false
	}

	checkDependenciesRules(r, moduleContents)
	checkFunctionsRules(r, moduleContents)
	checkContentsRules(r, moduleContents)
	checkNamingRules(r, moduleContents)

	return moduleContents
}

func checkNamingRules(r *archgo.Result, moduleContents map[string]bool) {
	if r.NamingRuleResult != nil {
		for _, dr := range r.NamingRuleResult.Results {
			for _, v := range dr.Verifications {
				updatePackage(moduleContents, v.Package)
			}
		}
	}
}

func checkContentsRules(r *archgo.Result, moduleContents map[string]bool) {
	if r.ContentsRuleResult != nil {
		for _, dr := range r.ContentsRuleResult.Results {
			for _, v := range dr.Verifications {
				updatePackage(moduleContents, v.Package)
			}
		}
	}
}

func checkFunctionsRules(r *archgo.Result, moduleContents map[string]bool) {
	if r.FunctionsRuleResult != nil {
		for _, dr := range r.FunctionsRuleResult.Results {
			for _, v := range dr.Verifications {
				updatePackage(moduleContents, v.Package)
			}
		}
	}
}

func checkDependenciesRules(r *archgo.Result, moduleContents map[string]bool) {
	if r.DependenciesRuleResult != nil {
		for _, dr := range r.DependenciesRuleResult.Results {
			for _, v := range dr.Verifications {
				updatePackage(moduleContents, v.Package)
			}
		}
	}
}

func updatePackage(moduleContents map[string]bool, pkg string) {
	_, ok := moduleContents[pkg]
	if ok {
		moduleContents[pkg] = true
	}
}
