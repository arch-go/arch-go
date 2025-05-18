package naming

import (
	"testing"

	"github.com/arch-go/arch-go/internal/common"
	"github.com/arch-go/arch-go/internal/utils/packages"
	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/api/configuration"
	"github.com/arch-go/arch-go/internal/model"
	"github.com/arch-go/arch-go/internal/utils/values"
)

func TestNamingVerifications(t *testing.T) {
	t.Run("analyzeStructs compliance case", func(t *testing.T) {
		interfaceName := "interfaceX"
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
				StructsThatImplement: configuration.StructsThatImplement{
					Internal: &interfaceName,
				},
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
		interfaceName := "interfaceX"
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
				StructsThatImplement: configuration.StructsThatImplement{
					Internal: &interfaceName,
				},
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
		interfaceName := "interfaceX"
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
				StructsThatImplement: configuration.StructsThatImplement{
					Internal: &interfaceName,
				},
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
		interfaceName := "interfaceX"
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
				StructsThatImplement: configuration.StructsThatImplement{
					Internal: &interfaceName,
				},
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

	t.Run("Pass checking internal interface", func(t *testing.T) {
		mainPackage := "github.com/arch-go/arch-go"
		pkgs, err := packages.GetBasicPackagesInfo(mainPackage, nil, common.Verbose)
		assert.NoError(t, err)

		module := model.ModuleInfo{
			MainPackage: mainPackage,
			Packages:    pkgs,
		}

		internalInterface := "Command"
		patternName := "Command"
		rule := configuration.NamingRule{
			Package: "**.arch-go.**",
			InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
				StructsThatImplement: configuration.StructsThatImplement{
					Internal: &internalInterface,
				},
				ShouldHaveSimpleNameEndingWith: &patternName,
			},
		}

		result := CheckRule(module, rule)

		assert.True(t, result.Passes)
	})

	t.Run("Fail checking internal interface", func(t *testing.T) {
		mainPackage := "github.com/arch-go/arch-go"
		pkgs, err := packages.GetBasicPackagesInfo(mainPackage, nil, common.Verbose)
		assert.NoError(t, err)

		module := model.ModuleInfo{
			MainPackage: mainPackage,
			Packages:    pkgs,
		}

		internalInterface := "Command"
		patternName := "CommandFoo"
		rule := configuration.NamingRule{
			Package: "**.arch-go.**",
			InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
				StructsThatImplement: configuration.StructsThatImplement{
					Internal: &internalInterface,
				},
				ShouldHaveSimpleNameEndingWith: &patternName,
			},
		}

		result := CheckRule(module, rule)

		assert.False(t, result.Passes)
	})

	t.Run("Pass checking external interface", func(t *testing.T) {
		mainPackage := "github.com/arch-go/arch-go"
		pkgs, err := packages.GetBasicPackagesInfo(mainPackage, nil, common.Verbose)
		assert.NoError(t, err)

		module := model.ModuleInfo{
			MainPackage: mainPackage,
			Packages:    pkgs,
		}

		patternName := "ExternalInterface"
		rule := configuration.NamingRule{
			Package: "**.arch-go.**",
			InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
				StructsThatImplement: configuration.StructsThatImplement{
					External: &configuration.PackageAndInterface{
						Package:   "github.com/stretchr/testify/require",
						Interface: "TestingT",
					},
				},
				ShouldHaveSimpleNameEndingWith: &patternName,
			},
		}

		result := CheckRule(module, rule)

		assert.True(t, result.Passes)
	})

	t.Run("Fail checking external interface", func(t *testing.T) {
		mainPackage := "github.com/arch-go/arch-go"
		pkgs, err := packages.GetBasicPackagesInfo(mainPackage, nil, common.Verbose)
		assert.NoError(t, err)

		module := model.ModuleInfo{
			MainPackage: mainPackage,
			Packages:    pkgs,
		}

		patternName := "ExternalInterfaceFoo"
		rule := configuration.NamingRule{
			Package: "**.arch-go.**",
			InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
				StructsThatImplement: configuration.StructsThatImplement{
					External: &configuration.PackageAndInterface{
						Package:   "github.com/stretchr/testify/require",
						Interface: "TestingT",
					},
				},
				ShouldHaveSimpleNameEndingWith: &patternName,
			},
		}

		result := CheckRule(module, rule)

		assert.False(t, result.Passes)
	})

	t.Run("Pass checking interface from go builtin package", func(t *testing.T) {
		mainPackage := "github.com/arch-go/arch-go"
		pkgs, err := packages.GetBasicPackagesInfo(mainPackage, nil, common.Verbose)
		assert.NoError(t, err)

		module := model.ModuleInfo{
			MainPackage: mainPackage,
			Packages:    pkgs,
		}

		patternName := "Err"
		rule := configuration.NamingRule{
			Package: "**.arch-go.**",
			InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
				StructsThatImplement: configuration.StructsThatImplement{
					Standard: &configuration.PackageAndInterface{
						Package:   "builtin",
						Interface: "error",
					},
				},
				ShouldHaveSimpleNameEndingWith: &patternName,
			},
		}

		result := CheckRule(module, rule)

		assert.True(t, result.Passes)
	})

	t.Run("Fail checking interface from go builtin package", func(t *testing.T) {
		mainPackage := "github.com/arch-go/arch-go"
		pkgs, err := packages.GetBasicPackagesInfo(mainPackage, nil, common.Verbose)
		assert.NoError(t, err)

		module := model.ModuleInfo{
			MainPackage: mainPackage,
			Packages:    pkgs,
		}

		patternName := "ErrFoo"
		rule := configuration.NamingRule{
			Package: "**.arch-go.**",
			InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
				StructsThatImplement: configuration.StructsThatImplement{
					Standard: &configuration.PackageAndInterface{
						Package:   "builtin",
						Interface: "error",
					},
				},
				ShouldHaveSimpleNameEndingWith: &patternName,
			},
		}

		result := CheckRule(module, rule)

		assert.False(t, result.Passes)
	})

	t.Run("Pass checking go std interface", func(t *testing.T) {
		mainPackage := "github.com/arch-go/arch-go"
		pkgs, err := packages.GetBasicPackagesInfo(mainPackage, nil, common.Verbose)
		assert.NoError(t, err)

		module := model.ModuleInfo{
			MainPackage: mainPackage,
			Packages:    pkgs,
		}

		patternName := "Reader"
		rule := configuration.NamingRule{
			Package: "**.arch-go.**",
			InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
				StructsThatImplement: configuration.StructsThatImplement{
					Standard: &configuration.PackageAndInterface{
						Package:   "io",
						Interface: "Reader",
					},
				},
				ShouldHaveSimpleNameEndingWith: &patternName,
			},
		}

		result := CheckRule(module, rule)

		assert.True(t, result.Passes)
	})

	t.Run("Fail checking go std interface", func(t *testing.T) {
		mainPackage := "github.com/arch-go/arch-go"
		pkgs, err := packages.GetBasicPackagesInfo(mainPackage, nil, common.Verbose)
		assert.NoError(t, err)

		module := model.ModuleInfo{
			MainPackage: mainPackage,
			Packages:    pkgs,
		}

		patternName := "ReaderFoo"
		rule := configuration.NamingRule{
			Package: "**.arch-go.**",
			InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
				StructsThatImplement: configuration.StructsThatImplement{
					Standard: &configuration.PackageAndInterface{
						Package:   "io",
						Interface: "Reader",
					},
				},
				ShouldHaveSimpleNameEndingWith: &patternName,
			},
		}

		result := CheckRule(module, rule)

		assert.False(t, result.Passes)
	})
}
