package dependencies

import (
	"regexp"

	"github.com/fdaines/arch-go/api/configuration"
	"github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/internal/utils/text"
)

func CheckRules(moduleInfo model.ModuleInfo, rules []*configuration.DependenciesRule) *RulesResult {
	result := &RulesResult{
		Passes: true,
	}

	for _, it := range rules {
		result.Results = append(result.Results, CheckRule(moduleInfo, *it))
	}

	// Update result.Passes based on each rule result
	for _, r := range result.Results {
		result.Passes = result.Passes && r.Passes
	}

	return result
}

func CheckRule(moduleInfo model.ModuleInfo, rule configuration.DependenciesRule) *RuleResult {
	result := &RuleResult{
		Rule:        rule,
		Description: resolveDescription(rule),
		Passes:      true,
	}

	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(rule.Package))
	for _, it := range moduleInfo.Packages {
		if it != nil && packageRegExp.MatchString(it.Path) {
			pass, details := checkDependencies(it, rule, moduleInfo)
			result.Passes = result.Passes && pass
			result.Verifications = append(
				result.Verifications,
				Verification{
					Package: it.Path,
					Passes:  pass,
					Details: details,
				},
			)
		}
	}

	return result
}

func checkDependencies(
	pkg *model.PackageInfo, rule configuration.DependenciesRule, moduleInfo model.ModuleInfo,
) (bool, []string) {
	pass1, details1 := checkAllowedDependencies(pkg, rule, moduleInfo)
	pass2, details2 := checkRestrictedDependencies(pkg, rule, moduleInfo)

	return pass1 && pass2, append(details1, details2...)
}

func checkAllowedDependencies(
	pkg *model.PackageInfo, rule configuration.DependenciesRule, moduleInfo model.ModuleInfo,
) (bool, []string) {
	result := true

	var details []string

	if rule.ShouldOnlyDependsOn != nil {
		for _, p := range pkg.PackageData.Imports {
			pass1, details1 := checkAllowedInternalImports(p, rule.ShouldOnlyDependsOn.Internal, moduleInfo)
			pass2, details2 := checkAllowedExternalImports(p, rule.ShouldOnlyDependsOn.External, moduleInfo)
			pass3, details3 := checkAllowedStandardImports(p, rule.ShouldOnlyDependsOn.Standard, moduleInfo)

			result = result && pass1 && pass2 && pass3

			details = append(details, append(details1, append(details2, details3...)...)...)
		}
	}

	return result, details
}

func checkRestrictedDependencies(
	pkg *model.PackageInfo, rule configuration.DependenciesRule, moduleInfo model.ModuleInfo,
) (bool, []string) {
	result := true

	var details []string

	if rule.ShouldNotDependsOn != nil {
		for _, p := range pkg.PackageData.Imports {
			pass1, details1 := checkRestrictedInternalImports(p, rule.ShouldNotDependsOn.Internal, moduleInfo)
			pass2, details2 := checkRestrictedExternalImports(p, rule.ShouldNotDependsOn.External, moduleInfo)
			pass3, details3 := checkRestrictedStandardImports(p, rule.ShouldNotDependsOn.Standard, moduleInfo)

			result = result && pass1 && pass2 && pass3

			details = append(details, append(details1, append(details2, details3...)...)...)
		}
	}

	return result, details
}
