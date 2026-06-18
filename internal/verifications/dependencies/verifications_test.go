package dependencies

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/v2/api/configuration"
	"github.com/arch-go/arch-go/v2/internal/model"
)

func TestCheckRuleWithInvalidPackageSelector(t *testing.T) {
	moduleInfo := model.ModuleInfo{
		MainPackage: "example.com/repro",
		Packages: []*model.PackageInfo{
			{Path: "example.com/repro/a"},
		},
	}
	rule := configuration.DependenciesRule{
		Package: "example.com/repro/**",
		ShouldNotDependsOn: &configuration.Dependencies{
			Internal: []string{"example.com/repro/forbidden"},
		},
	}

	result := CheckRule(moduleInfo, rule)

	assert.True(t, result.Passes)
	assert.Empty(t, result.Verifications)
}
