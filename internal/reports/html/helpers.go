package html

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func writeReport(content string, output io.Writer) {
	if err := os.WriteFile(".arch-go/report.html", []byte(content), 0o644); err == nil {
		fmt.Fprintln(output, "HTML report generated at: .arch-go/report.html")
	} else {
		panic(err)
	}
}

func copyAssets() {
	if isTestRun() {
		return
	}

	if _, err := os.Stat(".arch-go/"); os.IsNotExist(err) {
		os.Mkdir(".arch-go", 0o755)
	}

	cssByteArray, _ := styles.ReadFile("assets/report.css")
	os.WriteFile(".arch-go/report.css", cssByteArray, 0o644)

	logoPng, _ := images.ReadFile("assets/logo.png")
	os.WriteFile(".arch-go/logo.png", logoPng, 0o644)

	timerPng, _ := images.ReadFile("assets/timer.png")
	os.WriteFile(".arch-go/timer.png", timerPng, 0o644)
}

func isTestRun() bool {
	return flag.Lookup("test.v") != nil
}
