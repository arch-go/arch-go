package packages_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gopkg "golang.org/x/tools/go/packages"

	"github.com/arch-go/arch-go/internal/utils/packages"
)

func TestGetPackages(t *testing.T) {
	t.Run("Calls GetPackages function", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		loadPatch := gomonkey.ApplyFunc(gopkg.Load,
			func(_ *gopkg.Config, _ ...string) ([]*gopkg.Package, error) {
				return []*gopkg.Package{
					{
						PkgPath: "fmt",
					},
					{
						PkgPath: "io",
					},
					{
						PkgPath: "github.com/arch-go/arch-go/internal/reports/console",
					},
				}, nil
			},
		)

		defer loadPatch.Reset()

		expectedOutput := `Looking for packages.
3 packages found...
Loading package (1/3): fmt
Loading package (2/3): io
Loading package (3/3): github.com/arch-go/arch-go/internal/reports/console
`

		pkgs, _ := packages.GetBasicPackagesInfo("foo", outputBuffer, true)

		assert.Len(t, pkgs, 3)
		assert.Equal(t, "fmt", pkgs[0].Name)
		assert.Equal(t, "io", pkgs[1].Name)
		assert.Equal(t, "console", pkgs[2].Name)
		assert.Equal(t, "github.com/arch-go/arch-go/internal/reports/console", pkgs[2].Path)
		assert.Equal(t, expectedOutput, outputBuffer.String())
	})

	t.Run("Calls GetPackages function without printinfo", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		loadPatch := gomonkey.ApplyFunc(gopkg.Load,
			func(_ *gopkg.Config, _ ...string) ([]*gopkg.Package, error) {
				return []*gopkg.Package{
					{
						PkgPath: "fmt",
					},
					{
						PkgPath: "io",
					},
					{
						PkgPath: "github.com/arch-go/arch-go/internal/reports/console",
					},
				}, nil
			},
		)

		defer loadPatch.Reset()

		expectedOutput := ``

		pkgs, _ := packages.GetBasicPackagesInfo("foo", outputBuffer, false)

		assert.Len(t, pkgs, 3)
		assert.Equal(t, "fmt", pkgs[0].Name)
		assert.Equal(t, "io", pkgs[1].Name)
		assert.Equal(t, "console", pkgs[2].Name)
		assert.Equal(t, "github.com/arch-go/arch-go/internal/reports/console", pkgs[2].Path)
		assert.Equal(t, expectedOutput, outputBuffer.String())
	})

	t.Run("Calls GetPackages function", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		loadPatch := gomonkey.ApplyFuncReturn(gopkg.Load, nil, errors.New("test error"))
		expectedOutput := `Looking for packages.
`

		defer loadPatch.Reset()

		pkgs, err := packages.GetBasicPackagesInfo("foo", outputBuffer, true)

		require.Error(t, err)
		require.ErrorContains(t, err, "error: cannot load module packages: test error")
		assert.Empty(t, pkgs)
		assert.Equal(t, expectedOutput, outputBuffer.String())
	})

	t.Run("Get builtin package", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")

		pkgs, err := packages.GetBasicPackagesInfo("builtin", outputBuffer, false)

		assert.NoError(t, err)
		assert.Len(t, pkgs, 1)
		assert.Equal(t, "builtin", pkgs[0].Name)
		assert.Equal(t, "builtin", pkgs[0].Path)
		assert.Equal(t, "", outputBuffer.String())
	})
}
