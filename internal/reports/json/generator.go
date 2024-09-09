package json

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/arch-go/arch-go/internal/reports/model"
)

func GenerateReport(report *model.Report, output io.Writer) {
	bytes, err := json.Marshal(report)
	if err != nil {
		panic(err)
	}

	writeReport(bytes, output)
}

func writeReport(content []byte, output io.Writer) {
	_ = os.Mkdir(".arch-go", 0o755)

	if err := os.WriteFile(".arch-go/report.json", content, 0o600); err == nil {
		fmt.Fprintln(output, "JSON report generated at: .arch-go/report.json")
	} else {
		panic(err)
	}
}
