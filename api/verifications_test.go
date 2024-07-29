package api_test

import (
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"

	"github.com/fdaines/arch-go/api"
	"github.com/fdaines/arch-go/api/configuration"
	"github.com/fdaines/arch-go/internal/model"
)

func TestCheckArchitecture(t *testing.T) {
	mockTimeNow := gomonkey.ApplyFuncReturn(time.Now, time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC))
	defer mockTimeNow.Reset()

	t.Run("check passes", func(t *testing.T) {
		moduleInfo := model.ModuleInfo{
			MainPackage: "mymodule",
			Packages:    []*model.PackageInfo{},
		}
		config := configuration.Config{}

		result := api.CheckArchitecture(moduleInfo, config)

		assert.True(t, result.Passes)
	})
}
