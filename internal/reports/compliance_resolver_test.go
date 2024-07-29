package reports

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fdaines/arch-go/api"
	"github.com/fdaines/arch-go/api/configuration"
	"github.com/fdaines/arch-go/internal/reports/model"
	"github.com/fdaines/arch-go/internal/utils/values"
	"github.com/fdaines/arch-go/internal/verifications/contents"
	"github.com/fdaines/arch-go/internal/verifications/dependencies"
	"github.com/fdaines/arch-go/internal/verifications/functions"
	"github.com/fdaines/arch-go/internal/verifications/naming"
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
			Status: "PASS",
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
			Threshold:  100,
			Status:     "FAIL",
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
			Threshold:  51,
			Status:     "FAIL",
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
			Threshold:  50,
			Status:     "PASS",
			Violations: nil,
		}

		threshold := resolveCompliance(verificationResult, conf)

		assert.Equal(t, expectedResult, threshold)
	})
}
