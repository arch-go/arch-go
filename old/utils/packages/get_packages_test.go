package packages_test

import (
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	gopkg "golang.org/x/tools/go/packages"

	"github.com/fdaines/arch-go/old/utils/packages"
)

func TestGetPackages(t *testing.T) {
	t.Run("Calls GetPackages function", func(t *testing.T) {
		loadPatch := monkey.Patch(gopkg.Load, func(cfg *gopkg.Config, patterns ...string) ([]*gopkg.Package, error) {
			return []*gopkg.Package{
				&gopkg.Package{
					PkgPath: "fmt",
				},
				&gopkg.Package{
					PkgPath: "io",
				},
				&gopkg.Package{
					PkgPath: "github.com/fdaines/arch-go/old/report/console",
				},
			}, nil
		})
		defer loadPatch.Unpatch()

		pkgs, _ := packages.GetBasicPackagesInfo(true)

		assert.Equal(t, 3, len(pkgs))
		assert.Equal(t, "fmt", pkgs[0].Name)
		assert.Equal(t, "io", pkgs[1].Name)
		assert.Equal(t, "console", pkgs[2].Name)
		assert.Equal(t, "github.com/fdaines/arch-go/old/report/console", pkgs[2].Path)
	})
}
