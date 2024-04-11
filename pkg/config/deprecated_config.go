package config

type DeprecatedDependenciesRule struct {
	Package                     string   `yaml:"package,omitempty"`
	ShouldOnlyDependsOn         []string `yaml:"shouldOnlyDependsOn,omitempty"`
	ShouldNotDependsOn          []string `yaml:"shouldNotDependsOn,omitempty"`
	ShouldOnlyDependsOnExternal []string `yaml:"shouldOnlyDependsOnExternal,omitempty"`
	ShouldNotDependsOnExternal  []string `yaml:"shouldNotDependsOnExternal,omitempty"`
}

type DeprecatedConfig struct {
	DependenciesRules []*DeprecatedDependenciesRule `yaml:"dependenciesRules,omitempty"`
	ContentRules      []*ContentsRule               `yaml:"contentsRules,omitempty"`
	CyclesRules       []*CyclesRule                 `yaml:"cyclesRules,omitempty"`
	FunctionsRules    []*FunctionsRule              `yaml:"functionsRules,omitempty"`
	NamingRules       []*NamingRule                 `yaml:"namingRules,omitempty"`
}
