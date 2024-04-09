package describe

import (
	"bytes"
	monkey "github.com/agiledragon/gomonkey/v2"
	"github.com/fdaines/arch-go/old/config"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestDescribeCommand(t *testing.T) {

	t.Run("describe threshold", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		cp := 87
		cv := 34
		configuration := &config.Config{
			Threshold: &config.Threshold{
				Compliance: &cp,
				Coverage:   &cv,
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
					MaxParameters:            1,
					MaxReturnValues:          2,
					MaxLines:                 3,
					MaxPublicFunctionPerFile: 4,
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
		patch := monkey.ApplyFuncReturn(config.LoadConfig, configuration, nil)
		defer patch.Reset()

		expectedOutput := `Dependency Rules
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
	* The module must comply with at least 87% of the rules described above.
	* The rules described above must cover at least 34% of the packages in this module.

`

		command := NewCommand(outputBuffer)
		returnValue := runDescribeCommand(command)

		outputBytes, _ := io.ReadAll(outputBuffer)

		assert.Equal(t, 0, returnValue, "Unexpected Return Value")
		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("describe threshold", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		cp := 87
		cv := 34
		threshold := &config.Threshold{
			Compliance: &cp,
			Coverage:   &cv,
		}
		expectedOutput := `Threshold Rules
	* The module must comply with at least 87% of the rules described above.
	* The rules described above must cover at least 34% of the packages in this module.

`

		describeThresholdRules(threshold, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)

		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("empty thresold", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		expectedOutput := ``

		describeThresholdRules(nil, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)

		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("invalid configuration", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		patch := monkey.ApplyFuncReturn(config.LoadConfig, &config.Config{}, nil)
		defer patch.Reset()
		patchExit := monkey.ApplyFunc(os.Exit, func(c int) {})
		defer patchExit.Reset()

		expectedOutput := `Invalid Configuration: configuration file should have at least one rule
`

		NewCommand(outputBuffer).Run()

		outputBytes, _ := io.ReadAll(outputBuffer)

		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})
}
