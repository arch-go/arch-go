package utils

import "github.com/arch-go/arch-go/v2/internal/reports/model"

func ResolveStatus(result bool) string {
	if result {
		return passStatus
	}

	return failStatus
}

func ResolveGlobalStatus(compliance *model.ThresholdSummary, coverage *model.ThresholdSummary) bool {
	passCompliance := false
	if compliance == nil || compliance.Pass {
		passCompliance = true
	}

	passCoverage := false
	if coverage == nil || coverage.Pass {
		passCoverage = true
	}

	return passCompliance && passCoverage
}
