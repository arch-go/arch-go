package dependencies

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fdaines/arch-go/internal/model"
)

func TestCheckRestrictedDependencies(t *testing.T) {
	moduleInfo := model.ModuleInfo{
		MainPackage: "mymodule",
		Packages:    []*model.PackageInfo{},
	}

	t.Run("check restricted standard imports", func(t *testing.T) {
		testCases := []struct {
			inputPkg            string
			inputRestrictedPkgs []string
			expectedResult      bool
			expectedDetails     []string
		}{
			{
				"abc",
				[]string{"foo", "bar"},
				true,
				nil,
			},
			{
				"abc",
				[]string{"foo", "abc"},
				false,
				[]string{"ShouldNotDependsOn.Standard rule contains imported package 'abc'"},
			},
			{
				"abc",
				[]string{},
				true,
				nil,
			},
		}

		for idx, tc := range testCases {
			ok, details := checkRestrictedStandardImports(tc.inputPkg, tc.inputRestrictedPkgs, moduleInfo)
			assert.Equal(t, tc.expectedResult, ok, fmt.Sprintf("TestCase(%v) unexpected result.", idx+1))
			assert.Equal(t, tc.expectedDetails, details, fmt.Sprintf("TestCase(%v) unexpected details.", idx+1))
		}
	})

	t.Run("check restricted external imports", func(t *testing.T) {
		testCases := []struct {
			inputPkg            string
			inputRestrictedPkgs []string
			expectedResult      bool
			expectedDetails     []string
		}{
			{
				"foo.bar/blablabla",
				[]string{"xxx", "yyy"},
				true,
				nil,
			},
			{
				"foo.bar/blablabla",
				[]string{"foo", "abc"},
				false,
				[]string{"ShouldNotDependsOn.External rule contains imported package 'foo.bar/blablabla'"},
			},
			{
				"foo.bar/blablabla",
				[]string{},
				true,
				nil,
			},
		}

		for idx, tc := range testCases {
			ok, details := checkRestrictedExternalImports(tc.inputPkg, tc.inputRestrictedPkgs, moduleInfo)
			assert.Equal(t, tc.expectedResult, ok, fmt.Sprintf("TestCase(%v) unexpected result.", idx+1))
			assert.Equal(t, tc.expectedDetails, details, fmt.Sprintf("TestCase(%v) unexpected details.", idx+1))
		}
	})

	t.Run("check restricted internal imports", func(t *testing.T) {
		testCases := []struct {
			inputPkg            string
			inputRestrictedPkgs []string
			expectedResult      bool
			expectedDetails     []string
		}{
			{
				"mymodule/blablabla/src",
				[]string{"mymodule/blablabla/internal", "mymodule/blablabla/pkg"},
				true,
				nil,
			},
			{
				"mymodule/blablabla/pkg",
				[]string{"mymodule/blablabla/internal", "mymodule/blablabla/pkg"},
				false,
				[]string{"ShouldNotDependsOn.Internal rule contains imported package 'mymodule/blablabla/pkg'"},
			},
			{
				"mymodule/blablabla",
				[]string{},
				true,
				nil,
			},
		}

		for idx, tc := range testCases {
			ok, details := checkRestrictedInternalImports(tc.inputPkg, tc.inputRestrictedPkgs, moduleInfo)
			assert.Equal(t, tc.expectedResult, ok, fmt.Sprintf("TestCase(%v) unexpected result.", idx+1))
			assert.Equal(t, tc.expectedDetails, details, fmt.Sprintf("TestCase(%v) unexpected details.", idx+1))
		}
	})
}
