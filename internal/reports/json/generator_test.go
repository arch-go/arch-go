package json

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/v2/internal/reports/model"
	"github.com/arch-go/arch-go/v2/internal/utils/values"
)

func TestGenerateJsonReport(t *testing.T) {
	t.Run("generateJson with empty report", func(t *testing.T) {
		var filename, fileContent string

		output := bytes.NewBufferString("")

		mock := gomonkey.ApplyFunc(
			os.WriteFile,
			func(name string, data []byte, _ os.FileMode) error {
				filename = name
				fileContent = string(data)

				return nil
			})
		defer mock.Reset()

		report := &model.Report{}
		expectedFilename := ".arch-go/report.json"
		expectedFileContent := `{"version":"","summary":null,"compliance":{"pass":false,"rate":0,"threshold":null,"total":0,"passed":0,"failed":0},"coverage":{"pass":false,"rate":0,"threshold":null}}`

		GenerateReport(report, output)

		assert.Equal(t, expectedFilename, filename)
		assert.Equal(t, expectedFileContent, fileContent)
	})

	t.Run("generateJson with full report", func(t *testing.T) {
		var filename, fileContent string

		output := bytes.NewBufferString("")

		mock := gomonkey.ApplyFunc(
			os.WriteFile,
			func(name string, data []byte, _ os.FileMode) error {
				filename = name
				fileContent = string(data)

				return nil
			})
		defer mock.Reset()

		report := &model.Report{
			Summary: &model.Summary{
				Pass:     true,
				Time:     time.Time{},
				Duration: time.Duration(12345678),
			},
			Compliance: model.Compliance{
				Total:     100,
				Passed:    87,
				Failed:    13,
				Rate:      87,
				Threshold: values.GetIntRef(100),
				Pass:      false,
				Details: &model.ReportDetails{
					DependenciesVerificationDetails: model.Verification{
						Total:  1,
						Passed: 1,
						Failed: 0,
						Details: []model.VerificationDetails{
							{
								Rule:   "foobar rule dep",
								Pass:   true,
								Total:  1,
								Passed: 1,
								Failed: 0,
								PackageDetails: []model.PackageDetails{
									{
										Package: "my-package",
										Pass:    true,
									},
								},
							},
						},
					},
					FunctionsVerificationDetails: model.Verification{
						Total:  1,
						Passed: 1,
						Failed: 0,
						Details: []model.VerificationDetails{
							{
								Rule:   "foobar rule fn",
								Pass:   true,
								Total:  1,
								Passed: 1,
								Failed: 0,
								PackageDetails: []model.PackageDetails{
									{
										Package: "my-package",
										Pass:    true,
									},
								},
							},
						},
					},
					ContentsVerificationDetails: model.Verification{
						Total:  1,
						Passed: 1,
						Failed: 0,
						Details: []model.VerificationDetails{
							{
								Rule:   "foobar rule cn",
								Pass:   true,
								Total:  1,
								Passed: 1,
								Failed: 0,
								PackageDetails: []model.PackageDetails{
									{
										Package: "my-package",
										Pass:    true,
									},
								},
							},
						},
					},
					NamingVerificationDetails: model.Verification{
						Total:  1,
						Passed: 0,
						Failed: 1,
						Details: []model.VerificationDetails{
							{
								Rule:   "foobar rule nm",
								Pass:   false,
								Total:  1,
								Passed: 0,
								Failed: 1,
								PackageDetails: []model.PackageDetails{
									{
										Package: "my-package",
										Pass:    false,
										Details: []string{"foobar message"},
									},
								},
							},
						},
					},
				},
			},
			Coverage: model.Coverage{
				Rate:      80,
				Threshold: values.GetIntRef(60),
				Pass:      true,
				Uncovered: []string{"foobar"},
				Details: []model.CoverageDetails{
					{
						Package:           "foobar",
						ContentsRules:     0,
						DependenciesRules: 0,
						FunctionsRules:    0,
						NamingRules:       0,
						Covered:           false,
					},
					{
						Package:           "my-package1",
						ContentsRules:     1,
						DependenciesRules: 1,
						FunctionsRules:    1,
						NamingRules:       1,
						Covered:           true,
					},
					{
						Package:           "my-package2",
						ContentsRules:     1,
						DependenciesRules: 1,
						FunctionsRules:    1,
						NamingRules:       1,
						Covered:           true,
					},
					{
						Package:           "my-package3",
						ContentsRules:     1,
						DependenciesRules: 1,
						FunctionsRules:    1,
						NamingRules:       1,
						Covered:           true,
					},
					{
						Package:           "my-package4",
						ContentsRules:     1,
						DependenciesRules: 1,
						FunctionsRules:    1,
						NamingRules:       1,
						Covered:           true,
					},
				},
			},
		}

		expectedFilename := ".arch-go/report.json"
		expectedFileContent := `{"version":"","summary":{"pass":true,"timestamp":"0001-01-01T00:00:00Z","duration":12345678},"compliance":{"pass":false,"rate":87,"threshold":100,"total":100,"passed":87,"failed":13,"details":{"dependenciesRules":{"total":1,"passed":1,"failed":0,"details":[{"rule":"foobar rule dep","pass":true,"total":1,"passed":1,"failed":0,"packageDetails":[{"package":"my-package","pass":true}]}]},"functionsRules":{"total":1,"passed":1,"failed":0,"details":[{"rule":"foobar rule fn","pass":true,"total":1,"passed":1,"failed":0,"packageDetails":[{"package":"my-package","pass":true}]}]},"contentsRules":{"total":1,"passed":1,"failed":0,"details":[{"rule":"foobar rule cn","pass":true,"total":1,"passed":1,"failed":0,"packageDetails":[{"package":"my-package","pass":true}]}]},"namingRules":{"total":1,"passed":0,"failed":1,"details":[{"rule":"foobar rule nm","pass":false,"total":1,"passed":0,"failed":1,"packageDetails":[{"package":"my-package","pass":false,"details":["foobar message"]}]}]}}},"coverage":{"pass":true,"rate":80,"threshold":60,"uncoveredPackages":["foobar"],"details":[{"package":"foobar","contentsRules":0,"dependenciesRules":0,"functionsRules":0,"namingRules":0,"covered":false},{"package":"my-package1","contentsRules":1,"dependenciesRules":1,"functionsRules":1,"namingRules":1,"covered":true},{"package":"my-package2","contentsRules":1,"dependenciesRules":1,"functionsRules":1,"namingRules":1,"covered":true},{"package":"my-package3","contentsRules":1,"dependenciesRules":1,"functionsRules":1,"namingRules":1,"covered":true},{"package":"my-package4","contentsRules":1,"dependenciesRules":1,"functionsRules":1,"namingRules":1,"covered":true}]}}`

		GenerateReport(report, output)

		assert.Equal(t, expectedFilename, filename)
		assert.Equal(t, expectedFileContent, fileContent)
	})
}
