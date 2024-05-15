// Package configuration contains structures and functions to support loading architecture rules.
package configuration

// Config contains the architecture rules and the thresholds for coverage and compliance.
type Config struct {
	Version           int                 `yaml:"version,omitempty"`           // version of configuration
	Threshold         *Threshold          `yaml:"threshold,omitempty"`         // contains threshold values
	DependenciesRules []*DependenciesRule `yaml:"dependenciesRules,omitempty"` // contains a set of dependencies rules
	ContentRules      []*ContentsRule     `yaml:"contentsRules,omitempty"`     // contains a set of contents rules
	CyclesRules       []*CyclesRule       `yaml:"cyclesRules,omitempty"`       // contains a set of cycles rules (deprecated)
	FunctionsRules    []*FunctionsRule    `yaml:"functionsRules,omitempty"`    // contains a set of functions rules
	NamingRules       []*NamingRule       `yaml:"namingRules,omitempty"`       // contains a set of naming rules
}

// Threshold contains the compliance and coverage rate to consider the evaluations succeeded.
type Threshold struct {
	Compliance *int `yaml:"compliance,omitempty"` // Compliance threshold.
	Coverage   *int `yaml:"coverage,omitempty"`   // Coverage threshold.
}

// DependenciesRule represents a rule related to dependencies between packages.
type DependenciesRule struct {
	Package             string        `yaml:"package,omitempty"`             // the package pattern to be evaluated
	ShouldOnlyDependsOn *Dependencies `yaml:"shouldOnlyDependsOn,omitempty"` // packages should only use these dependencies
	ShouldNotDependsOn  *Dependencies `yaml:"shouldNotDependsOn,omitempty"`  // packages should not use these dependencies
}

// FunctionsRule represents a rule related to functions in packages.
type FunctionsRule struct {
	Package                  string `yaml:"package,omitempty"`                  // the package pattern to be evaluated
	MaxParameters            *int   `yaml:"maxParameters,omitempty"`            // the max number of parameters that the functions should contain
	MaxReturnValues          *int   `yaml:"maxReturnValues,omitempty"`          // the max number of values that the functions should return
	MaxLines                 *int   `yaml:"maxLines,omitempty"`                 // the max number of lines that the functions should contain
	MaxPublicFunctionPerFile *int   `yaml:"maxPublicFunctionPerFile,omitempty"` // the max number of
}

// ContentsRule represents a rule related to package contents.
type ContentsRule struct {
	Package                     string `yaml:"package,omitempty"`                     // the package pattern to be evaluated
	ShouldOnlyContainInterfaces bool   `yaml:"shouldOnlyContainInterfaces,omitempty"` // if true, then the packages should only contain interfaces
	ShouldOnlyContainStructs    bool   `yaml:"shouldOnlyContainStructs,omitempty"`    // if true, then the packages should only contain structs
	ShouldOnlyContainFunctions  bool   `yaml:"shouldOnlyContainFunctions,omitempty"`  // if true, then the packages should only contain functions
	ShouldOnlyContainMethods    bool   `yaml:"shouldOnlyContainMethods,omitempty"`    // if true, then the packages should only contain methods
	ShouldNotContainInterfaces  bool   `yaml:"shouldNotContainInterfaces,omitempty"`  // if true, then the packages should not contain interfaces
	ShouldNotContainStructs     bool   `yaml:"shouldNotContainStructs,omitempty"`     // if true, then the packages should not contain structs
	ShouldNotContainFunctions   bool   `yaml:"shouldNotContainFunctions,omitempty"`   // if true, then the packages should not contain functions
	ShouldNotContainMethods     bool   `yaml:"shouldNotContainMethods,omitempty"`     // if true, then the packages should not contain methods
}

// Deprecated: CyclesRule was deprecated in v1.4.0
type CyclesRule struct {
	Package                string `yaml:"package,omitempty"`                // the package pattern to be evaluated
	ShouldNotContainCycles bool   `yaml:"shouldNotContainCycles,omitempty"` // if true, then the packages should not contain cycles
}

// NamingRule represents a naming rule.
type NamingRule struct {
	Package                           string                       `yaml:"package"` // the package pattern to be evaluated
	InterfaceImplementationNamingRule *InterfaceImplementationRule `yaml:"interfaceImplementationNamingRule"`
}

// Dependencies contains dependencies grouped by origin.
type Dependencies struct {
	Internal []string `yaml:"internal,omitempty"` // contains a set of internal dependencies (same go module)
	External []string `yaml:"external,omitempty"` // contains a set of external dependencies
	Standard []string `yaml:"standard,omitempty"` // contains a set of standard dependencies
}

// InterfaceImplementationRule represents a naming rule related to interface implementation.
type InterfaceImplementationRule struct {
	StructsThatImplement             string  `yaml:"structsThatImplement"`             // the implemented interface
	ShouldHaveSimpleNameStartingWith *string `yaml:"shouldHaveSimpleNameStartingWith"` // the struct that implements the interface should have this prefix
	ShouldHaveSimpleNameEndingWith   *string `yaml:"shouldHaveSimpleNameEndingWith"`   // the struct that implements the interface should have this suffix
}
