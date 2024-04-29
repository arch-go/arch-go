package describe

import (
	"bytes"
	"io"
	"testing"

	"github.com/fdaines/arch-go/pkg/archgo/configuration"

	"github.com/stretchr/testify/assert"
)

func TestDescribeContentRules(t *testing.T) {
	t.Run("content rules", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rules := []*configuration.ContentsRule{
			{
				Package:                     "foobar1",
				ShouldOnlyContainInterfaces: true,
			},
			{
				Package:                  "foobar2",
				ShouldOnlyContainStructs: true,
			},
			{
				Package:                    "foobar3",
				ShouldOnlyContainFunctions: true,
			},
			{
				Package:                  "foobar4",
				ShouldOnlyContainMethods: true,
			},
			{
				Package:                   "foobar5",
				ShouldNotContainStructs:   true,
				ShouldNotContainFunctions: true,
				ShouldNotContainMethods:   true,
			},
			{
				Package:                    "foobar6",
				ShouldNotContainInterfaces: true,
				ShouldNotContainFunctions:  true,
				ShouldNotContainMethods:    true,
			},
		}
		expectedOutput := `Content Rules
	* Packages that match pattern 'foobar1' should only contain interfaces
	* Packages that match pattern 'foobar2' should only contain structs
	* Packages that match pattern 'foobar3' should only contain functions
	* Packages that match pattern 'foobar4' should only contain methods
	* Packages that match pattern 'foobar5' should not contain structs or functions or methods
	* Packages that match pattern 'foobar6' should not contain interfaces or functions or methods
`

		describeContentRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)

		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("empty rules", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		var rules []*configuration.ContentsRule
		expectedOutput := `Content Rules
	* No rules defined
`

		describeContentRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)

		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})
}
