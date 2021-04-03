package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type DependenciesRule struct {
	Package             string   `yaml:"package"`
	ShouldOnlyDependsOn []string `yaml:"shouldOnlyDependsOn"`
	ShouldNotDependsOn  []string `yaml:"shouldNotDependsOn"`
}

type FunctionsRule struct {
	Package                  string `yaml:"package"`
	MaxParameters            int    `yaml:"maxParameters"`
	MaxReturnValues          int    `yaml:"maxReturnValues"`
	MaxPublicFunctionPerFile int    `yaml:"maxPublicFunctionPerFile"`
}

type ContentsRule struct {
	Package                     string `yaml:"package"`
	ShouldOnlyContainInterfaces bool   `yaml:"shouldOnlyContainInterfaces"`
	ShouldOnlyContainTypes      bool   `yaml:"shouldOnlyContainTypes"`
	ShouldOnlyContainFunctions  bool   `yaml:"shouldOnlyContainFunctions"`
	ShouldOnlyContainMethods    bool   `yaml:"shouldOnlyContainMethods"`
	ShouldNotContainInterfaces  bool   `yaml:"shouldNotContainInterfaces"`
	ShouldNotContainTypes       bool   `yaml:"shouldNotContainTypes"`
	ShouldNotContainFunctions   bool   `yaml:"shouldNotContainFunctions"`
	ShouldNotContainMethods     bool   `yaml:"shouldNotContainMethods"`
}

type CyclesRule struct {
	Package                string `yaml:"package"`
	ShouldNotContainCycles bool   `yaml:"shouldNotContainCycles"`
}

type Config struct {
	DependenciesRules []DependenciesRule `yaml:"dependenciesRules"`
	ContentRules      []ContentsRule     `yaml:"contentsRules"`
	CyclesRules       []CyclesRule       `yaml:"cyclesRules"`
	FunctionsRules    []FunctionsRule    `yaml:"functionsRules"`
}

func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
