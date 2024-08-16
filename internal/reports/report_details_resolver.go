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

				if !dv.Passes {
					vFailed++
				}

				packageDetails = append(packageDetails, model.PackageDetails{
					Package: dv.Package,
					Pass:    dv.Passes,
					Details: dv.Details,
				})
			}

			vDetails = append(vDetails, model.VerificationDetails{
				Rule:           dr.Description,
				Pass:           vFailed == 0,
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

				if !fv.Passes {
					vFailed++
				}

				packageDetails = append(packageDetails, model.PackageDetails{
					Package: fv.Package,
					Pass:    fv.Passes,
					Details: fv.Details,
				})
			}

			vDetails = append(vDetails, model.VerificationDetails{
				Rule:           fr.Description,
				Pass:           vFailed == 0,
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

				if !cv.Passes {
					vFailed++
				}

				packageDetails = append(packageDetails, model.PackageDetails{
					Package: cv.Package,
					Pass:    cv.Passes,
					Details: cv.Details,
				})
			}

			vDetails = append(vDetails, model.VerificationDetails{
				Rule:           cr.Description,
				Pass:           vFailed == 0,
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

				if !nv.Passes {
					vFailed++
				}

				packageDetails = append(packageDetails, model.PackageDetails{
					Package: nv.Package,
					Pass:    nv.Passes,
					Details: nv.Details,
				})
			}

			vDetails = append(vDetails, model.VerificationDetails{
				Rule:           nr.Description,
				Pass:           vFailed == 0,
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
