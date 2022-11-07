package html

import (
	"bytes"
	"fmt"
	"github.com/fdaines/arch-go/internal/model/result"
	"strings"
)

const breakdownTemplate = `
<h3>Rules Summary</h3>
<table>
    <thead>
        <tr>
            <th style="width:200px;">Rule Type</th>
            <th style="width:120px;">Summary</th>
            <th style="width:100px;">Total</th>
            <th style="width:100px;">Succeed</th>
			<th style="width:100px;">Fail</th>
        </tr>
    </thead>
    <tbody>
		[RULES]
	</tbody>
</table>
[THRESHOLD_SUMMARY]
`

const ruleSummaryTemplate = `<tr>
	<td>%s</td>
	<td>
		<div class="result_bar">
			<div class="result_succeeded width-%d"></div>
			<div class="result_legend">%d/%d</div>
		</div>
	</td>
	<td style="text-align:center;">%d</td>
	<td style="text-align:center;">%d</td>
	<td style="text-align:center;">%d</td>
</tr>`

const thresholdTemplate = `<br/>
<table>
    <thead>
        <tr>
            <th colspan="4">Threshold Verification</th>
        </tr>
        <tr>
            <th style="width:200px;">Type</th>
            <th style="width:100px;">Current Rate</th>
            <th style="width:100px;">Threshold</th>
            <th style="width:100px;">Status</th>
        </tr>
    </thead>
    <tbody>
        [COMPLIANCE_THRESHOLD]
    </tbody>
</table>`

const thresholdRowTemplate = `<tr style="color:%s">
    <td>%s</td>
    <td style="text-align:center;">%d</td>
    <td style="text-align:center;">%d</td>
    <td style="text-align:center;font-weight:bold">%s</td>
</tr>`

func ruleList(summary result.RulesSummary) string {
	var buffer bytes.Buffer
	rules := []string{"DependenciesRule", "FunctionsRule", "ContentRule", "CycleRule", "NamingRule"}

	for _, r := range rules {
		var ratio int32
		if summary.Details[r].Total > 0 {
			ratio = 100 * summary.Details[r].Succeeded / summary.Details[r].Total
		}
		d := summary.Details[r]
		buffer.WriteString(fmt.Sprintf(ruleSummaryTemplate, r, ratio, d.Succeeded, d.Total, d.Total, d.Succeeded, d.Failed))
	}
	template := strings.Replace(breakdownTemplate, "[RULES]", buffer.String(), 1)
	template = thresholdVerification(template, summary)
	return template
}

func thresholdVerification(template string, summary result.RulesSummary) string {
	if summary.ComplianceThreshold == nil {
		return strings.Replace(template, "[THRESHOLD_SUMMARY]", "", 1)
	}

	var buffer bytes.Buffer
	statusColor := "red"
	if summary.ComplianceThreshold.Status == "Pass" {
		statusColor = "green"
	}

	buffer.WriteString(fmt.Sprintf(thresholdRowTemplate,
		statusColor,
		"Compliance",
		summary.ComplianceThreshold.Rate,
		summary.ComplianceThreshold.Threshold,
		summary.ComplianceThreshold.Status,
	))

	thresholdDetails := strings.Replace(thresholdTemplate, "[COMPLIANCE_THRESHOLD]", buffer.String(), 1)

	return strings.Replace(template, "[THRESHOLD_SUMMARY]", thresholdDetails, 1)
}
