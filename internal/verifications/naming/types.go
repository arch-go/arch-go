package naming

import (
	"github.com/fdaines/arch-go/api/configuration"
)

type RulesResult struct {
	Results []*RuleResult `json:"results"`
	Passes  bool          `json:"passes"`
}

type RuleResult struct {
	Rule          configuration.NamingRule `json:"rule"`
	Description   string                   `json:"description"`
	Verifications []Verification           `json:"verifications"`
	Passes        bool                     `json:"passes"`
}

type Verification struct {
	Package string   `json:"package"`
	Details []string `json:"details"`
	Passes  bool     `json:"passes"`
}

type InterfaceDescription struct {
	Name    string
	Methods []MethodDescription
}

type MethodDescription struct {
	Name         string
	Parameters   []string
	ReturnValues []string
}

type StructDescription struct {
	Name    string
	Methods []MethodDescription
}
