package describe

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/v2/api/configuration"
	"github.com/arch-go/arch-go/v2/internal/utils/values"
)

func TestDescribeFunctionRules(t *testing.T) {
	t.Run("function rules", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rules := []*configuration.FunctionsRule{
			{
				Package:                  "foobar",
				MaxLines:                 values.GetIntRef(123),
				MaxParameters:            values.GetIntRef(32),
				MaxPublicFunctionPerFile: values.GetIntRef(24),
				MaxReturnValues:          values.GetIntRef(3),
			},
		}
		expectedOutput := `Function Rules
	* Packages that match pattern 'foobar' should comply with the following rules:
		* Functions should not have more than 123 lines
		* Functions should not have more than 32 parameters
		* Functions should not have more than 3 return values
		* Files should not have more than 24 public functions
`

		describeFunctionRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)
		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("function rules with blanks - case 1", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rules := []*configuration.FunctionsRule{
			{
				Package:                  "foobar",
				MaxParameters:            values.GetIntRef(32),
				MaxPublicFunctionPerFile: values.GetIntRef(24),
				MaxReturnValues:          values.GetIntRef(3),
			},
		}
		expectedOutput := `Function Rules
	* Packages that match pattern 'foobar' should comply with the following rules:
		* Functions should not have more than 32 parameters
		* Functions should not have more than 3 return values
		* Files should not have more than 24 public functions
`

		describeFunctionRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)
		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("function rules with blanks - case 2", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rules := []*configuration.FunctionsRule{
			{
				Package:                  "foobar",
				MaxLines:                 values.GetIntRef(123),
				MaxPublicFunctionPerFile: values.GetIntRef(24),
				MaxReturnValues:          values.GetIntRef(3),
			},
		}
		expectedOutput := `Function Rules
	* Packages that match pattern 'foobar' should comply with the following rules:
		* Functions should not have more than 123 lines
		* Functions should not have more than 3 return values
		* Files should not have more than 24 public functions
`

		describeFunctionRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)
		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("function rules with blanks - case 3", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rules := []*configuration.FunctionsRule{
			{
				Package:         "foobar",
				MaxLines:        values.GetIntRef(123),
				MaxParameters:   values.GetIntRef(32),
				MaxReturnValues: values.GetIntRef(3),
			},
		}
		expectedOutput := `Function Rules
	* Packages that match pattern 'foobar' should comply with the following rules:
		* Functions should not have more than 123 lines
		* Functions should not have more than 32 parameters
		* Functions should not have more than 3 return values
`

		describeFunctionRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)
		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("function rules with blanks - case 4", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rules := []*configuration.FunctionsRule{
			{
				Package:                  "foobar",
				MaxLines:                 values.GetIntRef(123),
				MaxParameters:            values.GetIntRef(32),
				MaxPublicFunctionPerFile: values.GetIntRef(24),
			},
		}
		expectedOutput := `Function Rules
	* Packages that match pattern 'foobar' should comply with the following rules:
		* Functions should not have more than 123 lines
		* Functions should not have more than 32 parameters
		* Files should not have more than 24 public functions
`

		describeFunctionRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)
		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("empty rules", func(t *testing.T) {
		var rules []*configuration.FunctionsRule

		outputBuffer := bytes.NewBufferString("")
		expectedOutput := `Function Rules
	* No rules defined
`

		describeFunctionRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)
		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})
}
