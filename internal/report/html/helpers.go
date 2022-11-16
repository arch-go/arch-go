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
	cssByteArray, _ := styles.ReadFile("assets/report.css")
	os.WriteFile(".arch-go/report.css", cssByteArray, 0644)
	logoPng, _ := images.ReadFile("assets/logo.png")
	os.WriteFile(".arch-go/logo.png", logoPng, 0644)
}

func isTestRun() bool {
	return flag.Lookup("test.v") != nil
}
