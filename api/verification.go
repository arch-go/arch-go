package api

import (
	"github.com/fdaines/arch-go/api/configuration"
	"github.com/fdaines/arch-go/internal/model"
)

// CheckArchitecture runs the architecture analysis and return the Result.
func CheckArchitecture(moduleInfo model.ModuleInfo, config configuration.Config) *Result {
	architecture := newArchitectureAnalysis(moduleInfo, config)
	result, _ := architecture.Execute()

	return result
}
