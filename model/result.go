package model

type DependenciesRuleFailureDetail struct {
	Package string
	Details []string
}

type DependenciesRuleResult struct {
	Description string
	Passes bool
	Failures []*DependenciesRuleFailureDetail
}

type Result struct {
	DependenciesRulesResults []*DependenciesRuleResult
}
