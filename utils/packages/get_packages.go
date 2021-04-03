package packages

import (
	"fmt"
	"github.com/fdaines/arch-go/model"
	"github.com/fdaines/arch-go/utils/output"
	"go/build"
	"golang.org/x/tools/go/packages"
)

func GetBasicPackagesInfo() ([]*model.PackageInfo, error) {
	var packagesInfo []*model.PackageInfo
	var context = build.Default

	pkgs, err := getPackages()
	if err != nil {
		return nil, fmt.Errorf("Error: %v\n", err)
	} else {
		for index, packageName := range pkgs {
			output.PrintVerbose("Loading package (%d/%d): %s\n", index+1, len(pkgs), packageName)
			pkg, err := context.Import(packageName, "", 0)
			if err == nil {
				packagesInfo = append(packagesInfo, &model.PackageInfo{
					PackageData: pkg,
					Name:        pkg.Name,
					Path:        pkg.ImportPath,
				})
			}
		}
	}

	return packagesInfo, nil
}

func getPackages() ([]string, error) {
	output.Print("Looking for packages.")
	cfg := &packages.Config{}
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return nil, fmt.Errorf("Cannot load module packages: %+v", err)
	}
	var packages []string
	for _, pkg := range pkgs {
		packages = append(packages, pkg.PkgPath)
	}
	output.Printf("%v packages found...\n", len(packages))
	return packages, nil
}
