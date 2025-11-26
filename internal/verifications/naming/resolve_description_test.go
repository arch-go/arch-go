package naming

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/v2/api/configuration"
	"github.com/arch-go/arch-go/v2/internal/utils/values"
)

func TestResolveNamingRuleDescription(t *testing.T) {
	t.Run("rule includes starting with", func(t *testing.T) {
		interfaceName := "myInterface"
		rule := configuration.NamingRule{
			Package: "foobar",
			InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
				StructsThatImplement: configuration.StructsThatImplement{
					Internal: &interfaceName,
				},
				ShouldHaveSimpleNameStartingWith: values.GetStringRef("blabla"),
			},
		}
		expectedResult := `Packages matching pattern 'foobar' should comply with [structs that implement 'myInterface' should have simple name starting with 'blabla']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})

	t.Run("rule includes starting with - external interface", func(t *testing.T) {
		rule := configuration.NamingRule{
			Package: "foobar",
			InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
				StructsThatImplement: configuration.StructsThatImplement{
					External: &configuration.PackageAndInterface{
						Package:   "github.com/some/package",
						Interface: "InterfaceName",
					},
				},
				ShouldHaveSimpleNameStartingWith: values.GetStringRef("blabla"),
			},
		}
		expectedResult := `Packages matching pattern 'foobar' should comply with [structs that implement 'InterfaceName' from external package 'github.com/some/package' should have simple name starting with 'blabla']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})

	t.Run("rule includes starting with - standard interface", func(t *testing.T) {
		rule := configuration.NamingRule{
			Package: "foobar",
			InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
				StructsThatImplement: configuration.StructsThatImplement{
					Standard: &configuration.PackageAndInterface{
						Package:   "io",
						Interface: "Writer",
					},
				},
				ShouldHaveSimpleNameStartingWith: values.GetStringRef("blabla"),
			},
		}
		expectedResult := `Packages matching pattern 'foobar' should comply with [structs that implement 'Writer' from standard package 'io' should have simple name starting with 'blabla']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})

	t.Run("rule includes ending with", func(t *testing.T) {
		interfaceName := "myInterface"
		rule := configuration.NamingRule{
			Package: "foobar",
			InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
				StructsThatImplement: configuration.StructsThatImplement{
					Internal: &interfaceName,
				},
				ShouldHaveSimpleNameEndingWith: values.GetStringRef("blabla"),
			},
		}
		expectedResult := `Packages matching pattern 'foobar' should comply with [structs that implement 'myInterface' should have simple name ending with 'blabla']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})

	t.Run("rule includes ending with - external interface", func(t *testing.T) {
		rule := configuration.NamingRule{
			Package: "foobar",
			InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
				StructsThatImplement: configuration.StructsThatImplement{
					External: &configuration.PackageAndInterface{
						Package:   "github.com/some/package",
						Interface: "InterfaceName",
					},
				},
				ShouldHaveSimpleNameEndingWith: values.GetStringRef("blabla"),
			},
		}
		expectedResult := `Packages matching pattern 'foobar' should comply with [structs that implement 'InterfaceName' from external package 'github.com/some/package' should have simple name ending with 'blabla']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})

	t.Run("rule includes ending with - standard interface", func(t *testing.T) {
		rule := configuration.NamingRule{
			Package: "foobar",
			InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
				StructsThatImplement: configuration.StructsThatImplement{
					Standard: &configuration.PackageAndInterface{
						Package:   "io",
						Interface: "Writer",
					},
				},
				ShouldHaveSimpleNameEndingWith: values.GetStringRef("blabla"),
			},
		}
		expectedResult := `Packages matching pattern 'foobar' should comply with [structs that implement 'Writer' from standard package 'io' should have simple name ending with 'blabla']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})

	t.Run("rule does not include interface implementation rule", func(t *testing.T) {
		rule := configuration.NamingRule{
			Package: "foobar",
		}
		expectedResult := `Packages matching pattern 'foobar' should comply with []`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})
}
