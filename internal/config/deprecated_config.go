package config

type DeprecatedDependenciesRule struct {
	Package                     string   `yaml:"package"`
	ShouldOnlyDependsOn         []string `yaml:"shouldOnlyDependsOn"`
	ShouldNotDependsOn          []string `yaml:"shouldNotDependsOn"`
	ShouldOnlyDependsOnExternal []string `yaml:"shouldOnlyDependsOnExternal"`
	ShouldNotDependsOnExternal  []string `yaml:"shouldNotDependsOnExternal"`
}

type DeprecatedConfig struct {
	DependenciesRules []*DeprecatedDependenciesRule `yaml:"dependenciesRules"`
	ContentRules      []*ContentsRule               `yaml:"contentsRules"`
	CyclesRules       []*CyclesRule                 `yaml:"cyclesRules"`
	FunctionsRules    []*FunctionsRule              `yaml:"functionsRules"`
	NamingRules       []*NamingRule                 `yaml:"namingRules"`
}
