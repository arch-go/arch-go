package html

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportHtmlTemplates(t *testing.T) {
	t.Run("resolveTemplates", func(t *testing.T) {
		expectedTemplateFiles := []string{
			"compliance-rate.tmpl",
			"content.tmpl",
			"coverage-rate.tmpl",
			"footer.tmpl",
			"header.tmpl",
			"report.tmpl",
			"rules-summary.tmpl",
			"running.tmpl",
			"summary.tmpl",
			"tab-compliance.tmpl",
			"tab-content-rules.tmpl",
			"tab-coverage.tmpl",
			"tab-dependencies-rules.tmpl",
			"tab-functions-rules.tmpl",
			"tab-naming-rules.tmpl",
		}

		templates := resolveTemplates()

		var resultTemplateFiles []string
		for _, tmpl := range templates.Templates() {
			if strings.HasSuffix(tmpl.Name(), ".tmpl") {
				resultTemplateFiles = append(resultTemplateFiles, tmpl.Name())
			}
		}

		assert.Zero(t, templates.Name())
		assert.ElementsMatch(t, expectedTemplateFiles, resultTemplateFiles)
	})
}
