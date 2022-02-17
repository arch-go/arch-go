package html

import (
	"bytes"
	"fmt"
	"github.com/fdaines/arch-go/internal/impl/model"
	"strings"
)

const ruleDetailsBreakdownTemplate = `
<h3>Rules Details</h3>
<table border="1" frame="void"" rules="rows">
    <thead>
        <tr>
            <th style="width:200px;">Rule Type</th>
            <th>Rule Description</th>
            <th style="width:100px;">Result</th>
        </tr>
    </thead>
    <tbody>
		[DETAILS]
	</tbody>
</table>
`
const ruleDetailsTemplate = `<tr>
	<td style="font-weight:bold;vertical-align:top;" rowspan="%d">%s</td>
	<td style="font-weight:bold;">%s</td>
	<td style="font-weight:bold;vertical-align:top;text-align:center;color:%s">%v</td>
</tr>`
const ruleVerificationTemplate = `<tr style="color:%s">
	<td style="padding-left:10px;">* Package %s</td>
	<td style="text-align:center;vertical-align:top;">%v</td>
</tr>`

func ruleDetails(verifications []model.RuleVerification) string {
	var buffer bytes.Buffer
	for _, v := range verifications {
		buffer.WriteString(fmt.Sprintf(ruleDetailsTemplate, 1+len(v.GetVerifications()), v.Type(), v.Name(), resolveFontColor(v.Status()), resolveResult(v.Status())))
		for _, vx := range v.GetVerifications() {
			buffer.WriteString(fmt.Sprintf(ruleVerificationTemplate, resolveFontColor(vx.Passes), resolveNameAndDetails(vx), resolveResult(vx.Passes)))
		}
	}
	return strings.Replace(ruleDetailsBreakdownTemplate, "[DETAILS]", buffer.String(), 1)
}

func resolveResult(passes bool) string {
	if passes {
		return "Succeed"
	}
	return "Fail"
}

func resolveFontColor(passes bool) string {
	if passes {
		return "green"
	}
	return "red"
}

func resolveNameAndDetails(verification model.PackageVerification) string {
	var details bytes.Buffer
	for _, d := range verification.Details {
		details.WriteString(fmt.Sprintf("<br/>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;%s", d))
	}
	return fmt.Sprintf("%s%s", verification.Package.Path, details.String())
}