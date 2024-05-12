package api

import (
	"github.com/fdaines/arch-go/api/configuration"
	"github.com/fdaines/arch-go/internal/model"
)

func CheckArchitecture(moduleInfo model.ModuleInfo, config configuration.Config) *Result {
	architecture := NewArchitectureAnalysis(moduleInfo, config)
	result, _ := architecture.Execute()

	return result
}
