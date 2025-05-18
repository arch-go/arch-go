package describe

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/api/configuration"
	"github.com/arch-go/arch-go/internal/utils/values"
)

func TestDescribeNamingRules(t *testing.T) {
	t.Run("interface implementation naming rules", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		interfaceName1 := "QWERTY"
		interfaceName2 := "FOOBAR"
		rules := []*configuration.NamingRule{
			{
				Package: "foobar",
				InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
					StructsThatImplement: configuration.StructsThatImplement{
						Internal: &interfaceName1,
					},
					ShouldHaveSimpleNameStartingWith: values.GetStringRef("test"),
				},
			},
			{
				Package: "barfoo",
				InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
					StructsThatImplement: configuration.StructsThatImplement{
						Internal: &interfaceName2,
					},
					ShouldHaveSimpleNameEndingWith: values.GetStringRef("blablabla"),
				},
			},
		}
		expectedOutput := `Naming Rules
	* Packages that match pattern 'foobar' should comply with:
		* Structs that implement interfaces matching name 'QWERTY' should have simple name starting with 'test'
	* Packages that match pattern 'barfoo' should comply with:
		* Structs that implement interfaces matching name 'FOOBAR' should have simple name ending with 'blablabla'
`

		describeNamingRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)
		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("external interface implementation naming rules", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rules := []*configuration.NamingRule{
			{
				Package: "foobar",
				InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
					StructsThatImplement: configuration.StructsThatImplement{
						External: &configuration.PackageAndInterface{
							Package:   "github.com/some/package1",
							Interface: "InterfaceName1",
						},
					},
					ShouldHaveSimpleNameStartingWith: values.GetStringRef("test"),
				},
			},
			{
				Package: "barfoo",
				InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
					StructsThatImplement: configuration.StructsThatImplement{
						External: &configuration.PackageAndInterface{
							Package:   "github.com/some/package2",
							Interface: "InterfaceName2",
						},
					},
					ShouldHaveSimpleNameEndingWith: values.GetStringRef("blablabla"),
				},
			},
		}
		expectedOutput := `Naming Rules
	* Packages that match pattern 'foobar' should comply with:
		* Structs that implement interfaces matching name 'InterfaceName1' from external package 'github.com/some/package1' should have simple name starting with 'test'
	* Packages that match pattern 'barfoo' should comply with:
		* Structs that implement interfaces matching name 'InterfaceName2' from external package 'github.com/some/package2' should have simple name ending with 'blablabla'
`

		describeNamingRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)
		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("standard interface implementation naming rules", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rules := []*configuration.NamingRule{
			{
				Package: "foobar",
				InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
					StructsThatImplement: configuration.StructsThatImplement{
						Standard: &configuration.PackageAndInterface{
							Package:   "io",
							Interface: "Writer",
						},
					},
					ShouldHaveSimpleNameStartingWith: values.GetStringRef("test"),
				},
			},
			{
				Package: "barfoo",
				InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
					StructsThatImplement: configuration.StructsThatImplement{
						Standard: &configuration.PackageAndInterface{
							Package:   "builtin",
							Interface: "error",
						},
					},
					ShouldHaveSimpleNameEndingWith: values.GetStringRef("blablabla"),
				},
			},
		}
		expectedOutput := `Naming Rules
	* Packages that match pattern 'foobar' should comply with:
		* Structs that implement interfaces matching name 'Writer' from standard package 'io' should have simple name starting with 'test'
	* Packages that match pattern 'barfoo' should comply with:
		* Structs that implement interfaces matching name 'error' from standard package 'builtin' should have simple name ending with 'blablabla'
`

		describeNamingRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)
		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("empty rules", func(t *testing.T) {
		var rules []*configuration.NamingRule

		outputBuffer := bytes.NewBufferString("")
		expectedOutput := `Naming Rules
	* No rules defined
`

		describeNamingRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)
		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})
}
