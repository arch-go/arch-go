package packages

import (
	"errors"
	"fmt"
	"go/build"
	"golang.org/x/tools/go/packages"
)

func GetBasicPackagesInfo() []*PackageInfo {
	var packagesInfo []*PackageInfo
	var context = build.Default

	pkgs, err := getPackages()
	if err != nil {
		fmt.Printf(
			"Error: %v\n",
			err,
		)
	} else {
		for index, packageName := range pkgs {
			fmt.Printf("Loading package (%d/%d): %s", index+1, len(pkgs), packageName)
			pkg, err := context.Import(packageName, "", 0)
			if err == nil {
				packagesInfo = append(packagesInfo, &PackageInfo{
					PackageData: pkg,
					Name:        pkg.Name,
					Path:        pkg.ImportPath,
				})
			}
		}
	}

	return packagesInfo
}

func getPackages() ([]string, error) {
	fmt.Printf("Looking for packages.")
	cfg := &packages.Config{}
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot load module packages: %+v", err))
	}
	var packages []string
	for _, pkg := range pkgs {
		packages = append(packages, pkg.PkgPath)
	}
	fmt.Printf("%d packages found...", len(packages))
	return packages, nil
}
