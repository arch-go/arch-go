package contents

import (
	"testing"

	"github.com/fdaines/arch-go/pkg/config"

	"github.com/stretchr/testify/assert"
)

func TestCheckContentStructs(t *testing.T) {
	t.Run("check passes - case 1", func(t *testing.T) {
		input := &PackageContents{
			Structs: 10,
		}
		contentsRule := &config.ContentsRule{ShouldOnlyContainStructs: true}

		pass, details := checkStructs(input, contentsRule)

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("check passes - case 2", func(t *testing.T) {
		input := &PackageContents{
			Structs: 10,
		}
		contentsRule := &config.ContentsRule{
			ShouldNotContainInterfaces: true,
			ShouldNotContainFunctions:  true,
			ShouldNotContainMethods:    true,
		}

		pass, details := checkStructs(input, contentsRule)

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("check passes - case 3", func(t *testing.T) {
		input := &PackageContents{
			Structs: 0,
		}
		contentsRule := &config.ContentsRule{
			ShouldNotContainStructs: true,
		}

		pass, details := checkStructs(input, contentsRule)

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("check fails - case 0", func(t *testing.T) {
		input := &PackageContents{
			Structs: 10,
		}
		contentsRule := &config.ContentsRule{ShouldNotContainStructs: true}

		pass, details := checkStructs(input, contentsRule)

		expected := []string{
			"contains structs and it should not",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})
	t.Run("check fails - case 1", func(t *testing.T) {
		input := &PackageContents{
			Structs: 10,
		}
		contentsRule := &config.ContentsRule{ShouldOnlyContainFunctions: true}

		pass, details := checkStructs(input, contentsRule)

		expected := []string{
			"contains structs and should only contain functions",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})

	t.Run("check fails - case 2", func(t *testing.T) {
		input := &PackageContents{
			Structs: 10,
		}
		contentsRule := &config.ContentsRule{ShouldOnlyContainMethods: true}

		pass, details := checkStructs(input, contentsRule)

		expected := []string{
			"contains structs and should only contain methods",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})

	t.Run("check fails - case 3", func(t *testing.T) {
		input := &PackageContents{
			Structs: 10,
		}
		contentsRule := &config.ContentsRule{ShouldOnlyContainInterfaces: true}

		pass, details := checkStructs(input, contentsRule)

		expected := []string{
			"contains structs and should only contain interfaces",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})
}
