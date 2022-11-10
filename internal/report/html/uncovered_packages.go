package html

import (
	"bytes"
	"fmt"
	"github.com/fdaines/arch-go/internal/model/result"
	"strings"
)

func uncoveredPackages(summary result.RulesSummary) string {
	if summary.CoverageThreshold == nil || len(summary.CoverageThreshold.Violations) == 0 {
		return ""
	}

	var buffer bytes.Buffer
	buffer.WriteString("")
	for _, v := range summary.CoverageThreshold.Violations {
		buffer.WriteString(fmt.Sprintf("<li>%s</li>", v))
	}

	return strings.Replace(uncoveredPackagesTemplate, "[UNCOVERED_PACKAGES]", buffer.String(), 1)
}

const uncoveredPackagesTemplate = `
<h3>Uncovered Packages</h3>
<ul>
	[UNCOVERED_PACKAGES]
</ul>
`
