package html

import (
	"flag"
	"github.com/fdaines/arch-go/internal/impl/model"
	"github.com/fdaines/arch-go/internal/model/result"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	"bou.ke/monkey"
)

func init() {
	testing.Init()
	flag.Parse()
}

func TestGenerateHtmlReport(t *testing.T) {
	t.Run("Calls GenerateHtmlReport function", func(t *testing.T) {
		writeCalls := 0
		var content string
		writePatch := monkey.Patch(writeReport, func(c string) {
			writeCalls++
			content = c
		})
		defer writePatch.Unpatch()

		resultData := result.Report{
			TotalPackages: 5,
			Verifications: []model.RuleVerification{},
			Summary: result.RulesSummary{
				Total:     1,
				Succeeded: 1,
				Failed:    0,
			},
		}
		expected := "\n<!DOCTYPE html>\n<html>\n\t<head>\n\t\t<link href=\"http://maxcdn.bootstrapcdn.com/bootstrap/4.1.1/css/bootstrap.min.css\" rel=\"stylesheet\" id=\"bootstrap-css\" />\n\t\t<link rel=\"stylesheet\" type=\"text/css\" href=\"report.css\" />\n\t\t<title>Arch-Go Execution Report</title>\n\t</head>\n    <body>\n<h1>Arch-Go Verification Report</h1>\n<div class=\"container summary\">\n<div class=\"row\">\n    <div class=\"col-sm-6\">\n        <h3>Rules Summary</h3>\n        <table class=\"rules-summary\">\n    <thead>\n    <tr>\n        <th style=\"width:200px;\">Rule Type</th>\n        <th style=\"width:120px;\">Summary</th>\n        <th style=\"width:100px;\">Total</th>\n        <th style=\"width:100px;\">Succeed</th>\n        <th style=\"width:100px;\">Fail</th>\n    </tr>\n    </thead>\n    <tbody>\n    \n    <tr>\n        <td>Content Rule</td>\n        <td>\n            <div class=\"result_bar\">\n                <div class=\"result_succeeded width-0\"></div>\n                <div class=\"result_legend\">0/0</div>\n            </div>\n        </td>\n        <td style=\"text-align:center;\">0</td>\n        <td style=\"text-align:center;\">0</td>\n        <td style=\"text-align:center;\">0</td>\n    </tr>\n    \n    <tr>\n        <td>Dependency Rule</td>\n        <td>\n            <div class=\"result_bar\">\n                <div class=\"result_succeeded width-0\"></div>\n                <div class=\"result_legend\">0/0</div>\n            </div>\n        </td>\n        <td style=\"text-align:center;\">0</td>\n        <td style=\"text-align:center;\">0</td>\n        <td style=\"text-align:center;\">0</td>\n    </tr>\n    \n    <tr>\n        <td>Function Rule</td>\n        <td>\n            <div class=\"result_bar\">\n                <div class=\"result_succeeded width-0\"></div>\n                <div class=\"result_legend\">0/0</div>\n            </div>\n        </td>\n        <td style=\"text-align:center;\">0</td>\n        <td style=\"text-align:center;\">0</td>\n        <td style=\"text-align:center;\">0</td>\n    </tr>\n    \n    <tr>\n        <td>Naming Rule</td>\n        <td>\n            <div class=\"result_bar\">\n                <div class=\"result_succeeded width-0\"></div>\n                <div class=\"result_legend\">0/0</div>\n            </div>\n        </td>\n        <td style=\"text-align:center;\">0</td>\n        <td style=\"text-align:center;\">0</td>\n        <td style=\"text-align:center;\">0</td>\n    </tr>\n    \n    </tbody>\n</table>\n</div>\n    <div class=\"col-sm-3 threshold-badges\">\n    <div class=\"badge-progress badge-\">\n        <div class=\"badge-header\">\n            <div class=\"rate\">0%</div>\n        </div>\n        <div class=\"badge-progressbar\">\n            <span data-fill=\"0%\" style=\"width: 0%;\"></span>\n        </div>\n        <div class=\"badge-footer\">\n            <h3>Compliance Rate</h3>\n            <span>0/0 rules were successfully verified</span>\n            <span>[Threshold: 0%]</span>\n        </div>\n    </div>\n</div>\n<div class=\"col-sm-3 threshold-badges\">\n    <div class=\"badge-progress badge-\">\n        <div class=\"badge-header\">\n            <div class=\"rate\">0%</div>\n        </div>\n        <div class=\"badge-progressbar\">\n            <span data-fill=\"0%\" style=\"width: 0%;\"></span>\n        </div>\n        <div class=\"badge-footer\">\n            <h3>Coverage Rate</h3>\n            \n                <span>All the packages were considered by at least one rule.</span>\n            \n            <span>[Threshold: 0%]</span>\n        </div>\n    </div>\n</div>\n</div>\n</div>\n<br/>\n<br/>\n<div class=\"container\">\n    <div class=\"page-header\">\n        <h3>Verification Details</h3>\n    </div>\n    <div class=\"row\">\n        <div class=\"col-md-12\">\n            <nav>\n                <div class=\"nav nav-tabs\" id=\"nav-tab\" role=\"tablist\">\n                    <a class=\"nav-item nav-link active\" id=\"nav-compliance-tab\" data-toggle=\"tab\" href=\"#nav-compliance\" role=\"tab\" aria-controls=\"nav-compliance\" aria-selected=\"true\">Compliance</a>\n                    <a class=\"nav-item nav-link\" id=\"nav-coverage-tab\" data-toggle=\"tab\" href=\"#nav-coverage\" role=\"tab\" aria-controls=\"nav-coverage\" aria-selected=\"false\">Coverage</a>\n                </div>\n            </nav>\n            <div class=\"tab-content\" id=\"nav-tabContent\">\n                <div class=\"tab-pane fade show active\" id=\"nav-compliance\" role=\"tabpanel\" aria-labelledby=\"nav-compliance-tab\">\n                    <div class=\"container\">\n                        <h4>Rule Details</h4>\n                        <table class=\"rule-details\" border=\"1\" frame=\"void\" rules=\"rows\">\n    <thead>\n    <tr>\n        <th style=\"width:200px;\">Rule Type</th>\n        <th>Rule Description</th>\n        <th style=\"width:100px;\">Result</th>\n    </tr>\n    </thead>\n    <tbody>\n    \n    </tbody>\n</table>\n</div>\n                </div>\n                <div class=\"tab-pane fade\" id=\"nav-coverage\" role=\"tabpanel\" aria-labelledby=\"nav-coverage-tab\">\n                    <div class=\"container\">\n                        <h4>Uncovered Packages: 0</h4>\n                        \n                        <span>All the packages were considered by at least one rule.</span>\n                        \n                    </div>\n                </div>\n            </div>\n        </div>\n    </div>\n\n</div>\n<hr/>\n        Report generated by <a href='http://arch-go.org'>Arch-Go</a> v1.3.3\n        <script src=\"http://cdnjs.cloudflare.com/ajax/libs/jquery/3.2.1/jquery.min.js\"></script>\n        <script src=\"http://maxcdn.bootstrapcdn.com/bootstrap/4.1.1/js/bootstrap.min.js\"></script>\n    </body>\n</html>\n"

		GenerateHtmlReport(resultData)

		assert.Equal(t, 1, writeCalls)
		assert.Equal(t, content, expected)
		assert.Contains(t, content, expected)
	})

	t.Run("Calls writeReport function", func(t *testing.T) {
		statPatch := monkey.Patch(os.Stat, func(n string) (os.FileInfo, error) {
			return nil, nil
		})
		defer statPatch.Unpatch()
		var filename string
		var data string
		var perm os.FileMode
		writePatch := monkey.Patch(os.WriteFile, func(f string, d []byte, p os.FileMode) error {
			filename = f
			data = string(d)
			perm = p
			return nil
		})
		defer writePatch.Unpatch()

		content := "foobar"

		writeReport(content)

		assert.Equal(t, ".arch-go/report.html", filename)
		assert.Equal(t, content, data)
		assert.Equal(t, os.FileMode(0x1a4), perm)
	})
}
