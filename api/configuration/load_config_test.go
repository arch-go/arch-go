package configuration

import (
	"testing"

	"github.com/fdaines/arch-go/internal/utils/values"
	"github.com/stretchr/testify/assert"
)

func TestCheckThreshold(t *testing.T) {
	t.Run("nil threshold generates 100% coverage and compliance threshold", func(t *testing.T) {
		configuration := &Config{
			Threshold: nil,
		}
		expectedThreshold := &Threshold{
			Compliance: values.GetIntRef(100),
			Coverage:   values.GetIntRef(100),
		}

		checkThreshold(configuration)

		assert.Equal(t, expectedThreshold, configuration.Threshold)
	})

	t.Run("existing threshold generates does not change", func(t *testing.T) {
		configuration := &Config{
			Threshold: &Threshold{
				Compliance: values.GetIntRef(76),
				Coverage:   values.GetIntRef(36),
			},
		}
		expectedThreshold := &Threshold{
			Compliance: values.GetIntRef(76),
			Coverage:   values.GetIntRef(36),
		}

		checkThreshold(configuration)

		assert.Equal(t, expectedThreshold, configuration.Threshold)
	})
}
