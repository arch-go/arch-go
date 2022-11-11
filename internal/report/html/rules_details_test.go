package html

import (
	"github.com/fdaines/arch-go/internal/impl/dependencies"
	"github.com/fdaines/arch-go/internal/impl/model"
	baseModel "github.com/fdaines/arch-go/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRulesDetails(t *testing.T) {
	t.Run("Calls ruleDetails function with no verifications", func(t *testing.T) {
		verifications := []model.RuleVerification{}
		expected := "\n<h3>Rules Details</h3>\n<table border=\"1\" frame=\"void\"\" rules=\"rows\">\n    <thead>\n        <tr>\n            <th style=\"width:200px;\">Rule Type</th>\n            <th>Rule Description</th>\n            <th style=\"width:100px;\">Result</th>\n        </tr>\n    </thead>\n    <tbody>\n\t\t\n\t</tbody>\n</table>\n"

		result := ruleDetails(verifications)

		assert.Equal(t, expected, result)
	})

	t.Run("Calls ruleDetails function with verifications", func(t *testing.T) {
		verifications := []model.RuleVerification{
			&dependencies.DependencyRuleVerification{
				Module:      "foobar",
				Description: "Rule description",
				Passes:      true,
				PackageDetails: []baseModel.PackageVerification{
					{
						Package: &baseModel.PackageInfo{
							Path: "Foo/bar/path",
						},
						Details: []string{"foobar", "barfoo"},
						Passes:  true,
					},
					{
						Package: &baseModel.PackageInfo{
							Path: "dummy/sample/path",
						},
						Details: []string{"foo", "bar"},
						Passes:  false,
					},
				},
			},
		}
		expected := "\n<h3>Rules Details</h3>\n<table border=\"1\" frame=\"void\"\" rules=\"rows\">\n    <thead>\n        <tr>\n            <th style=\"width:200px;\">Rule Type</th>\n            <th>Rule Description</th>\n            <th style=\"width:100px;\">Result</th>\n        </tr>\n    </thead>\n    <tbody>\n\t\t<tr>\n\t<td style=\"font-weight:bold;vertical-align:top;\" rowspan=\"3\">DependenciesRule</td>\n\t<td style=\"font-weight:bold;\">Rule description</td>\n\t<td style=\"font-weight:bold;vertical-align:top;text-align:center;color:green\">Succeed</td>\n</tr><tr style=\"color:green\">\n\t<td style=\"padding-left:10px;\">* Package Foo/bar/path<br/>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;foobar<br/>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;barfoo</td>\n\t<td style=\"text-align:center;vertical-align:top;\">Succeed</td>\n</tr><tr style=\"color:red\">\n\t<td style=\"padding-left:10px;\">* Package dummy/sample/path<br/>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;foo<br/>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;bar</td>\n\t<td style=\"text-align:center;vertical-align:top;\">Fail</td>\n</tr>\n\t</tbody>\n</table>\n"

		result := ruleDetails(verifications)

		assert.Equal(t, expected, result)
	})
}
