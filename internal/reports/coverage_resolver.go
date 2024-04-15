package reports

import (
	"github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/pkg/config"
	"github.com/fdaines/arch-go/pkg/verifications"
)

func resolveCoverage(r *verifications.Result, m model.ModuleInfo, c config.Config) *ThresholdSummary {
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

	return &ThresholdSummary{
		Rate:       rate,
		Threshold:  threshold,
		Status:     status,
		Violations: uncoveredPackages,
	}
}

func checkPackagesCoverage(r *verifications.Result, m model.ModuleInfo) map[string]bool {
	moduleContents := make(map[string]bool)
	for _, pkg := range m.Packages {
		moduleContents[pkg.Path] = false
	}

	if r.DependenciesRuleResult != nil {
		for _, dr := range r.DependenciesRuleResult.Results {
			for _, v := range dr.Verifications {
				updatePackage(moduleContents, v.Package)
			}
		}
	}
	if r.FunctionsRuleResult != nil {
		for _, dr := range r.FunctionsRuleResult.Results {
			for _, v := range dr.Verifications {
				updatePackage(moduleContents, v.Package)
			}
		}
	}
	if r.ContentsRuleResult != nil {
		for _, dr := range r.ContentsRuleResult.Results {
			for _, v := range dr.Verifications {
				updatePackage(moduleContents, v.Package)
			}
		}
	}
	if r.NamingRuleResult != nil {
		for _, dr := range r.NamingRuleResult.Results {
			for _, v := range dr.Verifications {
				updatePackage(moduleContents, v.Package)
			}
		}
	}

	return moduleContents
}

func updatePackage(moduleContents map[string]bool, pkg string) {
	_, ok := moduleContents[pkg]
	if ok {
		moduleContents[pkg] = true
	}
}
