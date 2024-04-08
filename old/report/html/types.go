package html

type HtmlReport struct {
	Version           string
	ComplianceResult  ComplianceResult
	CoverageResult    CoverageResult
	UncoveredPackages []string
	RulesSummary      []RuleSummary
	RulesDetails      []RuleDetails
}

type RuleSummary struct {
	Type      string
	Succeeded int
	Failed    int
	Total     int
	Ratio     int
}

type RuleDetails struct {
	Type          string
	Description   string
	Status        string
	Color         string
	Verifications []RuleVerification
}

type RuleVerification struct {
	Package string
	Details []string
	Status  string
	Color   string
}

type ComplianceResult struct {
	Rate      int
	Threshold int
	Succeeded int
	Total     int
	Color     string
}

type CoverageResult struct {
	Rate      int
	Threshold int
	Uncovered int
	Total     int
	Color     string
}
