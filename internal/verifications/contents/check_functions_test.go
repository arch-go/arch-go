package contents

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/api/configuration"
)

func TestCheckContentFunctions(t *testing.T) {
	t.Run("check passes - case 1", func(t *testing.T) {
		input := &PackageContents{
			Functions: 10,
		}
		contentsRule := &configuration.ContentsRule{ShouldOnlyContainFunctions: true}

		pass, details := checkFunctions(input, contentsRule)

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("check passes - case 2", func(t *testing.T) {
		input := &PackageContents{
			Functions: 10,
		}
		contentsRule := &configuration.ContentsRule{
			ShouldNotContainMethods:    true,
			ShouldNotContainInterfaces: true,
			ShouldNotContainStructs:    true,
		}

		pass, details := checkFunctions(input, contentsRule)

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("check passes - case 3", func(t *testing.T) {
		input := &PackageContents{
			Functions: 0,
		}
		contentsRule := &configuration.ContentsRule{
			ShouldNotContainFunctions: true,
		}

		pass, details := checkFunctions(input, contentsRule)

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("check fails - case 0", func(t *testing.T) {
		input := &PackageContents{
			Functions: 10,
		}
		contentsRule := &configuration.ContentsRule{ShouldNotContainFunctions: true}

		pass, details := checkFunctions(input, contentsRule)

		expected := []string{
			"contains functions and it should not",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})

	t.Run("check fails - case 1", func(t *testing.T) {
		input := &PackageContents{
			Functions: 10,
		}
		contentsRule := &configuration.ContentsRule{ShouldOnlyContainInterfaces: true}

		pass, details := checkFunctions(input, contentsRule)

		expected := []string{
			"contains functions and should only contain interfaces",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})

	t.Run("check fails - case 2", func(t *testing.T) {
		input := &PackageContents{
			Functions: 10,
		}
		contentsRule := &configuration.ContentsRule{ShouldOnlyContainStructs: true}

		pass, details := checkFunctions(input, contentsRule)

		expected := []string{
			"contains functions and should only contain structs",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})

	t.Run("check fails - case 3", func(t *testing.T) {
		input := &PackageContents{
			Functions: 10,
		}
		contentsRule := &configuration.ContentsRule{ShouldOnlyContainMethods: true}

		pass, details := checkFunctions(input, contentsRule)

		expected := []string{
			"contains functions and should only contain methods",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})
}
