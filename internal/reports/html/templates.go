package html

import (
	"html/template"
)

func resolveTemplates() *template.Template {
	templates, _ := template.New("").Funcs(
		template.FuncMap{
			"ratio":       calculateRatio(),
			"formatDate":  formatDate(),
			"formatTime":  formatTime(),
			"toHumanTime": toHumanTime(),
			"toYesNo":     toYesNo(),
			"toPassFail":  toPassFail(),
		}).ParseFS(templateFiles, "templates/*.tmpl")

	return templates
}
