package configuration

import (
	"github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/internal/utils/output"
	"github.com/fdaines/arch-go/internal/utils/packages"
)

func Load(mainPackage string) model.ModuleInfo {
	packages, _ := packages.GetBasicPackagesInfo(mainPackage, output.CreateNilWriter(), false)

	return model.ModuleInfo{
		MainPackage: mainPackage,
		Packages:    packages,
	}
}
