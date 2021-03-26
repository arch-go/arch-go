package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type DependenciesRule struct {
	Package string `yaml:"package"`
	ShouldOnlyDependsOn []string `yaml:"shouldOnlyDependsOn"`
	ShouldNotDependsOn []string `yaml:"shouldNotDependsOn"`
}

type Config struct {
	DependenciesRules []DependenciesRule `yaml:"dependenciesRules"`
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
