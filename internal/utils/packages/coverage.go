package packages

import (
	"github.com/fdaines/arch-go/internal/impl/model"
	baseModel "github.com/fdaines/arch-go/internal/model"
	customTypes "github.com/fdaines/arch-go/internal/utils/types"
	"sort"
)

func ResolveCoveredPackages(verifications []model.RuleVerification) map[string]customTypes.Void {
	var member customTypes.Void
	set := make(map[string]customTypes.Void)

	for _, v := range verifications {
		for _, vx := range v.GetVerifications() {
			if vx.Package != nil {
				set[vx.Package.Path] = member
			}
		}
	}

	return set
}

func ResolveTotalPackages(pkgs []*baseModel.PackageInfo) map[string]customTypes.Void {
	var member customTypes.Void
	set := make(map[string]customTypes.Void)

	for _, p := range pkgs {
		set[p.Path] = member
	}

	return set
}

func ResolveUncoveredPackages(modulePackages, covered map[string]customTypes.Void) []string {
	var uncoveredPackages []string
	for c := range covered {
		delete(modulePackages, c)
	}
	for p := range modulePackages {
		uncoveredPackages = append(uncoveredPackages, p)
	}

	sort.Slice(uncoveredPackages, func(i, j int) bool {
		return uncoveredPackages[i] < uncoveredPackages[j]
	})

	return uncoveredPackages
}
