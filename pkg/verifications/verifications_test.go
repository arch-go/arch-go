package verifications_test

import (
	"testing"
	"time"

	"github.com/fdaines/arch-go/pkg/verifications"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/fdaines/arch-go/internal/model"
	"github.com/fdaines/arch-go/pkg/config"
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
		configuration := config.Config{}

		result := verifications.CheckArchitecture(moduleInfo, configuration)

		assert.True(t, result.Passes)
	})
}
