package html

import (
	"flag"
	"fmt"
	"github.com/fdaines/arch-go/internal/common"
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
	cssByteArray, _ := os.ReadFile(archGoModulePath() + "/internal/report/html/templates/report.css")
	os.WriteFile(".arch-go/report.css", cssByteArray, 0644)
}

func resolveTemplateRelativePath(template string) string {
	if isTestRun() {
		return "./templates/" + template
	}
	return archGoModulePath() + "/internal/report/html/templates/" + template
}

func archGoModulePath() string {
	return os.Getenv("GOPATH") + "/pkg/mod/github.com/fdaines/arch-go@v" + common.Version
}

func isTestRun() bool {
	return flag.Lookup("test.v") != nil
}
