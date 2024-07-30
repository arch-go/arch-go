package configuration

import (
	"github.com/arch-go/arch-go/internal/model"
	"github.com/arch-go/arch-go/internal/utils/output"
	"github.com/arch-go/arch-go/internal/utils/packages"
)

// Load takes the passed package as the main package and loads module information.
func Load(mainPackage string) model.ModuleInfo {
	packages, _ := packages.GetBasicPackagesInfo(mainPackage, output.CreateNilWriter(), false)

	return model.ModuleInfo{
		MainPackage: mainPackage,
		Packages:    packages,
	}
}
