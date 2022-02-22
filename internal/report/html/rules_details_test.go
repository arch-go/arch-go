package html

import (
	"github.com/fdaines/arch-go/internal/impl/dependencies"
	implModel "github.com/fdaines/arch-go/internal/impl/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRulesDetails(t *testing.T) {
	t.Run("Calls ruleDetails function with no verifications", func(t *testing.T) {
		verifications := []implModel.RuleVerification{
		}
		expected := "\n<h3>Rules Details</h3>\n<table border=\"1\" frame=\"void\"\" rules=\"rows\">\n    <thead>\n        <tr>\n            <th style=\"width:200px;\">Rule Type</th>\n            <th>Rule Description</th>\n            <th style=\"width:100px;\">Result</th>\n        </tr>\n    </thead>\n    <tbody>\n\t\t\n\t</tbody>\n</table>\n"

		result := ruleDetails(verifications)

		assert.Equal(t, expected, result)
	})

	t.Run("Calls ruleDetails function with verifications", func(t *testing.T) {
		verifications := []implModel.RuleVerification{
			&dependencies.DependencyRuleVerification{
				Module: "foobar",
				Description: "Rule description",
				Passes: true,
			},
		}
		expected := "\n<h3>Rules Details</h3>\n<table border=\"1\" frame=\"void\"\" rules=\"rows\">\n    <thead>\n        <tr>\n            <th style=\"width:200px;\">Rule Type</th>\n            <th>Rule Description</th>\n            <th style=\"width:100px;\">Result</th>\n        </tr>\n    </thead>\n    <tbody>\n\t\t<tr>\n\t<td style=\"font-weight:bold;vertical-align:top;\" rowspan=\"1\">DependenciesRule</td>\n\t<td style=\"font-weight:bold;\">Rule description</td>\n\t<td style=\"font-weight:bold;vertical-align:top;text-align:center;color:green\">Succeed</td>\n</tr>\n\t</tbody>\n</table>\n"

		result := ruleDetails(verifications)

		assert.Equal(t, expected, result)
	})
}