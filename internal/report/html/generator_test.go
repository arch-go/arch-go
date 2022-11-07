package html

import (
	"github.com/fdaines/arch-go/internal/impl/model"
	"github.com/fdaines/arch-go/internal/model/result"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"

	"bou.ke/monkey"
)

func TestGenerateHtmlReport(t *testing.T) {
	t.Run("Calls GenerateHtmlReport function", func(t *testing.T) {
		writeCalls := 0
		var content string
		writePatch := monkey.Patch(writeReport, func(c string) {
			writeCalls++
			content = c
		})
		defer writePatch.Unpatch()

		verifications := []model.RuleVerification{}
		summary := result.RulesSummary{
			Total:     1,
			Succeeded: 1,
			Failed:    0,
		}
		expected := "<!DOCTYPE html>\n<html>\n<head>\n    <link rel=\"stylesheet\" type=\"text/css\" href=\"style.css\">\n<style>\n\nblockquote,body,div,html,p,pre,span{margin:0;padding:0;border:0;outline:0;font-weight:inherit;font-style:inherit;font-size:100%;font-family:inherit;vertical-align:baseline}body{line-height:1;color:#000;background:#fff;margin-left:20px}.result_bar{display:inline-block;height:1.1em;width:130px;background:#faa;margin:0 5px;vertical-align:middle;border:1px solid #aaa;position:relative}.result_succeeded{display:inline-block;height:100%;background:#dfd;float:left}.result_legend{position:absolute;height:100%;width:100%;left:0;top:0;text-align:center}.width-1{width:1%}.width-2{width:2%}.width-3{width:3%}.width-4{width:4%}.width-5{width:5%}.width-6{width:6%}.width-7{width:7%}.width-8{width:8%}.width-9{width:9%}.width-10{width:10%}.width-11{width:11%}.width-12{width:12%}.width-13{width:13%}.width-14{width:14%}.width-15{width:15%}.width-16{width:16%}.width-17{width:17%}.width-18{width:18%}.width-19{width:19%}.width-20{width:20%}.width-21{width:21%}.width-22{width:22%}.width-23{width:23%}.width-24{width:24%}.width-25{width:25%}.width-26{width:26%}.width-27{width:27%}.width-28{width:28%}.width-29{width:29%}.width-30{width:30%}.width-31{width:31%}.width-32{width:32%}.width-33{width:33%}.width-34{width:34%}.width-35{width:35%}.width-36{width:36%}.width-37{width:37%}.width-38{width:38%}.width-39{width:39%}.width-40{width:40%}.width-41{width:41%}.width-42{width:42%}.width-43{width:43%}.width-44{width:44%}.width-45{width:45%}.width-46{width:46%}.width-47{width:47%}.width-48{width:48%}.width-49{width:49%}.width-50{width:50%}.width-51{width:51%}.width-52{width:52%}.width-53{width:53%}.width-54{width:54%}.width-55{width:55%}.width-56{width:56%}.width-57{width:57%}.width-58{width:58%}.width-59{width:59%}.width-60{width:60%}.width-61{width:61%}.width-62{width:62%}.width-63{width:63%}.width-64{width:64%}.width-65{width:65%}.width-66{width:66%}.width-67{width:67%}.width-68{width:68%}.width-69{width:69%}.width-70{width:70%}.width-71{width:71%}.width-72{width:72%}.width-73{width:73%}.width-74{width:74%}.width-75{width:75%}.width-76{width:76%}.width-77{width:77%}.width-78{width:78%}.width-79{width:79%}.width-80{width:80%}.width-81{width:81%}.width-82{width:82%}.width-83{width:83%}.width-84{width:84%}.width-85{width:85%}.width-86{width:86%}.width-87{width:87%}.width-88{width:88%}.width-89{width:89%}.width-90{width:90%}.width-91{width:91%}.width-92{width:92%}.width-93{width:93%}.width-94{width:94%}.width-95{width:95%}.width-96{width:96%}.width-97{width:97%}.width-98{width:98%}.width-99{width:99%}.width-100{width:100%}\n\n</style>\n</head>\n<body>\n\n<h1>Arch-Go Verification Report</h1>\n\n\n<h3>Rules Summary</h3>\n<table>\n    <thead>\n        <tr>\n            <th style=\"width:200px;\">Rule Type</th>\n            <th style=\"width:120px;\">Summary</th>\n            <th style=\"width:100px;\">Total</th>\n            <th style=\"width:100px;\">Succeed</th>\n\t\t\t<th style=\"width:100px;\">Fail</th>\n        </tr>\n    </thead>\n    <tbody>\n\t\t<tr>\n\t<td>DependenciesRule</td>\n\t<td>\n\t\t<div class=\"result_bar\">\n\t\t\t<div class=\"result_succeeded width-0\"></div>\n\t\t\t<div class=\"result_legend\">0/0</div>\n\t\t</div>\n\t</td>\n\t<td style=\"text-align:center;\">0</td>\n\t<td style=\"text-align:center;\">0</td>\n\t<td style=\"text-align:center;\">0</td>\n</tr><tr>\n\t<td>FunctionsRule</td>\n\t<td>\n\t\t<div class=\"result_bar\">\n\t\t\t<div class=\"result_succeeded width-0\"></div>\n\t\t\t<div class=\"result_legend\">0/0</div>\n\t\t</div>\n\t</td>\n\t<td style=\"text-align:center;\">0</td>\n\t<td style=\"text-align:center;\">0</td>\n\t<td style=\"text-align:center;\">0</td>\n</tr><tr>\n\t<td>ContentRule</td>\n\t<td>\n\t\t<div class=\"result_bar\">\n\t\t\t<div class=\"result_succeeded width-0\"></div>\n\t\t\t<div class=\"result_legend\">0/0</div>\n\t\t</div>\n\t</td>\n\t<td style=\"text-align:center;\">0</td>\n\t<td style=\"text-align:center;\">0</td>\n\t<td style=\"text-align:center;\">0</td>\n</tr><tr>\n\t<td>CycleRule</td>\n\t<td>\n\t\t<div class=\"result_bar\">\n\t\t\t<div class=\"result_succeeded width-0\"></div>\n\t\t\t<div class=\"result_legend\">0/0</div>\n\t\t</div>\n\t</td>\n\t<td style=\"text-align:center;\">0</td>\n\t<td style=\"text-align:center;\">0</td>\n\t<td style=\"text-align:center;\">0</td>\n</tr><tr>\n\t<td>NamingRule</td>\n\t<td>\n\t\t<div class=\"result_bar\">\n\t\t\t<div class=\"result_succeeded width-0\"></div>\n\t\t\t<div class=\"result_legend\">0/0</div>\n\t\t</div>\n\t</td>\n\t<td style=\"text-align:center;\">0</td>\n\t<td style=\"text-align:center;\">0</td>\n\t<td style=\"text-align:center;\">0</td>\n</tr>\n\t</tbody>\n</table>\n\n\n<br/>\n\n<h3>Rules Details</h3>\n<table border=\"1\" frame=\"void\"\" rules=\"rows\">\n    <thead>\n        <tr>\n            <th style=\"width:200px;\">Rule Type</th>\n            <th>Rule Description</th>\n            <th style=\"width:100px;\">Result</th>\n        </tr>\n    </thead>\n    <tbody>\n\t\t\n\t</tbody>\n</table>\n\n<br/>\n<hr/>\nReport generated by <a href='http://arch-go.org'>Arch-Go</a> v1.0.2\n</body>\n</html>"

		GenerateHtmlReport(verifications, summary)

		assert.Equal(t, 1, writeCalls)
		assert.True(t, strings.Contains(content, expected))
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
