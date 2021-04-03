package result

type FunctionsRuleResultDetail struct {
	Package string
	File    string
	Name    string
	Details []string
}

type FunctionsRuleResult struct {
	Description string
	Passes      bool
	Failures    []*FunctionsRuleResultDetail
}
