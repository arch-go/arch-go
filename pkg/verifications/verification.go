package verifications

import (
	"github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/pkg/config"
)

func CheckArchitecture(moduleInfo model.ModuleInfo, config config.Config) *Result {
	architecture := NewArchitectureAnalysis(moduleInfo, config)
	result, _ := architecture.Execute()

	return result
}
