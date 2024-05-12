package functions

import (
	"testing"

	"github.com/fdaines/arch-go/internal/utils/values"

	"github.com/stretchr/testify/assert"
)

func TestCheckFunctionMaxLines(t *testing.T) {
	t.Run("check passes", func(t *testing.T) {
		pass, details := checkMaxLines(functionTestDetails, values.GetIntRef(200))

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("check passes when max value is nil", func(t *testing.T) {
		pass, details := checkMaxLines(functionTestDetails, nil)

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("check doesn't pass", func(t *testing.T) {
		pass, details := checkMaxLines(functionTestDetails, values.GetIntRef(100))

		expected := []string{
			"Function myfunction23 in file /foo/bar/myfile2.go has too many lines (200)",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})

	t.Run("when there is no functions", func(t *testing.T) {
		var emptyFunctions []*FunctionDetails

		pass, details := checkMaxLines(emptyFunctions, values.GetIntRef(100))

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("when all function fails", func(t *testing.T) {
		pass, details := checkMaxLines(functionTestDetails, values.GetIntRef(9))

		expected := []string{
			"Function myfunction1 in file /foo/bar/myfile.go has too many lines (10)",
			"Function myfunction2 in file /foo/bar/myfile.go has too many lines (100)",
			"Function myfunction21 in file /foo/bar/myfile2.go has too many lines (15)",
			"Function myfunction22 in file /foo/bar/myfile2.go has too many lines (100)",
			"Function myfunction23 in file /foo/bar/myfile2.go has too many lines (200)",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})
}
