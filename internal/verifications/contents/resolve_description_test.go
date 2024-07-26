package contents

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fdaines/arch-go/api/configuration"
)

func TestResolveContentsRuleDescription(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		rule := configuration.ContentsRule{
			Package:                    "foobar",
			ShouldOnlyContainFunctions: true,
		}
		expectedResult := `Packages matching pattern 'foobar' should complies with ['should only contain functions']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})

	t.Run("case 2", func(t *testing.T) {
		rule := configuration.ContentsRule{
			Package:                  "foobar",
			ShouldOnlyContainMethods: true,
		}
		expectedResult := `Packages matching pattern 'foobar' should complies with ['should only contain methods']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})

	t.Run("case 3", func(t *testing.T) {
		rule := configuration.ContentsRule{
			Package:                     "foobar",
			ShouldOnlyContainInterfaces: true,
		}
		expectedResult := `Packages matching pattern 'foobar' should complies with ['should only contain interfaces']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})

	t.Run("case 4", func(t *testing.T) {
		rule := configuration.ContentsRule{
			Package:                  "foobar",
			ShouldOnlyContainStructs: true,
		}
		expectedResult := `Packages matching pattern 'foobar' should complies with ['should only contain structs']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})

	t.Run("case 5", func(t *testing.T) {
		rule := configuration.ContentsRule{
			Package:                   "foobar",
			ShouldNotContainMethods:   true,
			ShouldNotContainStructs:   true,
			ShouldNotContainFunctions: true,
		}
		expectedResult := `Packages matching pattern 'foobar' should complies with ['should not contain structs','should not contain functions','should not contain methods']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})

	t.Run("case 6", func(t *testing.T) {
		rule := configuration.ContentsRule{
			Package:                    "foobar",
			ShouldNotContainMethods:    true,
			ShouldNotContainInterfaces: true,
			ShouldNotContainFunctions:  true,
		}
		expectedResult := `Packages matching pattern 'foobar' should complies with ['should not contain interfaces','should not contain functions','should not contain methods']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})
}
