package configuration

// DeprecatedConfig represents a configuration in a deprecated format.
// Deprecated: Use Config instead.
type DeprecatedConfig struct {
	// DependenciesRules contains a set of dependencies rules.
	DependenciesRules []*DeprecatedDependenciesRule `yaml:"dependenciesRules,omitempty"`

	// ContentRules contains a set of contents rules.
	ContentRules []*ContentsRule `yaml:"contentsRules,omitempty"`

	// CyclesRules contains a set of cycles rules.
	CyclesRules []*CyclesRule `yaml:"cyclesRules,omitempty"`

	// FunctionsRules contains a set of functions rules.
	FunctionsRules []*FunctionsRule `yaml:"functionsRules,omitempty"`

	// NamingRules contains a set of naming rules.
	NamingRules []*NamingRule `yaml:"namingRules,omitempty"`
}

// DeprecatedDependenciesRule represents a dependencies rule in a deprecated format.
// Deprecated: Use DependenciesRule instead.
type DeprecatedDependenciesRule struct {
	// Package is the package pattern to be evaluated.
	Package string `yaml:"package,omitempty"`

	// ShouldOnlyDependsOn packages should only use these internal dependencies.
	ShouldOnlyDependsOn []string `yaml:"shouldOnlyDependsOn,omitempty"`

	// ShouldNotDependsOn packages should not use these internal dependencies.
	ShouldNotDependsOn []string `yaml:"shouldNotDependsOn,omitempty"`

	// ShouldOnlyDependsOnExternal packages should only use these external dependencies.
	ShouldOnlyDependsOnExternal []string `yaml:"shouldOnlyDependsOnExternal,omitempty"`

	// ShouldNotDependsOnExternal packages should not use these external dependencies.
	ShouldNotDependsOnExternal []string `yaml:"shouldNotDependsOnExternal,omitempty"`
}
