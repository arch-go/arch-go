package archgo_test

import (
	"testing"
	"time"

	"github.com/fdaines/arch-go/pkg/archgo"
	"github.com/fdaines/arch-go/pkg/archgo/configuration"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/fdaines/arch-go/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestCheckArchitecture(t *testing.T) {
	mockTimeNow := gomonkey.ApplyFuncReturn(time.Now, time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC))
	defer mockTimeNow.Reset()

	t.Run("check passes", func(t *testing.T) {
		moduleInfo := model.ModuleInfo{
			MainPackage: "mymodule",
			Packages:    []*model.PackageInfo{},
		}
		configuration := configuration.Config{}

		result := archgo.CheckArchitecture(moduleInfo, configuration)

		assert.True(t, result.Passes)
	})
}
