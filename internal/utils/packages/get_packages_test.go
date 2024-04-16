package packages_test

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"

	"github.com/stretchr/testify/assert"
	gopkg "golang.org/x/tools/go/packages"

	"github.com/fdaines/arch-go/internal/utils/packages"
)

func TestGetPackages(t *testing.T) {
	t.Run("Calls GetPackages function", func(t *testing.T) {
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

		pkgs, _ := packages.GetBasicPackagesInfo("foo", true)

		assert.Equal(t, 3, len(pkgs))
		assert.Equal(t, "fmt", pkgs[0].Name)
		assert.Equal(t, "io", pkgs[1].Name)
		assert.Equal(t, "console", pkgs[2].Name)
		assert.Equal(t, "github.com/fdaines/arch-go/internal/reports/console", pkgs[2].Path)
	})
}
