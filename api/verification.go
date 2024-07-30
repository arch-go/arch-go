package api

import (
	"github.com/arch-go/arch-go/api/configuration"
	"github.com/arch-go/arch-go/internal/model"
)

// CheckArchitecture runs the architecture analysis and return the Result.
func CheckArchitecture(moduleInfo model.ModuleInfo, config configuration.Config) *Result {
	architecture := newArchitectureAnalysis(moduleInfo, config)
	result, _ := architecture.Execute()

	return result
}
