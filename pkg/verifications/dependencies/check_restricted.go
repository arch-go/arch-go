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
	failure := false
	if !strings.HasPrefix(pkg, moduleInfo.MainPackage) && packages.IsStandardPackage(pkg) {
		for _, restrictedImport := range restricted {
			restrictedImportRegexp, _ := regexp.Compile(text.PreparePackageRegexp(restrictedImport))
			failure = failure || restrictedImportRegexp.MatchString(pkg)
		}
		if failure {
			details = append(details, fmt.Sprintf("ShouldNotDependsOn.Standard rule contains imported package '%s'", pkg))
		}
	}

	return !failure, details
}

func checkRestrictedExternalImports(pkg string, restricted []string, moduleInfo model.ModuleInfo) (bool, []string) {
	if len(restricted) == 0 {
		return true, nil
	}
	var details []string
	failure := false
	if !strings.HasPrefix(pkg, moduleInfo.MainPackage) && packages.IsExternalPackage(pkg) {
		for _, restrictedImport := range restricted {
			restrictedImportRegexp, _ := regexp.Compile(text.PreparePackageRegexp(restrictedImport))
			failure = failure || restrictedImportRegexp.MatchString(pkg)
		}
		if failure {
			details = append(details, fmt.Sprintf("ShouldNotDependsOn.External rule contains imported package '%s'", pkg))
		}
	}

	return !failure, details
}

func checkRestrictedInternalImports(pkg string, restricted []string, moduleInfo model.ModuleInfo) (bool, []string) {
	if len(restricted) == 0 {
		return true, nil
	}
	var details []string
	failure := false
	if strings.HasPrefix(pkg, moduleInfo.MainPackage) {
		for _, restrictedImport := range restricted {
			restrictedImportRegexp, _ := regexp.Compile(text.PreparePackageRegexp(restrictedImport))
			failure = failure || restrictedImportRegexp.MatchString(pkg)
		}
		if failure {
			details = append(details, fmt.Sprintf("ShouldNotDependsOn.Internal rule contains imported package '%s'", pkg))
		}
	}

	return !failure, details
}
