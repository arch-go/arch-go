package migrate_config

import (
	"testing"

	"github.com/fdaines/arch-go/api/configuration"
	"github.com/fdaines/arch-go/internal/utils/values"
	"github.com/stretchr/testify/assert"
)

func TestMigrationCommand(t *testing.T) {
	t.Run("test migrate empty configuration", func(t *testing.T) {
		originalConfig := &configuration.DeprecatedConfig{}

		expectedConfig := &configuration.Config{
			Version: 1,
		}

		result := migrateRules(originalConfig)

		assert.Equal(t, expectedConfig, result)
	})

	t.Run("test migrate configuration with cycles rules", func(t *testing.T) {
		originalConfig := &configuration.DeprecatedConfig{
			DependenciesRules: []*configuration.DeprecatedDependenciesRule{
				{
					Package:             "foobar",
					ShouldOnlyDependsOn: []string{"a", "b"},
				},
			},
			FunctionsRules: []*configuration.FunctionsRule{
				{
					Package:                  "foobar",
					MaxReturnValues:          values.GetIntRef(2),
					MaxParameters:            values.GetIntRef(3),
					MaxLines:                 values.GetIntRef(35),
					MaxPublicFunctionPerFile: values.GetIntRef(5),
				},
			},
			ContentRules: []*configuration.ContentsRule{
				{
					Package:                     "foobar",
					ShouldOnlyContainInterfaces: true,
				},
			},
			NamingRules: []*configuration.NamingRule{
				{
					Package: "foobar",
					InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
						StructsThatImplement:           "blabla",
						ShouldHaveSimpleNameEndingWith: values.GetStringRef("jojo"),
					},
				},
			},
			CyclesRules: []*configuration.CyclesRule{
				{
					Package:                "foobar",
					ShouldNotContainCycles: true,
				},
			},
		}

		expectedConfig := &configuration.Config{
			Version: 1,
			DependenciesRules: []*configuration.DependenciesRule{
				{
					Package: "foobar",
					ShouldOnlyDependsOn: &configuration.Dependencies{
						Internal: []string{"a", "b"},
					},
				},
			},
			FunctionsRules: []*configuration.FunctionsRule{
				{
					Package:                  "foobar",
					MaxReturnValues:          values.GetIntRef(2),
					MaxParameters:            values.GetIntRef(3),
					MaxLines:                 values.GetIntRef(35),
					MaxPublicFunctionPerFile: values.GetIntRef(5),
				},
			},
			ContentRules: []*configuration.ContentsRule{
				{
					Package:                     "foobar",
					ShouldOnlyContainInterfaces: true,
				},
			},
			NamingRules: []*configuration.NamingRule{
				{
					Package: "foobar",
					InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
						StructsThatImplement:           "blabla",
						ShouldHaveSimpleNameEndingWith: values.GetStringRef("jojo"),
					},
				},
			},
		}

		result := migrateRules(originalConfig)

		assert.Equal(t, expectedConfig, result)
	})

	t.Run("test migrate configuration without cycles rules", func(t *testing.T) {
		originalConfig := &configuration.DeprecatedConfig{
			DependenciesRules: []*configuration.DeprecatedDependenciesRule{
				{
					Package:             "foobar",
					ShouldOnlyDependsOn: []string{"a", "b"},
				},
			},
			FunctionsRules: []*configuration.FunctionsRule{
				{
					Package:                  "foobar",
					MaxReturnValues:          values.GetIntRef(2),
					MaxParameters:            values.GetIntRef(3),
					MaxLines:                 values.GetIntRef(35),
					MaxPublicFunctionPerFile: values.GetIntRef(5),
				},
			},
			ContentRules: []*configuration.ContentsRule{
				{
					Package:                     "foobar",
					ShouldOnlyContainInterfaces: true,
				},
			},
			NamingRules: []*configuration.NamingRule{
				{
					Package: "foobar",
					InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
						StructsThatImplement:           "blabla",
						ShouldHaveSimpleNameEndingWith: values.GetStringRef("jojo"),
					},
				},
			},
		}

		expectedConfig := &configuration.Config{
			Version: 1,
			DependenciesRules: []*configuration.DependenciesRule{
				{
					Package: "foobar",
					ShouldOnlyDependsOn: &configuration.Dependencies{
						Internal: []string{"a", "b"},
					},
				},
			},
			FunctionsRules: []*configuration.FunctionsRule{
				{
					Package:                  "foobar",
					MaxReturnValues:          values.GetIntRef(2),
					MaxParameters:            values.GetIntRef(3),
					MaxLines:                 values.GetIntRef(35),
					MaxPublicFunctionPerFile: values.GetIntRef(5),
				},
			},
			ContentRules: []*configuration.ContentsRule{
				{
					Package:                     "foobar",
					ShouldOnlyContainInterfaces: true,
				},
			},
			NamingRules: []*configuration.NamingRule{
				{
					Package: "foobar",
					InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
						StructsThatImplement:           "blabla",
						ShouldHaveSimpleNameEndingWith: values.GetStringRef("jojo"),
					},
				},
			},
		}

		result := migrateRules(originalConfig)

		assert.Equal(t, expectedConfig, result)
	})
}
