package html

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUncoveredPackages(t *testing.T) {
	t.Run("Returns empty string when there is no coverage analysis", func(t *testing.T) {
		summary := createRulesSummaryMock(false, false)
		expected := ""

		result := uncoveredPackages(summary)

		assert.Equal(t, expected, result)
	})

	t.Run("Returns empty string when there are no uncovered packages", func(t *testing.T) {
		summary := createRulesSummaryMock(false, true)
		summary.CoverageThreshold.Violations = []string{}
		expected := ""

		result := uncoveredPackages(summary)

		assert.Equal(t, expected, result)
	})

	t.Run("Returns uncovered packages section when there are uncovered packages", func(t *testing.T) {
		summary := createRulesSummaryMock(false, true)
		expected := "\n<h3>Uncovered Packages</h3>\n<ul>\n\t<li>foobar</li><li>barfoo</li>\n</ul>\n"

		result := uncoveredPackages(summary)

		assert.Equal(t, expected, result)
	})
}
