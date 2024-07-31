package functions_test

import (
	"encoding/json"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/api/configuration"
	"github.com/arch-go/arch-go/internal/model"
	"github.com/arch-go/arch-go/internal/utils/values"
	"github.com/arch-go/arch-go/internal/verifications/functions"
)

func TestCheckFunctionRules(t *testing.T) {
	t.Parallel()
	t.Run("check passes", func(t *testing.T) {
		functionDetails := []*functions.FunctionDetails{
			{
				FilePath:   "filepath1",
				File:       "file01",
				Name:       "function1",
				IsPublic:   true,
				NumParams:  2,
				NumReturns: 1,
				NumLines:   5,
			},
			{
				FilePath:   "filepath2",
				File:       "file02",
				Name:       "function2",
				IsPublic:   true,
				NumParams:  0,
				NumReturns: 2,
				NumLines:   10,
			},
			{
				FilePath:   "filepath3",
				File:       "file03",
				Name:       "function3",
				IsPublic:   true,
				NumParams:  5,
				NumReturns: 2,
				NumLines:   10,
			},
		}
		mock := gomonkey.ApplyFuncReturn(functions.RetrieveFunctions, functionDetails, nil)

		defer mock.Reset()

		moduleInfo := model.ModuleInfo{
			MainPackage: "mymodule",
			Packages: []*model.PackageInfo{
				{
					Name: "foobar1",
					Path: "barfoo1",
				},
				{
					Name: "foobar2",
					Path: "barfoo2",
				},
			},
		}
		functionRules := []*configuration.FunctionsRule{
			{
				Package:                  "barfoo",
				MaxLines:                 values.GetIntRef(10),
				MaxParameters:            values.GetIntRef(5),
				MaxReturnValues:          values.GetIntRef(2),
				MaxPublicFunctionPerFile: values.GetIntRef(1),
			},
		}
		fx := functionRules[0]

		expectedResult := &functions.RulesResult{
			Passes: true,
			Results: []*functions.RuleResult{
				{
					Rule:        *fx,
					Description: "Functions in packages matching pattern 'barfoo' should have ['at most 5 parameters','at most 2 return values','at most 10 lines','no more than 1 public functions per file']",
					Passes:      true,
					Verifications: []functions.Verification{
						{
							Package: "barfoo1",
							Passes:  true,
						},
						{
							Package: "barfoo2",
							Passes:  true,
						},
					},
				},
			},
		}

		result := functions.CheckRules(moduleInfo, functionRules)
		jsonExpectedResult, _ := json.Marshal(expectedResult)
		jsonResult, _ := json.Marshal(result)

		assert.Equal(t, jsonExpectedResult, jsonResult)
	})

	t.Run("check fails", func(t *testing.T) {
		index := 1
		functionDetails1 := []*functions.FunctionDetails{
			{
				FilePath:   "filepath1",
				File:       "file01",
				Name:       "function1",
				IsPublic:   true,
				NumParams:  20,
				NumReturns: 1,
				NumLines:   5,
			},
			{
				FilePath:   "filepath2",
				File:       "file02",
				Name:       "function2",
				IsPublic:   true,
				NumParams:  0,
				NumReturns: 2,
				NumLines:   10,
			},
			{
				FilePath:   "filepath3",
				File:       "file03",
				Name:       "function3",
				IsPublic:   true,
				NumParams:  5,
				NumReturns: 2,
				NumLines:   10,
			},
		}
		functionDetails2 := []*functions.FunctionDetails{
			{
				FilePath:   "filepath100",
				File:       "file01",
				Name:       "function1",
				IsPublic:   true,
				NumParams:  2,
				NumReturns: 1,
				NumLines:   5,
			},
			{
				FilePath:   "filepath300",
				File:       "file03",
				Name:       "function3",
				IsPublic:   true,
				NumParams:  5,
				NumReturns: 2,
				NumLines:   10,
			},
		}
		mock := gomonkey.ApplyFunc(
			functions.RetrieveFunctions,
			func(_ *model.PackageInfo, _ string) ([]*functions.FunctionDetails, error) {
				if index == 1 {
					index++
					return functionDetails1, nil
				}

				return functionDetails2, nil
			})

		defer mock.Reset()

		moduleInfo := model.ModuleInfo{
			MainPackage: "mymodule",
			Packages: []*model.PackageInfo{
				{
					Name: "foobar1",
					Path: "barfoo1",
				},
				{
					Name: "foobar2",
					Path: "barfoo2",
				},
			},
		}
		functionRules := []*configuration.FunctionsRule{
			{
				Package:                  "barfoo",
				MaxLines:                 values.GetIntRef(10),
				MaxParameters:            values.GetIntRef(5),
				MaxReturnValues:          values.GetIntRef(2),
				MaxPublicFunctionPerFile: values.GetIntRef(1),
			},
		}
		fx := functionRules[0]

		expectedResult := &functions.RulesResult{
			Passes: false,
			Results: []*functions.RuleResult{
				{
					Rule:        *fx,
					Description: "Functions in packages matching pattern 'barfoo' should have ['at most 5 parameters','at most 2 return values','at most 10 lines','no more than 1 public functions per file']",
					Passes:      false,
					Verifications: []functions.Verification{
						{
							Package: "barfoo1",
							Details: []string{"Function function1 in file filepath1 receive too many parameters (20)"},
							Passes:  false,
						},
						{
							Package: "barfoo2",
							Passes:  true,
						},
					},
				},
			},
		}

		result := functions.CheckRules(moduleInfo, functionRules)

		jsonExpectedResult, _ := json.Marshal(expectedResult)
		jsonResult, _ := json.Marshal(result)
		assert.Equal(t, jsonExpectedResult, jsonResult)
	})
}
