package naming

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fdaines/arch-go/api/configuration"
	"github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/internal/utils/values"
)

func TestNamingVerifications(t *testing.T) {
	t.Run("analyzeStructs compliance case", func(t *testing.T) {
		inputInterfaces := []InterfaceDescription{
			{
				Name: "interfaceX",
				Methods: []MethodDescription{
					{
						Name:         "method01",
						Parameters:   []string{},
						ReturnValues: []string{},
					},
				},
			},
		}
		inputPackage := &model.PackageInfo{Path: "foo/bar"}
		inputDetails := []string{}
		inputRule := configuration.NamingRule{
			InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
				StructsThatImplement:           "interfaceX",
				ShouldHaveSimpleNameEndingWith: values.GetStringRef("Foobar"),
			},
		}
		inputStructs := []StructDescription{
			{
				Name: "structFoobar",
				Methods: []MethodDescription{
					{
						Name:         "method01",
						Parameters:   []string{},
						ReturnValues: []string{},
					},
					{
						Name:         "method02",
						Parameters:   []string{},
						ReturnValues: []string{},
					},
				},
			},
		}

		passes, details := analyzeStructs(inputInterfaces, inputPackage, inputDetails, inputRule, inputStructs)

		assert.True(t, passes)
		assert.Equal(t, []string{}, details)
	})

	t.Run("analyzeStructs compliance case 2 when there is no struct that implements interface", func(t *testing.T) {
		inputInterfaces := []InterfaceDescription{
			{
				Name: "interfaceX",
				Methods: []MethodDescription{
					{
						Name:         "blablabla",
						Parameters:   []string{},
						ReturnValues: []string{},
					},
				},
			},
		}
		inputPackage := &model.PackageInfo{Path: "foo/bar"}
		inputDetails := []string{}
		inputRule := configuration.NamingRule{
			InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
				StructsThatImplement:           "interfaceX",
				ShouldHaveSimpleNameEndingWith: values.GetStringRef("Foobar"),
			},
		}
		inputStructs := []StructDescription{
			{
				Name: "structFoobar",
				Methods: []MethodDescription{
					{
						Name:         "method01",
						Parameters:   []string{},
						ReturnValues: []string{},
					},
					{
						Name:         "method02",
						Parameters:   []string{},
						ReturnValues: []string{},
					},
				},
			},
		}

		passes, details := analyzeStructs(inputInterfaces, inputPackage, inputDetails, inputRule, inputStructs)

		assert.True(t, passes)
		assert.Equal(t, []string{}, details)
	})

	t.Run("analyzeStructs non compliance case 1", func(t *testing.T) {
		inputInterfaces := []InterfaceDescription{
			{
				Name: "interfaceX",
				Methods: []MethodDescription{
					{
						Name:         "method01",
						Parameters:   []string{},
						ReturnValues: []string{},
					},
				},
			},
		}
		inputPackage := &model.PackageInfo{Path: "foo/bar"}
		inputDetails := []string{}
		inputRule := configuration.NamingRule{
			InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
				StructsThatImplement:           "interfaceX",
				ShouldHaveSimpleNameEndingWith: values.GetStringRef("Foobar"),
			},
		}
		inputStructs := []StructDescription{
			{
				Name: "struct1",
				Methods: []MethodDescription{
					{
						Name:         "method01",
						Parameters:   []string{},
						ReturnValues: []string{},
					},
					{
						Name:         "method02",
						Parameters:   []string{},
						ReturnValues: []string{},
					},
				},
			},
		}

		passes, details := analyzeStructs(inputInterfaces, inputPackage, inputDetails, inputRule, inputStructs)

		assert.False(t, passes)
		assert.Equal(t, []string{"Struct [struct1] in Package [foo/bar] does not match Naming Rule"}, details)
	})

	t.Run("analyzeStructs non compliance case 2", func(t *testing.T) {
		inputInterfaces := []InterfaceDescription{
			{
				Name: "interfaceX",
				Methods: []MethodDescription{
					{
						Name:         "method01",
						Parameters:   []string{},
						ReturnValues: []string{},
					},
				},
			},
		}
		inputPackage := &model.PackageInfo{Path: "foo/bar"}
		inputDetails := []string{}
		inputRule := configuration.NamingRule{
			InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
				StructsThatImplement:             "interfaceX",
				ShouldHaveSimpleNameStartingWith: values.GetStringRef("Foobar"),
			},
		}
		inputStructs := []StructDescription{
			{
				Name: "struct1",
				Methods: []MethodDescription{
					{
						Name:         "method01",
						Parameters:   []string{},
						ReturnValues: []string{},
					},
					{
						Name:         "method02",
						Parameters:   []string{},
						ReturnValues: []string{},
					},
				},
			},
		}

		passes, details := analyzeStructs(inputInterfaces, inputPackage, inputDetails, inputRule, inputStructs)

		assert.False(t, passes)
		assert.Equal(t, []string{"Struct [struct1] in Package [foo/bar] does not match Naming Rule"}, details)
	})
}
