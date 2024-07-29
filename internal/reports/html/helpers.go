package html

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func writeReport(content string, output io.Writer) {
	if err := os.WriteFile(".arch-go/report.html", []byte(content), 0o600); err == nil {
		_, _ = fmt.Fprintln(output, "HTML report generated at: .arch-go/report.html")
	} else {
		panic(err)
	}
}

func copyAssets() {
	if isTestRun() {
		return
	}

	if _, err := os.Stat(".arch-go/"); os.IsNotExist(err) {
		_ = os.Mkdir(".arch-go", 0o755)
	}

	cssByteArray, _ := styles.ReadFile("assets/report.css")
	_ = os.WriteFile(".arch-go/report.css", cssByteArray, 0o600)

	logoPng, _ := images.ReadFile("assets/logo.png")
	_ = os.WriteFile(".arch-go/logo.png", logoPng, 0o600)

	timerPng, _ := images.ReadFile("assets/timer.png")
	_ = os.WriteFile(".arch-go/timer.png", timerPng, 0o600)
}

func isTestRun() bool {
	return flag.Lookup("test.v") != nil
}
