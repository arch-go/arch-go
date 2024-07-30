package describe

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	monkey "github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/api/configuration"
	"github.com/arch-go/arch-go/internal/utils/values"
)

func TestDescribeCommand(t *testing.T) {
	t.Run("describe rules", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		configuration := &configuration.Config{
			Threshold: &configuration.Threshold{
				Compliance: values.GetIntRef(87),
				Coverage:   values.GetIntRef(34),
			},
			DependenciesRules: []*configuration.DependenciesRule{
				{
					Package: "foobar",
					ShouldOnlyDependsOn: &configuration.Dependencies{
						Internal: []string{"foo"},
						External: []string{"bar"},
						Standard: []string{"foobar"},
					},
				},
				{
					Package: "barfoo",
					ShouldNotDependsOn: &configuration.Dependencies{
						Internal: []string{"oof"},
						External: []string{"rab"},
						Standard: []string{"raboof"},
					},
				},
			},
			ContentRules: []*configuration.ContentsRule{
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
			FunctionsRules: []*configuration.FunctionsRule{
				{
					Package:                  "function-package",
					MaxParameters:            values.GetIntRef(1),
					MaxReturnValues:          values.GetIntRef(2),
					MaxLines:                 values.GetIntRef(3),
					MaxPublicFunctionPerFile: values.GetIntRef(4),
				},
			},
			NamingRules: []*configuration.NamingRule{
				{
					Package: "foobar",
					InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
						StructsThatImplement:             "foo",
						ShouldHaveSimpleNameStartingWith: values.GetStringRef("bla"),
					},
				},
				{
					Package: "barfoo",
					InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
						StructsThatImplement:           "foo",
						ShouldHaveSimpleNameEndingWith: values.GetStringRef("anything"),
					},
				},
			},
		}

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

		command := NewCommand(configuration, outputBuffer)
		returnValue := runDescribeCommand(command.(describeCommand)) //nolint: forcetypeassert
		outputBytes, _ := io.ReadAll(outputBuffer)

		assert.Equal(t, 0, returnValue, "Unexpected Return Value")
		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("describe threshold", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		cp := 87
		cv := 34
		threshold := &configuration.Threshold{
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
		fmt.Println("INvalid case")

		outputBuffer := bytes.NewBufferString("")
		patchExit := monkey.ApplyFunc(os.Exit, func(_ int) {})

		defer patchExit.Reset()

		expectedOutput := `Invalid Configuration: configuration file should have at least one rule
`

		fmt.Println("INvalid case2")
		NewCommand(&configuration.Config{}, outputBuffer).Run()
		fmt.Println("INvalid case3")

		outputBytes, _ := io.ReadAll(outputBuffer)
		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})
}
