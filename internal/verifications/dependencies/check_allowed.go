package dependencies

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/arch-go/arch-go/v2/internal/model"
	"github.com/arch-go/arch-go/v2/internal/utils/packages"
	"github.com/arch-go/arch-go/v2/internal/utils/text"
)

func checkAllowedStandardImports(pkg string, allowed []string, moduleInfo model.ModuleInfo) (bool, []string) {
	if len(allowed) == 0 {
		return true, nil
	}

	var details []string

	success := true

	if !strings.HasPrefix(pkg, moduleInfo.MainPackage) && packages.IsStandardPackage(pkg) {
		success = false

		for _, allowedImport := range allowed {
			allowedImportRegexp, _ := regexp.Compile(text.PreparePackageRegexp(allowedImport))
			success = success || allowedImportRegexp.MatchString(pkg)
		}

		if !success {
			details = append(details,
				fmt.Sprintf("ShouldOnlyDependsOn.Standard rule doesn't contains imported package '%s'", pkg))
		}
	}

	return success, details
}

func checkAllowedExternalImports(pkg string, allowed []string, moduleInfo model.ModuleInfo) (bool, []string) {
	if len(allowed) == 0 {
		return true, nil
	}

	var details []string

	success := true

	if !strings.HasPrefix(pkg, moduleInfo.MainPackage) && packages.IsExternalPackage(pkg) {
		success = false

		for _, allowedImport := range allowed {
			allowedImportRegexp, _ := regexp.Compile(text.PreparePackageRegexp(allowedImport))
			success = success || allowedImportRegexp.MatchString(pkg)
		}

		if !success {
			details = append(details,
				fmt.Sprintf("ShouldOnlyDependsOn.External rule doesn't contains imported package '%s'", pkg))
		}
	}

	return success, details
}

func checkAllowedInternalImports(pkg string, allowed []string, moduleInfo model.ModuleInfo) (bool, []string) {
	if len(allowed) == 0 {
		return true, nil
	}

	var details []string

	success := true

	if strings.HasPrefix(pkg, moduleInfo.MainPackage) {
		success = false

		for _, allowedImport := range allowed {
			allowedImportRegexp, _ := regexp.Compile(text.PreparePackageRegexp(allowedImport))
			success = success || allowedImportRegexp.MatchString(pkg)
		}

		if !success {
			details = append(details,
				fmt.Sprintf("ShouldOnlyDependsOn.Internal rule doesn't contains imported package '%s'", pkg))
		}
	}

	return success, details
}
