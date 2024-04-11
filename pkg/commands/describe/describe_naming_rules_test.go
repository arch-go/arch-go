package describe

import (
	"bytes"
	"io"
	"testing"

	"github.com/fdaines/arch-go/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestDescribeNamingRules(t *testing.T) {
	t.Run("interface implementation naming rules", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rules := []*config.NamingRule{
			{
				Package: "foobar",
				InterfaceImplementationNamingRule: &config.InterfaceImplementationRule{
					StructsThatImplement:             "QWERTY",
					ShouldHaveSimpleNameStartingWith: "test",
				},
			},
			{
				Package: "barfoo",
				InterfaceImplementationNamingRule: &config.InterfaceImplementationRule{
					StructsThatImplement:           "FOOBAR",
					ShouldHaveSimpleNameEndingWith: "blablabla",
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
		outputBuffer := bytes.NewBufferString("")
		var rules []*config.NamingRule
		expectedOutput := `Naming Rules
	* No rules defined
`

		describeNamingRules(rules, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)

		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})
}
