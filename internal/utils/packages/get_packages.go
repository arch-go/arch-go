package packages

import (
	"fmt"
	"go/build"

	"github.com/fdaines/arch-go/internal/model"
	"golang.org/x/tools/go/packages"
)

func GetBasicPackagesInfo(mainPackage string, printInfo bool) ([]*model.PackageInfo, error) {
	var packagesInfo []*model.PackageInfo
	context := build.Default

	pkgs, err := getPackages(mainPackage, printInfo)
	if err != nil {
		return nil, fmt.Errorf("Error: %v\n", err)
	} else {
		for index, packageName := range pkgs {
			if printInfo {
				fmt.Printf("Loading package (%d/%d): %s\n", index+1, len(pkgs), packageName)
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

func getPackages(mainPackage string, printInfo bool) ([]string, error) {
	if printInfo {
		fmt.Print("Looking for packages.\n")
	}
	pkgs, err := packages.Load(&packages.Config{Tests: false}, mainPackage+"/...")
	if err != nil {
		return nil, fmt.Errorf("Cannot load module packages: %+v", err)
	}
	var packages []string
	for _, pkg := range pkgs {
		packages = append(packages, pkg.PkgPath)
	}
	if printInfo {
		fmt.Printf("%v packages found...\n", len(packages))
	}
	return packages, nil
}
