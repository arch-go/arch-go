package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fdaines/arch-go/internal/utils/values"
)

func TestCheckFunctionMaxPublicFunctions(t *testing.T) {
	t.Run("check passes", func(t *testing.T) {
		pass, details := checkMaxPublicFunctions(functionTestDetails, values.GetIntRef(5))

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("check passes when max value is nil", func(t *testing.T) {
		pass, details := checkMaxPublicFunctions(functionTestDetails, nil)

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("check doesn't pass", func(t *testing.T) {
		pass, details := checkMaxPublicFunctions(functionTestDetails, values.GetIntRef(1))

		expected := []string{
			"File /foo/bar/myfile.go has too many public functions (2)",
			"File /foo/bar/myfile2.go has too many public functions (2)",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})

	t.Run("when there is no functions", func(t *testing.T) {
		var emptyFunctions []*FunctionDetails

		pass, details := checkMaxPublicFunctions(emptyFunctions, values.GetIntRef(5))

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("when all function fails", func(t *testing.T) {
		pass, details := checkMaxPublicFunctions(functionTestDetails, values.GetIntRef(1))

		expected := []string{
			"File /foo/bar/myfile.go has too many public functions (2)",
			"File /foo/bar/myfile2.go has too many public functions (2)",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})
}
