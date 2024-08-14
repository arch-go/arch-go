package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/internal/reports/model"
)

func TestReportsUtilStatus(t *testing.T) {
	t.Run("ResolveRuleStatus", func(t *testing.T) {
		result1 := ResolveRuleStatus(0)
		assert.Equal(t, "PASS", result1)

		result2 := ResolveRuleStatus(1)
		assert.Equal(t, "FAIL", result2)

		result3 := ResolveRuleStatus(100)
		assert.Equal(t, "FAIL", result3)
	})

	t.Run("ResolveGlobalStatus", func(t *testing.T) {
		status := ResolveGlobalStatus(nil, nil)
		assert.Equal(t, "PASS", status)

		status = ResolveGlobalStatus(&model.ThresholdSummary{Pass: true}, nil)
		assert.Equal(t, "PASS", status)

		status = ResolveGlobalStatus(&model.ThresholdSummary{Pass: false}, nil)
		assert.Equal(t, "FAIL", status)

		status = ResolveGlobalStatus(nil, &model.ThresholdSummary{Pass: true})
		assert.Equal(t, "PASS", status)

		status = ResolveGlobalStatus(nil, &model.ThresholdSummary{Pass: false})
		assert.Equal(t, "FAIL", status)

		status = ResolveGlobalStatus(&model.ThresholdSummary{Pass: true}, &model.ThresholdSummary{Pass: true})
		assert.Equal(t, "PASS", status)

		status = ResolveGlobalStatus(&model.ThresholdSummary{Pass: true}, &model.ThresholdSummary{Pass: false})
		assert.Equal(t, "FAIL", status)

		status = ResolveGlobalStatus(&model.ThresholdSummary{Pass: false}, &model.ThresholdSummary{Pass: true})
		assert.Equal(t, "FAIL", status)

		status = ResolveGlobalStatus(&model.ThresholdSummary{Pass: false}, &model.ThresholdSummary{Pass: false})
		assert.Equal(t, "FAIL", status)
	})
}
