package naming

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/fdaines/arch-go/internal/config"
	"github.com/fdaines/arch-go/internal/impl/model"
	baseModel "github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/internal/utils/text"
	"regexp"
	"strings"
)

type NamingRuleVerification struct {
	Module         string
	Description    string
	Rule           *config.NamingRule
	PackageDetails []baseModel.PackageVerification
	Passes         bool
}

func NewNamingRuleVerification(module string, rule *config.NamingRule) *NamingRuleVerification {
	description := resolveRuleDescription(rule)

	return &NamingRuleVerification{
		Module:      module,
		Rule:        rule,
		Description: description,
		Passes:      true,
	}
}

func (d *NamingRuleVerification) Verify() bool {
	d.Passes = true

	if d.Rule.InterfaceImplementationNamingRule != nil {
		interfaces, _ := getInterfacesMatching(d.Module, d.Rule.InterfaceImplementationNamingRule.StructsThatImplement)
		d.checkInterfaceImplementationNamingRule(interfaces)
	}
	return d.Passes
}

func (d *NamingRuleVerification) checkInterfaceImplementationNamingRule(interfaces []InterfaceDescription) {
	for index, pd := range d.PackageDetails {
		packagePasses := true
		structs, _ := getStructsWithMethods(d.Module, pd)
		if len(structs) == 0 {
			d.PackageDetails[index].Details = append(d.PackageDetails[index].Details, "Package has no structs type definitions")
			d.PackageDetails[index].Passes = packagePasses
			continue
		}
		matchInterface := false
		for _, s := range structs {
			for _, i := range interfaces {
				matchInterface, packagePasses = d.checkNamingRule(s, i, matchInterface, packagePasses, index)
			}
		}
		if !matchInterface {
			d.PackageDetails[index].Details = append(d.PackageDetails[index].Details, "No structs implements interfaces matching naming rule")
		}
		d.PackageDetails[index].Passes = packagePasses
		d.Passes = d.Passes && packagePasses
	}
}

func (d *NamingRuleVerification) checkNamingRule(s StructDescription, i InterfaceDescription, matchInterface bool, packagePasses bool, index int) (bool, bool) {
	if implementsInterface(s, i) {
		matchInterface = true
		detail := fmt.Sprintf("Struct '%s' implements '%s' and complies with naming rule", s.Name, i.Name)
		if d.Rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith != "" {
			if !strings.HasSuffix(s.Name, d.Rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameEndingWith) {
				packagePasses = false
				detail = fmt.Sprintf("Struct '%s' implements '%s' but doesn't comply with naming rule", s.Name, i.Name)
			}
		}
		if d.Rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith != "" {
			if !strings.HasPrefix(s.Name, d.Rule.InterfaceImplementationNamingRule.ShouldHaveSimpleNameStartingWith) {
				packagePasses = false
				detail = fmt.Sprintf("Struct '%s' implements '%s' but doesn't comply with naming rule", s.Name, i.Name)
			}
		}
		d.PackageDetails[index].Details = append(d.PackageDetails[index].Details, detail)
	}
	return matchInterface, packagePasses
}

func (d *NamingRuleVerification) Type() string {
	return model.NamingRule
}

func (d *NamingRuleVerification) Name() string {
	return d.Description
}

func (d *NamingRuleVerification) Status() bool {
	return d.Passes
}

func (d *NamingRuleVerification) ValidatePatterns() bool {
	_, err := regexp.Compile(text.PreparePackageRegexp(d.Rule.Package))
	if err != nil {
		color.Red("[%s] - Invalid Package Pattern: %s\n", d.Description, d.Rule.Package)
		return false
	}
	return true
}

func (d *NamingRuleVerification) PrintResults() {
	if d.Passes {
		color.Green("[PASS] - %s\n", d.Description)
	} else {
		color.Red("[FAIL] - %s\n", d.Description)
	}
	for _, p := range d.PackageDetails {
		if p.Passes {
			color.Green("\tPackage '%s' passes\n", p.Package.Path)
			for _, str := range p.Details {
				color.Green("\t\t%s\n", str)
			}
		} else {
			color.Red("\tPackage '%s' fails\n", p.Package.Path)
			for _, str := range p.Details {
				color.Red("\t\t%s\n", str)
			}
		}
	}
}

func (d *NamingRuleVerification) GetVerifications() []baseModel.PackageVerification {
	return d.PackageDetails
}
