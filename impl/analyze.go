package impl

import (
	"github.com/fatih/color"
	"github.com/fdaines/arch-go/config"
	"github.com/fdaines/arch-go/impl/dependencies"
	"github.com/fdaines/arch-go/model"
	"github.com/fdaines/arch-go/utils/output"
	"github.com/fdaines/arch-go/utils/packages"
	"os"
)

func CheckArchitecture(config *config.Config, mainPackage string, pkgs []*packages.PackageInfo) {
	output.Printf("Module: %s\n", mainPackage)
	dependenciesResult := checkDependencies(config.DependenciesRules, mainPackage, pkgs)
	output.Print("--------------------------------------")
	success := true
	for _,dr := range dependenciesResult {
		if dr.Passes {
			color.Green("[PASS] - %s\n", dr.Description)
		} else {
			success = false
			color.Red("[FAIL] - %s\n", dr.Description)
			for _,fd := range dr.Failures {
				color.Red("\tPackage '%s'\n", fd.Package)
				for _,str := range fd.Details {
					color.Red("\t\t%s\n", str)
				}
			}
		}
	}

	if !success {
		os.Exit(1)
	}
}

func checkDependencies(rules []config.DependenciesRule, mainPackage string, pkgs []*packages.PackageInfo) []*model.DependenciesRuleResult {
	results := []*model.DependenciesRuleResult{}
	for _, r := range rules {
		results  = dependencies.CheckDependenciesRule(results, r, mainPackage, pkgs)
	}
	return results
}
