package dependencies

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/v2/api/configuration"
)

func TestResolveNamingRuleDescription(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		rule := configuration.DependenciesRule{
			Package: "foobar",
			ShouldOnlyDependsOn: &configuration.Dependencies{
				Internal: []string{"foo", "bar"},
				External: []string{"ex1", "ex2"},
				Standard: []string{"st1", "st45"},
			},
		}
		expectedResult := `Packages matching pattern 'foobar' should ['only depend on internal packages that matches [[foo bar]]','only depend on external packages that matches [[ex1 ex2]]','only depend on standard packages that matches [[st1 st45]]']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})

	t.Run("case 2", func(t *testing.T) {
		rule := configuration.DependenciesRule{
			Package: "foobar",
			ShouldNotDependsOn: &configuration.Dependencies{
				Internal: []string{"foo", "bar"},
				External: []string{"ex1", "ex2"},
				Standard: []string{"st1", "st45"},
			},
		}
		expectedResult := `Packages matching pattern 'foobar' should ['not depend on internal packages that matches [[foo bar]]','not depend on external packages that matches [[ex1 ex2]]','not depend on standard packages that matches [[st1 st45]]']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})

	t.Run("case 3", func(t *testing.T) {
		rule := configuration.DependenciesRule{
			Package: "foobar",
			ShouldOnlyDependsOn: &configuration.Dependencies{
				Internal: []string{"foo", "bar"},
			},
		}
		expectedResult := `Packages matching pattern 'foobar' should ['only depend on internal packages that matches [[foo bar]]']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})

	t.Run("case 4", func(t *testing.T) {
		rule := configuration.DependenciesRule{
			Package: "foobar",
			ShouldOnlyDependsOn: &configuration.Dependencies{
				External: []string{"x1", "x2"},
			},
		}
		expectedResult := `Packages matching pattern 'foobar' should ['only depend on external packages that matches [[x1 x2]]']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})

	t.Run("case 5", func(t *testing.T) {
		rule := configuration.DependenciesRule{
			Package: "foobar",
			ShouldOnlyDependsOn: &configuration.Dependencies{
				Standard: []string{"s100", "s200"},
			},
		}
		expectedResult := `Packages matching pattern 'foobar' should ['only depend on standard packages that matches [[s100 s200]]']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})

	t.Run("case 6", func(t *testing.T) {
		rule := configuration.DependenciesRule{
			Package: "foobar",
			ShouldNotDependsOn: &configuration.Dependencies{
				Internal: []string{"foo", "bar"},
			},
		}
		expectedResult := `Packages matching pattern 'foobar' should ['not depend on internal packages that matches [[foo bar]]']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})

	t.Run("case 7", func(t *testing.T) {
		rule := configuration.DependenciesRule{
			Package: "foobar",
			ShouldNotDependsOn: &configuration.Dependencies{
				External: []string{"x1", "x2"},
			},
		}
		expectedResult := `Packages matching pattern 'foobar' should ['not depend on external packages that matches [[x1 x2]]']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})

	t.Run("case 8", func(t *testing.T) {
		rule := configuration.DependenciesRule{
			Package: "foobar",
			ShouldNotDependsOn: &configuration.Dependencies{
				Standard: []string{"s100", "s200"},
			},
		}
		expectedResult := `Packages matching pattern 'foobar' should ['not depend on standard packages that matches [[s100 s200]]']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})
}
