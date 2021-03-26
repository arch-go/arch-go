package packages

import (
	"fmt"
	"go/build"
	"golang.org/x/tools/go/packages"
)

func GetBasicPackagesInfo() ([]*PackageInfo, error) {
	var packagesInfo []*PackageInfo
	var context = build.Default

	pkgs, err := getPackages()
	if err != nil {
		return nil, fmt.Errorf("Error: %v\n", err)
	} else {
		for index, packageName := range pkgs {
			fmt.Printf("Loading package (%d/%d): %s\n", index+1, len(pkgs), packageName)
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

	return packagesInfo, nil
}

func getPackages() ([]string, error) {
	fmt.Printf("Looking for packages.\n")
	cfg := &packages.Config{}
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return nil, fmt.Errorf("Cannot load module packages: %+v", err)
	}
	var packages []string
	for _, pkg := range pkgs {
		packages = append(packages, pkg.PkgPath)
	}
	fmt.Printf("%d packages found...\n", len(packages))
	return packages, nil
}
