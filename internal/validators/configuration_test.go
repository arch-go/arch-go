package validators

import (
	"testing"

	"github.com/fdaines/arch-go/api/configuration"

	"github.com/fdaines/arch-go/internal/utils/values"

	"github.com/stretchr/testify/assert"
)

func TestValidateConfiguration(t *testing.T) {
	t.Run("valid configuration", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*configuration.DependenciesRule{},
			ContentRules: []*configuration.ContentsRule{
				{
					Package:                  "foobar1",
					ShouldOnlyContainStructs: true,
				},
				{
					Package:                     "foobar2",
					ShouldOnlyContainInterfaces: true,
				},
				{
					Package:                    "foobar3",
					ShouldOnlyContainFunctions: true,
				},
				{
					Package:                  "foobar4",
					ShouldOnlyContainMethods: true,
				},
				{
					Package:                    "foobar5",
					ShouldNotContainStructs:    true,
					ShouldNotContainInterfaces: true,
					ShouldNotContainMethods:    true,
				},
			},
			FunctionsRules: []*configuration.FunctionsRule{
				{
					Package:                  "foobar0",
					MaxLines:                 values.GetIntRef(0),
					MaxParameters:            values.GetIntRef(0),
					MaxReturnValues:          values.GetIntRef(0),
					MaxPublicFunctionPerFile: values.GetIntRef(0),
				},
				{
					Package:  "foobar1",
					MaxLines: values.GetIntRef(1),
				},
				{
					Package:       "foobar2",
					MaxParameters: values.GetIntRef(1),
				},
				{
					Package:         "foobar3",
					MaxReturnValues: values.GetIntRef(1),
				},
				{
					Package:                  "foobar4",
					MaxPublicFunctionPerFile: values.GetIntRef(1),
				},
			},
			NamingRules: []*configuration.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Nil(t, result, "Valid configuration should not generate an error")
	})

	t.Run("valid configuration - only dependencies rules", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:   1,
			Threshold: nil,
			DependenciesRules: []*configuration.DependenciesRule{
				{
					Package: "foobar",
					ShouldOnlyDependsOn: &configuration.Dependencies{
						Internal: []string{"time"},
					},
				},
			},
		}

		result := ValidateConfiguration(configuration)
		assert.Nil(t, result, "Valid configuration should not generate an error")
	})

	t.Run("valid configuration - only functions rules", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:   1,
			Threshold: nil,
			FunctionsRules: []*configuration.FunctionsRule{
				{
					Package:                  "foobar4",
					MaxPublicFunctionPerFile: values.GetIntRef(1),
				},
			},
		}

		result := ValidateConfiguration(configuration)
		assert.Nil(t, result, "Valid configuration should not generate an error")
	})

	t.Run("valid configuration - only contents rules", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:   1,
			Threshold: nil,
			ContentRules: []*configuration.ContentsRule{
				{
					Package:                  "foobar1",
					ShouldOnlyContainStructs: true,
				},
			},
		}

		result := ValidateConfiguration(configuration)
		assert.Nil(t, result, "Valid configuration should not generate an error")
	})

	t.Run("valid configuration - only naming rules", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:   1,
			Threshold: nil,
			NamingRules: []*configuration.NamingRule{
				{
					Package: "foobar",
					InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
						StructsThatImplement:           "bla",
						ShouldHaveSimpleNameEndingWith: values.GetStringRef("foo"),
					},
				},
			},
		}

		result := ValidateConfiguration(configuration)
		assert.Nil(t, result, "Valid configuration should not generate an error")
	})

	t.Run("invalid configuration - nil object", func(t *testing.T) {
		result := ValidateConfiguration(nil)
		assert.Equal(t, "configuration file not found", result.Error(), "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - no rules", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*configuration.DependenciesRule{},
			ContentRules:      []*configuration.ContentsRule{},
			FunctionsRules:    []*configuration.FunctionsRule{},
			NamingRules:       []*configuration.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "configuration file should have at least one rule", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - function rules case 1", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*configuration.DependenciesRule{},
			ContentRules:      []*configuration.ContentsRule{},
			FunctionsRules: []*configuration.FunctionsRule{
				{
					Package: "",
				},
			},
			NamingRules: []*configuration.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "function rule - empty package", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - function rules case 2", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*configuration.DependenciesRule{},
			ContentRules:      []*configuration.ContentsRule{},
			FunctionsRules: []*configuration.FunctionsRule{
				{
					Package:  "foobar",
					MaxLines: values.GetIntRef(-1),
				},
			},
			NamingRules: []*configuration.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "function rule - MaxLines is less than zero", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - function rules case 3", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*configuration.DependenciesRule{},
			ContentRules:      []*configuration.ContentsRule{},
			FunctionsRules: []*configuration.FunctionsRule{
				{
					Package:       "foobar",
					MaxParameters: values.GetIntRef(-1),
				},
			},
			NamingRules: []*configuration.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "function rule - MaxParameters is less than zero", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - function rules case 4", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*configuration.DependenciesRule{},
			ContentRules:      []*configuration.ContentsRule{},
			FunctionsRules: []*configuration.FunctionsRule{
				{
					Package:         "foobar",
					MaxReturnValues: values.GetIntRef(-1),
				},
			},
			NamingRules: []*configuration.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "function rule - MaxReturnValues is less than zero", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - function rules case 5", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*configuration.DependenciesRule{},
			ContentRules:      []*configuration.ContentsRule{},
			FunctionsRules: []*configuration.FunctionsRule{
				{
					Package:                  "foobar",
					MaxPublicFunctionPerFile: values.GetIntRef(-1),
				},
			},
			NamingRules: []*configuration.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "function rule - MaxPublicFunctionPerFile is less than zero", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - function rules case 6", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*configuration.DependenciesRule{},
			ContentRules:      []*configuration.ContentsRule{},
			FunctionsRules: []*configuration.FunctionsRule{
				{
					Package: "foobar",
				},
			},
			NamingRules: []*configuration.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "function rule - At least one criteria should be set", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - content rules case 1", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*configuration.DependenciesRule{},
			ContentRules: []*configuration.ContentsRule{
				{},
			},
			FunctionsRules: []*configuration.FunctionsRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "content rule - empty package", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - content rules case 2", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*configuration.DependenciesRule{},
			ContentRules: []*configuration.ContentsRule{
				{
					Package:                     "foobar",
					ShouldOnlyContainInterfaces: true,
					ShouldOnlyContainMethods:    true,
				},
			},
			FunctionsRules: []*configuration.FunctionsRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "content rule - if ShouldOnlyContainMethods is set, then it should be the only parameter", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - content rules case 3", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*configuration.DependenciesRule{},
			ContentRules: []*configuration.ContentsRule{
				{
					Package:                     "foobar",
					ShouldOnlyContainInterfaces: true,
					ShouldOnlyContainStructs:    true,
				},
			},
			FunctionsRules: []*configuration.FunctionsRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "content rule - if ShouldOnlyContainStructs is set, then it should be the only parameter", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - content rules case 4", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*configuration.DependenciesRule{},
			ContentRules: []*configuration.ContentsRule{
				{
					Package:                     "foobar",
					ShouldOnlyContainInterfaces: true,
					ShouldOnlyContainFunctions:  true,
				},
			},
			FunctionsRules: []*configuration.FunctionsRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "content rule - if ShouldOnlyContainFunctions is set, then it should be the only parameter", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - content rules case 5", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*configuration.DependenciesRule{},
			ContentRules: []*configuration.ContentsRule{
				{
					Package:                     "foobar",
					ShouldOnlyContainInterfaces: true,
					ShouldNotContainInterfaces:  true,
				},
			},
			FunctionsRules: []*configuration.FunctionsRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "content rule - if ShouldOnlyContainInterfaces is set, then it should be the only parameter", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - content rules case 6", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*configuration.DependenciesRule{},
			ContentRules: []*configuration.ContentsRule{
				{
					Package: "foobar",
				},
			},
			FunctionsRules: []*configuration.FunctionsRule{},
			NamingRules:    []*configuration.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "content rule - At least one criteria should be set", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - dependencies rules case 1", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:   1,
			Threshold: nil,
			DependenciesRules: []*configuration.DependenciesRule{
				{
					Package: "",
				},
			},
			ContentRules:   []*configuration.ContentsRule{},
			FunctionsRules: []*configuration.FunctionsRule{},
			NamingRules:    []*configuration.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "dependencies rule - empty package", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - dependencies rules case 2", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:   1,
			Threshold: nil,
			DependenciesRules: []*configuration.DependenciesRule{
				{
					Package: "foobar",
				},
			},
			ContentRules:   []*configuration.ContentsRule{},
			FunctionsRules: []*configuration.FunctionsRule{},
			NamingRules:    []*configuration.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "dependencies rule - Should contain one of 'ShouldOnlyDependsOn' or 'ShouldNotDependsOn'", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - dependencies rules case 3", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:   1,
			Threshold: nil,
			DependenciesRules: []*configuration.DependenciesRule{
				{
					Package:             "foobar",
					ShouldNotDependsOn:  &configuration.Dependencies{},
					ShouldOnlyDependsOn: &configuration.Dependencies{},
				},
			},
			ContentRules:   []*configuration.ContentsRule{},
			FunctionsRules: []*configuration.FunctionsRule{},
			NamingRules:    []*configuration.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "dependencies rule - Should contain only one of 'ShouldOnlyDependsOn' or 'ShouldNotDependsOn'", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - dependencies rules case 4", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:   1,
			Threshold: nil,
			DependenciesRules: []*configuration.DependenciesRule{
				{
					Package:            "foobar",
					ShouldNotDependsOn: &configuration.Dependencies{},
				},
			},
			ContentRules:   []*configuration.ContentsRule{},
			FunctionsRules: []*configuration.FunctionsRule{},
			NamingRules:    []*configuration.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "dependencies rule - ShouldNotDependsOn needs at least one of 'External', 'Internal' or 'Standard'", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - dependencies rules case 5", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:   1,
			Threshold: nil,
			DependenciesRules: []*configuration.DependenciesRule{
				{
					Package:             "foobar",
					ShouldOnlyDependsOn: &configuration.Dependencies{},
				},
			},
			ContentRules:   []*configuration.ContentsRule{},
			FunctionsRules: []*configuration.FunctionsRule{},
			NamingRules:    []*configuration.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "dependencies rule - ShouldOnlyDependsOn needs at least one of 'External', 'Internal' or 'Standard'", "Invalid configuration should return an error")
	})

	t.Run("test count rules", func(t *testing.T) {
		configuration := &configuration.Config{
			Version:   1,
			Threshold: nil,
			DependenciesRules: []*configuration.DependenciesRule{
				{
					Package: "foobar",
					ShouldOnlyDependsOn: &configuration.Dependencies{
						Internal: []string{"time"},
					},
				},
			},
			ContentRules: []*configuration.ContentsRule{
				{
					Package:                  "foobar1",
					ShouldOnlyContainStructs: true,
				},
			},
			FunctionsRules: []*configuration.FunctionsRule{
				{
					Package:                  "foobar4",
					MaxPublicFunctionPerFile: values.GetIntRef(1),
				},
			},
			NamingRules: []*configuration.NamingRule{
				{
					Package: "foobar",
					InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
						StructsThatImplement:           "bla",
						ShouldHaveSimpleNameEndingWith: values.GetStringRef("foo"),
					},
				},
			},
		}

		result := countRules(configuration)
		assert.Equal(t, 4, result, "Expects 4 rules.")
	})
}
