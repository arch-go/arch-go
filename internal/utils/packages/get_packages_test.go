package packages_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/agiledragon/gomonkey/v2"

	"github.com/stretchr/testify/assert"
	gopkg "golang.org/x/tools/go/packages"

	"github.com/fdaines/arch-go/internal/utils/packages"
)

func TestGetPackages(t *testing.T) {
	t.Run("Calls GetPackages function", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		loadPatch := gomonkey.ApplyFunc(gopkg.Load, func(cfg *gopkg.Config, patterns ...string) ([]*gopkg.Package, error) {
			return []*gopkg.Package{
				{
					PkgPath: "fmt",
				},
				{
					PkgPath: "io",
				},
				{
					PkgPath: "github.com/fdaines/arch-go/internal/reports/console",
				},
			}, nil
		})
		defer loadPatch.Reset()
		expectedOutput := `Looking for packages.
3 packages found...
Loading package (1/3): fmt
Loading package (2/3): io
Loading package (3/3): github.com/fdaines/arch-go/internal/reports/console
`

		pkgs, _ := packages.GetBasicPackagesInfo("foo", outputBuffer, true)

		assert.Equal(t, 3, len(pkgs))
		assert.Equal(t, "fmt", pkgs[0].Name)
		assert.Equal(t, "io", pkgs[1].Name)
		assert.Equal(t, "console", pkgs[2].Name)
		assert.Equal(t, "github.com/fdaines/arch-go/internal/reports/console", pkgs[2].Path)
		assert.Equal(t, expectedOutput, outputBuffer.String())
	})

	t.Run("Calls GetPackages function without printinfo", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		loadPatch := gomonkey.ApplyFunc(gopkg.Load, func(cfg *gopkg.Config, patterns ...string) ([]*gopkg.Package, error) {
			return []*gopkg.Package{
				{
					PkgPath: "fmt",
				},
				{
					PkgPath: "io",
				},
				{
					PkgPath: "github.com/fdaines/arch-go/internal/reports/console",
				},
			}, nil
		})
		defer loadPatch.Reset()
		expectedOutput := ``

		pkgs, _ := packages.GetBasicPackagesInfo("foo", outputBuffer, false)

		assert.Equal(t, 3, len(pkgs))
		assert.Equal(t, "fmt", pkgs[0].Name)
		assert.Equal(t, "io", pkgs[1].Name)
		assert.Equal(t, "console", pkgs[2].Name)
		assert.Equal(t, "github.com/fdaines/arch-go/internal/reports/console", pkgs[2].Path)
		assert.Equal(t, expectedOutput, outputBuffer.String())
	})

	t.Run("Calls GetPackages function", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		loadPatch := gomonkey.ApplyFuncReturn(gopkg.Load, nil, errors.New("load error"))
		defer loadPatch.Reset()
		expectedOutput := `Looking for packages.
`

		pkgs, err := packages.GetBasicPackagesInfo("foo", outputBuffer, true)

		assert.Equal(t, 0, len(pkgs))
		assert.NotNil(t, err)
		assert.Equal(t, "Error: Cannot load module packages: load error\n", err.Error())
		assert.Equal(t, expectedOutput, outputBuffer.String())
	})
}
