package api

import (
	"github.com/arch-go/arch-go/v2/api/configuration"
	"github.com/arch-go/arch-go/v2/internal/model"
)

// CheckArchitecture runs the architecture analysis and return the Result.
func CheckArchitecture(moduleInfo model.ModuleInfo, config configuration.Config) *Result {
	architecture := newArchitectureAnalysis(moduleInfo, config)
	result, _ := architecture.Execute()

	return result
}
