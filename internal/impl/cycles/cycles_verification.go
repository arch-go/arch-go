package cycles

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/fdaines/arch-go/internal/config"
	"github.com/fdaines/arch-go/internal/impl/model"
	baseModel "github.com/fdaines/arch-go/internal/model"
)

type CyclesRuleVerification struct {
	Module         string
	ModulePackages []*baseModel.PackageInfo
	Description    string
	Rule           *config.CyclesRule
	PackageDetails []model.PackageVerification
	Passes         bool
}

func NewCyclesRuleVerification(module string, packages []*baseModel.PackageInfo, rule *config.CyclesRule) *CyclesRuleVerification {
	return &CyclesRuleVerification{
		Module:         module,
		ModulePackages: packages,
		Rule:           rule,
		Description:    fmt.Sprintf("Packages matching pattern '%s' should not have cycles", rule.Package),
		Passes:         true,
	}
}

func (d *CyclesRuleVerification) Verify() bool {
	pkgsMap := makePackageInfoMap(d.ModulePackages)
	result := true
	for index, pd := range d.PackageDetails {
		packagePasses := true
		hasCycles, cycle := searchForCycles(pd.Package, d.Module, pkgsMap)
		if hasCycles {
			packagePasses = false
			d.PackageDetails[index].Details = append(d.PackageDetails[index].Details, cycle...)
		}
		d.PackageDetails[index].Passes = packagePasses
		result = result && packagePasses
	}
	d.Passes = result
	return d.Passes
}

func (d *CyclesRuleVerification) Type() string {
	return "CycleRule"
}

func (d *CyclesRuleVerification) Status() bool {
	return d.Passes
}

func (d *CyclesRuleVerification) Name() string {
	return d.Description
}

func (d *CyclesRuleVerification) PrintResults() {
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
