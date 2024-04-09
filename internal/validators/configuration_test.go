package validators

import (
	"github.com/fdaines/arch-go/old/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateConfiguration(t *testing.T) {

	t.Run("valid configuration", func(t *testing.T) {
		configuration := &config.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*config.DependenciesRule{},
			ContentRules: []*config.ContentsRule{
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
			FunctionsRules: []*config.FunctionsRule{
				{
					Package:  "foobar1",
					MaxLines: 1,
				},
				{
					Package:       "foobar2",
					MaxParameters: 1,
				},
				{
					Package:         "foobar3",
					MaxReturnValues: 1,
				},
				{
					Package:                  "foobar4",
					MaxPublicFunctionPerFile: 1,
				},
			},
			NamingRules: []*config.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Nil(t, result, "Valid configuration should not generate an error")
	})

	t.Run("invalid configuration - no rules", func(t *testing.T) {
		configuration := &config.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*config.DependenciesRule{},
			ContentRules:      []*config.ContentsRule{},
			FunctionsRules:    []*config.FunctionsRule{},
			NamingRules:       []*config.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "configuration file should have at least one rule", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - function rules case 1", func(t *testing.T) {
		configuration := &config.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*config.DependenciesRule{},
			ContentRules:      []*config.ContentsRule{},
			FunctionsRules: []*config.FunctionsRule{
				{
					Package: "",
				},
			},
			NamingRules: []*config.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "function rule - empty package", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - function rules case 2", func(t *testing.T) {
		configuration := &config.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*config.DependenciesRule{},
			ContentRules:      []*config.ContentsRule{},
			FunctionsRules: []*config.FunctionsRule{
				{
					Package:  "foobar",
					MaxLines: -1,
				},
			},
			NamingRules: []*config.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "function rule - MaxLines is less than zero", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - function rules case 3", func(t *testing.T) {
		configuration := &config.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*config.DependenciesRule{},
			ContentRules:      []*config.ContentsRule{},
			FunctionsRules: []*config.FunctionsRule{
				{
					Package:       "foobar",
					MaxParameters: -1,
				},
			},
			NamingRules: []*config.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "function rule - MaxParameters is less than zero", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - function rules case 4", func(t *testing.T) {
		configuration := &config.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*config.DependenciesRule{},
			ContentRules:      []*config.ContentsRule{},
			FunctionsRules: []*config.FunctionsRule{
				{
					Package:         "foobar",
					MaxReturnValues: -1,
				},
			},
			NamingRules: []*config.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "function rule - MaxReturnValues is less than zero", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - function rules case 5", func(t *testing.T) {
		configuration := &config.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*config.DependenciesRule{},
			ContentRules:      []*config.ContentsRule{},
			FunctionsRules: []*config.FunctionsRule{
				{
					Package:                  "foobar",
					MaxPublicFunctionPerFile: -1,
				},
			},
			NamingRules: []*config.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "function rule - MaxPublicFunctionPerFile is less than zero", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - function rules case 6", func(t *testing.T) {
		configuration := &config.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*config.DependenciesRule{},
			ContentRules:      []*config.ContentsRule{},
			FunctionsRules: []*config.FunctionsRule{
				{
					Package: "foobar",
				},
			},
			NamingRules: []*config.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "function rule - At least one criteria should be set", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - content rules case 1", func(t *testing.T) {
		configuration := &config.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*config.DependenciesRule{},
			ContentRules: []*config.ContentsRule{
				{},
			},
			FunctionsRules: []*config.FunctionsRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "content rule - empty package", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - content rules case 2", func(t *testing.T) {
		configuration := &config.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*config.DependenciesRule{},
			ContentRules: []*config.ContentsRule{
				{
					Package:                     "foobar",
					ShouldOnlyContainInterfaces: true,
					ShouldOnlyContainMethods:    true,
				},
			},
			FunctionsRules: []*config.FunctionsRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "content rule - if ShouldOnlyContainMethods is set, then it should be the only parameter", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - content rules case 3", func(t *testing.T) {
		configuration := &config.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*config.DependenciesRule{},
			ContentRules: []*config.ContentsRule{
				{
					Package:                     "foobar",
					ShouldOnlyContainInterfaces: true,
					ShouldOnlyContainStructs:    true,
				},
			},
			FunctionsRules: []*config.FunctionsRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "content rule - if ShouldOnlyContainStructs is set, then it should be the only parameter", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - content rules case 4", func(t *testing.T) {
		configuration := &config.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*config.DependenciesRule{},
			ContentRules: []*config.ContentsRule{
				{
					Package:                     "foobar",
					ShouldOnlyContainInterfaces: true,
					ShouldOnlyContainFunctions:  true,
				},
			},
			FunctionsRules: []*config.FunctionsRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "content rule - if ShouldOnlyContainFunctions is set, then it should be the only parameter", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - content rules case 5", func(t *testing.T) {
		configuration := &config.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*config.DependenciesRule{},
			ContentRules: []*config.ContentsRule{
				{
					Package:                     "foobar",
					ShouldOnlyContainInterfaces: true,
					ShouldNotContainInterfaces:  true,
				},
			},
			FunctionsRules: []*config.FunctionsRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "content rule - if ShouldOnlyContainInterfaces is set, then it should be the only parameter", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - content rules case 6", func(t *testing.T) {
		configuration := &config.Config{
			Version:           1,
			Threshold:         nil,
			DependenciesRules: []*config.DependenciesRule{},
			ContentRules: []*config.ContentsRule{
				{
					Package: "foobar",
				},
			},
			FunctionsRules: []*config.FunctionsRule{},
			NamingRules:    []*config.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "content rule - At least one criteria should be set", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - dependencies rules case 1", func(t *testing.T) {
		configuration := &config.Config{
			Version:   1,
			Threshold: nil,
			DependenciesRules: []*config.DependenciesRule{
				{
					Package: "",
				},
			},
			ContentRules:   []*config.ContentsRule{},
			FunctionsRules: []*config.FunctionsRule{},
			NamingRules:    []*config.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "dependencies rule - empty package", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - dependencies rules case 2", func(t *testing.T) {
		configuration := &config.Config{
			Version:   1,
			Threshold: nil,
			DependenciesRules: []*config.DependenciesRule{
				{
					Package: "foobar",
				},
			},
			ContentRules:   []*config.ContentsRule{},
			FunctionsRules: []*config.FunctionsRule{},
			NamingRules:    []*config.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "dependencies rule - Should contain one of 'ShouldOnlyDependsOn' or 'ShouldNotDependsOn'", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - dependencies rules case 3", func(t *testing.T) {
		configuration := &config.Config{
			Version:   1,
			Threshold: nil,
			DependenciesRules: []*config.DependenciesRule{
				{
					Package:             "foobar",
					ShouldNotDependsOn:  &config.Dependencies{},
					ShouldOnlyDependsOn: &config.Dependencies{},
				},
			},
			ContentRules:   []*config.ContentsRule{},
			FunctionsRules: []*config.FunctionsRule{},
			NamingRules:    []*config.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "dependencies rule - Should contain only one of 'ShouldOnlyDependsOn' or 'ShouldNotDependsOn'", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - dependencies rules case 4", func(t *testing.T) {
		configuration := &config.Config{
			Version:   1,
			Threshold: nil,
			DependenciesRules: []*config.DependenciesRule{
				{
					Package:            "foobar",
					ShouldNotDependsOn: &config.Dependencies{},
				},
			},
			ContentRules:   []*config.ContentsRule{},
			FunctionsRules: []*config.FunctionsRule{},
			NamingRules:    []*config.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "dependencies rule - ShouldNotDependsOn needs at least one of 'External', 'Internal' or 'Standard'", "Invalid configuration should return an error")
	})

	t.Run("invalid configuration - dependencies rules case 5", func(t *testing.T) {
		configuration := &config.Config{
			Version:   1,
			Threshold: nil,
			DependenciesRules: []*config.DependenciesRule{
				{
					Package:             "foobar",
					ShouldOnlyDependsOn: &config.Dependencies{},
				},
			},
			ContentRules:   []*config.ContentsRule{},
			FunctionsRules: []*config.FunctionsRule{},
			NamingRules:    []*config.NamingRule{},
		}

		result := ValidateConfiguration(configuration)
		assert.Equal(t, result.Error(), "dependencies rule - ShouldOnlyDependsOn needs at least one of 'External', 'Internal' or 'Standard'", "Invalid configuration should return an error")
	})

}
