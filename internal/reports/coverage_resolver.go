package reports

import (
	"github.com/arch-go/arch-go/api"
	"github.com/arch-go/arch-go/api/configuration"
	"github.com/arch-go/arch-go/internal/model"
	model2 "github.com/arch-go/arch-go/internal/reports/model"
	"github.com/arch-go/arch-go/internal/utils/values"
)

func resolveCoverage(
	result *api.Result, moduleInfo model.ModuleInfo, conf configuration.Config,
) *model2.ThresholdSummary {
	var uncoveredPackages []string

	moduleContents := checkPackagesCoverage(result, moduleInfo)
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

	threshold := values.GetIntRef(0)
	if conf.Threshold != nil && conf.Threshold.Coverage != nil {
		threshold = conf.Threshold.Coverage
	}

	return &model2.ThresholdSummary{
		Rate:       rate,
		Threshold:  threshold,
		Pass:       threshold == nil || rate >= *threshold,
		Violations: uncoveredPackages,
	}
}

func checkPackagesCoverage(result *api.Result, moduleInfo model.ModuleInfo) map[string]bool {
	moduleContents := make(map[string]bool)
	for _, pkg := range moduleInfo.Packages {
		moduleContents[pkg.Path] = false
	}

	checkDependenciesRules(result, moduleContents)
	checkFunctionsRules(result, moduleContents)
	checkContentsRules(result, moduleContents)
	checkNamingRules(result, moduleContents)

	return moduleContents
}

func checkNamingRules(r *api.Result, moduleContents map[string]bool) {
	if r.NamingRuleResult != nil {
		for _, dr := range r.NamingRuleResult.Results {
			for _, v := range dr.Verifications {
				updatePackage(moduleContents, v.Package)
			}
		}
	}
}

func checkContentsRules(r *api.Result, moduleContents map[string]bool) {
	if r.ContentsRuleResult != nil {
		for _, dr := range r.ContentsRuleResult.Results {
			for _, v := range dr.Verifications {
				updatePackage(moduleContents, v.Package)
			}
		}
	}
}

func checkFunctionsRules(r *api.Result, moduleContents map[string]bool) {
	if r.FunctionsRuleResult != nil {
		for _, dr := range r.FunctionsRuleResult.Results {
			for _, v := range dr.Verifications {
				updatePackage(moduleContents, v.Package)
			}
		}
	}
}

func checkDependenciesRules(r *api.Result, moduleContents map[string]bool) {
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
