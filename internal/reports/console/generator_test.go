package console

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/internal/reports/model"
	"github.com/arch-go/arch-go/internal/utils/values"
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
			Compliance: model.Compliance{
				Rate:      87,
				Threshold: values.GetIntRef(100),
				Pass:      false,
				Details:   &model.ReportDetails{},
			},
			Coverage: model.Coverage{
				Rate:      87,
				Threshold: values.GetIntRef(60),
				Pass:      true,
			},
		}

		expectedOutput := `+---+--------------------+-------+--------+--------+
| # | RULE TYPE          | TOTAL | PASSED | FAILED |
+---+--------------------+-------+--------+--------+
| 1 | Dependencies Rules |     0 |      0 |      0 |
| 2 | Functions Rules    |     0 |      0 |      0 |
| 3 | Contents Rules     |     0 |      0 |      0 |
| 4 | Naming Rules       |     0 |      0 |      0 |
+---+--------------------+-------+--------+--------+
|   | COMPLIANCE RATE    |  87% [FAIL]             |
|   | COVERAGE RATE      |  87% [PASS]             |
+---+--------------------+-------------------------+
`

		GenerateConsoleReport(report, outputBuffer)

		assert.Equal(t, expectedOutput, outputBuffer.String())
	})
}
