package configuration

import (
	"bytes"
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"

	"github.com/fdaines/arch-go/internal/utils/values"
)

func TestDeprecationCheck(t *testing.T) {
	t.Run("nil configuration", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		color.Output = outputBuffer

		checkForDeprecatedConfiguration(nil)
		assert.Equal(t, "", outputBuffer.String())
	})

	t.Run("configuration with cycles rules", func(t *testing.T) {
		configurationRules := Config{
			FunctionsRules: []*FunctionsRule{
				{
					Package:  "**.qwerty.**",
					MaxLines: values.GetIntRef(123),
				},
			},
			ContentRules: []*ContentsRule{
				{
					Package:                    "**.blablabla.**",
					ShouldNotContainInterfaces: true,
				},
			},
			NamingRules: []*NamingRule{
				{
					Package: "**.foobar.**",
					InterfaceImplementationNamingRule: &InterfaceImplementationRule{
						StructsThatImplement:           "*Command",
						ShouldHaveSimpleNameEndingWith: values.GetStringRef("Foobar"),
					},
				},
			},
			CyclesRules: []*CyclesRule{
				{
					Package:                "foobar",
					ShouldNotContainCycles: true,
				},
			},
		}
		outputBuffer := bytes.NewBufferString("")
		color.Output = outputBuffer

		checkForDeprecatedConfiguration(&configurationRules)

		assert.Nil(t, configurationRules.CyclesRules)
		assert.Equal(t, "[WARNING] - Cycle Rules were deprecated on Arch-Go v1.4.0\n", outputBuffer.String())
	})

	t.Run("configuration without cycles rules", func(t *testing.T) {
		configurationRules := Config{
			FunctionsRules: []*FunctionsRule{
				{
					Package:  "**.qwerty.**",
					MaxLines: values.GetIntRef(123),
				},
			},
			ContentRules: []*ContentsRule{
				{
					Package:                    "**.blablabla.**",
					ShouldNotContainInterfaces: true,
				},
			},
			NamingRules: []*NamingRule{
				{
					Package: "**.foobar.**",
					InterfaceImplementationNamingRule: &InterfaceImplementationRule{
						StructsThatImplement:           "*Command",
						ShouldHaveSimpleNameEndingWith: values.GetStringRef("Foobar"),
					},
				},
			},
		}

		outputBuffer := bytes.NewBufferString("")
		color.Output = outputBuffer

		checkForDeprecatedConfiguration(&configurationRules)

		assert.Nil(t, configurationRules.CyclesRules)
		assert.Equal(t, "", outputBuffer.String())
	})
}
