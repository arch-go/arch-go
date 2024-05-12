package reports

import "github.com/fdaines/arch-go/internal/reports/model"

func resolveRuleStatus(failed int) string {
	if failed > 0 {
		return failStatus
	}
	return passStatus
}

func checkVerificationStatus(passes bool, vFailed *int) string {
	if !passes {
		*vFailed++
		return failStatus
	}
	return passStatus
}

func resolveVerificationStatus(passes bool, verificationDetails *model.Verification) {
	if passes {
		verificationDetails.Passed++
	} else {
		verificationDetails.Failed++
	}
	verificationDetails.Total++
}

func resolveGlobalStatus(compliance *model.ThresholdSummary, coverage *model.ThresholdSummary) string {
	passCompliance := false
	if compliance == nil || compliance.Status == passStatus {
		passCompliance = true
	}
	passCoverage := false
	if coverage == nil || coverage.Status == passStatus {
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
