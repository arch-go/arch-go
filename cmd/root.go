package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/fdaines/arch-go/common"
	"github.com/fdaines/arch-go/config"
	"github.com/fdaines/arch-go/impl"
	"github.com/fdaines/arch-go/impl/model"
	"github.com/fdaines/arch-go/utils/output"
	"github.com/fdaines/arch-go/utils/packages"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:     "arch-go",
	Version: common.Version,
	Short:   "Architecture checks for Go",
	Long: `Architecture checks for Go:
* Dependencies
* Package contents
* Cyclic dependencies`,
	Run: runCommand,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&common.Verbose, "verbose", "v", false, "Verbose Output")
}

func runCommand(cmd *cobra.Command, args []string) {
	configuration, err := config.LoadConfig("arch-go.yml")
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	} else {
		mainPackage, _ := packages.GetMainPackage()
		pkgs, _ := packages.GetBasicPackagesInfo()
		result := impl.CheckArchitecture(configuration, mainPackage, pkgs)
		checkResult(result)
	}
}

func checkResult(result *model.Result) {
	var rules, success, fails int
	output.Print("--------------------------------------")
	for _, dr := range result.DependenciesRulesResults {
		rules++
		if dr.Passes {
			success++
			color.Green("[PASS] - %s\n", dr.Description)
		} else {
			fails++
			color.Red("[FAIL] - %s\n", dr.Description)
			for _, fd := range dr.Failures {
				color.Red("\tPackage '%s' fails\n", fd.Package)
				for _, str := range fd.Details {
					color.Red("\t\t%s\n", str)
				}
			}
		}
	}
	for _, cr := range result.ContentsRuleResults {
		rules++
		if cr.Passes {
			success++
			color.Green("[PASS] - %s\n", cr.Description)
		} else {
			fails++
			color.Red("[FAIL] - %s\n", cr.Description)
		}
	}
	for _, cr := range result.CyclesRuleResults {
		rules++
		if cr.Passes {
			success++
			color.Green("[PASS] - %s\n", cr.Description)
		} else {
			fails++
			color.Red("[FAIL] - %s\n", cr.Description)
			for _, fd := range cr.Failures {
				color.Red("\tPackage '%s' fails\n", fd.Package)
				for idx, str := range fd.Details {
					spaces := strings.Repeat(" ", idx+1)
					color.Red("\t%s + imports %s\n", spaces, str)
				}
			}
		}
	}

	output.Print("--------------------------------------")
	output.Printf("Total Rules: \t%d\n", rules)
	output.Printf("Succeeded: \t%d\n", success)
	output.Printf("Failed: \t%d\n", fails)
	output.Printf("Time: \t%.3f s\n", 1.0)
	if fails > 0 {
		os.Exit(1)
	}
}
