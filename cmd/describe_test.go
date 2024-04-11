package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/fdaines/arch-go/internal/utils/values"

	"github.com/spf13/viper"

	monkey "github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"

	"github.com/fdaines/arch-go/pkg/config"
)

func TestDescribeCommand(t *testing.T) {
	viper.AddConfigPath("../test/")

	t.Run("when arch-go.yaml has no rules", func(t *testing.T) {
		var exitCode int
		cmd := NewDescribeCommand()
		patch := monkey.ApplyFuncReturn(config.LoadConfig, &config.Config{}, nil)
		defer patch.Reset()
		patchExit := monkey.ApplyFunc(os.Exit, func(c int) { exitCode = c })
		defer patchExit.Reset()

		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		cmd.Execute()
		out, _ := io.ReadAll(b)

		expected := `Invalid Configuration: configuration file should have at least one rule
`
		assert.Equal(t, expected, string(out))
		assert.Equal(t, 1, exitCode)
	})

	t.Run("when arch-go.yaml has rules", func(t *testing.T) {
		var exitCode int
		cmd := NewDescribeCommand()
		patch := monkey.ApplyFuncReturn(config.LoadConfig, configLoaderMockWithRules(), nil)
		defer patch.Reset()
		patchExit := monkey.ApplyFunc(os.Exit, func(c int) { exitCode = c })
		defer patchExit.Reset()

		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		cmd.Execute()
		out, _ := io.ReadAll(b)

		expected := `Dependency Rules
	* Packages that match pattern 'foobar',
		* Should only depends on:
			* Internal Packages that matches:
				- 'foo'
			* External Packages that matches:
				- 'bar'
			* StandardLib Packages that matches:
				- 'foobar'
	* Packages that match pattern 'barfoo',
		* Should not depends on:
			* Internal Packages that matches:
				- 'oof'
			* External Packages that matches:
				- 'rab'
			* StandardLib Packages that matches:
				- 'raboof'
Function Rules
	* Packages that match pattern 'function-package' should comply with the following rules:
		* Functions should not have more than 3 lines
		* Functions should not have more than 1 parameters
		* Functions should not have more than 2 return values
		* Files should not have more than 4 public functions
Content Rules
	* Packages that match pattern 'package1' should only contain interfaces
	* Packages that match pattern 'package2' should only contain structs
	* Packages that match pattern 'package3' should only contain functions
	* Packages that match pattern 'package4' should only contain methods
	* Packages that match pattern 'package5' should not contain interfaces
	* Packages that match pattern 'package6' should not contain structs
	* Packages that match pattern 'package7' should not contain functions
	* Packages that match pattern 'package8' should not contain methods
Naming Rules
	* Packages that match pattern 'foobar' should comply with:
		* Structs that implement interfaces matching name 'foo' should have simple name starting with 'bla'
	* Packages that match pattern 'barfoo' should comply with:
		* Structs that implement interfaces matching name 'foo' should have simple name ending with 'anything'
Threshold Rules
	* The module must comply with at least 98% of the rules described above.
	* The rules described above must cover at least 80% of the packages in this module.

`
		assert.Equal(t, expected, string(out))
		assert.Equal(t, 0, exitCode)
	})

	t.Run("when arch-go.yaml does not exist", func(t *testing.T) {
		var exitCode int
		configLoaderPatch := monkey.ApplyFuncReturn(config.LoadConfig, nil, fmt.Errorf("dummy error"))
		defer configLoaderPatch.Reset()
		patchExit := monkey.ApplyFunc(os.Exit, func(c int) { exitCode = c })
		defer patchExit.Reset()

		cmd := NewDescribeCommand()

		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		cmd.Execute()
		out, _ := io.ReadAll(b)

		expected := `Error: dummy error
`
		assert.Equal(t, expected, string(out))
		assert.Equal(t, 1, exitCode)
	})
}

func configLoaderMockWithRules() *config.Config {
	return &config.Config{
		Threshold: &config.Threshold{
			Compliance: values.GetIntRef(98),
			Coverage:   values.GetIntRef(80),
		},
		DependenciesRules: []*config.DependenciesRule{
			{
				Package: "foobar",
				ShouldOnlyDependsOn: &config.Dependencies{
					Internal: []string{"foo"},
					External: []string{"bar"},
					Standard: []string{"foobar"},
				},
			},
			{
				Package: "barfoo",
				ShouldNotDependsOn: &config.Dependencies{
					Internal: []string{"oof"},
					External: []string{"rab"},
					Standard: []string{"raboof"},
				},
			},
		},
		ContentRules: []*config.ContentsRule{
			{
				Package:                     "package1",
				ShouldOnlyContainInterfaces: true,
			},
			{
				Package:                  "package2",
				ShouldOnlyContainStructs: true,
			},
			{
				Package:                    "package3",
				ShouldOnlyContainFunctions: true,
			},
			{
				Package:                  "package4",
				ShouldOnlyContainMethods: true,
			},
			{
				Package:                    "package5",
				ShouldNotContainInterfaces: true,
			},
			{
				Package:                 "package6",
				ShouldNotContainStructs: true,
			},
			{
				Package:                   "package7",
				ShouldNotContainFunctions: true,
			},
			{
				Package:                 "package8",
				ShouldNotContainMethods: true,
			},
		},
		FunctionsRules: []*config.FunctionsRule{
			{
				Package:                  "function-package",
				MaxParameters:            values.GetIntRef(1),
				MaxReturnValues:          values.GetIntRef(2),
				MaxLines:                 values.GetIntRef(3),
				MaxPublicFunctionPerFile: values.GetIntRef(4),
			},
		},
		NamingRules: []*config.NamingRule{
			{
				Package: "foobar",
				InterfaceImplementationNamingRule: &config.InterfaceImplementationRule{
					StructsThatImplement:             "foo",
					ShouldHaveSimpleNameStartingWith: "bla",
				},
			},
			{
				Package: "barfoo",
				InterfaceImplementationNamingRule: &config.InterfaceImplementationRule{
					StructsThatImplement:           "foo",
					ShouldHaveSimpleNameEndingWith: "anything",
				},
			},
		},
	}
}
