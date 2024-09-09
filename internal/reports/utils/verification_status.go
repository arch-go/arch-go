package utils

import "github.com/arch-go/arch-go/internal/reports/model"

func ResolveVerificationStatus(passes bool, verificationDetails *model.Verification) {
	if passes {
		verificationDetails.Passed++
	} else {
		verificationDetails.Failed++
	}

	verificationDetails.Total++
}
