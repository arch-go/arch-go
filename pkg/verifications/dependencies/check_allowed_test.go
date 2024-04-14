package dependencies

import (
	"fmt"
	"testing"

	"github.com/fdaines/arch-go/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestCheckAllowedDependencies(t *testing.T) {
	moduleInfo := model.ModuleInfo{
		MainPackage: "mymodule",
		Packages:    []*model.PackageInfo{},
	}

	t.Run("check allowed standard imports", func(t *testing.T) {
		testCases := []struct {
			inputPkg         string
			inputAllowedPkgs []string
			expectedResult   bool
			expectedDetails  []string
		}{
			{
				"abc",
				[]string{"foo", "bar"},
				false,
				[]string{"ShouldOnlyDependsOn.Standard rule doesn't contains imported package 'abc'"},
			},
			{
				"abc",
				[]string{"foo", "abc"},
				true,
				nil,
			},
			{
				"abc",
				[]string{},
				true,
				nil,
			},
		}

		for idx, tc := range testCases {
			ok, details := checkAllowedStandardImports(tc.inputPkg, tc.inputAllowedPkgs, moduleInfo)
			assert.Equal(t, tc.expectedResult, ok, fmt.Sprintf("TestCase(%v) unexpected result.", idx+1))
			assert.Equal(t, tc.expectedDetails, details, fmt.Sprintf("TestCase(%v) unexpected details.", idx+1))
		}
	})

	t.Run("check allowed external imports", func(t *testing.T) {
		testCases := []struct {
			inputPkg         string
			inputAllowedPkgs []string
			expectedResult   bool
			expectedDetails  []string
		}{
			{
				"foo.bar/blablabla",
				[]string{"xxx", "yyy"},
				false,
				[]string{"ShouldOnlyDependsOn.External rule doesn't contains imported package 'foo.bar/blablabla'"},
			},
			{
				"foo.bar/blablabla",
				[]string{"foo", "abc"},
				true,
				nil,
			},
			{
				"foo.bar/blablabla",
				[]string{},
				true,
				nil,
			},
		}

		for idx, tc := range testCases {
			ok, details := checkAllowedExternalImports(tc.inputPkg, tc.inputAllowedPkgs, moduleInfo)
			assert.Equal(t, tc.expectedResult, ok, fmt.Sprintf("TestCase(%v) unexpected result.", idx+1))
			assert.Equal(t, tc.expectedDetails, details, fmt.Sprintf("TestCase(%v) unexpected details.", idx+1))
		}
	})

	t.Run("check allowed internal imports", func(t *testing.T) {
		testCases := []struct {
			inputPkg         string
			inputAllowedPkgs []string
			expectedResult   bool
			expectedDetails  []string
		}{
			{
				"mymodule/blablabla/src",
				[]string{"mymodule/blablabla/internal", "mymodule/blablabla/pkg"},
				false,
				[]string{"ShouldOnlyDependsOn.Internal rule doesn't contains imported package 'mymodule/blablabla/src'"},
			},
			{
				"mymodule/blablabla/pkg",
				[]string{"mymodule/blablabla/internal", "mymodule/blablabla/pkg"},
				true,
				nil,
			},
			{
				"mymodule/blablabla",
				[]string{},
				true,
				nil,
			},
		}

		for idx, tc := range testCases {
			ok, details := checkAllowedInternalImports(tc.inputPkg, tc.inputAllowedPkgs, moduleInfo)
			assert.Equal(t, tc.expectedResult, ok, fmt.Sprintf("TestCase(%v) unexpected result.", idx+1))
			assert.Equal(t, tc.expectedDetails, details, fmt.Sprintf("TestCase(%v) unexpected details.", idx+1))
		}
	})
}
