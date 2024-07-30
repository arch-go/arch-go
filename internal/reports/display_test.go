package reports

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/internal/reports/model"
)

func TestDisplay(t *testing.T) {
	t.Run("empty report summary", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		summary := &model.ReportSummary{}

		expectedOutput := `--------------------------------------
	Execution Summary
--------------------------------------
Total Rules: 	0
Succeeded: 	0
Failed: 	0
--------------------------------------
`

		displaySummary(summary, outputBuffer)

		assert.Equal(t, expectedOutput, outputBuffer.String())
	})

	t.Run("minimal report summary", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		summary := &model.ReportSummary{
			Status: "OK",
			Time:   time.Now(),
			Total:  10,
			Passed: 8,
			Failed: 2,
		}

		expectedOutput := `--------------------------------------
	Execution Summary
--------------------------------------
Total Rules: 	10
Succeeded: 	8
Failed: 	2
--------------------------------------
`

		displaySummary(summary, outputBuffer)

		assert.Equal(t, expectedOutput, outputBuffer.String())
	})

	t.Run("full report summary failing compliance and coverage", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		summary := &model.ReportSummary{
			Status: "OK",
			Time:   time.Now(),
			Total:  10,
			Passed: 8,
			Failed: 2,
			ComplianceThreshold: &model.ThresholdSummary{
				Rate:      80,
				Threshold: 100,
				Status:    "FAIL",
			},
			CoverageThreshold: &model.ThresholdSummary{
				Rate:      85,
				Threshold: 90,
				Status:    "FAIL",
			},
		}

		expectedOutput := `--------------------------------------
	Execution Summary
--------------------------------------
Total Rules: 	10
Succeeded: 	8
Failed: 	2
--------------------------------------
Compliance:       80% (FAIL)
Coverage:         85% (FAIL)
`

		displaySummary(summary, outputBuffer)

		assert.Equal(t, expectedOutput, outputBuffer.String())
	})

	t.Run("full report summary passing compliance and coverage", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		summary := &model.ReportSummary{
			Status: "OK",
			Time:   time.Now(),
			Total:  10,
			Passed: 8,
			Failed: 2,
			ComplianceThreshold: &model.ThresholdSummary{
				Rate:      100,
				Threshold: 100,
				Status:    "PASS",
			},
			CoverageThreshold: &model.ThresholdSummary{
				Rate:      90,
				Threshold: 90,
				Status:    "PASS",
			},
		}

		expectedOutput := `--------------------------------------
	Execution Summary
--------------------------------------
Total Rules: 	10
Succeeded: 	8
Failed: 	2
--------------------------------------
Compliance:      100% (PASS)
Coverage:         90% (PASS)
`

		displaySummary(summary, outputBuffer)

		assert.Equal(t, expectedOutput, outputBuffer.String())
	})
}
