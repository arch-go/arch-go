package utils

import "github.com/arch-go/arch-go/internal/reports/model"

func ResolveStatus(result bool) string {
	if result {
		return passStatus
	}

	return failStatus
}

func ResolveRuleStatus(failed int) string {
	if failed > 0 {
		return failStatus
	}

	return passStatus
}

func CheckVerificationStatus(passes bool, vFailed *int) string {
	if !passes {
		*vFailed++

		return failStatus
	}

	return passStatus
}

func ResolveVerificationStatus(passes bool, verificationDetails *model.Verification) {
	if passes {
		verificationDetails.Passed++
	} else {
		verificationDetails.Failed++
	}

	verificationDetails.Total++
}

func ResolveGlobalStatus(compliance *model.ThresholdSummary, coverage *model.ThresholdSummary) string {
	passCompliance := false
	if compliance == nil || compliance.Pass {
		passCompliance = true
	}

	passCoverage := false
	if coverage == nil || coverage.Pass {
		passCoverage = true
	}

	if passCompliance && passCoverage {
		return passStatus
	}

	return failStatus
}

const (
	failStatus = "FAIL"
	passStatus = "PASS"
)
