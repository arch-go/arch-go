package api

import (
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/fdaines/arch-go/api/configuration"
	"github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/internal/utils/values"
	"github.com/fdaines/arch-go/internal/verifications/contents"
	"github.com/fdaines/arch-go/internal/verifications/functions"
	"github.com/fdaines/arch-go/internal/verifications/naming"
)

func TestArchitecture(t *testing.T) {
	mockTimeNow := gomonkey.ApplyFuncReturn(time.Now, time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC))

	defer mockTimeNow.Reset()

	moduleInfo := model.ModuleInfo{
		MainPackage: "mymodule",
		Packages: []*model.PackageInfo{
			{
				Name: "foobar1",
				Path: "barfoo1",
			},
			{
				Name: "foobar2",
				Path: "barfoo2",
			},
		},
	}
	configurationRules := configuration.Config{
		FunctionsRules: []*configuration.FunctionsRule{
			{
				Package:  "**.qwerty.**",
				MaxLines: values.GetIntRef(123),
			},
		},
		ContentRules: []*configuration.ContentsRule{
			{
				Package:                    "**.blablabla.**",
				ShouldNotContainInterfaces: true,
			},
		},
		NamingRules: []*configuration.NamingRule{
			{
				Package: "**.foobar.**",
				InterfaceImplementationNamingRule: &configuration.InterfaceImplementationRule{
					StructsThatImplement:           "*Command",
					ShouldHaveSimpleNameEndingWith: values.GetStringRef("Foobar"),
				},
			},
		},
	}

	t.Run("all verification passes", func(t *testing.T) {
		expectedFunctionsResult := &functions.RulesResult{Passes: true}
		expectedContentsResult := &contents.RulesResult{Passes: true}
		expectedNamingResult := &naming.RulesResult{Passes: true}

		expectedResult := &Result{
			Time:                time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Duration:            0,
			Passes:              true,
			FunctionsRuleResult: expectedFunctionsResult,
			ContentsRuleResult:  expectedContentsResult,
			NamingRuleResult:    expectedNamingResult,
		}
		functionRulesVerification := func(_ model.ModuleInfo, _ []*configuration.FunctionsRule) *functions.RulesResult {
			return expectedFunctionsResult
		}
		contentsRulesVerification := func(_ model.ModuleInfo, _ []*configuration.ContentsRule) *contents.RulesResult {
			return expectedContentsResult
		}
		namingRulesVerification := func(_ model.ModuleInfo, _ []*configuration.NamingRule) *naming.RulesResult {
			return expectedNamingResult
		}

		architectureAnalysis := newArchitectureAnalysis(moduleInfo, configurationRules).
			withFunctionRulesVerification(functionRulesVerification).
			withContentsRulesVerification(contentsRulesVerification).
			withNamingRulesVerification(namingRulesVerification)

		result, err := architectureAnalysis.Execute()

		require.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("function verification fails", func(t *testing.T) {
		expectedFunctionsResult := &functions.RulesResult{Passes: false}
		expectedContentsResult := &contents.RulesResult{Passes: true}
		expectedNamingResult := &naming.RulesResult{Passes: true}

		expectedResult := &Result{
			Time:                time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Duration:            0,
			Passes:              false,
			FunctionsRuleResult: expectedFunctionsResult,
			ContentsRuleResult:  expectedContentsResult,
			NamingRuleResult:    expectedNamingResult,
		}
		functionRulesVerification := func(_ model.ModuleInfo, _ []*configuration.FunctionsRule) *functions.RulesResult {
			return expectedFunctionsResult
		}
		contentsRulesVerification := func(_ model.ModuleInfo, _ []*configuration.ContentsRule) *contents.RulesResult {
			return expectedContentsResult
		}
		namingRulesVerification := func(_ model.ModuleInfo, _ []*configuration.NamingRule) *naming.RulesResult {
			return expectedNamingResult
		}

		architectureAnalysis := newArchitectureAnalysis(moduleInfo, configurationRules).
			withFunctionRulesVerification(functionRulesVerification).
			withContentsRulesVerification(contentsRulesVerification).
			withNamingRulesVerification(namingRulesVerification)

		result, err := architectureAnalysis.Execute()

		require.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("contents verification fails", func(t *testing.T) {
		expectedFunctionsResult := &functions.RulesResult{Passes: true}
		expectedContentsResult := &contents.RulesResult{Passes: false}
		expectedNamingResult := &naming.RulesResult{Passes: true}

		expectedResult := &Result{
			Time:                time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Duration:            0,
			Passes:              false,
			FunctionsRuleResult: expectedFunctionsResult,
			ContentsRuleResult:  expectedContentsResult,
			NamingRuleResult:    expectedNamingResult,
		}
		functionRulesVerification := func(_ model.ModuleInfo, _ []*configuration.FunctionsRule) *functions.RulesResult {
			return expectedFunctionsResult
		}
		contentsRulesVerification := func(_ model.ModuleInfo, _ []*configuration.ContentsRule) *contents.RulesResult {
			return expectedContentsResult
		}
		namingRulesVerification := func(_ model.ModuleInfo, _ []*configuration.NamingRule) *naming.RulesResult {
			return expectedNamingResult
		}

		architectureAnalysis := newArchitectureAnalysis(moduleInfo, configurationRules).
			withFunctionRulesVerification(functionRulesVerification).
			withContentsRulesVerification(contentsRulesVerification).
			withNamingRulesVerification(namingRulesVerification)

		result, err := architectureAnalysis.Execute()

		require.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("naming verification fails", func(t *testing.T) {
		expectedFunctionsResult := &functions.RulesResult{Passes: true}
		expectedContentsResult := &contents.RulesResult{Passes: true}
		expectedNamingResult := &naming.RulesResult{Passes: false}

		expectedResult := &Result{
			Time:                time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Duration:            0,
			Passes:              false,
			FunctionsRuleResult: expectedFunctionsResult,
			ContentsRuleResult:  expectedContentsResult,
			NamingRuleResult:    expectedNamingResult,
		}
		functionRulesVerification := func(_ model.ModuleInfo, _ []*configuration.FunctionsRule) *functions.RulesResult {
			return expectedFunctionsResult
		}
		contentsRulesVerification := func(_ model.ModuleInfo, _ []*configuration.ContentsRule) *contents.RulesResult {
			return expectedContentsResult
		}
		namingRulesVerification := func(_ model.ModuleInfo, _ []*configuration.NamingRule) *naming.RulesResult {
			return expectedNamingResult
		}

		architectureAnalysis := newArchitectureAnalysis(moduleInfo, configurationRules).
			withFunctionRulesVerification(functionRulesVerification).
			withContentsRulesVerification(contentsRulesVerification).
			withNamingRulesVerification(namingRulesVerification)

		result, err := architectureAnalysis.Execute()

		require.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("passes when there are no packages and no rules", func(t *testing.T) {
		moduleInfo := model.ModuleInfo{
			MainPackage: "mymodule",
			Packages:    []*model.PackageInfo{},
		}
		localConfiguration := configuration.Config{}

		expectedResult := &Result{
			Time:     time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Duration: 0,
			Passes:   true,
		}

		architectureAnalysis := newArchitectureAnalysis(moduleInfo, localConfiguration)

		result, err := architectureAnalysis.Execute()

		require.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("passes when there are no packages", func(t *testing.T) {
		moduleInfo := model.ModuleInfo{
			MainPackage: "mymodule",
			Packages:    []*model.PackageInfo{},
		}
		expectedFunctionsResult := &functions.RulesResult{
			Passes: true,
			Results: []*functions.RuleResult{
				{
					Rule:        *configurationRules.FunctionsRules[0],
					Description: "Functions in packages matching pattern '**.qwerty.**' should have ['at most 123 lines']",
					Passes:      true,
				},
			},
		}
		expectedContentsResult := &contents.RulesResult{
			Passes: true,
			Results: []*contents.RuleResult{
				{
					Rule:        *configurationRules.ContentRules[0],
					Description: "Packages matching pattern '**.blablabla.**' should complies with ['should not contain interfaces']",
					Passes:      true,
				},
			},
		}
		expectedNamingResult := &naming.RulesResult{
			Passes: true,
			Results: []*naming.RuleResult{
				{
					Rule:        *configurationRules.NamingRules[0],
					Description: "Packages matching pattern '**.foobar.**' should comply with [structs that implement '*Command' should have simple name ending with 'Foobar']",
					Passes:      true,
				},
			},
		}

		expectedResult := &Result{
			Time:                time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Duration:            0,
			Passes:              true,
			FunctionsRuleResult: expectedFunctionsResult,
			ContentsRuleResult:  expectedContentsResult,
			NamingRuleResult:    expectedNamingResult,
		}

		architectureAnalysis := newArchitectureAnalysis(moduleInfo, configurationRules)

		result, err := architectureAnalysis.Execute()

		require.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("passes when there are no rules", func(t *testing.T) {
		localConfiguration := configuration.Config{}

		expectedResult := &Result{
			Time:     time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Duration: 0,
			Passes:   true,
		}

		architectureAnalysis := newArchitectureAnalysis(moduleInfo, localConfiguration)

		result, err := architectureAnalysis.Execute()

		require.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})
}
