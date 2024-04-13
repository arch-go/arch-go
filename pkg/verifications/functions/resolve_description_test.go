package functions

import (
	"testing"

	"github.com/fdaines/arch-go/internal/utils/values"

	"github.com/fdaines/arch-go/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestResolveNamingRuleDescription(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		rule := config.FunctionsRule{
			Package:                  "foobar",
			MaxLines:                 values.GetIntRef(10),
			MaxParameters:            values.GetIntRef(4),
			MaxReturnValues:          values.GetIntRef(2),
			MaxPublicFunctionPerFile: values.GetIntRef(5),
		}
		expectedResult := `Functions in packages matching pattern 'foobar' should have ['at most 4 parameters','at most 2 return values','at most 10 lines','no more than 5 public functions per file']`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})

	t.Run("case 2", func(t *testing.T) {
		rule := config.FunctionsRule{
			Package: "foobar",
		}
		expectedResult := `Functions in packages matching pattern 'foobar' should have []`

		description := resolveDescription(rule)

		assert.Equal(t, expectedResult, description)
	})
}
