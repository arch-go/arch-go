package reports

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/internal/reports/model"
)

func TestDisplayRules(t *testing.T) {
	t.Run("display rules for report details", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		report := &model.Report{
			ArchGoVersion: "vfoobar",
			Summary:       &model.ReportSummary{},
			Details: &model.ReportDetails{
				DependenciesVerificationDetails: model.Verification{
					Total:  1,
					Passed: 1,
					Failed: 0,
					Details: []model.VerificationDetails{
						{
							Rule:   "foobar rule dep",
							Status: "PASS",
							Total:  1,
							Passed: 1,
							Failed: 0,
							PackageDetails: []model.PackageDetails{
								{
									Package: "my-package",
									Status:  "PASS",
								},
							},
						},
					},
				},
				FunctionsVerificationDetails: model.Verification{
					Total:  1,
					Passed: 1,
					Failed: 0,
					Details: []model.VerificationDetails{
						{
							Rule:   "foobar rule fn",
							Status: "PASS",
							Total:  1,
							Passed: 1,
							Failed: 0,
							PackageDetails: []model.PackageDetails{
								{
									Package: "my-package",
									Status:  "PASS",
								},
							},
						},
					},
				},
				ContentsVerificationDetails: model.Verification{
					Total:  1,
					Passed: 1,
					Failed: 0,
					Details: []model.VerificationDetails{
						{
							Rule:   "foobar rule cn",
							Status: "PASS",
							Total:  1,
							Passed: 1,
							Failed: 0,
							PackageDetails: []model.PackageDetails{
								{
									Package: "my-package",
									Status:  "PASS",
								},
							},
						},
					},
				},
				NamingVerificationDetails: model.Verification{
					Total:  1,
					Passed: 0,
					Failed: 1,
					Details: []model.VerificationDetails{
						{
							Rule:   "foobar rule nm",
							Status: "FAIL",
							Total:  1,
							Passed: 0,
							Failed: 1,
							PackageDetails: []model.PackageDetails{
								{
									Package: "my-package",
									Status:  "FAIL",
									Details: []string{"foobar message"},
								},
							},
						},
					},
				},
			},
			CoverageInfo: []model.CoverageInfo{},
		}
		expectedOutput := `[PASS] - foobar rule cn
	Package 'my-package' passes
[PASS] - foobar rule dep
	Package 'my-package' passes
[PASS] - foobar rule fn
	Package 'my-package' passes
[FAIL] - foobar rule nm
	Package 'my-package' fails
		foobar message
`

		displayRules(report, outputBuffer)

		assert.Equal(t, expectedOutput, outputBuffer.String())
	})

	t.Run("empty report summary", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		report := &model.Report{
			ArchGoVersion: "vfoobar",
			Summary:       &model.ReportSummary{},
			Details:       &model.ReportDetails{},
			CoverageInfo:  []model.CoverageInfo{},
		}
		expectedOutput := ``

		displayRules(report, outputBuffer)

		assert.Equal(t, expectedOutput, outputBuffer.String())
	})
}
