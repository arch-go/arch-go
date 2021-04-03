package result

type DependenciesRuleFailureDetail struct {
	Package string
	Details []string
}

type DependenciesRuleResult struct {
	Description string
	Passes      bool
	Failures    []*DependenciesRuleFailureDetail
}
