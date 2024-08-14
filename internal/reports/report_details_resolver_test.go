package reports

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/api"
	"github.com/arch-go/arch-go/api/configuration"
	"github.com/arch-go/arch-go/internal/reports/model"
	"github.com/arch-go/arch-go/internal/verifications/contents"
	"github.com/arch-go/arch-go/internal/verifications/dependencies"
	"github.com/arch-go/arch-go/internal/verifications/functions"
	"github.com/arch-go/arch-go/internal/verifications/naming"
)

func TestResolveReportDetails(t *testing.T) {
	t.Run("resolveReportDetails", func(t *testing.T) {
		result := &api.Result{
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

		expectedResult := &model.ReportDetails{
			DependenciesVerificationDetails: model.Verification{
				Total:  1,
				Passed: 1,
				Failed: 0,
				Details: []model.VerificationDetails{
					{
						Rule:   "foobar rule dep",
						Pass:   true,
						Total:  1,
						Passed: 1,
						Failed: 0,
						PackageDetails: []model.PackageDetails{
							{
								Package: "my-package",
								Pass:    true,
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
						Pass:   true,
						Total:  1,
						Passed: 1,
						Failed: 0,
						PackageDetails: []model.PackageDetails{
							{
								Package: "my-package",
								Pass:    true,
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
						Pass:   true,
						Total:  1,
						Passed: 1,
						Failed: 0,
						PackageDetails: []model.PackageDetails{
							{
								Package: "my-package",
								Pass:    true,
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
						Total:  1,
						Passed: 0,
						Failed: 1,
						PackageDetails: []model.PackageDetails{
							{
								Package: "my-package",
								Pass:    false,
								Details: []string{"foobar message"},
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
