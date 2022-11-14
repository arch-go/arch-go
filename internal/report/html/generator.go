package html

import (
	"bytes"
	"fmt"
	"github.com/fdaines/arch-go/internal/common"
	"github.com/fdaines/arch-go/internal/model/result"
	"html/template"
	"os"
)

func GenerateHtmlReport(resultData result.Report) {
	html := generateHtml(resultData)
	copyAssets()
	writeReport(html)
}

func generateHtml(report result.Report) string {
	htmlReport := mapToHtmlReport(report)

	var processed bytes.Buffer
	templates := resolveTemplates()
	templates.ExecuteTemplate(&processed, "report", htmlReport)

	return string(processed.Bytes())
}

func resolveTemplates() *template.Template {
	allTemplateFiles := []string{
		"content.tmpl",
		"rules-summary.tmpl",
		"rules-details.tmpl",
		"compliance-rate.tmpl",
		"coverage-rate.tmpl",
		"summary.tmpl",
		"footer.tmpl",
		"header.tmpl",
		"report.tmpl",
	}

	var allTemplatePaths []string
	for _, tmpl := range allTemplateFiles {
		allTemplatePaths = append(allTemplatePaths, "./internal/report/html/templates/"+tmpl)
	}

	templates, _ := template.New("").Funcs(
		template.FuncMap{
			"inc": func(number int) int {
				return 1 + number
			},
		}).ParseFiles(allTemplatePaths...)
	return templates
}

func mapToHtmlReport(report result.Report) HtmlReport {
	htmlReport := HtmlReport{
		Version: common.Version,
	}
	resolveRulesSummary(report, &htmlReport)
	resolveRulesDetails(report, &htmlReport)
	resolveComplianceResults(report, &htmlReport)
	resolveCoverageResults(report, &htmlReport)

	return htmlReport
}

func resolveRulesDetails(report result.Report, htmlReport *HtmlReport) {
	var details []RuleDetails
	for _, verification := range report.Verifications {
		ruleDetails := RuleDetails{
			Type:        verification.Type(),
			Description: verification.Name(),
			Color:       "red",
			Status:      "Fail",
		}
		if verification.Status() {
			ruleDetails.Color = "green"
			ruleDetails.Status = "Succeed"
		}
		for _, v := range verification.GetVerifications() {
			ruleVerification := RuleVerification{
				Package: v.Package.Path,
				Status:  "Fail",
				Color:   "red",
				Details: v.Details,
			}
			if v.Passes {
				ruleVerification.Status = "Succeed"
				ruleVerification.Color = "green"
			}
			ruleDetails.Verifications = append(ruleDetails.Verifications, ruleVerification)
		}
		details = append(details, ruleDetails)
	}
	htmlReport.RulesDetails = details
}

func resolveRulesSummary(report result.Report, htmlReport *HtmlReport) {
	var rules []RuleSummary

	ruleTypes := []string{"DependenciesRule", "FunctionsRule", "ContentRule", "CycleRule", "NamingRule"}
	for _, rt := range ruleTypes {
		var ratio int
		ruleSummary := report.Summary.Details[rt]
		if ruleSummary.Total > 0 {
			ratio = 100 * ruleSummary.Succeeded / ruleSummary.Total
		}
		rules = append(rules, RuleSummary{
			Type:      rt,
			Succeeded: ruleSummary.Succeeded,
			Failed:    ruleSummary.Failed,
			Total:     ruleSummary.Total,
			Ratio:     ratio,
		})
	}
	htmlReport.RulesSummary = rules
}

func resolveComplianceResults(report result.Report, htmlReport *HtmlReport) {
	if report.Summary.ComplianceThreshold != nil {
		htmlReport.ComplianceResult = ComplianceResult{
			Rate:      report.Summary.ComplianceThreshold.Rate,
			Succeeded: report.Summary.Succeeded,
			Total:     report.Summary.Total,
			Threshold: report.Summary.ComplianceThreshold.Threshold,
			Color:     "red",
		}
		if report.Summary.ComplianceThreshold.Status == "Pass" {
			htmlReport.ComplianceResult.Color = "green"
		}
	}
}

func resolveCoverageResults(report result.Report, htmlReport *HtmlReport) {
	if report.Summary.CoverageThreshold != nil {
		htmlReport.CoverageResult = CoverageResult{
			Rate:      report.Summary.CoverageThreshold.Rate,
			Uncovered: len(report.Summary.CoverageThreshold.Violations),
			Total:     report.TotalPackages,
			Threshold: report.Summary.CoverageThreshold.Threshold,
			Color:     "red",
		}
		if report.Summary.CoverageThreshold.Status == "Pass" {
			htmlReport.CoverageResult.Color = "green"
		}
		htmlReport.UncoveredPackages = report.Summary.CoverageThreshold.Violations
	}
}

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
	if _, err := os.Stat(".arch-go/"); os.IsNotExist(err) {
		os.Mkdir(".arch-go", 0755)
	}
	css, _ := os.ReadFile("./internal/report/html/templates/report.css")
	cssByteArray := []byte(css)
	os.WriteFile(".arch-go/report.css", cssByteArray, 0644)
}
