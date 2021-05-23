package impl

import (
	"fmt"
	"github.com/fdaines/arch-go/config"
	"github.com/fdaines/arch-go/impl/dependencies"
	"github.com/fdaines/arch-go/impl/model"
	baseModel "github.com/fdaines/arch-go/model"
	"github.com/fdaines/arch-go/utils"
	"github.com/fdaines/arch-go/utils/packages"
	"github.com/fdaines/arch-go/utils/text"
	"os"
	"regexp"
)

func CheckArchitecture() bool {
	returnValue := true
	utils.ExecuteWithTimer(func() {
		configuration, err := config.LoadConfig("arch-go.yml")
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
			os.Exit(1)
		} else {
			mainPackage, _ := packages.GetMainPackage()
			pkgs, _ := packages.GetBasicPackagesInfo()
			moduleInfo := &baseModel.ModuleInfo{
				MainPackage: mainPackage,
				Packages:    pkgs,
			}

			verifications := resolveVerifications(configuration, moduleInfo)
			fmt.Printf("%+v\n", verifications)
			for _,v := range verifications {
				v.Verify()
			}
			for _,v := range verifications {
				v.PrintResults()
			}
//			result := checkArchitectureRules(configuration, moduleInfo)
//			returnValue = checkResult(result)
		}
	})
	return returnValue
}

func resolveVerifications(configuration *config.Config, moduleInfo *baseModel.ModuleInfo) []model.RuleVerification {
	var verifications []model.RuleVerification
	for _,dependencyRule := range configuration.DependenciesRules {
		verificationInstance := &dependencies.DependencyRuleVerification{
			Module: moduleInfo.MainPackage,
			Rule: dependencyRule,
			Passes: true,
		}
		packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(dependencyRule.Package))
		for _, pkg := range moduleInfo.Packages {
			if packageRegExp.MatchString(pkg.Path) {
				verificationInstance.PackageDetails = append(verificationInstance.PackageDetails, model.PackageVerification{
					Package: pkg,
					Passes: false,
				})
			}
		}
		verifications = append(verifications, verificationInstance)
	}
	/*
	for _,contentRule := range configuration.ContentRules {
		packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(contentRule.Package))
		for _, pkg := range moduleInfo.Packages {
			if packageRegExp.MatchString(pkg.Path) {
				verifications = append(verifications, &model.ContentsRuleVerification{
					Module: moduleInfo.MainPackage,
					Package: pkg,
					Rule: contentRule,
				})
			}
		}
	}
	for _,cycleRule := range configuration.CyclesRules {
		packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(cycleRule.Package))
		for _, pkg := range moduleInfo.Packages {
			if packageRegExp.MatchString(pkg.Path) {
				verifications = append(verifications, &model.CyclesRuleVerification{
					Module: moduleInfo.MainPackage,
					Package: pkg,
					Rule: cycleRule,
				})
			}
		}
	}
	for _,functionRule := range configuration.FunctionsRules {
		packageRegExp, _ := regexp.Compile(text.PreparePackageRegexp(functionRule.Package))
		for _, pkg := range moduleInfo.Packages {
			if packageRegExp.MatchString(pkg.Path) {
				verifications = append(verifications, &model.FunctionsRuleVerification{
					Module: moduleInfo.MainPackage,
					Package: pkg,
					Rule: functionRule,
				})
			}
		}
	}
	*/
	return verifications
}
