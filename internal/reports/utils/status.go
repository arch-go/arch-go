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
