package html

import (
	"bytes"
	"fmt"
	"github.com/fdaines/arch-go/internal/model/result"
	"strings"
)

const breakdownTemplate = `
<h3>Breakdown by Rule</h3>
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
`

const ruleDetailTemplate = `<tr>
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

func ruleList(summary result.RulesSummary) string {
	var buffer bytes.Buffer
	rules := []string{"DependenciesRule", "FunctionsRule", "ContentRule", "CycleRule", "NamingRule"}

	for _,r := range rules {
		var ratio int32
		if summary.Details[r].Total > 0 {
			ratio = 100 * summary.Details[r].Succeeded / summary.Details[r].Total
		}
		d := summary.Details[r]
		buffer.WriteString(fmt.Sprintf(ruleDetailTemplate, r, ratio, d.Succeeded, d.Total, d.Total, d.Succeeded, d.Failed))
	}
	return strings.Replace(breakdownTemplate, "[RULES]", buffer.String(), 1)
}

