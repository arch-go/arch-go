package config

type DependenciesRule struct {
	Package                     string   `yaml:"packageX"`
	ShouldOnlyDependsOn         []string `yaml:"shouldOnlyDependsOn"`
	ShouldNotDependsOn          []string `yaml:"shouldNotDependsOn"`
	ShouldOnlyDependsOnExternal []string `yaml:"shouldOnlyDependsOnExternal"`
	ShouldNotDependsOnExternal  []string `yaml:"shouldNotDependsOnExternal"`
}

type FunctionsRule struct {
	Package                  string `yaml:"package"`
	MaxParameters            int    `yaml:"maxParameters"`
	MaxReturnValues          int    `yaml:"maxReturnValues"`
	MaxLines                 int    `yaml:"maxLines"`
	MaxPublicFunctionPerFile int    `yaml:"maxPublicFunctionPerFile"`
}

type ContentsRule struct {
	Package                     string `yaml:"package"`
	ShouldOnlyContainInterfaces bool   `yaml:"shouldOnlyContainInterfaces"`
	ShouldOnlyContainStructs    bool   `yaml:"shouldOnlyContainStructs"`
	ShouldOnlyContainFunctions  bool   `yaml:"shouldOnlyContainFunctions"`
	ShouldOnlyContainMethods    bool   `yaml:"shouldOnlyContainMethods"`
	ShouldNotContainInterfaces  bool   `yaml:"shouldNotContainInterfaces"`
	ShouldNotContainStructs     bool   `yaml:"shouldNotContainStructs"`
	ShouldNotContainFunctions   bool   `yaml:"shouldNotContainFunctions"`
	ShouldNotContainMethods     bool   `yaml:"shouldNotContainMethods"`
}

type CyclesRule struct {
	Package                string `yaml:"package"`
	ShouldNotContainCycles bool   `yaml:"shouldNotContainCycles"`
}

type Config struct {
	Version           *int8               `yaml:"version"`
	DependenciesRules []*DependenciesRule `yaml:"dependenciesRules"`
	ContentRules      []*ContentsRule     `yaml:"contentsRules"`
	CyclesRules       []*CyclesRule       `yaml:"cyclesRules"`
	FunctionsRules    []*FunctionsRule    `yaml:"functionsRules"`
	NamingRules       []*NamingRule       `yaml:"namingRules"`
}
