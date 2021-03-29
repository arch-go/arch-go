package contents

import (
	"github.com/fdaines/arch-go/config"
	"github.com/fdaines/arch-go/impl/model"
	"github.com/fdaines/arch-go/utils/output"
	"github.com/fdaines/arch-go/utils/packages"
	"github.com/fdaines/arch-go/utils/text"
	"regexp"
)

func CheckRule(results []*model.ContentsRuleResult, rule config.ContentsRule, mainPackage string, pkgs []*packages.PackageInfo) []*model.ContentsRuleResult {
	packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(rule.Package))
	for _, p := range pkgs {
		if packageRegExp.MatchString(p.Path) {
			output.Printf("Check package2: %s\n", p.Path)
			contents, _ := retrieveContents(p, mainPackage)
			results = checkInterfaces(results, contents, rule)
			results = checkTypes(results, contents, rule)
			results = checkMethods(results, contents, rule)
			results = checkFunctions(results, contents, rule)
		}
	}

	return results
}

func checkInterfaces(results []*model.ContentsRuleResult, contents *PackageContents, rule config.ContentsRule) []*model.ContentsRuleResult {
	output.Printf("Check Interfaces: %d\n", contents.Interfaces)
	if contents.Interfaces > 0 {
		if rule.ShouldNotContainInterfaces {
			return appendError(results, rule.Package, "should not contain interfaces")
		}
		if rule.ShouldOnlyContainTypes {
			return appendError(results, rule.Package, "should only contain types")
		}
		if rule.ShouldOnlyContainMethods {
			return appendError(results, rule.Package, "should only contain methods")
		}
		if rule.ShouldOnlyContainFunctions {
			return appendError(results, rule.Package, "should only contain functions")
		}
		if rule.ShouldOnlyContainInterfaces {
			return appendSuccess(results, rule.Package, "should only contain interfaces")
		}
	}
	if rule.ShouldNotContainInterfaces {
		return appendSuccess(results, rule.Package, "should not contain interfaces")
	}
	return results
}

func checkTypes(results []*model.ContentsRuleResult, contents *PackageContents, rule config.ContentsRule) []*model.ContentsRuleResult {
	output.Printf("Check Types: %d\n", contents.Types)
	if contents.Types > 0 {
		if rule.ShouldNotContainTypes {
			return appendError(results, rule.Package, "should not contain types")
		}
		if rule.ShouldOnlyContainInterfaces {
			return appendError(results, rule.Package, "should only contain interfaces")
		}
		if rule.ShouldOnlyContainMethods {
			return appendError(results, rule.Package, "should only contain methods")
		}
		if rule.ShouldOnlyContainFunctions {
			return appendError(results, rule.Package, "should only contain functions")
		}
		if rule.ShouldOnlyContainTypes {
			return appendSuccess(results, rule.Package, "should only contain types")
		}
	}
	if rule.ShouldNotContainTypes {
		return appendSuccess(results, rule.Package, "should not contain types")
	}
	return results
}

func checkMethods(results []*model.ContentsRuleResult, contents *PackageContents, rule config.ContentsRule) []*model.ContentsRuleResult {
	output.Printf("Check Methods: %d\n", contents.Methods)
	if contents.Methods > 0 {
		if rule.ShouldNotContainMethods {
			return appendError(results, rule.Package, "should not contain methods")
		}
		if rule.ShouldOnlyContainTypes {
			return appendError(results, rule.Package, "should only contain types")
		}
		if rule.ShouldOnlyContainInterfaces {
			return appendError(results, rule.Package, "should only contain interfaces")
		}
		if rule.ShouldOnlyContainFunctions {
			return appendError(results, rule.Package, "should only contain functions")
		}
		if rule.ShouldOnlyContainMethods {
			return appendSuccess(results, rule.Package, "should only contain methods")
		}
	}
	if rule.ShouldNotContainMethods {
		return appendSuccess(results, rule.Package, "should not contain methods")
	}
	return results
}

func checkFunctions(results []*model.ContentsRuleResult, contents *PackageContents, rule config.ContentsRule) []*model.ContentsRuleResult {
	output.Printf("Check Functions: %d\n", contents.Functions)
	if contents.Functions > 0 {
		if rule.ShouldNotContainFunctions {
			return appendError(results, rule.Package, "should not contain functions")
		}
		if rule.ShouldOnlyContainTypes {
			return appendError(results, rule.Package, "should only contain types")
		}
		if rule.ShouldOnlyContainMethods {
			return appendError(results, rule.Package, "should only contain methods")
		}
		if rule.ShouldOnlyContainInterfaces {
			return appendError(results, rule.Package, "should only contain interfaces")
		}
		if rule.ShouldOnlyContainFunctions {
			return appendSuccess(results, rule.Package, "should only contain functions")
		}
	}
	if rule.ShouldNotContainFunctions {
		return appendSuccess(results, rule.Package, "should not contain functions")
	}
	return results
}
