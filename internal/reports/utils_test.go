package reports

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fdaines/arch-go/internal/reports/model"
	"github.com/fdaines/arch-go/internal/utils/values"
)

func TestReportsUtils(t *testing.T) {
	t.Run("resolveRuleStatus", func(t *testing.T) {
		result1 := resolveRuleStatus(0)
		assert.Equal(t, "PASS", result1)

		result2 := resolveRuleStatus(1)
		assert.Equal(t, "FAIL", result2)

		result3 := resolveRuleStatus(100)
		assert.Equal(t, "FAIL", result3)
	})

	t.Run("checkVerificationStatus", func(t *testing.T) {
		total1 := values.GetIntRef(0)
		result1 := checkVerificationStatus(true, total1)
		assert.Equal(t, "PASS", result1)
		assert.Equal(t, 0, *total1)

		total2 := values.GetIntRef(0)
		result2 := checkVerificationStatus(false, total2)
		assert.Equal(t, "FAIL", result2)
		assert.Equal(t, 1, *total2)
	})

	t.Run("resolveVerificationStatus", func(t *testing.T) {
		v1 := &model.Verification{}
		resolveVerificationStatus(true, v1)
		assert.Equal(t, 1, v1.Passed)
		assert.Equal(t, 0, v1.Failed)
		assert.Equal(t, 1, v1.Total)

		v2 := &model.Verification{}
		resolveVerificationStatus(false, v2)
		assert.Equal(t, 0, v2.Passed)
		assert.Equal(t, 1, v2.Failed)
		assert.Equal(t, 1, v2.Total)
	})

	t.Run("resolveGlobalStatus", func(t *testing.T) {
		status := resolveGlobalStatus(nil, nil)
		assert.Equal(t, "PASS", status)

		status = resolveGlobalStatus(&model.ThresholdSummary{Status: "PASS"}, nil)
		assert.Equal(t, "PASS", status)

		status = resolveGlobalStatus(&model.ThresholdSummary{Status: "FAIL"}, nil)
		assert.Equal(t, "FAIL", status)

		status = resolveGlobalStatus(nil, &model.ThresholdSummary{Status: "PASS"})
		assert.Equal(t, "PASS", status)

		status = resolveGlobalStatus(nil, &model.ThresholdSummary{Status: "FAIL"})
		assert.Equal(t, "FAIL", status)

		status = resolveGlobalStatus(&model.ThresholdSummary{Status: "PASS"}, &model.ThresholdSummary{Status: "PASS"})
		assert.Equal(t, "PASS", status)

		status = resolveGlobalStatus(&model.ThresholdSummary{Status: "PASS"}, &model.ThresholdSummary{Status: "FAIL"})
		assert.Equal(t, "FAIL", status)

		status = resolveGlobalStatus(&model.ThresholdSummary{Status: "FAIL"}, &model.ThresholdSummary{Status: "PASS"})
		assert.Equal(t, "FAIL", status)

		status = resolveGlobalStatus(&model.ThresholdSummary{Status: "FAIL"}, &model.ThresholdSummary{Status: "FAIL"})
		assert.Equal(t, "FAIL", status)
	})
}
