package contents

import (
	"regexp"

	"github.com/fdaines/arch-go/api/configuration"
	"github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/internal/utils/text"
)

func CheckRules(moduleInfo model.ModuleInfo, rules []*configuration.ContentsRule) *RulesResult {
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

func CheckRule(moduleInfo model.ModuleInfo, contentsRule configuration.ContentsRule) *RuleResult {
	result := &RuleResult{
		Rule:        contentsRule,
		Description: resolveDescription(contentsRule),
		Passes:      true,
	}

	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(contentsRule.Package))
	for _, it := range moduleInfo.Packages {
		if it != nil && packageRegExp.MatchString(it.Path) {
			contents, _ := retrieveContents(it)
			pass, details := checkContentsRule(contents, contentsRule)
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

func checkContentsRule(contents *PackageContents, functionRule configuration.ContentsRule) (bool, []string) {
	pass1, details1 := checkFunctions(contents, &functionRule)
	pass2, details2 := checkMethods(contents, &functionRule)
	pass3, details3 := checkInterfaces(contents, &functionRule)
	pass4, details4 := checkStructs(contents, &functionRule)

	return pass2 && pass1 && pass3 && pass4,
		append(details1, append(details2, append(details3, details4...)...)...)
}
