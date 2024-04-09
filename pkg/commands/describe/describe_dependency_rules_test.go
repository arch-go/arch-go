package describe

import (
	"bytes"
	"github.com/fdaines/arch-go/old/config"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestDescribeDependencyRules(t *testing.T) {

	t.Run("dependency rules", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rules := []*config.DependenciesRule{
			{
				Package: "foobar",
				ShouldOnlyDependsOn: &config.Dependencies{
					Internal: []string{"foo", "bar"},
					External: []string{"xyz", "abc"},
					Standard: []string{"aaa", "bbb"},
				},
			},
			{
				Package: "foobar1",
				ShouldOnlyDependsOn: &config.Dependencies{
					Internal: []string{"foo", "bar"},
				},
			},
			{
				Package: "foobar2",
				ShouldOnlyDependsOn: &config.Dependencies{
					External: []string{"xyz", "abc"},
				},
			},
			{
				Package: "foobar3",
				ShouldOnlyDependsOn: &config.Dependencies{
					Standard: []string{"aaa", "bbb"},
				},
			},
			{
				Package: "barfoo",
				ShouldNotDependsOn: &config.Dependencies{
					Internal: []string{"i1", "i2"},
					External: []string{"e1", "e2"},
					Standard: []string{"s1", "s2"},
				},
			},
			{
				Package: "barfoo1",
				ShouldNotDependsOn: &config.Dependencies{
					Internal: []string{"i1", "i2"},
				},
			},
			{
				Package: "barfoo2",
				ShouldNotDependsOn: &config.Dependencies{
					External: []string{"e1", "e2"},
				},
			},
			{
				Package: "barfoo3",
				ShouldNotDependsOn: &config.Dependencies{
					Standard: []string{"s1", "s2"},
				},
			},
		}
		expectedOutput := `Dependency Rules
	* Packages that match pattern 'foobar',
		* Should only depends on:
			* Internal Packages that matches:
				- 'foo'
				- 'bar'
			* External Packages that matches:
				- 'xyz'
				- 'abc'
			* StandardLib Packages that matches:
				- 'aaa'
				- 'bbb'
	* Packages that match pattern 'foobar1',
		* Should only depends on:
			* Internal Packages that matches:
				- 'foo'
				- 'bar'
	* Packages that match pattern 'foobar2',
		* Should only depends on:
			* External Packages that matches:
				- 'xyz'
				- 'abc'
	* Packages that match pattern 'foobar3',
		* Should only depends on:
			* StandardLib Packages that matches:
				- 'aaa'
				- 'bbb'
	* Packages that match pattern 'barfoo',
		* Should not depends on:
			* Internal Packages that matches:
				- 'i1'
				- 'i2'
			* External Packages that matches:
				- 'e1'
				- 'e2'
			* StandardLib Packages that matches:
				- 's1'
				- 's2'
	* Packages that match pattern 'barfoo1',
		* Should not depends on:
			* Internal Packages that matches:
				- 'i1'
				- 'i2'
	* Packages that match pattern 'barfoo2',
		* Should not depends on:
			* External Packages that matches:
				- 'e1'
				- 'e2'
	* Packages that match pattern 'barfoo3',
		* Should not depends on:
			* StandardLib Packages that matches:
				- 's1'
				- 's2'
`

		describeDependencyRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)

		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("empty rules", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		var rules []*config.DependenciesRule
		expectedOutput := `Dependency Rules
	* No rules defined
`

		describeDependencyRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)

		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})
}
