package html

import (
	"bytes"
	"fmt"
	"io"

	"github.com/fdaines/arch-go/internal/reports/model"
)

func GenerateHtmlReport(report *model.Report, output io.Writer) {
	html := generateHtml(report, output)
	copyAssets()
	writeReport(html, output)
}

func generateHtml(report *model.Report, output io.Writer) string {
	var processed bytes.Buffer
	templates := resolveTemplates()

	err := templates.ExecuteTemplate(&processed, "report", report)
	if err != nil {
		fmt.Fprintf(output, "Error: %+v\n", err)
	}
	return string(processed.Bytes())
}
