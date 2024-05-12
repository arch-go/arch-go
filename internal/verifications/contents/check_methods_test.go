package contents

import (
	"testing"

	"github.com/fdaines/arch-go/api/configuration"

	"github.com/stretchr/testify/assert"
)

func TestCheckContentMethods(t *testing.T) {
	t.Run("check passes - case 1", func(t *testing.T) {
		input := &PackageContents{
			Methods: 10,
		}
		contentsRule := &configuration.ContentsRule{ShouldOnlyContainMethods: true}

		pass, details := checkMethods(input, contentsRule)

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("check passes - case 2", func(t *testing.T) {
		input := &PackageContents{
			Methods: 10,
		}
		contentsRule := &configuration.ContentsRule{
			ShouldNotContainInterfaces: true,
			ShouldNotContainFunctions:  true,
			ShouldNotContainStructs:    true,
		}

		pass, details := checkMethods(input, contentsRule)

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("check passes - case 3", func(t *testing.T) {
		input := &PackageContents{
			Methods: 0,
		}
		contentsRule := &configuration.ContentsRule{
			ShouldNotContainMethods: true,
		}

		pass, details := checkMethods(input, contentsRule)

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("check fails - case 0", func(t *testing.T) {
		input := &PackageContents{
			Methods: 10,
		}
		contentsRule := &configuration.ContentsRule{ShouldNotContainMethods: true}

		pass, details := checkMethods(input, contentsRule)

		expected := []string{
			"contains methods and it should not",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})

	t.Run("check fails - case 1", func(t *testing.T) {
		input := &PackageContents{
			Methods: 10,
		}
		contentsRule := &configuration.ContentsRule{ShouldOnlyContainFunctions: true}

		pass, details := checkMethods(input, contentsRule)

		expected := []string{
			"contains methods and should only contain functions",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})

	t.Run("check fails - case 2", func(t *testing.T) {
		input := &PackageContents{
			Methods: 10,
		}
		contentsRule := &configuration.ContentsRule{ShouldOnlyContainStructs: true}

		pass, details := checkMethods(input, contentsRule)

		expected := []string{
			"contains methods and should only contain structs",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})

	t.Run("check fails - case 3", func(t *testing.T) {
		input := &PackageContents{
			Methods: 10,
		}
		contentsRule := &configuration.ContentsRule{ShouldOnlyContainInterfaces: true}

		pass, details := checkMethods(input, contentsRule)

		expected := []string{
			"contains methods and should only contain interfaces",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})
}
