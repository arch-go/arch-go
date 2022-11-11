package html

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRulesList(t *testing.T) {
	t.Run("Calls RuleList function", func(t *testing.T) {
		summary := createRulesSummaryMock(true, true)
		expected := "\n<h3>Rules Summary</h3>\n<table>\n    <thead>\n        <tr>\n            <th style=\"width:200px;\">Rule Type</th>\n            <th style=\"width:120px;\">Summary</th>\n            <th style=\"width:100px;\">Total</th>\n            <th style=\"width:100px;\">Succeed</th>\n\t\t\t<th style=\"width:100px;\">Fail</th>\n        </tr>\n    </thead>\n    <tbody>\n\t\t<tr>\n\t<td>DependenciesRule</td>\n\t<td>\n\t\t<div class=\"result_bar\">\n\t\t\t<div class=\"result_succeeded width-0\"></div>\n\t\t\t<div class=\"result_legend\">0/0</div>\n\t\t</div>\n\t</td>\n\t<td style=\"text-align:center;\">0</td>\n\t<td style=\"text-align:center;\">0</td>\n\t<td style=\"text-align:center;\">0</td>\n</tr><tr>\n\t<td>FunctionsRule</td>\n\t<td>\n\t\t<div class=\"result_bar\">\n\t\t\t<div class=\"result_succeeded width-0\"></div>\n\t\t\t<div class=\"result_legend\">0/0</div>\n\t\t</div>\n\t</td>\n\t<td style=\"text-align:center;\">0</td>\n\t<td style=\"text-align:center;\">0</td>\n\t<td style=\"text-align:center;\">0</td>\n</tr><tr>\n\t<td>ContentRule</td>\n\t<td>\n\t\t<div class=\"result_bar\">\n\t\t\t<div class=\"result_succeeded width-0\"></div>\n\t\t\t<div class=\"result_legend\">0/0</div>\n\t\t</div>\n\t</td>\n\t<td style=\"text-align:center;\">0</td>\n\t<td style=\"text-align:center;\">0</td>\n\t<td style=\"text-align:center;\">0</td>\n</tr><tr>\n\t<td>CycleRule</td>\n\t<td>\n\t\t<div class=\"result_bar\">\n\t\t\t<div class=\"result_succeeded width-0\"></div>\n\t\t\t<div class=\"result_legend\">0/0</div>\n\t\t</div>\n\t</td>\n\t<td style=\"text-align:center;\">0</td>\n\t<td style=\"text-align:center;\">0</td>\n\t<td style=\"text-align:center;\">0</td>\n</tr><tr>\n\t<td>NamingRule</td>\n\t<td>\n\t\t<div class=\"result_bar\">\n\t\t\t<div class=\"result_succeeded width-0\"></div>\n\t\t\t<div class=\"result_legend\">0/0</div>\n\t\t</div>\n\t</td>\n\t<td style=\"text-align:center;\">0</td>\n\t<td style=\"text-align:center;\">0</td>\n\t<td style=\"text-align:center;\">0</td>\n</tr>\n\t</tbody>\n</table>\n<br/>\n<table>\n    <thead>\n        <tr>\n            <th colspan=\"4\">Threshold Verification</th>\n        </tr>\n        <tr>\n            <th style=\"width:200px;\">Type</th>\n            <th style=\"width:100px;\">Current Rate</th>\n            <th style=\"width:100px;\">Threshold</th>\n            <th style=\"width:100px;\">Status</th>\n        </tr>\n    </thead>\n    <tbody>\n        <tr style=\"color:green\">\n    <td>Compliance</td>\n    <td style=\"text-align:center;\">90%</td>\n    <td style=\"text-align:center;\">90%</td>\n    <td style=\"text-align:center;font-weight:bold\">Pass</td>\n</tr><tr style=\"color:red\">\n    <td>Coverage</td>\n    <td style=\"text-align:center;\">85%</td>\n    <td style=\"text-align:center;\">90%</td>\n    <td style=\"text-align:center;font-weight:bold\">Fail</td>\n</tr>\n    </tbody>\n</table>\n"

		result := ruleList(summary)

		assert.Equal(t, expected, result)
	})
}

func TestThresholdVerification(t *testing.T) {
	t.Run("Calls ThresholdVerification function", func(t *testing.T) {
		summary := createRulesSummaryMock(true, true)
		template := "blabla[THRESHOLD_SUMMARY]"
		expected := "blabla<br/>\n<table>\n    <thead>\n        <tr>\n            <th colspan=\"4\">Threshold Verification</th>\n        </tr>\n        <tr>\n            <th style=\"width:200px;\">Type</th>\n            <th style=\"width:100px;\">Current Rate</th>\n            <th style=\"width:100px;\">Threshold</th>\n            <th style=\"width:100px;\">Status</th>\n        </tr>\n    </thead>\n    <tbody>\n        <tr style=\"color:green\">\n    <td>Compliance</td>\n    <td style=\"text-align:center;\">90%</td>\n    <td style=\"text-align:center;\">90%</td>\n    <td style=\"text-align:center;font-weight:bold\">Pass</td>\n</tr><tr style=\"color:red\">\n    <td>Coverage</td>\n    <td style=\"text-align:center;\">85%</td>\n    <td style=\"text-align:center;\">90%</td>\n    <td style=\"text-align:center;font-weight:bold\">Fail</td>\n</tr>\n    </tbody>\n</table>"

		result := thresholdVerification(template, summary)

		assert.Equal(t, expected, result)
	})
}
