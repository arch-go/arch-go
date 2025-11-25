package utils

import "github.com/arch-go/arch-go/v2/internal/reports/model"

func ResolveVerificationStatus(passes bool, verificationDetails *model.Verification) {
	if passes {
		verificationDetails.Passed++
	} else {
		verificationDetails.Failed++
	}

	verificationDetails.Total++
}
