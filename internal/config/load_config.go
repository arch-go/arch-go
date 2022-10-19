package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		_, err := LoadDeprecatedConfig(configPath)
		if err == nil {
			return nil, fmt.Errorf("A deprecated architecture description was provided. To update the arch-go.yml file please run './arch-go migrate-config'")
		}
		return nil, err
	}

	return config, nil
}

func LoadDeprecatedConfig(configPath string) (*DeprecatedConfig, error) {
	config := &DeprecatedConfig{}
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
