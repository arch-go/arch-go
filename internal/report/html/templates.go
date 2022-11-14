package html

import (
	"html/template"
	"path/filepath"
)

func resolveTemplates() *template.Template {
	allTemplateFiles := []string{
		"content.tmpl",
		"rules-summary.tmpl",
		"rules-details.tmpl",
		"compliance-rate.tmpl",
		"coverage-rate.tmpl",
		"summary.tmpl",
		"footer.tmpl",
		"header.tmpl",
		"report.tmpl",
	}

	var allTemplatePaths []string

	for _, tmpl := range allTemplateFiles {
		absolutePath, _ := filepath.Abs(resolveTemplateRelativePath(tmpl))

		allTemplatePaths = append(allTemplatePaths, absolutePath)
	}
	templates, _ := template.New("").Funcs(
		template.FuncMap{
			"inc": func(number int) int {
				return 1 + number
			},
		}).ParseFiles(allTemplatePaths...)
	return templates
}
