package describe

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fdaines/arch-go/api/configuration"
	"github.com/fdaines/arch-go/internal/utils/values"
)

func TestDescribeNamingRules(t *testing.T) {
	t.Run("interface implementation naming rules", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rules := []*configuration.NamingRule{
			{
				Package: "foobar",
				InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
					StructsThatImplement:             "QWERTY",
					ShouldHaveSimpleNameStartingWith: values.GetStringRef("test"),
				},
			},
			{
				Package: "barfoo",
				InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
					StructsThatImplement:           "FOOBAR",
					ShouldHaveSimpleNameEndingWith: values.GetStringRef("blablabla"),
				},
			},
		}
		expectedOutput := `Naming Rules
	* Packages that match pattern 'foobar' should comply with:
		* Structs that implement interfaces matching name 'QWERTY' should have simple name starting with 'test'
	* Packages that match pattern 'barfoo' should comply with:
		* Structs that implement interfaces matching name 'FOOBAR' should have simple name ending with 'blablabla'
`

		describeNamingRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)
		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("empty rules", func(t *testing.T) {
		var rules []*configuration.NamingRule

		outputBuffer := bytes.NewBufferString("")
		expectedOutput := `Naming Rules
	* No rules defined
`

		describeNamingRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)
		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})
}
