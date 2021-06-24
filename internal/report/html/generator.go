package html

import (
	"fmt"
	"github.com/fdaines/arch-go/internal/impl/model"
	"github.com/fdaines/arch-go/internal/model/result"
	"io/ioutil"
	"os"
	"strings"
)

func GenerateHtmlReport(verifications []model.RuleVerification, summary result.RulesSummary) {
	rules := ruleList(summary)
	html := strings.Replace(htmlTemplate(), "[RULE_LIST]", rules, 1)
	writeReport(html)
}

func writeReport(content string) {
	if _, err := os.Stat(".arch-go/"); os.IsNotExist(err) {
		os.Mkdir(".arch-go", 0755)
	}
	htmlByteArray := []byte(content)
	err := ioutil.WriteFile(".arch-go/report.html", htmlByteArray, 0644)
	if err == nil {
		fmt.Println("HTML report generated at: .arch-go/report.html")
	} else {
		panic(err)
	}
}