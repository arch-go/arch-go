package packages_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"

	"github.com/stretchr/testify/assert"

	"github.com/fdaines/arch-go/internal/utils/packages"
)

func TestGetMainPackage(t *testing.T) {
	t.Run("Calls GetMainPackage function", func(t *testing.T) {
		gomodFile := `module github.com/fdaines/my-golang-module

go 1.18

require (
	foobar 0.0.1
)`
		readFilePatch := gomonkey.ApplyFunc(os.ReadFile, func(fn string) ([]byte, error) {
			return []byte(gomodFile), nil
		})
		defer readFilePatch.Reset()
		statPatch := gomonkey.ApplyFuncReturn(os.Stat, nil, nil)
		defer statPatch.Reset()

		expected := "github.com/fdaines/my-golang-module"
		modulePath, _ := packages.GetMainPackage()

		assert.Equal(t, expected, modulePath)
	})

	t.Run("Calls GetMainPackage function and go.mod file doesnt exists", func(t *testing.T) {
		statPatch := gomonkey.ApplyFuncReturn(os.Stat, nil, fmt.Errorf("Error"))
		defer statPatch.Reset()

		expected := ""
		modulepath, err := packages.GetMainPackage()

		assert.Equal(t, expected, modulepath)
		assert.Equal(t, "Could not load go.mod file. Error\n", err.Error())
	})
}
