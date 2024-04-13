package verifications

import (
	"time"

	"github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/pkg/config"
	"github.com/fdaines/arch-go/pkg/verifications/contents"
	"github.com/fdaines/arch-go/pkg/verifications/functions"
	"github.com/fdaines/arch-go/pkg/verifications/naming"
)

type architectureAnalysis struct {
	moduleInfo         model.ModuleInfo
	configuration      config.Config
	checkFunctionRules func(model.ModuleInfo, []*config.FunctionsRule) *functions.RulesResult
	checkContentsRules func(model.ModuleInfo, []*config.ContentsRule) *contents.RulesResult
	checkNamingRules   func(model.ModuleInfo, []*config.NamingRule) *naming.RulesResult
}

func NewArchitectureAnalysis(m model.ModuleInfo, c config.Config) *architectureAnalysis {
	return &architectureAnalysis{
		moduleInfo:         m,
		configuration:      c,
		checkFunctionRules: functions.CheckRules,
		checkContentsRules: contents.CheckRules,
		checkNamingRules:   naming.CheckRules,
	}
}

func (a *architectureAnalysis) Run() (*Result, error) {
	verificationResult := &Result{
		Time:   time.Now(),
		Passes: true,
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
	f func(model.ModuleInfo, []*config.FunctionsRule) *functions.RulesResult,
) *architectureAnalysis {
	a.checkFunctionRules = f
	return a
}

func (a *architectureAnalysis) withContentsRulesVerification(
	f func(model.ModuleInfo, []*config.ContentsRule) *contents.RulesResult,
) *architectureAnalysis {
	a.checkContentsRules = f
	return a
}

func (a *architectureAnalysis) withNamingRulesVerification(
	f func(model.ModuleInfo, []*config.NamingRule) *naming.RulesResult,
) *architectureAnalysis {
	a.checkNamingRules = f
	return a
}
