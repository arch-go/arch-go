package console

import (
	"bytes"
	"testing"

	"github.com/fdaines/arch-go/internal/reports/model"

	"github.com/stretchr/testify/assert"
)

func TestConsoleReportGenerator(t *testing.T) {
	t.Run("Empty Report", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		report := &model.Report{}

		expectedOutput := ``

		GenerateConsoleReport(report, outputBuffer)

		assert.Equal(t, expectedOutput, outputBuffer.String())
	})

	t.Run("Full Report", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		report := &model.Report{
			Summary: &model.ReportSummary{
				ComplianceThreshold: &model.ThresholdSummary{
					Rate:      87,
					Threshold: 100,
					Status:    "FAIL",
				},
				CoverageThreshold: &model.ThresholdSummary{
					Rate:      87,
					Threshold: 60,
					Status:    "PASS",
				},
			},
			Details: &model.ReportDetails{},
		}

		expectedOutput := `+---+--------------------+-------------+-------------+-------------+
| # | RULE TYPE          |       TOTAL |      PASSED |      FAILED |
+---+--------------------+-------------+-------------+-------------+
| 1 | Dependencies Rules |           0 |           0 |           0 |
| 2 | Functions Rules    |           0 |           0 |           0 |
| 3 | Contents Rules     |           0 |           0 |           0 |
| 4 | Naming Rules       |           0 |           0 |           0 |
+---+--------------------+-------------+-------------+-------------+
|   | COMPLIANCE RATE    |  87% [FAIL]                             |
|   | COVERAGE RATE      |  87% [PASS]                             |
+---+--------------------+-----------------------------------------+
`

		GenerateConsoleReport(report, outputBuffer)

		assert.Equal(t, expectedOutput, outputBuffer.String())
	})
}
