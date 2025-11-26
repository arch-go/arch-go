package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/v2/internal/reports/model"
)

func TestReportsUtilStatus(t *testing.T) {
	t.Run("ResolveGlobalStatus", func(t *testing.T) {
		status := ResolveGlobalStatus(nil, nil)
		assert.True(t, status)

		status = ResolveGlobalStatus(&model.ThresholdSummary{Pass: true}, nil)
		assert.True(t, status)

		status = ResolveGlobalStatus(&model.ThresholdSummary{Pass: false}, nil)
		assert.False(t, status)

		status = ResolveGlobalStatus(nil, &model.ThresholdSummary{Pass: true})
		assert.True(t, status)

		status = ResolveGlobalStatus(nil, &model.ThresholdSummary{Pass: false})
		assert.False(t, status)

		status = ResolveGlobalStatus(&model.ThresholdSummary{Pass: true}, &model.ThresholdSummary{Pass: true})
		assert.True(t, status)

		status = ResolveGlobalStatus(&model.ThresholdSummary{Pass: true}, &model.ThresholdSummary{Pass: false})
		assert.False(t, status)

		status = ResolveGlobalStatus(&model.ThresholdSummary{Pass: false}, &model.ThresholdSummary{Pass: true})
		assert.False(t, status)

		status = ResolveGlobalStatus(&model.ThresholdSummary{Pass: false}, &model.ThresholdSummary{Pass: false})
		assert.False(t, status)
	})
}
