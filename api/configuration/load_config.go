package configuration

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

// LoadConfig loads configuration struct from a YAML file.
func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	d := yaml.NewDecoder(file)
	if err1 := d.Decode(&config); err1 != nil {
		_, err2 := LoadDeprecatedConfig(configPath)
		if err2 == nil {
			return nil, errors.New("a deprecated architecture description was provided." +
				" To update the arch-go.yml file please run 'arch-go migrate-configuration'")
		}

		return nil, err1
	}

	checkThreshold(config)
	checkForDeprecatedConfiguration(config)

	return config, nil
}

// LoadDeprecatedConfig loads configuration struct from a YAML file that contains a deprecated format.
func LoadDeprecatedConfig(configPath string) (*DeprecatedConfig, error) {
	config := &DeprecatedConfig{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	d := yaml.NewDecoder(file)
	if err = d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func checkThreshold(config *Config) {
	if config.Threshold == nil {
		config.Threshold = &Threshold{}
	}

	maxThreshold := 100
	if config.Threshold.Compliance == nil {
		config.Threshold.Compliance = &maxThreshold
	}

	if config.Threshold.Coverage == nil {
		config.Threshold.Coverage = &maxThreshold
	}
}
