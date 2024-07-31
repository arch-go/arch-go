package reports

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/api"
	"github.com/arch-go/arch-go/api/configuration"
	"github.com/arch-go/arch-go/internal/model"
	model2 "github.com/arch-go/arch-go/internal/reports/model"
	"github.com/arch-go/arch-go/internal/utils/values"
	"github.com/arch-go/arch-go/internal/verifications/contents"
	"github.com/arch-go/arch-go/internal/verifications/dependencies"
	"github.com/arch-go/arch-go/internal/verifications/functions"
	"github.com/arch-go/arch-go/internal/verifications/naming"
)

func TestCoverageResolver(t *testing.T) {
	moduleInfo := model.ModuleInfo{
		Packages: []*model.PackageInfo{
			{Path: "foo/bar1"},
			{Path: "foo/bar2"},
			{Path: "foo/bar3"},
			{Path: "foo/bar4"},
			{Path: "foo/bar5"},
		},
	}
	rulesConfiguration := configuration.Config{}

	t.Run("checkPackagesCoverage case 1", func(t *testing.T) {
		verificationResult := &api.Result{}

		expectedResult := map[string]bool{
			"foo/bar1": false,
			"foo/bar2": false,
			"foo/bar3": false,
			"foo/bar4": false,
			"foo/bar5": false,
		}

		result := checkPackagesCoverage(verificationResult, moduleInfo)

		assert.Equal(t, expectedResult, result)
	})

	t.Run("checkPackagesCoverage case 2", func(t *testing.T) {
		verificationResult := &api.Result{
			DependenciesRuleResult: &dependencies.RulesResult{
				Results: []*dependencies.RuleResult{
					{
						Verifications: []dependencies.Verification{
							{Package: "foo/bar5"},
							{Package: "foo/bar2"},
							{Package: "other/package"},
						},
					},
					{Passes: false},
				},
			},
		}

		expectedResult := map[string]bool{
			"foo/bar1": false,
			"foo/bar2": true,
			"foo/bar3": false,
			"foo/bar4": false,
			"foo/bar5": true,
		}

		result := checkPackagesCoverage(verificationResult, moduleInfo)

		assert.Equal(t, expectedResult, result)
	})

	t.Run("checkPackagesCoverage case 3", func(t *testing.T) {
		verificationResult := &api.Result{
			FunctionsRuleResult: &functions.RulesResult{
				Results: []*functions.RuleResult{
					{
						Verifications: []functions.Verification{
							{Package: "foo/bar5"},
							{Package: "foo/bar2"},
							{Package: "other/package"},
						},
					},
					{Passes: false},
				},
			},
		}

		expectedResult := map[string]bool{
			"foo/bar1": false,
			"foo/bar2": true,
			"foo/bar3": false,
			"foo/bar4": false,
			"foo/bar5": true,
		}

		result := checkPackagesCoverage(verificationResult, moduleInfo)

		assert.Equal(t, expectedResult, result)
	})

	t.Run("checkPackagesCoverage case 4", func(t *testing.T) {
		verificationResult := &api.Result{
			ContentsRuleResult: &contents.RulesResult{
				Results: []*contents.RuleResult{
					{
						Verifications: []contents.Verification{
							{Package: "foo/bar5"},
							{Package: "foo/bar2"},
							{Package: "other/package"},
						},
					},
					{Passes: false},
				},
			},
		}

		expectedResult := map[string]bool{
			"foo/bar1": false,
			"foo/bar2": true,
			"foo/bar3": false,
			"foo/bar4": false,
			"foo/bar5": true,
		}

		result := checkPackagesCoverage(verificationResult, moduleInfo)

		assert.Equal(t, expectedResult, result)
	})

	t.Run("checkPackagesCoverage case 5", func(t *testing.T) {
		verificationResult := &api.Result{
			NamingRuleResult: &naming.RulesResult{
				Results: []*naming.RuleResult{
					{
						Verifications: []naming.Verification{
							{Package: "foo/bar5"},
							{Package: "foo/bar2"},
							{Package: "other/package"},
						},
					},
					{Passes: false},
				},
			},
		}

		expectedResult := map[string]bool{
			"foo/bar1": false,
			"foo/bar2": true,
			"foo/bar3": false,
			"foo/bar4": false,
			"foo/bar5": true,
		}

		result := checkPackagesCoverage(verificationResult, moduleInfo)

		assert.Equal(t, expectedResult, result)
	})

	t.Run("checkPackagesCoverage case 6", func(t *testing.T) {
		verificationResult := &api.Result{
			DependenciesRuleResult: &dependencies.RulesResult{
				Results: []*dependencies.RuleResult{
					{
						Verifications: []dependencies.Verification{
							{Package: "foo/bar5"},
							{Package: "foo/bar2"},
							{Package: "other/package"},
						},
					},
					{Passes: false},
				},
			},
			FunctionsRuleResult: &functions.RulesResult{
				Results: []*functions.RuleResult{
					{
						Verifications: []functions.Verification{
							{Package: "foo/bar4"},
							{Package: "foo/bar2"},
							{Package: "other/package"},
						},
					},
					{Passes: false},
				},
			},
			ContentsRuleResult: &contents.RulesResult{
				Results: []*contents.RuleResult{
					{
						Verifications: []contents.Verification{
							{Package: "foo/bar1"},
							{Package: "foo/bar2"},
							{Package: "other/package"},
						},
					},
					{Passes: false},
				},
			},
			NamingRuleResult: &naming.RulesResult{
				Results: []*naming.RuleResult{
					{
						Verifications: []naming.Verification{
							{Package: "foo/bar5"},
							{Package: "foo/bar3"},
							{Package: "other/packageX"},
						},
					},
					{Passes: false},
				},
			},
		}

		expectedResult := map[string]bool{
			"foo/bar1": true,
			"foo/bar2": true,
			"foo/bar3": true,
			"foo/bar4": true,
			"foo/bar5": true,
		}

		result := checkPackagesCoverage(verificationResult, moduleInfo)

		assert.Equal(t, expectedResult, result)
	})

	t.Run("resolveCoverage case 1", func(t *testing.T) {
		verificationResult := &api.Result{}

		expectedViolations := []string{"foo/bar1", "foo/bar2", "foo/bar3", "foo/bar4", "foo/bar5"}

		result := resolveCoverage(verificationResult, moduleInfo, rulesConfiguration)

		assert.True(t, result.Pass)
		assert.ElementsMatch(t, expectedViolations, result.Violations)
	})

	t.Run("resolveCoverage case 2", func(t *testing.T) {
		verificationResult := &api.Result{}
		config := configuration.Config{
			Threshold: &configuration.Threshold{},
		}

		expectedResult := &model2.ThresholdSummary{
			Rate:       0,
			Threshold:  0,
			Pass:       true,
			Violations: []string{"foo/bar1", "foo/bar2", "foo/bar3", "foo/bar4", "foo/bar5"},
		}

		result := resolveCoverage(verificationResult, moduleInfo, config)

		assert.Equal(t, expectedResult.Rate, result.Rate)
		assert.Equal(t, expectedResult.Threshold, result.Threshold)
		assert.Equal(t, expectedResult.Pass, result.Pass)
		assert.ElementsMatch(t, expectedResult.Violations, result.Violations)
	})

	t.Run("resolveCoverage case 3", func(t *testing.T) {
		verificationResult := &api.Result{}
		config := configuration.Config{
			Threshold: &configuration.Threshold{
				Coverage: values.GetIntRef(100),
			},
		}

		expectedResult := &model2.ThresholdSummary{
			Rate:       0,
			Threshold:  100,
			Pass:       false,
			Violations: []string{"foo/bar1", "foo/bar2", "foo/bar3", "foo/bar4", "foo/bar5"},
		}

		result := resolveCoverage(verificationResult, moduleInfo, config)

		assert.Equal(t, expectedResult.Rate, result.Rate)
		assert.Equal(t, expectedResult.Threshold, result.Threshold)
		assert.Equal(t, expectedResult.Pass, result.Pass)
		assert.ElementsMatch(t, expectedResult.Violations, result.Violations)
	})

	t.Run("resolveCoverage case 4", func(t *testing.T) {
		verificationResult := &api.Result{
			DependenciesRuleResult: &dependencies.RulesResult{
				Results: []*dependencies.RuleResult{
					{
						Verifications: []dependencies.Verification{
							{Package: "foo/bar1"},
							{Package: "foo/bar2"},
							{Package: "foo/bar3"},
							{Package: "foo/bar4"},
							{Package: "foo/bar5"},
						},
					},
				},
			},
		}
		config := configuration.Config{
			Threshold: &configuration.Threshold{
				Coverage: values.GetIntRef(100),
			},
		}

		expectedResult := &model2.ThresholdSummary{
			Rate:       100,
			Threshold:  100,
			Pass:       true,
			Violations: nil,
		}

		result := resolveCoverage(verificationResult, moduleInfo, config)

		assert.Equal(t, expectedResult.Rate, result.Rate)
		assert.Equal(t, expectedResult.Threshold, result.Threshold)
		assert.Equal(t, expectedResult.Pass, result.Pass)
		assert.ElementsMatch(t, expectedResult.Violations, result.Violations)
	})

	t.Run("resolveCoverage case 5", func(t *testing.T) {
		verificationResult := &api.Result{
			DependenciesRuleResult: &dependencies.RulesResult{
				Results: []*dependencies.RuleResult{
					{
						Verifications: []dependencies.Verification{
							{Package: "foo/bar1"},
							{Package: "foo/bar2"},
							{Package: "foo/bar4"},
							{Package: "foo/bar5"},
						},
					},
				},
			},
		}
		config := configuration.Config{
			Threshold: &configuration.Threshold{
				Coverage: values.GetIntRef(78),
			},
		}

		expectedResult := &model2.ThresholdSummary{
			Rate:       80,
			Threshold:  78,
			Pass:       true,
			Violations: []string{"foo/bar3"},
		}

		result := resolveCoverage(verificationResult, moduleInfo, config)

		assert.Equal(t, expectedResult.Rate, result.Rate)
		assert.Equal(t, expectedResult.Threshold, result.Threshold)
		assert.Equal(t, expectedResult.Pass, result.Pass)
		assert.ElementsMatch(t, expectedResult.Violations, result.Violations)
	})
}
