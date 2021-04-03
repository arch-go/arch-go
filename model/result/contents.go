package result

type ContentsRuleFailureDetail struct {
	Package string
	Details []string
}

type ContentsRuleResult struct {
	Description string
	Passes      bool
	Failures    []*ContentsRuleFailureDetail
}
