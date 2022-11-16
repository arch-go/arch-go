package html

import (
	"html/template"
)

func resolveTemplates() *template.Template {
	templates, _ := template.New("").Funcs(
		template.FuncMap{
			"inc": func(number int) int {
				return 1 + number
			},
		}).ParseFS(templateFiles, "templates/*.tmpl")

	return templates
}
