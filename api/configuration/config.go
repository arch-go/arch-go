package configuration

type Dependencies struct {
	Internal []string `yaml:"internal,omitempty"`
	External []string `yaml:"external,omitempty"`
	Standard []string `yaml:"standard,omitempty"`
}

type DependenciesRule struct {
	Package             string        `yaml:"package,omitempty"`
	ShouldOnlyDependsOn *Dependencies `yaml:"shouldOnlyDependsOn,omitempty"`
	ShouldNotDependsOn  *Dependencies `yaml:"shouldNotDependsOn,omitempty"`
}

type FunctionsRule struct {
	Package                  string `yaml:"package,omitempty"`
	MaxParameters            *int   `yaml:"maxParameters,omitempty"`
	MaxReturnValues          *int   `yaml:"maxReturnValues,omitempty"`
	MaxLines                 *int   `yaml:"maxLines,omitempty"`
	MaxPublicFunctionPerFile *int   `yaml:"maxPublicFunctionPerFile,omitempty"`
}

type ContentsRule struct {
	Package                     string `yaml:"package,omitempty"`
	ShouldOnlyContainInterfaces bool   `yaml:"shouldOnlyContainInterfaces,omitempty"`
	ShouldOnlyContainStructs    bool   `yaml:"shouldOnlyContainStructs,omitempty"`
	ShouldOnlyContainFunctions  bool   `yaml:"shouldOnlyContainFunctions,omitempty"`
	ShouldOnlyContainMethods    bool   `yaml:"shouldOnlyContainMethods,omitempty"`
	ShouldNotContainInterfaces  bool   `yaml:"shouldNotContainInterfaces,omitempty"`
	ShouldNotContainStructs     bool   `yaml:"shouldNotContainStructs,omitempty"`
	ShouldNotContainFunctions   bool   `yaml:"shouldNotContainFunctions,omitempty"`
	ShouldNotContainMethods     bool   `yaml:"shouldNotContainMethods,omitempty"`
}

type CyclesRule struct {
	Package                string `yaml:"package,omitempty"`
	ShouldNotContainCycles bool   `yaml:"shouldNotContainCycles,omitempty"`
}

type NamingRule struct {
	Package                           string                       `yaml:"package"`
	InterfaceImplementationNamingRule *InterfaceImplementationRule `yaml:"interfaceImplementationNamingRule"`
}

type InterfaceImplementationRule struct {
	StructsThatImplement             string  `yaml:"structsThatImplement"`
	ShouldHaveSimpleNameStartingWith *string `yaml:"shouldHaveSimpleNameStartingWith"`
	ShouldHaveSimpleNameEndingWith   *string `yaml:"shouldHaveSimpleNameEndingWith"`
}

type Threshold struct {
	Compliance *int `yaml:"compliance,omitempty"`
	Coverage   *int `yaml:"coverage,omitempty"`
}

type Config struct {
	Version           int                 `yaml:"version,omitempty"`
	Threshold         *Threshold          `yaml:"threshold,omitempty"`
	DependenciesRules []*DependenciesRule `yaml:"dependenciesRules,omitempty"`
	ContentRules      []*ContentsRule     `yaml:"contentsRules,omitempty"`
	CyclesRules       []*CyclesRule       `yaml:"cyclesRules,omitempty"`
	FunctionsRules    []*FunctionsRule    `yaml:"functionsRules,omitempty"`
	NamingRules       []*NamingRule       `yaml:"namingRules,omitempty"`
}
