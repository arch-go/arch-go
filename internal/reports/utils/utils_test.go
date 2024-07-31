package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/internal/reports/model"
	"github.com/arch-go/arch-go/internal/utils/values"
)

func TestReportsUtils(t *testing.T) {
	t.Run("ResolveRuleStatus", func(t *testing.T) {
		result1 := ResolveRuleStatus(0)
		assert.Equal(t, "PASS", result1)

		result2 := ResolveRuleStatus(1)
		assert.Equal(t, "FAIL", result2)

		result3 := ResolveRuleStatus(100)
		assert.Equal(t, "FAIL", result3)
	})

	t.Run("CheckVerificationStatus", func(t *testing.T) {
		total1 := values.GetIntRef(0)
		result1 := CheckVerificationStatus(true, total1)
		assert.Equal(t, "PASS", result1)
		assert.Equal(t, 0, *total1)

		total2 := values.GetIntRef(0)
		result2 := CheckVerificationStatus(false, total2)
		assert.Equal(t, "FAIL", result2)
		assert.Equal(t, 1, *total2)
	})

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
