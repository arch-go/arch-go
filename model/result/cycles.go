package result

type CyclesRuleResultDetail struct {
	Package string
	Details []string
}

type CyclesRuleResult struct {
	Description string
	Passes      bool
	Failures    []*CyclesRuleResultDetail
}
