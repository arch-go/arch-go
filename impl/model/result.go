package model

type Result struct {
	DependenciesRulesResults []*DependenciesRuleResult
	ContentsRuleResults      []*ContentsRuleResult
	CyclesRuleResults        []*CyclesRuleResult
	FunctionsRulesResults    []*FunctionsRuleResult
}
