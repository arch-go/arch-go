package naming

import (
	"github.com/fatih/color"
	"github.com/fdaines/arch-go/internal/config"
	"github.com/fdaines/arch-go/internal/impl/model"
)

type NamingRuleVerification struct {
	Module         string
	Description    string
	Rule           *config.NamingRule
	PackageDetails []model.PackageVerification
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

func (d *NamingRuleVerification) Verify() {
	d.Passes = true
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
		} else {
			color.Red("\tPackage '%s' fails\n", p.Package.Path)
			for _, str := range p.Details {
				color.Red("\t\t%s\n", str)
			}
		}
	}
}
