package json

import (
	"testing"
	"time"

	"github.com/arch-go/arch-go/internal/reports/model"
	"github.com/arch-go/arch-go/internal/utils/values"

	"github.com/stretchr/testify/assert"
)

func TestGenerateJsonReport(t *testing.T) {
	t.Run("generateJson with empty report", func(t *testing.T) {
		report := &model.Report{}
		expected := `{"version":"","summary":null,"compliance":{"pass":false,"rate":0,"threshold":null,"total":0,"passed":0,"failed":0},"coverage":{"pass":false,"rate":0,"threshold":null}}`

		bytes, err := generateJson(report)
		assert.Nil(t, err)
		assert.Equal(t, expected, string(bytes))
	})

	t.Run("generateJson with full report", func(t *testing.T) {
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
		expected := `{"version":"","summary":{"pass":true,"timestamp":"0001-01-01T00:00:00Z","duration":12345678},"compliance":{"pass":false,"rate":87,"threshold":100,"total":100,"passed":87,"failed":13,"details":{"dependencies_rules":{"total":1,"passed":1,"failed":0,"details":[{"rule":"foobar rule dep","pass":true,"total":1,"passed":1,"failed":0,"package_details":[{"package":"my-package","pass":true}]}]},"functions_rules":{"total":1,"passed":1,"failed":0,"details":[{"rule":"foobar rule fn","pass":true,"total":1,"passed":1,"failed":0,"package_details":[{"package":"my-package","pass":true}]}]},"contents_rules":{"total":1,"passed":1,"failed":0,"details":[{"rule":"foobar rule cn","pass":true,"total":1,"passed":1,"failed":0,"package_details":[{"package":"my-package","pass":true}]}]},"naming_rules":{"total":1,"passed":0,"failed":1,"details":[{"rule":"foobar rule nm","pass":false,"total":1,"passed":0,"failed":1,"package_details":[{"package":"my-package","pass":false,"details":["foobar message"]}]}]}}},"coverage":{"pass":true,"rate":80,"threshold":60,"uncovered_packages":["foobar"],"details":[{"package":"foobar","contents_rules":0,"dependencies_rules":0,"functions_rules":0,"naming_rules":0,"covered":false},{"package":"my-package1","contents_rules":1,"dependencies_rules":1,"functions_rules":1,"naming_rules":1,"covered":true},{"package":"my-package2","contents_rules":1,"dependencies_rules":1,"functions_rules":1,"naming_rules":1,"covered":true},{"package":"my-package3","contents_rules":1,"dependencies_rules":1,"functions_rules":1,"naming_rules":1,"covered":true},{"package":"my-package4","contents_rules":1,"dependencies_rules":1,"functions_rules":1,"naming_rules":1,"covered":true}]}}`

		bytes, err := generateJson(report)
		assert.Nil(t, err)
		assert.Equal(t, expected, string(bytes))
	})
}
