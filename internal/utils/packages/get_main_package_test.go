package packages_test

import (
	"errors"
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/internal/utils/packages"
)

func TestGetMainPackage(t *testing.T) {
	t.Run("Calls GetMainPackage function", func(t *testing.T) {
		gomodFile := `module github.com/arch-go/my-golang-module

go 1.18

require (
	foobar 0.0.1
)`
		readFilePatch := gomonkey.ApplyFunc(os.ReadFile, func(_ string) ([]byte, error) {
			return []byte(gomodFile), nil
		})
		statPatch := gomonkey.ApplyFuncReturn(os.Stat, nil, nil)

		defer readFilePatch.Reset()
		defer statPatch.Reset()

		expected := "github.com/arch-go/my-golang-module"
		modulePath, _ := packages.GetMainPackage()

		assert.Equal(t, expected, modulePath)
	})

	t.Run("Calls GetMainPackage function and go.mod file doesnt exists", func(t *testing.T) {
		statPatch := gomonkey.ApplyFuncReturn(os.Stat, nil, errors.New("test error"))
		defer statPatch.Reset()

		expected := ""
		modulepath, err := packages.GetMainPackage()

		assert.Equal(t, expected, modulepath)
		assert.Equal(t, "could not load go.mod file. test error", err.Error())
	})
}
