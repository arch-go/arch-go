package html

import (
	"fmt"
	"html/template"
	"time"
)

func resolveTemplates() *template.Template {
	templates, _ := template.New("").Funcs(
		template.FuncMap{
			"inc": func(number int) int {
				return 1 + number
			},
			"passes": func(status string) bool {
				return status == "PASS" || status == "YES"
			},
			"ratio": func(num int, den int) int {
				if den == 0 {
					return 100
				}
				return 100 * num / den
			},
			"formatDate": func(t time.Time) string {
				return t.Format("2006/01/02 15:04:05")
			},
			"toHumanTime": func(d time.Duration) string {
				if d.Seconds() > 0.9 {
					return fmt.Sprintf("%v[s]", d.Seconds())
				}
				if d.Milliseconds() > 0 {
					return fmt.Sprintf("%v[ms]", d.Milliseconds())
				}
				if d.Microseconds() > 0 {
					return fmt.Sprintf("%v[μs]", d.Microseconds())
				}
				return fmt.Sprintf("%v[ns]", d.Nanoseconds())
			},
		}).ParseFS(templateFiles, "templates/*.tmpl")

	return templates
}
