package archgo

import (
	"github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/pkg/archgo/configuration"
)

func CheckArchitecture(moduleInfo model.ModuleInfo, config configuration.Config) *Result {
	architecture := NewArchitectureAnalysis(moduleInfo, config)
	result, _ := architecture.Execute()

	return result
}
