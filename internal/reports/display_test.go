package reports

import (
	"bytes"
	"io"
	"os"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/v2/internal/common"
	"github.com/arch-go/arch-go/v2/internal/reports/html"
	"github.com/arch-go/arch-go/v2/internal/reports/json"
	"github.com/arch-go/arch-go/v2/internal/reports/model"
	"github.com/arch-go/arch-go/v2/internal/utils/values"
)

func TestDisplay(t *testing.T) {
	t.Run("empty report summary", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		report := &model.Report{}

		expectedOutput := `--------------------------------------
	Execution Summary
--------------------------------------
Total Rules: 	0
Succeeded: 	0
Failed: 	0
--------------------------------------
`

		displaySummary(report, outputBuffer)

		assert.Equal(t, expectedOutput, outputBuffer.String())
	})

	t.Run("minimal report summary", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		report := &model.Report{
			Summary: &model.Summary{
				Pass: true,
				Time: time.Now(),
			},
			Compliance: model.Compliance{
				Total:  10,
				Passed: 8,
				Failed: 2,
			},
		}

		expectedOutput := `--------------------------------------
	Execution Summary
--------------------------------------
Total Rules: 	10
Succeeded: 	8
Failed: 	2
--------------------------------------
`

		displaySummary(report, outputBuffer)

		assert.Equal(t, expectedOutput, outputBuffer.String())
	})

	t.Run("full report summary failing compliance and coverage", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		report := &model.Report{
			Summary: &model.Summary{
				Pass: true,
				Time: time.Now(),
			},
			Compliance: model.Compliance{
				Total:     10,
				Passed:    8,
				Failed:    2,
				Rate:      80,
				Threshold: values.GetIntRef(100),
				Pass:      false,
			},
			Coverage: model.Coverage{
				Rate:      85,
				Threshold: values.GetIntRef(90),
				Pass:      false,
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

		displaySummary(report, outputBuffer)

		assert.Equal(t, expectedOutput, outputBuffer.String())
	})

	t.Run("full report summary passing compliance and coverage", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		report := &model.Report{
			Summary: &model.Summary{
				Pass: true,
				Time: time.Now(),
			},
			Compliance: model.Compliance{
				Total:     10,
				Passed:    8,
				Failed:    2,
				Rate:      100,
				Threshold: values.GetIntRef(100),
				Pass:      true,
			},
			Coverage: model.Coverage{
				Rate:      90,
				Threshold: values.GetIntRef(90),
				Pass:      true,
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

		displaySummary(report, outputBuffer)

		assert.Equal(t, expectedOutput, outputBuffer.String())
	})

	t.Run("exports results", func(t *testing.T) {
		var htmlCalled, jsonCalled bool

		mock := gomonkey.ApplyFunc(
			os.WriteFile,
			func(_ string, _ []byte, _ os.FileMode) error {
				return nil
			})
		defer mock.Reset()

		htmlMock := gomonkey.ApplyFunc(
			html.GenerateHTMLReport,
			func(_ *model.Report, _ io.Writer) {
				htmlCalled = true
			})
		defer htmlMock.Reset()

		jsonMock := gomonkey.ApplyFunc(
			json.GenerateReport,
			func(_ *model.Report, _ io.Writer) {
				jsonCalled = true
			})
		defer jsonMock.Reset()

		common.HTML = true
		common.JSON = true
		outputBuffer := bytes.NewBufferString("")
		report := &model.Report{}

		expectedOutput := `--------------------------------------
	Execution Summary
--------------------------------------
Total Rules: 	0
Succeeded: 	0
Failed: 	0
--------------------------------------
`

		DisplayResult(report, outputBuffer)

		assert.True(t, htmlCalled)
		assert.True(t, jsonCalled)
		assert.Equal(t, expectedOutput, outputBuffer.String())
	})
}
