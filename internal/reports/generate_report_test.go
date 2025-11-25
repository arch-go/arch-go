package reports

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/v2/api"
	"github.com/arch-go/arch-go/v2/api/configuration"
	"github.com/arch-go/arch-go/v2/internal/common"
	"github.com/arch-go/arch-go/v2/internal/model"
	reportModel "github.com/arch-go/arch-go/v2/internal/reports/model"
	"github.com/arch-go/arch-go/v2/internal/utils/values"
	"github.com/arch-go/arch-go/v2/internal/verifications/contents"
	"github.com/arch-go/arch-go/v2/internal/verifications/dependencies"
	"github.com/arch-go/arch-go/v2/internal/verifications/functions"
	"github.com/arch-go/arch-go/v2/internal/verifications/naming"
)

func TestGenerateReport(t *testing.T) {
	t.Run("Empty Result", func(t *testing.T) {
		apiResult := &api.Result{}
		module := model.ModuleInfo{}
		config := configuration.Config{}

		expectedResult := &reportModel.Report{
			ArchGoVersion: common.Version,
			Summary: &reportModel.Summary{
				Time:     time.Time{},
				Duration: time.Duration(0),
				Pass:     true,
			},
			Compliance: reportModel.Compliance{
				Total:     0,
				Passed:    0,
				Failed:    0,
				Rate:      0,
				Threshold: values.GetIntRef(0),
				Pass:      true,
				Details:   &reportModel.ReportDetails{},
			},
			Coverage: reportModel.Coverage{
				Rate:      0,
				Threshold: values.GetIntRef(0),
				Pass:      true,
			},
		}

		result := GenerateReport(apiResult, module, config)

		assert.Equal(t, expectedResult, result)
	})

	t.Run("Full Result", func(t *testing.T) {
		apiResult := &api.Result{
			Time:     time.Time{},
			Duration: time.Duration(123456789),
			Pass:     false,
			DependenciesRuleResult: &dependencies.RulesResult{
				Passes: true,
				Results: []*dependencies.RuleResult{
					{
						Rule:        configuration.DependenciesRule{},
						Description: "foobar rule dep",
						Passes:      true,
						Verifications: []dependencies.Verification{
							{
								Package: "my-package/pkg1",
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
								Package: "my-package/pkg1",
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
								Package: "my-package/pkg1",
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
								Package: "my-package/pkg1",
								Passes:  false,
								Details: []string{"foobar message"},
							},
						},
					},
				},
			},
		}
		module := model.ModuleInfo{
			MainPackage: "foobar",
			Packages: []*model.PackageInfo{
				{
					Name: "pkg1",
					Path: "foobar/pkg1",
				},
				{
					Name: "pkgX",
					Path: "my-package/pkg1",
				},
			},
		}
		config := configuration.Config{}

		expectedResult := &reportModel.Report{
			ArchGoVersion: common.Version,
			Summary: &reportModel.Summary{
				Time:     time.Time{},
				Duration: time.Duration(123456789),
				Pass:     true,
			},
			Compliance: reportModel.Compliance{
				Total:     4,
				Passed:    3,
				Failed:    1,
				Rate:      75,
				Threshold: values.GetIntRef(0),
				Pass:      true,
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
								Pass:   true,
								PackageDetails: []reportModel.PackageDetails{
									{
										Package: "my-package/pkg1",
										Pass:    true,
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
								Pass:   true,
								PackageDetails: []reportModel.PackageDetails{
									{
										Package: "my-package/pkg1",
										Pass:    true,
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
								Pass:   true,
								PackageDetails: []reportModel.PackageDetails{
									{
										Package: "my-package/pkg1",
										Pass:    true,
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
								Pass:   false,
								PackageDetails: []reportModel.PackageDetails{
									{
										Package: "my-package/pkg1",
										Pass:    false,
										Details: []string{"foobar message"},
									},
								},
							},
						},
					},
				},
			},
			Coverage: reportModel.Coverage{
				Rate:      50,
				Threshold: values.GetIntRef(0),
				Pass:      true,
				Uncovered: []string{"foobar/pkg1"},
				Details: []reportModel.CoverageDetails{
					{
						Package:           "foobar/pkg1",
						ContentsRules:     0,
						DependenciesRules: 0,
						FunctionsRules:    0,
						NamingRules:       0,
						Covered:           false,
					},
					{
						Package:           "my-package/pkg1",
						ContentsRules:     1,
						DependenciesRules: 1,
						FunctionsRules:    1,
						NamingRules:       1,
						Covered:           true,
					},
				},
			},
		}

		result := GenerateReport(apiResult, module, config)
		resultJSONBytes, _ := json.Marshal(result)
		expectedResultJSONBytes, _ := json.Marshal(expectedResult)

		assert.Equal(t, string(expectedResultJSONBytes), string(resultJSONBytes))
	})
}
