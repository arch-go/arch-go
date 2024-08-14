package html

import (
	"html/template"
)

func resolveTemplates() *template.Template {
	templates, _ := template.New("").Funcs(
		template.FuncMap{
			"inc":            increment(),
			"ratio":          calculateRatio(),
			"formatDateTime": formatDateTime(),
			"formatDate":     formatDate(),
			"formatTime":     formatTime(),
			"toHumanTime":    toHumanTime(),
			"toYesNo":        toYesNo(),
			"toPassFail":     toPassFail(),
		}).ParseFS(templateFiles, "templates/*.tmpl")

	return templates
}
