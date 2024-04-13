package impl

import (
	"github.com/fatih/color"
	"github.com/fdaines/arch-go/internal/common"
	"github.com/fdaines/arch-go/internal/utils/timer"
	"github.com/fdaines/arch-go/old/config"
	"github.com/fdaines/arch-go/old/impl/dependencies"
	"github.com/fdaines/arch-go/old/impl/model"
	baseModel "github.com/fdaines/arch-go/old/model"
	"github.com/fdaines/arch-go/old/model/result"
	"github.com/fdaines/arch-go/old/report/console"
	"github.com/fdaines/arch-go/old/report/html"
	"github.com/fdaines/arch-go/old/utils/packages"
	"github.com/fdaines/arch-go/old/utils/text"
	"os"
	"regexp"
)

func CheckArchitecture() bool {
	returnValue := true
	timer.ExecuteWithTimer(func() {
		configuration, err := config.LoadConfig("arch-go.yml")
		if err != nil {
			color.Red("Error: %+v\n", err)
			os.Exit(1)
		} else {
			mainPackage, _ := packages.GetMainPackage()
			pkgs, _ := packages.GetBasicPackagesInfo(true)
			moduleInfo := &baseModel.ModuleInfo{
				MainPackage: mainPackage,
				Packages:    pkgs,
			}

			verifications := resolveVerifications(configuration, moduleInfo)
			validateVerifications(verifications)

			for _, v := range verifications {
				v.Verify()
				v.PrintResults()
			}

			resultData := result.ResolveReport(pkgs, verifications, configuration)
			returnValue = resultData.Summary.Status

			if common.Html {
				html.GenerateHtmlReport(resultData)
			} else {
				console.GenerateConsoleReport(resultData)
			}
			resultData.Print()

		}
	})
	return returnValue
}

func validateVerifications(verifications []model.RuleVerification) {
	ok := true
	for _, v := range verifications {
		ok = ok && v.ValidatePatterns()
	}
	if !ok {
		color.Red("Some package patterns are invalid, please check documentation: https://github.com/fdaines/arch-go\n")
		os.Exit(1)
	}
}

func resolveVerifications(configuration *config.Config, moduleInfo *baseModel.ModuleInfo) []model.RuleVerification {
	var verifications []model.RuleVerification
	verifications = resolveDependencyRules(configuration, moduleInfo, verifications)

	return verifications
}

func resolveDependencyRules(configuration *config.Config, moduleInfo *baseModel.ModuleInfo, verifications []model.RuleVerification) []model.RuleVerification {
	for _, dependencyRule := range configuration.DependenciesRules {
		verificationInstance := dependencies.NewDependencyRuleVerification(moduleInfo.MainPackage, dependencyRule)
		packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(dependencyRule.Package))
		for _, pkg := range moduleInfo.Packages {
			if packageRegExp.MatchString(pkg.Path) {
				verificationInstance.PackageDetails = append(verificationInstance.PackageDetails, baseModel.PackageVerification{
					Package: pkg,
					Passes:  false,
				})
			}
		}
		verifications = append(verifications, verificationInstance)
	}
	return verifications
}
