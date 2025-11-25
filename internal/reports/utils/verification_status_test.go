package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/v2/internal/reports/model"
)

func TestReportsUtilVerificationStatus(t *testing.T) {
	t.Run("ResolveVerificationStatus", func(t *testing.T) {
		v1 := &model.Verification{}
		ResolveVerificationStatus(true, v1)
		assert.Equal(t, 1, v1.Passed)
		assert.Equal(t, 0, v1.Failed)
		assert.Equal(t, 1, v1.Total)

		v2 := &model.Verification{}
		ResolveVerificationStatus(false, v2)
		assert.Equal(t, 0, v2.Passed)
		assert.Equal(t, 1, v2.Failed)
		assert.Equal(t, 1, v2.Total)
	})
}
