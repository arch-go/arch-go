// Package configuration contains structures and functions to support loading architecture rules.
package configuration

// Config contains the architecture rules and the thresholds for coverage and compliance.
type Config struct {
	// Version is the version of configuration.
	Version int `yaml:"version,omitempty"`

	// Threshold contains threshold values.
	Threshold *Threshold `yaml:"threshold,omitempty"`

	// DependenciesRules contains a set of dependencies rules.
	DependenciesRules []*DependenciesRule `yaml:"dependenciesRules,omitempty"`

	// ContentRules contains a set of contents rules.
	ContentRules []*ContentsRule `yaml:"contentsRules,omitempty"`

	// CyclesRules contains a set of cycles rules.
	// Deprecated
	CyclesRules []*CyclesRule `yaml:"cyclesRules,omitempty"`

	// FunctionsRules contains a set of functions rules.
	FunctionsRules []*FunctionsRule `yaml:"functionsRules,omitempty"`

	// NamingRules contains a set of naming rules.
	NamingRules []*NamingRule `yaml:"namingRules,omitempty"`
}

// Threshold contains the compliance and coverage rate to consider the evaluations succeeded.
type Threshold struct {
	// Compliance threshold.
	Compliance *int `yaml:"compliance,omitempty"`

	// Coverage threshold.
	Coverage *int `yaml:"coverage,omitempty"`
}

// DependenciesRule represents a rule related to dependencies between packages.
type DependenciesRule struct {
	// Package is the package pattern to be evaluated.
	Package string `yaml:"package,omitempty"`

	// ShouldOnlyDependsOn packages should only use these dependencies.
	ShouldOnlyDependsOn *Dependencies `yaml:"shouldOnlyDependsOn,omitempty"`

	// ShouldNotDependsOn packages should not use these dependencies.
	ShouldNotDependsOn *Dependencies `yaml:"shouldNotDependsOn,omitempty"`
}

// FunctionsRule represents a rule related to functions in packages.
type FunctionsRule struct {
	// Package is the package pattern to be evaluated.
	Package string `yaml:"package,omitempty"`

	// MaxParameters are the max number of parameters that the functions should contain.
	MaxParameters *int `yaml:"maxParameters,omitempty"`

	// MaxReturnValues are the max number of values that the functions should return.
	MaxReturnValues *int `yaml:"maxReturnValues,omitempty"`

	// MaxLines are the max number of lines that the functions should contain.
	MaxLines *int `yaml:"maxLines,omitempty"`

	// MaxPublicFunctionPerFile are the max number of functions allowed per file.
	MaxPublicFunctionPerFile *int `yaml:"maxPublicFunctionPerFile,omitempty"`
}

// ContentsRule represents a rule related to package contents.
type ContentsRule struct {
	// Package are the package pattern to be evaluated.
	Package string `yaml:"package,omitempty"`

	// ShouldOnlyContainInterfaces if true, then the packages should only contain interfaces.
	ShouldOnlyContainInterfaces bool `yaml:"shouldOnlyContainInterfaces,omitempty"`

	// ShouldOnlyContainStructs if true, then the packages should only contain structs.
	ShouldOnlyContainStructs bool `yaml:"shouldOnlyContainStructs,omitempty"`

	// ShouldOnlyContainFunctions if true, then the packages should only contain functions.
	ShouldOnlyContainFunctions bool `yaml:"shouldOnlyContainFunctions,omitempty"`

	// ShouldOnlyContainMethods if true, then the packages should only contain methods.
	ShouldOnlyContainMethods bool `yaml:"shouldOnlyContainMethods,omitempty"`

	// ShouldNotContainInterfaces if true, then the packages should not contain interfaces.
	ShouldNotContainInterfaces bool `yaml:"shouldNotContainInterfaces,omitempty"`

	// ShouldNotContainStructs if true, then the packages should not contain structs.
	ShouldNotContainStructs bool `yaml:"shouldNotContainStructs,omitempty"`

	// ShouldNotContainFunctions if true, then the packages should not contain functions.
	ShouldNotContainFunctions bool `yaml:"shouldNotContainFunctions,omitempty"`

	// ShouldNotContainMethods if true, then the packages should not contain methods.
	ShouldNotContainMethods bool `yaml:"shouldNotContainMethods,omitempty"`
}

// Deprecated: CyclesRule was deprecated in v1.4.0.
type CyclesRule struct {
	// Package is the package pattern to be evaluated.
	Package string `yaml:"package,omitempty"`

	// ShouldNotContainCycles if true, then the packages should not contain cycles.
	ShouldNotContainCycles bool `yaml:"shouldNotContainCycles,omitempty"`
}

// NamingRule represents a naming rule.
type NamingRule struct {
	// Package the package pattern to be evaluated.
	Package string `yaml:"package"`

	// InterfaceImplementationNamingRule.
	InterfaceImplementationNamingRule *InterfaceImplementationRule `yaml:"interfaceImplementationNamingRule"`
}

// Dependencies contains dependencies grouped by origin.
type Dependencies struct {
	// Internal contains a set of internal dependencies (same go module).
	Internal []string `yaml:"internal,omitempty"`

	// External contains a set of external dependencies.
	External []string `yaml:"external,omitempty"`

	// Standard contains a set of standard dependencies.
	Standard []string `yaml:"standard,omitempty"`
}

// InterfaceImplementationRule represents a naming rule related to interface implementation.
type InterfaceImplementationRule struct {
	// StructsThatImplement the implemented interface.
	StructsThatImplement StructsThatImplement `yaml:"structsThatImplement"`

	// ShouldHaveSimpleNameStartingWith the struct that implements the interface should have this prefix.
	ShouldHaveSimpleNameStartingWith *string `yaml:"shouldHaveSimpleNameStartingWith"`

	// ShouldHaveSimpleNameEndingWith the struct that implements the interface should have this suffix.
	ShouldHaveSimpleNameEndingWith *string `yaml:"shouldHaveSimpleNameEndingWith"`
}

// StructsThatImplement tells where the interface to be implemented is defined and its name.
type StructsThatImplement struct {
	// Internal contains the interface definition in the project.
	Internal *string `yaml:"internal,omitempty"`

	// External contains the interface definition in an external package.
	External *PackageAndInterface `yaml:"external,omitempty"`

	// Standard contains the interface definition in a standard Go package.
	Standard *PackageAndInterface `yaml:"standard,omitempty"`
}

// PackageAndInterface contains the package and interface name to be implemented.
type PackageAndInterface struct {
	Package   string `yaml:"package"`
	Interface string `yaml:"interface"`
}
