package reports

import (
	"github.com/fdaines/arch-go/internal/reports/model"
	"github.com/fdaines/arch-go/pkg/archgo"
)

func resolveReportDetails(result *archgo.Result) *model.ReportDetails {
	return &model.ReportDetails{
		DependenciesVerificationDetails: resolveDependenciesDetails(result),
		FunctionsVerificationDetails:    resolveFunctionsDetails(result),
		ContentsVerificationDetails:     resolveContentsDetails(result),
		NamingVerificationDetails:       resolveNamingDetails(result),
	}
}

func resolveDependenciesDetails(result *archgo.Result) model.Verification {
	verificationDetails := model.Verification{}
	var vDetails []model.VerificationDetails
	if result.DependenciesRuleResult != nil {
		for _, dr := range result.DependenciesRuleResult.Results {
			resolveVerificationStatus(dr.Passes, &verificationDetails)

			var packageDetails []model.PackageDetails
			vTotal, vFailed := 0, 0
			for _, dv := range dr.Verifications {
				vTotal++
				status := checkVerificationStatus(dv.Passes, &vFailed)
				packageDetails = append(packageDetails, model.PackageDetails{
					Package: dv.Package,
					Status:  status,
					Details: dv.Details,
				})
			}
			ruleStatus := resolveRuleStatus(vFailed)

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

func resolveFunctionsDetails(result *archgo.Result) model.Verification {
	verificationDetails := model.Verification{}
	var vDetails []model.VerificationDetails
	if result.FunctionsRuleResult != nil {
		for _, fr := range result.FunctionsRuleResult.Results {
			resolveVerificationStatus(fr.Passes, &verificationDetails)

			var packageDetails []model.PackageDetails
			vTotal, vFailed := 0, 0
			for _, fv := range fr.Verifications {
				vTotal++
				status := checkVerificationStatus(fv.Passes, &vFailed)
				packageDetails = append(packageDetails, model.PackageDetails{
					Package: fv.Package,
					Status:  status,
					Details: fv.Details,
				})
			}
			ruleStatus := resolveRuleStatus(vFailed)

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

func resolveContentsDetails(result *archgo.Result) model.Verification {
	verificationDetails := model.Verification{}
	var vDetails []model.VerificationDetails
	if result.ContentsRuleResult != nil {
		for _, cr := range result.ContentsRuleResult.Results {
			resolveVerificationStatus(cr.Passes, &verificationDetails)

			var packageDetails []model.PackageDetails
			vTotal, vFailed := 0, 0
			for _, cv := range cr.Verifications {
				vTotal++
				status := checkVerificationStatus(cv.Passes, &vFailed)
				packageDetails = append(packageDetails, model.PackageDetails{
					Package: cv.Package,
					Status:  status,
					Details: cv.Details,
				})
			}
			ruleStatus := resolveRuleStatus(vFailed)

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

func resolveNamingDetails(result *archgo.Result) model.Verification {
	verificationDetails := model.Verification{}
	var vDetails []model.VerificationDetails
	if result.NamingRuleResult != nil {
		for _, nr := range result.NamingRuleResult.Results {
			resolveVerificationStatus(nr.Passes, &verificationDetails)

			var packageDetails []model.PackageDetails
			vTotal, vFailed := 0, 0
			for _, nv := range nr.Verifications {
				vTotal++
				status := checkVerificationStatus(nv.Passes, &vFailed)
				packageDetails = append(packageDetails, model.PackageDetails{
					Package: nv.Package,
					Status:  status,
					Details: nv.Details,
				})
			}
			ruleStatus := resolveRuleStatus(vFailed)

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
