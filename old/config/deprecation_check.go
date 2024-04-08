package config

import "github.com/fatih/color"

func checkForDeprecatedConfiguration(configuration *Config) {
	if configuration == nil {
		return
	}
	if configuration.CyclesRules != nil && len(configuration.CyclesRules) > 0 {
		color.Yellow("[WARNING] - Cycle Rules were deprecated on Arch-Go v1.4.0")
		configuration.CyclesRules = nil
	}
}
