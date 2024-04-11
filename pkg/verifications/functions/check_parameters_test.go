package functions

import (
	"testing"

	"github.com/fdaines/arch-go/internal/utils/values"

	"github.com/stretchr/testify/assert"
)

func TestCheckFunctionParameters(t *testing.T) {
	t.Run("check passes", func(t *testing.T) {
		pass, details := checkMaxParameters(functionTestDetails, values.GetIntRef(100))

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("check doesn't pass", func(t *testing.T) {
		pass, details := checkMaxParameters(functionTestDetails, values.GetIntRef(1))

		expected := []string{
			"Function myfunction1 in file /foo/bar/myfile.go receive too many parameters (2)",
			"Function myfunction2 in file /foo/bar/myfile.go receive too many parameters (2)",
			"Function myfunction22 in file /foo/bar/myfile2.go receive too many parameters (2)",
			"Function myfunction23 in file /foo/bar/myfile2.go receive too many parameters (20)",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})

	t.Run("when there is no functions", func(t *testing.T) {
		var emptyFunctions []*FunctionDetails

		pass, details := checkMaxParameters(emptyFunctions, values.GetIntRef(100))

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("when all function fails", func(t *testing.T) {
		pass, details := checkMaxParameters(functionTestDetails, values.GetIntRef(0))

		expected := []string{
			"Function myfunction1 in file /foo/bar/myfile.go receive too many parameters (2)",
			"Function myfunction2 in file /foo/bar/myfile.go receive too many parameters (2)",
			"Function myfunction21 in file /foo/bar/myfile2.go receive too many parameters (1)",
			"Function myfunction22 in file /foo/bar/myfile2.go receive too many parameters (2)",
			"Function myfunction23 in file /foo/bar/myfile2.go receive too many parameters (20)",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})
}
