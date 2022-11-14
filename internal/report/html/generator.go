package html

import (
	"bytes"
	"fmt"
	"github.com/fdaines/arch-go/internal/model/result"
)

func GenerateHtmlReport(resultData result.Report) {
	html := generateHtml(resultData)
	copyAssets()
	writeReport(html)
}

func generateHtml(report result.Report) string {
	htmlReport := mapToHtmlReport(report)

	var processed bytes.Buffer
	templates := resolveTemplates()

	err := templates.ExecuteTemplate(&processed, "report", htmlReport)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}
	return string(processed.Bytes())
}
