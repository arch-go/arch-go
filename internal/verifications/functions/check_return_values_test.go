package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/internal/utils/values"
)

func TestCheckFunctionReturnValues(t *testing.T) {
	t.Run("check passes", func(t *testing.T) {
		pass, details := checkMaxReturnValues(functionTestDetails, values.GetIntRef(100))

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("check passes when max value is nil", func(t *testing.T) {
		pass, details := checkMaxReturnValues(functionTestDetails, nil)

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("check doesn't pass", func(t *testing.T) {
		pass, details := checkMaxReturnValues(functionTestDetails, values.GetIntRef(1))

		expected := []string{
			"Function myfunction1 in file /foo/bar/myfile.go returns too many values (2)",
			"Function myfunction21 in file /foo/bar/myfile2.go returns too many values (2)",
			"Function myfunction22 in file /foo/bar/myfile2.go returns too many values (5)",
			"Function myfunction23 in file /foo/bar/myfile2.go returns too many values (21)",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})

	t.Run("when there is no functions", func(t *testing.T) {
		var emptyFunctions []*FunctionDetails

		pass, details := checkMaxReturnValues(emptyFunctions, values.GetIntRef(100))

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("when all function fails", func(t *testing.T) {
		pass, details := checkMaxReturnValues(functionTestDetails, values.GetIntRef(0))

		expected := []string{
			"Function myfunction1 in file /foo/bar/myfile.go returns too many values (2)",
			"Function myfunction2 in file /foo/bar/myfile.go returns too many values (1)",
			"Function myfunction21 in file /foo/bar/myfile2.go returns too many values (2)",
			"Function myfunction22 in file /foo/bar/myfile2.go returns too many values (5)",
			"Function myfunction23 in file /foo/bar/myfile2.go returns too many values (21)",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})
}
