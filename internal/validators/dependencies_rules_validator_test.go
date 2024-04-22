package validators

import (
	"testing"

	"github.com/fdaines/arch-go/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestDependenciesRulesValidator(t *testing.T) {
	t.Run("dependenciesSize", func(t *testing.T) {
		dependencies := &config.Dependencies{
			Internal: []string{"foo", "bar"},
			External: []string{"blablabla"},
			Standard: []string{"std1", "std2"},
		}

		result := dependenciesSize(dependencies)
		assert.Equal(t, 5, result)
	})
}
