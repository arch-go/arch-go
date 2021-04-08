package impl

import (
	"fmt"
	"github.com/fdaines/arch-go/config"
	"github.com/fdaines/arch-go/model"
	"github.com/fdaines/arch-go/model/result"
	"github.com/fdaines/arch-go/utils"
	"github.com/fdaines/arch-go/utils/packages"
	"os"
)

func CheckArchitecture() bool {
	returnValue := true
	utils.ExecuteWithTimer(func() {
		configuration, err := config.LoadConfig("arch-go.yml")
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
			os.Exit(1)
		} else {
			mainPackage, _ := packages.GetMainPackage()
			pkgs, _ := packages.GetBasicPackagesInfo()
			moduleInfo := &model.ModuleInfo{
				MainPackage: mainPackage,
				Packages:    pkgs,
			}
			result := checkArchitectureRules(configuration, moduleInfo)
			returnValue = checkResult(result)
		}
	})
	return returnValue
}

func checkResult(result *result.Result) bool {
	summary := result.Print()
	summary.Print()
	return summary.Status()
}
