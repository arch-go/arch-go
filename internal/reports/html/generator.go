package html

import (
	"bytes"
	"fmt"
	"io"

	"github.com/arch-go/arch-go/internal/reports/model"
)

func GenerateHTMLReport(report *model.Report, output io.Writer) {
	html := generateHTML(report, output)

	copyAssets()
	writeReport(html, output)
}

func generateHTML(report *model.Report, output io.Writer) string {
	var processed bytes.Buffer

	templates := resolveTemplates()

	if err := templates.ExecuteTemplate(&processed, "report", report); err != nil {
		fmt.Fprintf(output, "Error: %+v\n", err)
	}

	return processed.String()
}
