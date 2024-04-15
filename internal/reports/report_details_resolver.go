package reports

import (
	"github.com/fdaines/arch-go/pkg/verifications"
)

func resolveReportDetails(result *verifications.Result) *ReportDetails {
	return &ReportDetails{
		DependenciesVerificationDetails: resolveDependenciesDetails(result),
		FunctionsVerificationDetails:    resolveFunctionsDetails(result),
		ContentsVerificationDetails:     resolveContentsDetails(result),
		NamingVerificationDetails:       resolveNamingDetails(result),
	}
}

func resolveDependenciesDetails(result *verifications.Result) Verification {
	verificationDetails := Verification{}
	var vDetails []VerificationDetails
	for _, r := range result.DependenciesRuleResult.Results {
		resolveVerificationStatus(r.Passes, &verificationDetails)

		var packageDetails []PackageDetails
		vTotal, vFailed := 0, 0
		for _, v := range r.Verifications {
			vTotal++
			status := checkVerificationStatus(v.Passes, &vFailed)
			packageDetails = append(packageDetails, PackageDetails{
				Package: v.Package,
				Status:  status,
				Details: v.Details,
			})
		}
		ruleStatus := resolveRuleStatus(vFailed)

		vDetails = append(vDetails, VerificationDetails{
			Rule:           r.Description,
			Status:         ruleStatus,
			Passed:         vTotal - vFailed,
			Failed:         vFailed,
			Total:          vTotal,
			PackageDetails: packageDetails,
		})
		verificationDetails.Details = vDetails
	}

	return verificationDetails
}

func resolveFunctionsDetails(result *verifications.Result) Verification {
	verificationDetails := Verification{}
	var vDetails []VerificationDetails
	for _, r := range result.FunctionsRuleResult.Results {
		resolveVerificationStatus(r.Passes, &verificationDetails)

		var packageDetails []PackageDetails
		vTotal, vFailed := 0, 0
		for _, v := range r.Verifications {
			vTotal++
			status := checkVerificationStatus(v.Passes, &vFailed)
			packageDetails = append(packageDetails, PackageDetails{
				Package: v.Package,
				Status:  status,
				Details: v.Details,
			})
		}
		ruleStatus := resolveRuleStatus(vFailed)

		vDetails = append(vDetails, VerificationDetails{
			Rule:           r.Description,
			Status:         ruleStatus,
			Passed:         vTotal - vFailed,
			Failed:         vFailed,
			Total:          vTotal,
			PackageDetails: packageDetails,
		})
		verificationDetails.Details = vDetails
	}

	return verificationDetails
}

func resolveContentsDetails(result *verifications.Result) Verification {
	verificationDetails := Verification{}
	var vDetails []VerificationDetails
	for _, r := range result.ContentsRuleResult.Results {
		resolveVerificationStatus(r.Passes, &verificationDetails)

		var packageDetails []PackageDetails
		vTotal, vFailed := 0, 0
		for _, v := range r.Verifications {
			vTotal++
			status := checkVerificationStatus(v.Passes, &vFailed)
			packageDetails = append(packageDetails, PackageDetails{
				Package: v.Package,
				Status:  status,
				Details: v.Details,
			})
		}
		ruleStatus := resolveRuleStatus(vFailed)

		vDetails = append(vDetails, VerificationDetails{
			Rule:           r.Description,
			Status:         ruleStatus,
			Passed:         vTotal - vFailed,
			Failed:         vFailed,
			Total:          vTotal,
			PackageDetails: packageDetails,
		})
		verificationDetails.Details = vDetails
	}

	return verificationDetails
}

func resolveNamingDetails(result *verifications.Result) Verification {
	verificationDetails := Verification{}
	var vDetails []VerificationDetails
	for _, r := range result.NamingRuleResult.Results {
		resolveVerificationStatus(r.Passes, &verificationDetails)

		var packageDetails []PackageDetails
		vTotal, vFailed := 0, 0
		for _, v := range r.Verifications {
			vTotal++
			status := checkVerificationStatus(v.Passes, &vFailed)
			packageDetails = append(packageDetails, PackageDetails{
				Package: v.Package,
				Status:  status,
				Details: v.Details,
			})
		}
		ruleStatus := resolveRuleStatus(vFailed)

		vDetails = append(vDetails, VerificationDetails{
			Rule:           r.Description,
			Status:         ruleStatus,
			Passed:         vTotal - vFailed,
			Failed:         vFailed,
			Total:          vTotal,
			PackageDetails: packageDetails,
		})
		verificationDetails.Details = vDetails
	}

	return verificationDetails
}
