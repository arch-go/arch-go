package reports

import (
	"github.com/arch-go/arch-go/api"
	"github.com/arch-go/arch-go/internal/reports/model"
	"github.com/arch-go/arch-go/internal/reports/utils"
)

func resolveReportDetails(result *api.Result) *model.ReportDetails {
	return &model.ReportDetails{
		DependenciesVerificationDetails: resolveDependenciesDetails(result),
		FunctionsVerificationDetails:    resolveFunctionsDetails(result),
		ContentsVerificationDetails:     resolveContentsDetails(result),
		NamingVerificationDetails:       resolveNamingDetails(result),
	}
}

func resolveDependenciesDetails(result *api.Result) model.Verification {
	var vDetails []model.VerificationDetails

	verificationDetails := model.Verification{}

	if result.DependenciesRuleResult != nil {
		for _, dr := range result.DependenciesRuleResult.Results {
			var packageDetails []model.PackageDetails

			utils.ResolveVerificationStatus(dr.Passes, &verificationDetails)

			vTotal, vFailed := 0, 0

			for _, dv := range dr.Verifications {
				vTotal++
				status := utils.CheckVerificationStatus(dv.Passes, &vFailed)
				packageDetails = append(packageDetails, model.PackageDetails{
					Package: dv.Package,
					Status:  status,
					Details: dv.Details,
				})
			}

			ruleStatus := utils.ResolveRuleStatus(vFailed)

			vDetails = append(vDetails, model.VerificationDetails{
				Rule:           dr.Description,
				Status:         ruleStatus,
				Passed:         vTotal - vFailed,
				Failed:         vFailed,
				Total:          vTotal,
				PackageDetails: packageDetails,
			})
			verificationDetails.Details = vDetails
		}
	}

	return verificationDetails
}

func resolveFunctionsDetails(result *api.Result) model.Verification {
	var vDetails []model.VerificationDetails

	verificationDetails := model.Verification{}

	if result.FunctionsRuleResult != nil {
		for _, fr := range result.FunctionsRuleResult.Results {
			var packageDetails []model.PackageDetails

			utils.ResolveVerificationStatus(fr.Passes, &verificationDetails)

			vTotal, vFailed := 0, 0

			for _, fv := range fr.Verifications {
				vTotal++
				status := utils.CheckVerificationStatus(fv.Passes, &vFailed)
				packageDetails = append(packageDetails, model.PackageDetails{
					Package: fv.Package,
					Status:  status,
					Details: fv.Details,
				})
			}

			ruleStatus := utils.ResolveRuleStatus(vFailed)

			vDetails = append(vDetails, model.VerificationDetails{
				Rule:           fr.Description,
				Status:         ruleStatus,
				Passed:         vTotal - vFailed,
				Failed:         vFailed,
				Total:          vTotal,
				PackageDetails: packageDetails,
			})
			verificationDetails.Details = vDetails
		}
	}

	return verificationDetails
}

func resolveContentsDetails(result *api.Result) model.Verification {
	var vDetails []model.VerificationDetails

	verificationDetails := model.Verification{}

	if result.ContentsRuleResult != nil {
		for _, cr := range result.ContentsRuleResult.Results {
			var packageDetails []model.PackageDetails

			utils.ResolveVerificationStatus(cr.Passes, &verificationDetails)

			vTotal, vFailed := 0, 0

			for _, cv := range cr.Verifications {
				vTotal++
				status := utils.CheckVerificationStatus(cv.Passes, &vFailed)
				packageDetails = append(packageDetails, model.PackageDetails{
					Package: cv.Package,
					Status:  status,
					Details: cv.Details,
				})
			}

			ruleStatus := utils.ResolveRuleStatus(vFailed)

			vDetails = append(vDetails, model.VerificationDetails{
				Rule:           cr.Description,
				Status:         ruleStatus,
				Passed:         vTotal - vFailed,
				Failed:         vFailed,
				Total:          vTotal,
				PackageDetails: packageDetails,
			})
			verificationDetails.Details = vDetails
		}
	}

	return verificationDetails
}

func resolveNamingDetails(result *api.Result) model.Verification {
	var vDetails []model.VerificationDetails

	verificationDetails := model.Verification{}

	if result.NamingRuleResult != nil {
		for _, nr := range result.NamingRuleResult.Results {
			var packageDetails []model.PackageDetails

			utils.ResolveVerificationStatus(nr.Passes, &verificationDetails)

			vTotal, vFailed := 0, 0

			for _, nv := range nr.Verifications {
				vTotal++
				status := utils.CheckVerificationStatus(nv.Passes, &vFailed)
				packageDetails = append(packageDetails, model.PackageDetails{
					Package: nv.Package,
					Status:  status,
					Details: nv.Details,
				})
			}

			ruleStatus := utils.ResolveRuleStatus(vFailed)

			vDetails = append(vDetails, model.VerificationDetails{
				Rule:           nr.Description,
				Status:         ruleStatus,
				Passed:         vTotal - vFailed,
				Failed:         vFailed,
				Total:          vTotal,
				PackageDetails: packageDetails,
			})
			verificationDetails.Details = vDetails
		}
	}

	return verificationDetails
}
