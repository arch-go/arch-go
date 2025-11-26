package packages

import (
	"fmt"
	"go/build"
	"io"

	"golang.org/x/tools/go/packages"

	"github.com/arch-go/arch-go/v2/internal/model"
)

func GetBasicPackagesInfo(mainPackage string, out io.Writer, printInfo bool) ([]*model.PackageInfo, error) {
	var packagesInfo []*model.PackageInfo

	context := build.Default

	pkgs, err := getPackages(mainPackage, out, printInfo)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	for index, packageName := range pkgs {
		if printInfo {
			fmt.Fprintf(out, "Loading package (%d/%d): %s\n", index+1, len(pkgs), packageName)
		}

		pkg, errX := context.Import(packageName, "", 0)
		if errX == nil {
			packagesInfo = append(packagesInfo, &model.PackageInfo{
				PackageData: pkg,
				Name:        pkg.Name,
				Path:        pkg.ImportPath,
			})
		}
	}

	return packagesInfo, nil
}

func getPackages(mainPackage string, out io.Writer, printInfo bool) ([]string, error) {
	if printInfo {
		fmt.Fprint(out, "Looking for packages.\n")
	}

	pattern := mainPackage + "/..."
	if mainPackage == "builtin" {
		pattern = mainPackage
	}

	pkgs, err := packages.Load(&packages.Config{Tests: false}, pattern)
	if err != nil {
		return nil, fmt.Errorf("cannot load module packages: %w", err)
	}

	packages := make([]string, len(pkgs))
	for i, pkg := range pkgs {
		packages[i] = pkg.PkgPath
	}

	if printInfo {
		fmt.Fprintf(out, "%v packages found...\n", len(packages))
	}

	return packages, nil
}
