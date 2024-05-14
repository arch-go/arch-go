package reports

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/fdaines/arch-go/api"
	"github.com/fdaines/arch-go/api/configuration"
	"github.com/fdaines/arch-go/internal/verifications/contents"
	"github.com/fdaines/arch-go/internal/verifications/dependencies"
	"github.com/fdaines/arch-go/internal/verifications/functions"
	"github.com/fdaines/arch-go/internal/verifications/naming"

	"github.com/fdaines/arch-go/internal/model"
	reportModel "github.com/fdaines/arch-go/internal/reports/model"

	"github.com/stretchr/testify/assert"
)

func TestGenerateReport(t *testing.T) {
	t.Run("Empty Result", func(t *testing.T) {
		apiResult := &api.Result{}
		module := model.ModuleInfo{}
		config := configuration.Config{}

		expectedResult := &reportModel.Report{
			ArchGoVersion: "1.5.0",
			Summary: &reportModel.ReportSummary{
				Time:     time.Time{},
				Duration: time.Duration(0),
				Status:   "PASS",
				Total:    0,
				Passed:   0,
				Failed:   0,
				ComplianceThreshold: &reportModel.ThresholdSummary{
					Rate:      0,
					Threshold: 0,
					Status:    "PASS",
				},
				CoverageThreshold: &reportModel.ThresholdSummary{
					Rate:      0,
					Threshold: 0,
					Status:    "PASS",
				},
			},
			Details: &reportModel.ReportDetails{},
		}

		result := GenerateReport(apiResult, module, config)

		assert.Equal(t, expectedResult, result)
	})

	t.Run("Full Result", func(t *testing.T) {
		apiResult := &api.Result{
			Time:     time.Time{},
			Duration: time.Duration(123456789),
			Passes:   false,
			DependenciesRuleResult: &dependencies.RulesResult{
				Passes: true,
				Results: []*dependencies.RuleResult{
					{
						Rule:        configuration.DependenciesRule{},
						Description: "foobar rule dep",
						Passes:      true,
						Verifications: []dependencies.Verification{
							{
								Package: "my-package",
								Passes:  true,
							},
						},
					},
				},
			},
			FunctionsRuleResult: &functions.RulesResult{
				Passes: true,
				Results: []*functions.RuleResult{
					{
						Rule:        configuration.FunctionsRule{},
						Description: "foobar rule fn",
						Passes:      true,
						Verifications: []functions.Verification{
							{
								Package: "my-package",
								Passes:  true,
							},
						},
					},
				},
			},
			ContentsRuleResult: &contents.RulesResult{
				Passes: true,
				Results: []*contents.RuleResult{
					{
						Rule:        configuration.ContentsRule{},
						Description: "foobar rule cn",
						Passes:      true,
						Verifications: []contents.Verification{
							{
								Package: "my-package",
								Passes:  true,
							},
						},
					},
				},
			},
			NamingRuleResult: &naming.RulesResult{
				Passes: false,
				Results: []*naming.RuleResult{
					{
						Rule:        configuration.NamingRule{},
						Description: "foobar rule nm",
						Passes:      false,
						Verifications: []naming.Verification{
							{
								Package: "my-package",
								Passes:  false,
								Details: []string{"foobar message"},
							},
						},
					},
				},
			},
		}
		module := model.ModuleInfo{}
		config := configuration.Config{}

		expectedResult := &reportModel.Report{
			ArchGoVersion: "1.5.0",
			Summary: &reportModel.ReportSummary{
				Time:     time.Time{},
				Duration: time.Duration(123456789),
				Status:   "PASS",
				Total:    4,
				Passed:   3,
				Failed:   1,
				ComplianceThreshold: &reportModel.ThresholdSummary{
					Rate:      75,
					Threshold: 0,
					Status:    "PASS",
				},
				CoverageThreshold: &reportModel.ThresholdSummary{
					Rate:      0,
					Threshold: 0,
					Status:    "PASS",
				},
			},
			Details: &reportModel.ReportDetails{
				DependenciesVerificationDetails: reportModel.Verification{
					Total:  1,
					Passed: 1,
					Failed: 0,
					Details: []reportModel.VerificationDetails{
						{
							Rule:   "foobar rule dep",
							Total:  1,
							Passed: 1,
							Failed: 0,
							Status: "PASS",
							PackageDetails: []reportModel.PackageDetails{
								{
									Package: "my-package",
									Status:  "PASS",
								},
							},
						},
					},
				},
				FunctionsVerificationDetails: reportModel.Verification{
					Total:  1,
					Passed: 1,
					Failed: 0,
					Details: []reportModel.VerificationDetails{
						{
							Rule:   "foobar rule fn",
							Total:  1,
							Passed: 1,
							Failed: 0,
							Status: "PASS",
							PackageDetails: []reportModel.PackageDetails{
								{
									Package: "my-package",
									Status:  "PASS",
								},
							},
						},
					},
				},
				ContentsVerificationDetails: reportModel.Verification{
					Total:  1,
					Passed: 1,
					Failed: 0,
					Details: []reportModel.VerificationDetails{
						{
							Rule:   "foobar rule cn",
							Total:  1,
							Passed: 1,
							Failed: 0,
							Status: "PASS",
							PackageDetails: []reportModel.PackageDetails{
								{
									Package: "my-package",
									Status:  "PASS",
								},
							},
						},
					},
				},
				NamingVerificationDetails: reportModel.Verification{
					Total:  1,
					Passed: 0,
					Failed: 1,
					Details: []reportModel.VerificationDetails{
						{
							Rule:   "foobar rule nm",
							Total:  1,
							Passed: 0,
							Failed: 1,
							Status: "FAIL",
							PackageDetails: []reportModel.PackageDetails{
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
		}

		result := GenerateReport(apiResult, module, config)

		resultJsonBytes, _ := json.Marshal(result)
		expectedResultJsonBytes, _ := json.Marshal(expectedResult)

		assert.Equal(t, string(expectedResultJsonBytes), string(resultJsonBytes))
	})
}
