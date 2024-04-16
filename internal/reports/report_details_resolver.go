package reports

import (
	"github.com/fdaines/arch-go/internal/reports/model"
	"github.com/fdaines/arch-go/pkg/verifications"
)

func resolveReportDetails(result *verifications.Result) *model.ReportDetails {
	return &model.ReportDetails{
		DependenciesVerificationDetails: resolveDependenciesDetails(result),
		FunctionsVerificationDetails:    resolveFunctionsDetails(result),
		ContentsVerificationDetails:     resolveContentsDetails(result),
		NamingVerificationDetails:       resolveNamingDetails(result),
	}
}

func resolveDependenciesDetails(result *verifications.Result) model.Verification {
	verificationDetails := model.Verification{}
	var vDetails []model.VerificationDetails
	if result.DependenciesRuleResult != nil {
		for _, r := range result.DependenciesRuleResult.Results {
			resolveVerificationStatus(r.Passes, &verificationDetails)

			var packageDetails []model.PackageDetails
			vTotal, vFailed := 0, 0
			for _, v := range r.Verifications {
				vTotal++
				status := checkVerificationStatus(v.Passes, &vFailed)
				packageDetails = append(packageDetails, model.PackageDetails{
					Package: v.Package,
					Status:  status,
					Details: v.Details,
				})
			}
			ruleStatus := resolveRuleStatus(vFailed)

			vDetails = append(vDetails, model.VerificationDetails{
				Rule:           r.Description,
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

func resolveFunctionsDetails(result *verifications.Result) model.Verification {
	verificationDetails := model.Verification{}
	var vDetails []model.VerificationDetails
	if result.FunctionsRuleResult != nil {
		for _, r := range result.FunctionsRuleResult.Results {
			resolveVerificationStatus(r.Passes, &verificationDetails)

			var packageDetails []model.PackageDetails
			vTotal, vFailed := 0, 0
			for _, v := range r.Verifications {
				vTotal++
				status := checkVerificationStatus(v.Passes, &vFailed)
				packageDetails = append(packageDetails, model.PackageDetails{
					Package: v.Package,
					Status:  status,
					Details: v.Details,
				})
			}
			ruleStatus := resolveRuleStatus(vFailed)

			vDetails = append(vDetails, model.VerificationDetails{
				Rule:           r.Description,
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

func resolveContentsDetails(result *verifications.Result) model.Verification {
	verificationDetails := model.Verification{}
	var vDetails []model.VerificationDetails
	if result.ContentsRuleResult != nil {
		for _, r := range result.ContentsRuleResult.Results {
			resolveVerificationStatus(r.Passes, &verificationDetails)

			var packageDetails []model.PackageDetails
			vTotal, vFailed := 0, 0
			for _, v := range r.Verifications {
				vTotal++
				status := checkVerificationStatus(v.Passes, &vFailed)
				packageDetails = append(packageDetails, model.PackageDetails{
					Package: v.Package,
					Status:  status,
					Details: v.Details,
				})
			}
			ruleStatus := resolveRuleStatus(vFailed)

			vDetails = append(vDetails, model.VerificationDetails{
				Rule:           r.Description,
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

func resolveNamingDetails(result *verifications.Result) model.Verification {
	verificationDetails := model.Verification{}
	var vDetails []model.VerificationDetails
	if result.NamingRuleResult != nil {
		for _, r := range result.NamingRuleResult.Results {
			resolveVerificationStatus(r.Passes, &verificationDetails)

			var packageDetails []model.PackageDetails
			vTotal, vFailed := 0, 0
			for _, v := range r.Verifications {
				vTotal++
				status := checkVerificationStatus(v.Passes, &vFailed)
				packageDetails = append(packageDetails, model.PackageDetails{
					Package: v.Package,
					Status:  status,
					Details: v.Details,
				})
			}
			ruleStatus := resolveRuleStatus(vFailed)

			vDetails = append(vDetails, model.VerificationDetails{
				Rule:           r.Description,
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
