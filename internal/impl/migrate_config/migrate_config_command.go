package migrate_config

import (
	"fmt"
	"github.com/fdaines/arch-go/internal/config"
	"github.com/fdaines/arch-go/internal/utils"
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

func MigrateConfiguration(out io.Writer) {
	utils.ExecuteWithTimer(func() {
		deprecatedConfiguration, err := config.LoadDeprecatedConfig("arch-go.yml")
		if err != nil {
			fmt.Fprintf(out, "Error: %+v\n", err)
			os.Exit(1)
		} else {
			fmt.Fprintf(out, "Migrating deprecated configuration to current schema.\n")
			configuration := migrateRules(deprecatedConfiguration)
			yamlData, err := yaml.Marshal(&deprecatedConfiguration)
			if err != nil {
				fmt.Fprintf(out, "Error while Marshaling. %+v\n", err)
			}
			ok := writeConfiguration(yamlData, "old-arch-go.yml")
			if ok {
				fmt.Fprintf(out, "Deprecated configuration backup at: old-arch-go.yml\n")
			}
			yamlData, err = yaml.Marshal(&configuration)
			if err != nil {
				fmt.Fprintf(out, "Error while Marshaling. %+v\n", err)
			}
			ok = writeConfiguration(yamlData, "arch-go.yml")
			if ok {
				fmt.Fprintf(out, "Configuration saved at: arch-go.yml\n")
			}
		}
	})
}

func writeConfiguration(data []byte, filename string) bool {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		panic(err)
	}
	return true
}

func migrateRules(deprecatedConfig *config.DeprecatedConfig) *config.Config {
	return &config.Config{
		Version:           1,
		DependenciesRules: migrateDependencyRules(deprecatedConfig.DependenciesRules),
		CyclesRules:       deprecatedConfig.CyclesRules,
		ContentRules:      deprecatedConfig.ContentRules,
		FunctionsRules:    deprecatedConfig.FunctionsRules,
		NamingRules:       deprecatedConfig.NamingRules,
	}
}

func migrateDependencyRules(rules []*config.DeprecatedDependenciesRule) []*config.DependenciesRule {
	var dependencyRules []*config.DependenciesRule
	for _, r := range rules {
		dependencyRules = append(dependencyRules, &config.DependenciesRule{
			Package:             r.Package,
			ShouldOnlyDependsOn: resolveAllowedDependencies(r),
			ShouldNotDependsOn:  resolveRestrictedDependencies(r),
		})
	}

	return dependencyRules
}

func resolveAllowedDependencies(r *config.DeprecatedDependenciesRule) *config.Dependencies {
	if len(r.ShouldOnlyDependsOn)+len(r.ShouldOnlyDependsOnExternal) > 0 {
		return &config.Dependencies{
			Internal: r.ShouldOnlyDependsOn,
			External: r.ShouldOnlyDependsOnExternal,
		}
	}
	return nil
}

func resolveRestrictedDependencies(r *config.DeprecatedDependenciesRule) *config.Dependencies {
	if len(r.ShouldNotDependsOn)+len(r.ShouldNotDependsOnExternal) > 0 {
		return &config.Dependencies{
			Internal: r.ShouldNotDependsOn,
			External: r.ShouldNotDependsOnExternal,
		}
	}
	return nil
}
