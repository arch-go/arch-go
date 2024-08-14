package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/internal/reports/model"
)

func TestReportsUtilStatus(t *testing.T) {
	t.Run("ResolveGlobalStatus", func(t *testing.T) {
		status := ResolveGlobalStatus(nil, nil)
		assert.Equal(t, true, status)

		status = ResolveGlobalStatus(&model.ThresholdSummary{Pass: true}, nil)
		assert.Equal(t, true, status)

		status = ResolveGlobalStatus(&model.ThresholdSummary{Pass: false}, nil)
		assert.Equal(t, false, status)

		status = ResolveGlobalStatus(nil, &model.ThresholdSummary{Pass: true})
		assert.Equal(t, true, status)

		status = ResolveGlobalStatus(nil, &model.ThresholdSummary{Pass: false})
		assert.Equal(t, false, status)

		status = ResolveGlobalStatus(&model.ThresholdSummary{Pass: true}, &model.ThresholdSummary{Pass: true})
		assert.Equal(t, true, status)

		status = ResolveGlobalStatus(&model.ThresholdSummary{Pass: true}, &model.ThresholdSummary{Pass: false})
		assert.Equal(t, false, status)

		status = ResolveGlobalStatus(&model.ThresholdSummary{Pass: false}, &model.ThresholdSummary{Pass: true})
		assert.Equal(t, false, status)

		status = ResolveGlobalStatus(&model.ThresholdSummary{Pass: false}, &model.ThresholdSummary{Pass: false})
		assert.Equal(t, false, status)
	})
}
