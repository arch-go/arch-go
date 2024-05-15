package configuration

// Deprecated: DeprecatedConfig represents a configuration in a deprecated format.
type DeprecatedConfig struct {
	DependenciesRules []*DeprecatedDependenciesRule `yaml:"dependenciesRules,omitempty"` // contains a set of dependencies rules
	ContentRules      []*ContentsRule               `yaml:"contentsRules,omitempty"`     // contains a set of contents rules
	CyclesRules       []*CyclesRule                 `yaml:"cyclesRules,omitempty"`       // contains a set of cycles rules
	FunctionsRules    []*FunctionsRule              `yaml:"functionsRules,omitempty"`    // contains a set of functions rules
	NamingRules       []*NamingRule                 `yaml:"namingRules,omitempty"`       // contains a set of naming rules
}

// Deprecated: DeprecatedDependenciesRule represents a dependencies rule in a deprecated format.
type DeprecatedDependenciesRule struct {
	Package                     string   `yaml:"package,omitempty"`                     // the package pattern to be evaluated
	ShouldOnlyDependsOn         []string `yaml:"shouldOnlyDependsOn,omitempty"`         // packages should only use these internal dependencies
	ShouldNotDependsOn          []string `yaml:"shouldNotDependsOn,omitempty"`          // packages should not use these internal dependencies
	ShouldOnlyDependsOnExternal []string `yaml:"shouldOnlyDependsOnExternal,omitempty"` // packages should only use these external dependencies
	ShouldNotDependsOnExternal  []string `yaml:"shouldNotDependsOnExternal,omitempty"`  // packages should not use these external dependencies
}
