package reports

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/v2/api"
	"github.com/arch-go/arch-go/v2/api/configuration"
	"github.com/arch-go/arch-go/v2/internal/reports/model"
	"github.com/arch-go/arch-go/v2/internal/verifications/contents"
	"github.com/arch-go/arch-go/v2/internal/verifications/dependencies"
	"github.com/arch-go/arch-go/v2/internal/verifications/functions"
	"github.com/arch-go/arch-go/v2/internal/verifications/naming"
)

func TestResolveReportDetails(t *testing.T) {
	t.Run("resolveReportDetails", func(t *testing.T) {
		result := &api.Result{
			Time:     time.Time{},
			Duration: time.Duration(123456789),
			Pass:     false,
			DependenciesRuleResult: &dependencies.RulesResult{
				Passes: true,
				Results: []*dependencies.RuleResult{
					{
						Rule:        configuration.DependenciesRule{},
						Description: "foobar rule dep",
						Passes:      false,
						Verifications: []dependencies.Verification{
							{
								Package: "my-package",
								Passes:  true,
							},
							{
								Package: "my-package-2",
								Passes:  false,
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
						Passes:      false,
						Verifications: []functions.Verification{
							{
								Package: "my-package",
								Passes:  true,
							},
							{
								Package: "my-package-2",
								Passes:  false,
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
						Passes:      false,
						Verifications: []contents.Verification{
							{
								Package: "my-package",
								Passes:  true,
							},
							{
								Package: "my-package-2",
								Passes:  false,
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
							{
								Package: "my-package-2",
								Passes:  true,
							},
						},
					},
				},
			},
		}

		expectedResult := &model.ReportDetails{
			DependenciesVerificationDetails: model.Verification{
				Total:  1,
				Passed: 0,
				Failed: 1,
				Details: []model.VerificationDetails{
					{
						Rule:   "foobar rule dep",
						Pass:   false,
						Total:  2,
						Passed: 1,
						Failed: 1,
						PackageDetails: []model.PackageDetails{
							{
								Package: "my-package",
								Pass:    true,
							},
							{
								Package: "my-package-2",
								Pass:    false,
							},
						},
					},
				},
			},
			FunctionsVerificationDetails: model.Verification{
				Total:  1,
				Passed: 0,
				Failed: 1,
				Details: []model.VerificationDetails{
					{
						Rule:   "foobar rule fn",
						Pass:   false,
						Total:  2,
						Passed: 1,
						Failed: 1,
						PackageDetails: []model.PackageDetails{
							{
								Package: "my-package",
								Pass:    true,
							},
							{
								Package: "my-package-2",
								Pass:    false,
							},
						},
					},
				},
			},
			ContentsVerificationDetails: model.Verification{
				Total:  1,
				Passed: 0,
				Failed: 1,
				Details: []model.VerificationDetails{
					{
						Rule:   "foobar rule cn",
						Pass:   false,
						Total:  2,
						Passed: 1,
						Failed: 1,
						PackageDetails: []model.PackageDetails{
							{
								Package: "my-package",
								Pass:    true,
							},
							{
								Package: "my-package-2",
								Pass:    false,
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
						Pass:   false,
						Total:  2,
						Passed: 1,
						Failed: 1,
						PackageDetails: []model.PackageDetails{
							{
								Package: "my-package",
								Pass:    false,
								Details: []string{"foobar message"},
							},
							{
								Package: "my-package-2",
								Pass:    true,
							},
						},
					},
				},
			},
		}

		details := resolveReportDetails(result)

		assert.Equal(t, expectedResult, details)
	})
}
