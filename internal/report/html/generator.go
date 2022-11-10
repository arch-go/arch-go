package html

import (
	"fmt"
	"github.com/fdaines/arch-go/internal/impl/model"
	"github.com/fdaines/arch-go/internal/model/result"
	"os"
	"strings"
)

func GenerateHtmlReport(verifications []model.RuleVerification, summary result.RulesSummary) {
	html := generateHtml(summary, verifications)
	writeReport(html)
}

func generateHtml(summary result.RulesSummary, verifications []model.RuleVerification) string {
	rules := ruleList(summary)
	html := strings.Replace(htmlTemplate(), "[RULE_LIST]", rules, 1)
	details := ruleDetails(verifications)
	html = strings.Replace(html, "[RULE_DETAILS]", details, 1)
	uncoveredPackages := uncoveredPackages(summary)
	html = strings.Replace(html, "[UNCOVERED_PACKAGES]", uncoveredPackages, 1)
	return html
}

func writeReport(content string) {
	if _, err := os.Stat(".arch-go/"); os.IsNotExist(err) {
		os.Mkdir(".arch-go", 0755)
	}
	htmlByteArray := []byte(content)
	err := os.WriteFile(".arch-go/report.html", htmlByteArray, 0644)
	if err == nil {
		fmt.Println("HTML report generated at: .arch-go/report.html")
	} else {
		panic(err)
	}
}
