package impl

import (
	"github.com/fatih/color"
	"github.com/fdaines/arch-go/internal/common"
	"github.com/fdaines/arch-go/internal/utils/timer"
	"github.com/fdaines/arch-go/old/config"
	"github.com/fdaines/arch-go/old/impl/model"
	"github.com/fdaines/arch-go/old/model/result"
	"github.com/fdaines/arch-go/old/report/console"
	"github.com/fdaines/arch-go/old/report/html"
	"github.com/fdaines/arch-go/old/utils/packages"
	"os"
)

func CheckArchitecture() bool {
	returnValue := true
	timer.ExecuteWithTimer(func() {
		configuration, err := config.LoadConfig("arch-go.yml")
		if err != nil {
			color.Red("Error: %+v\n", err)
			os.Exit(1)
		} else {
			pkgs, _ := packages.GetBasicPackagesInfo(true)

			var verifications []model.RuleVerification
			validateVerifications(verifications)

			for _, v := range verifications {
				v.Verify()
				v.PrintResults()
			}

			resultData := result.ResolveReport(pkgs, verifications, configuration)
			returnValue = resultData.Summary.Status

			if common.Html {
				html.GenerateHtmlReport(resultData)
			} else {
				console.GenerateConsoleReport(resultData)
			}
			resultData.Print()

		}
	})
	return returnValue
}

func validateVerifications(verifications []model.RuleVerification) {
	ok := true
	for _, v := range verifications {
		ok = ok && v.ValidatePatterns()
	}
	if !ok {
		color.Red("Some package patterns are invalid, please check documentation: https://github.com/fdaines/arch-go\n")
		os.Exit(1)
	}
}
