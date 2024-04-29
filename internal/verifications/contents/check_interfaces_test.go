package contents

import (
	"testing"

	"github.com/fdaines/arch-go/pkg/archgo/configuration"

	"github.com/stretchr/testify/assert"
)

func TestCheckContentInterfaces(t *testing.T) {
	t.Run("check passes - case 1", func(t *testing.T) {
		input := &PackageContents{
			Interfaces: 10,
		}
		contentsRule := &configuration.ContentsRule{ShouldOnlyContainInterfaces: true}

		pass, details := checkInterfaces(input, contentsRule)

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("check passes - case 2", func(t *testing.T) {
		input := &PackageContents{
			Interfaces: 10,
		}
		contentsRule := &configuration.ContentsRule{
			ShouldNotContainMethods:   true,
			ShouldNotContainFunctions: true,
			ShouldNotContainStructs:   true,
		}

		pass, details := checkInterfaces(input, contentsRule)

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("check passes - case 3", func(t *testing.T) {
		input := &PackageContents{
			Interfaces: 0,
		}
		contentsRule := &configuration.ContentsRule{
			ShouldNotContainInterfaces: true,
		}

		pass, details := checkInterfaces(input, contentsRule)

		var expected []string

		assert.ElementsMatch(t, expected, details)
		assert.True(t, pass)
	})

	t.Run("check fails - case 0", func(t *testing.T) {
		input := &PackageContents{
			Interfaces: 10,
		}
		contentsRule := &configuration.ContentsRule{ShouldNotContainInterfaces: true}

		pass, details := checkInterfaces(input, contentsRule)

		expected := []string{
			"contains interfaces and it should not",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})

	t.Run("check fails - case 1", func(t *testing.T) {
		input := &PackageContents{
			Interfaces: 10,
		}
		contentsRule := &configuration.ContentsRule{ShouldOnlyContainFunctions: true}

		pass, details := checkInterfaces(input, contentsRule)

		expected := []string{
			"contains interfaces and should only contain functions",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})

	t.Run("check fails - case 2", func(t *testing.T) {
		input := &PackageContents{
			Interfaces: 10,
		}
		contentsRule := &configuration.ContentsRule{ShouldOnlyContainStructs: true}

		pass, details := checkInterfaces(input, contentsRule)

		expected := []string{
			"contains interfaces and should only contain structs",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})

	t.Run("check fails - case 3", func(t *testing.T) {
		input := &PackageContents{
			Interfaces: 10,
		}
		contentsRule := &configuration.ContentsRule{ShouldOnlyContainMethods: true}

		pass, details := checkInterfaces(input, contentsRule)

		expected := []string{
			"contains interfaces and should only contain methods",
		}

		assert.ElementsMatch(t, expected, details)
		assert.False(t, pass)
	})
}
