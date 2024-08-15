package json

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/arch-go/arch-go/internal/reports/model"
)

func GenerateJsonReport(report *model.Report, output io.Writer) {
	bytes, err := generateJson(report)
	if err != nil {
		panic(err)
	}
	writeReport(bytes, output)
}

func generateJson(report *model.Report) ([]byte, error) {
	return json.Marshal(report)
}

func writeReport(content []byte, output io.Writer) {
	if err := os.WriteFile(".arch-go/report.json", content, 0o600); err == nil {
		fmt.Fprintln(output, "JSON report generated at: .arch-go/report.json")
	} else {
		panic(err)
	}
}