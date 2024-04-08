package result

import (
	"github.com/fdaines/arch-go/old/config"
	"github.com/fdaines/arch-go/old/impl/model"
	baseModel "github.com/fdaines/arch-go/old/model"
)

func ResolveReport(pkgs []*baseModel.PackageInfo, verifications []model.RuleVerification, configuration *config.Config) Report {
	report := Report{
		TotalPackages: len(pkgs),
		Summary:       ResolveRulesSummary(pkgs, verifications, configuration),
		Verifications: verifications,
	}

	return report
}

type Report struct {
	TotalPackages int
	Summary       RulesSummary
	Verifications []model.RuleVerification
}

func (r Report) Print() {
	r.Summary.Print()
}
