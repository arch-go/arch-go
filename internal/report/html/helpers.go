package html

import (
	"flag"
	"fmt"
	"os"
)

func writeReport(content string) {
	htmlByteArray := []byte(content)
	err := os.WriteFile(".arch-go/report.html", htmlByteArray, 0644)
	if err == nil {
		fmt.Println("HTML report generated at: .arch-go/report.html")
	} else {
		panic(err)
	}
}

func copyAssets() {
	if isTestRun() {
		return
	}
	if _, err := os.Stat(".arch-go/"); os.IsNotExist(err) {
		os.Mkdir(".arch-go", 0755)
	}
	css, _ := os.ReadFile("./internal/report/html/templates/report.css")
	cssByteArray := []byte(css)
	os.WriteFile(".arch-go/report.css", cssByteArray, 0644)
}

func resolveTemplateRelativePath(template string) string {
	if isTestRun() {
		return "./templates/" + template
	}
	return "./internal/report/html/templates/" + template
}

func isTestRun() bool {
	return flag.Lookup("test.v") != nil
}
