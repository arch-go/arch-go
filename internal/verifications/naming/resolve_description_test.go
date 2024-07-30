package naming

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/api/configuration"
	"github.com/arch-go/arch-go/internal/utils/values"
)

func TestResolveNamingRuleDescription(t *testing.T) {
	t.Run("rule includes starting with", func(t *testing.T) {
		rule := configuration.NamingRule{
			Package: "foobar",
			InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
				StructsThatImplement:             "myInterface",
				ShouldHaveSimpleNameStartingWith: values.GetStringRef("blabla"),
			},
		}
		expectedResult := `Packages matching pattern 'foobar' should comply with [structs that implement 'myInterface' should have simple name starting with 'blabla']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})

	t.Run("rule includes ending with", func(t *testing.T) {
		rule := configuration.NamingRule{
			Package: "foobar",
			InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
				StructsThatImplement:           "myInterface",
				ShouldHaveSimpleNameEndingWith: values.GetStringRef("blabla"),
			},
		}
		expectedResult := `Packages matching pattern 'foobar' should comply with [structs that implement 'myInterface' should have simple name ending with 'blabla']`

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
