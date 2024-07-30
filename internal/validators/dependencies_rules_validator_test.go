package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/api/configuration"
)

func TestDependenciesRulesValidator(t *testing.T) {
	t.Run("dependenciesSize", func(t *testing.T) {
		dependencies := &configuration.Dependencies{
			Internal: []string{"foo", "bar"},
			External: []string{"blablabla"},
			Standard: []string{"std1", "std2"},
		}

		result := dependenciesSize(dependencies)
		assert.Equal(t, 5, result)
	})
}
