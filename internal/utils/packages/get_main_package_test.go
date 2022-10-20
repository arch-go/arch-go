package packages_test

import (
	"fmt"
	"os"
	"testing"

	"bou.ke/monkey"
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
		readFilePatch := monkey.Patch(os.ReadFile, func(fn string) ([]byte, error) {
			return []byte(gomodFile), nil
		})
		defer readFilePatch.Unpatch()
		statPatch := monkey.Patch(os.Stat, func(string) (os.FileInfo, error) {
			return nil, nil
		})
		defer statPatch.Unpatch()

		expected := "github.com/fdaines/my-golang-module"
		modulePath, _ := packages.GetMainPackage()

		assert.Equal(t, expected, modulePath)
	})

	t.Run("Calls GetMainPackage function and go.mod file doesnt exists", func(t *testing.T) {
		statPatch := monkey.Patch(os.Stat, func(string) (os.FileInfo, error) {
			return nil, fmt.Errorf("Error")
		})
		defer statPatch.Unpatch()

		expected := ""
		modulepath, err := packages.GetMainPackage()

		assert.Equal(t, expected, modulepath)
		assert.Equal(t, "Could not load go.mod file. Error\n", err.Error())
	})
}
