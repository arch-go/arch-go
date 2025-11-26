package reports

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/v2/api"
	"github.com/arch-go/arch-go/v2/api/configuration"
	"github.com/arch-go/arch-go/v2/internal/reports/model"
	"github.com/arch-go/arch-go/v2/internal/utils/values"
	"github.com/arch-go/arch-go/v2/internal/verifications/contents"
	"github.com/arch-go/arch-go/v2/internal/verifications/dependencies"
	"github.com/arch-go/arch-go/v2/internal/verifications/functions"
	"github.com/arch-go/arch-go/v2/internal/verifications/naming"
)

func TestComplianceResolver(t *testing.T) {
	t.Run("resolveTotals", func(t *testing.T) {
		passes, totals := resolveTotals(&api.Result{})
		assert.Equal(t, 0, passes)
		assert.Equal(t, 0, totals)

		passes, totals = resolveTotals(&api.Result{
			DependenciesRuleResult: &dependencies.RulesResult{
				Results: []*dependencies.RuleResult{
					{Passes: true},
					{Passes: false},
				},
			},
		})
		assert.Equal(t, 1, passes)
		assert.Equal(t, 2, totals)

		passes, totals = resolveTotals(&api.Result{
			FunctionsRuleResult: &functions.RulesResult{
				Results: []*functions.RuleResult{
					{Passes: true},
					{Passes: false},
					{Passes: true},
				},
			},
		})
		assert.Equal(t, 2, passes)
		assert.Equal(t, 3, totals)

		passes, totals = resolveTotals(&api.Result{
			ContentsRuleResult: &contents.RulesResult{
				Results: []*contents.RuleResult{
					{Passes: false},
				},
			},
		})
		assert.Equal(t, 0, passes)
		assert.Equal(t, 1, totals)

		passes, totals = resolveTotals(&api.Result{
			NamingRuleResult: &naming.RulesResult{
				Results: []*naming.RuleResult{
					{Passes: true},
					{Passes: true},
					{Passes: true},
				},
			},
		})
		assert.Equal(t, 3, passes)
		assert.Equal(t, 3, totals)

		passes, totals = resolveTotals(&api.Result{
			DependenciesRuleResult: &dependencies.RulesResult{
				Results: []*dependencies.RuleResult{
					{Passes: true},
					{Passes: true},
				},
			},
			FunctionsRuleResult: &functions.RulesResult{
				Results: []*functions.RuleResult{
					{Passes: false},
					{Passes: false},
					{Passes: false},
				},
			},
			ContentsRuleResult: &contents.RulesResult{
				Results: []*contents.RuleResult{
					{Passes: true},
				},
			},
			NamingRuleResult: &naming.RulesResult{
				Results: []*naming.RuleResult{
					{Passes: true},
					{Passes: true},
					{Passes: false},
				},
			},
		})
		assert.Equal(t, 5, passes)
		assert.Equal(t, 9, totals)
	})

	t.Run("resolveCompliance case 1", func(t *testing.T) {
		verificationResult := &api.Result{}
		conf := configuration.Config{}
		expectedResult := &model.ThresholdSummary{
			Threshold: values.GetIntRef(0),
			Pass:      true,
		}

		threshold := resolveCompliance(verificationResult, conf)

		assert.Equal(t, expectedResult, threshold)
	})

	t.Run("resolveCompliance case 2", func(t *testing.T) {
		verificationResult := &api.Result{}
		conf := configuration.Config{
			Threshold: &configuration.Threshold{
				Compliance: values.GetIntRef(100),
			},
		}
		expectedResult := &model.ThresholdSummary{
			Rate:       0,
			Threshold:  values.GetIntRef(100),
			Pass:       false,
			Violations: []string{""},
		}

		threshold := resolveCompliance(verificationResult, conf)

		assert.Equal(t, expectedResult, threshold)
	})

	t.Run("resolveCompliance case 3", func(t *testing.T) {
		verificationResult := &api.Result{
			DependenciesRuleResult: &dependencies.RulesResult{
				Results: []*dependencies.RuleResult{
					{Passes: true},
					{Passes: false},
				},
			},
		}
		conf := configuration.Config{
			Threshold: &configuration.Threshold{
				Compliance: values.GetIntRef(51),
			},
		}
		expectedResult := &model.ThresholdSummary{
			Rate:       50,
			Threshold:  values.GetIntRef(51),
			Pass:       false,
			Violations: []string{""},
		}

		threshold := resolveCompliance(verificationResult, conf)

		assert.Equal(t, expectedResult, threshold)
	})

	t.Run("resolveCompliance case 5", func(t *testing.T) {
		verificationResult := &api.Result{
			DependenciesRuleResult: &dependencies.RulesResult{
				Results: []*dependencies.RuleResult{
					{Passes: true},
					{Passes: false},
				},
			},
		}
		conf := configuration.Config{
			Threshold: &configuration.Threshold{
				Compliance: values.GetIntRef(50),
			},
		}
		expectedResult := &model.ThresholdSummary{
			Rate:       50,
			Threshold:  values.GetIntRef(50),
			Pass:       true,
			Violations: nil,
		}

		threshold := resolveCompliance(verificationResult, conf)

		assert.Equal(t, expectedResult, threshold)
	})
}
