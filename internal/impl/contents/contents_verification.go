package contents

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/fdaines/arch-go/internal/config"
	"github.com/fdaines/arch-go/internal/impl/model"
	"strings"
)

type ContentsRuleVerification struct {
	Module         string
	Description    string
	Rule           *config.ContentsRule
	PackageDetails []model.PackageVerification
	Passes         bool
}

func NewContentsRuleVerification(module string, rule *config.ContentsRule) *ContentsRuleVerification {
	var ruleDescriptions []string
	if rule.ShouldOnlyContainStructs {
		ruleDescriptions = append(ruleDescriptions, "'should only contain structs'")
	}
	if rule.ShouldOnlyContainInterfaces {
		ruleDescriptions = append(ruleDescriptions, "'should only contain interfaces'")
	}
	if rule.ShouldOnlyContainFunctions {
		ruleDescriptions = append(ruleDescriptions, "'should only contain functions'")
	}
	if rule.ShouldOnlyContainMethods {
		ruleDescriptions = append(ruleDescriptions, "'should only contain methods'")
	}
	if rule.ShouldNotContainStructs {
		ruleDescriptions = append(ruleDescriptions, "'should not contain structs'")
	}
	if rule.ShouldNotContainInterfaces {
		ruleDescriptions = append(ruleDescriptions, "'should not contain interfaces'")
	}
	if rule.ShouldNotContainFunctions {
		ruleDescriptions = append(ruleDescriptions, "'should not contain functions'")
	}
	if rule.ShouldNotContainMethods {
		ruleDescriptions = append(ruleDescriptions, "'should not contain methods'")
	}
	description := fmt.Sprintf("Packages matching pattern '%s' should complies with [%s]", rule.Package, strings.Join(ruleDescriptions, ","))

	return &ContentsRuleVerification{
		Module:      module,
		Rule:        rule,
		Description: description,
		Passes:      true,
	}
}

func (d *ContentsRuleVerification) Verify() {
	result := true
	for index, pd := range d.PackageDetails {
		packagePasses := true
		contents, _ := retrieveContents(pd.Package, d.Module)
		fmt.Printf("Package: %+v\n", pd.Package.Path)
		fmt.Printf("Contents: %+v\n", contents)
		fmt.Printf("Rule: %+v\n", d.Rule)

		ruleResult, ruleDetails := check_interfaces(contents, d.Rule)
		packagePasses = packagePasses && ruleResult
		d.PackageDetails[index].Details = append(d.PackageDetails[index].Details, ruleDetails...)
		fmt.Printf("packagePasses1: %+v\n", packagePasses)

		ruleResult, ruleDetails = check_structs(contents, d.Rule)
		packagePasses = packagePasses && ruleResult
		d.PackageDetails[index].Details = append(d.PackageDetails[index].Details, ruleDetails...)
		fmt.Printf("packagePasses2: %+v\n", packagePasses)

		ruleResult, ruleDetails = check_functions(contents, d.Rule)
		packagePasses = packagePasses && ruleResult
		d.PackageDetails[index].Details = append(d.PackageDetails[index].Details, ruleDetails...)
		fmt.Printf("packagePasses3: %+v\n", packagePasses)

		ruleResult, ruleDetails = check_methods(contents, d.Rule)
		packagePasses = packagePasses && ruleResult
		d.PackageDetails[index].Details = append(d.PackageDetails[index].Details, ruleDetails...)
		fmt.Printf("packagePasses4: %+v\n", packagePasses)

		d.PackageDetails[index].Passes = packagePasses
		result = result && packagePasses
	}
	d.Passes = result
}

func (d *ContentsRuleVerification) PrintResults() {
	if d.Passes {
		color.Green("[PASS] - %s\n", d.Description)
	} else {
		color.Red("[FAIL] - %s\n", d.Description)
	}
	for _, p := range d.PackageDetails {
		if p.Passes {
			color.Green("\tPackage '%s' passes\n", p.Package.Path)
		} else {
			color.Red("\tPackage '%s' fails\n", p.Package.Path)
			for _, str := range p.Details {
				color.Red("\t\t%s\n", str)
			}
		}
	}
}
