package config

type NamingRule struct {
	Package                           string                       `yaml:"package"`
	InterfaceImplementationNamingRule *InterfaceImplementationRule `yaml:"interfaceImplementationNamingRule"`
}

type InterfaceImplementationRule struct {
	StructsThatImplement             string `yaml:"structsThatImplement"`
	ShouldHaveSimpleNameStartingWith string `yaml:"shouldHaveSimpleNameStartingWith"`
	ShouldHaveSimpleNameEndingWith   string `yaml:"shouldHaveSimpleNameEndingWith"`
}
