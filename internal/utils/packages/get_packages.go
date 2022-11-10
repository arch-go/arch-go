package packages

import (
	"fmt"
	"github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/internal/utils/output"
	"go/build"
	"golang.org/x/tools/go/packages"
)

func GetBasicPackagesInfo(printInfo bool) ([]*model.PackageInfo, error) {
	var packagesInfo []*model.PackageInfo
	var context = build.Default

	pkgs, err := getPackages(printInfo)
	if err != nil {
		return nil, fmt.Errorf("Error: %v\n", err)
	} else {
		for index, packageName := range pkgs {
			if printInfo {
				output.PrintVerbose("Loading package (%d/%d): %s\n", index+1, len(pkgs), packageName)
			}
			pkg, errX := context.Import(packageName, "", 0)
			if errX == nil {
				packagesInfo = append(packagesInfo, &model.PackageInfo{
					PackageData: pkg,
					Name:        pkg.Name,
					Path:        pkg.ImportPath,
					Covered:     false,
				})
			}
		}
	}

	return packagesInfo, nil
}

func getPackages(printInfo bool) ([]string, error) {
	if printInfo {
		output.Print("Looking for packages.\n")
	}
	pkgs, err := packages.Load(nil, "./...")
	if err != nil {
		return nil, fmt.Errorf("Cannot load module packages: %+v", err)
	}
	var packages []string
	for _, pkg := range pkgs {
		packages = append(packages, pkg.PkgPath)
	}
	if printInfo {
		output.Printf("%v packages found...\n", len(packages))
	}
	return packages, nil
}
