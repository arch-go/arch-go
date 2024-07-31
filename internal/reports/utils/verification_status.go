package utils

import "github.com/arch-go/arch-go/internal/reports/model"

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
