package dependencies

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/internal/utils/packages"
	"github.com/fdaines/arch-go/internal/utils/text"
)

func checkRestrictedStandardImports(pkg string, restricted []string, moduleInfo model.ModuleInfo) (bool, []string) {
	if len(restricted) == 0 {
		return true, nil
	}
	var details []string
	fails := false
	if !strings.HasPrefix(pkg, moduleInfo.MainPackage) && packages.IsStandardPackage(pkg) {
		fails = false
		for _, restrictedImport := range restricted {
			restrictedImportRegexp, _ := regexp.Compile(text.PreparePackageRegexp(restrictedImport))
			fails = fails || restrictedImportRegexp.MatchString(pkg)
		}
		if fails {
			details = append(details, fmt.Sprintf("ShouldNotDependsOn.Standard rule contains imported package '%s'", pkg))
		}
	}

	return !fails, details
}

func checkRestrictedExternalImports(pkg string, restricted []string, moduleInfo model.ModuleInfo) (bool, []string) {
	if len(restricted) == 0 {
		return true, nil
	}
	var details []string
	fails := false
	if !strings.HasPrefix(pkg, moduleInfo.MainPackage) && packages.IsExternalPackage(pkg) {
		fails = false
		for _, restrictedImport := range restricted {
			restrictedImportRegexp, _ := regexp.Compile(text.PreparePackageRegexp(restrictedImport))
			fails = fails || restrictedImportRegexp.MatchString(pkg)
		}
		if fails {
			details = append(details, fmt.Sprintf("ShouldNotDependsOn.External rule contains imported package '%s'", pkg))
		}
	}

	return !fails, details
}

func checkRestrictedInternalImports(pkg string, restricted []string, moduleInfo model.ModuleInfo) (bool, []string) {
	if len(restricted) == 0 {
		return true, nil
	}
	var details []string
	fails := false
	if strings.HasPrefix(pkg, moduleInfo.MainPackage) {
		fails = false
		for _, restrictedImport := range restricted {
			restrictedImportRegexp, _ := regexp.Compile(text.PreparePackageRegexp(restrictedImport))
			fails = fails || restrictedImportRegexp.MatchString(pkg)
		}
		if fails {
			details = append(details, fmt.Sprintf("ShouldNotDependsOn.Internal rule contains imported package '%s'", pkg))
		}
	}

	return !fails, details
}
