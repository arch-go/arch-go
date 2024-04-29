package archgo

import (
	"time"

	contents2 "github.com/fdaines/arch-go/internal/verifications/contents"
	dependencies2 "github.com/fdaines/arch-go/internal/verifications/dependencies"
	functions2 "github.com/fdaines/arch-go/internal/verifications/functions"
	naming2 "github.com/fdaines/arch-go/internal/verifications/naming"
	"github.com/fdaines/arch-go/pkg/archgo/configuration"

	"github.com/fdaines/arch-go/internal/model"
)

type architectureAnalysis struct {
	moduleInfo             model.ModuleInfo
	configuration          configuration.Config
	checkDependenciesRules func(model.ModuleInfo, []*configuration.DependenciesRule) *dependencies2.RulesResult
	checkFunctionRules     func(model.ModuleInfo, []*configuration.FunctionsRule) *functions2.RulesResult
	checkContentsRules     func(model.ModuleInfo, []*configuration.ContentsRule) *contents2.RulesResult
	checkNamingRules       func(model.ModuleInfo, []*configuration.NamingRule) *naming2.RulesResult
}

func NewArchitectureAnalysis(m model.ModuleInfo, c configuration.Config) *architectureAnalysis {
	return &architectureAnalysis{
		moduleInfo:             m,
		configuration:          c,
		checkDependenciesRules: dependencies2.CheckRules,
		checkFunctionRules:     functions2.CheckRules,
		checkContentsRules:     contents2.CheckRules,
		checkNamingRules:       naming2.CheckRules,
	}
}

func (a *architectureAnalysis) Execute() (*Result, error) {
	verificationResult := &Result{
		Time:   time.Now(),
		Passes: true,
	}
	if len(a.configuration.DependenciesRules) > 0 {
		verificationResult.DependenciesRuleResult = a.checkDependenciesRules(a.moduleInfo, a.configuration.DependenciesRules)
		verificationResult.Passes = verificationResult.Passes && verificationResult.DependenciesRuleResult.Passes
	}
	if len(a.configuration.FunctionsRules) > 0 {
		verificationResult.FunctionsRuleResult = a.checkFunctionRules(a.moduleInfo, a.configuration.FunctionsRules)
		verificationResult.Passes = verificationResult.Passes && verificationResult.FunctionsRuleResult.Passes
	}
	if len(a.configuration.ContentRules) > 0 {
		verificationResult.ContentsRuleResult = a.checkContentsRules(a.moduleInfo, a.configuration.ContentRules)
		verificationResult.Passes = verificationResult.Passes && verificationResult.ContentsRuleResult.Passes
	}
	if len(a.configuration.NamingRules) > 0 {
		verificationResult.NamingRuleResult = a.checkNamingRules(a.moduleInfo, a.configuration.NamingRules)
		verificationResult.Passes = verificationResult.Passes && verificationResult.NamingRuleResult.Passes
	}
	endMoment := time.Now()

	verificationResult.Duration = endMoment.Sub(verificationResult.Time)

	return verificationResult, nil
}

func (a *architectureAnalysis) withFunctionRulesVerification(
	f func(model.ModuleInfo, []*configuration.FunctionsRule) *functions2.RulesResult,
) *architectureAnalysis {
	a.checkFunctionRules = f
	return a
}

func (a *architectureAnalysis) withContentsRulesVerification(
	f func(model.ModuleInfo, []*configuration.ContentsRule) *contents2.RulesResult,
) *architectureAnalysis {
	a.checkContentsRules = f
	return a
}

func (a *architectureAnalysis) withNamingRulesVerification(
	f func(model.ModuleInfo, []*configuration.NamingRule) *naming2.RulesResult,
) *architectureAnalysis {
	a.checkNamingRules = f
	return a
}
