package cycles

import (
	"github.com/fdaines/arch-go/model"
	"github.com/fdaines/arch-go/utils/arrays"
	"strings"
)

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
