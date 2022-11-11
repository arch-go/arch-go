package packages_test

import (
	"github.com/fdaines/arch-go/internal/impl/model"
	baseModel "github.com/fdaines/arch-go/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fdaines/arch-go/internal/utils/packages"
)

func TestPackageCoverage(t *testing.T) {
	t.Run("Calls ResolveCoveredPackages function", func(t *testing.T) {
		verifications := generateVerifications()

		coveredPackages := packages.ResolveCoveredPackages(verifications)

		assert.Equal(t, 3, len(coveredPackages))
	})

	t.Run("Calls ResolveTotalPackages function", func(t *testing.T) {
		packageList := generatePackageList()

		totalPackages := packages.ResolveTotalPackages(packageList)

		assert.Equal(t, 5, len(totalPackages))
	})

	t.Run("Calls ResolveUncoveredPackages function", func(t *testing.T) {
		packageList := generatePackageList()
		verifications := generateVerifications()

		totalPackages := packages.ResolveTotalPackages(packageList)
		coveredPackages := packages.ResolveCoveredPackages(verifications)

		uncovered := packages.ResolveUncoveredPackages(totalPackages, coveredPackages)

		assert.Equal(t, 2, len(uncovered))
	})
}

func generatePackageList() []*baseModel.PackageInfo {
	return []*baseModel.PackageInfo{
		{Path: "foo/bar"},
		{Path: "foo"},
		{Path: "bar"},
		{Path: "foobar/barfoo"},
		{Path: "blablabla/dummy/package"},
	}
}

func generateVerifications() []model.RuleVerification {
	return []model.RuleVerification{
		DummyVerification{},
	}
}

type DummyVerification struct{}

func (d DummyVerification) Type() string           { return "" }
func (d DummyVerification) Status() bool           { return true }
func (d DummyVerification) Name() string           { return "" }
func (d DummyVerification) Verify() bool           { return true }
func (d DummyVerification) PrintResults()          { return }
func (d DummyVerification) ValidatePatterns() bool { return true }
func (d DummyVerification) GetVerifications() []baseModel.PackageVerification {
	return []baseModel.PackageVerification{
		{Package: &baseModel.PackageInfo{Path: "foo"}},
		{Package: &baseModel.PackageInfo{Path: "bar"}},
		{Package: &baseModel.PackageInfo{Path: "foo/bar"}},
	}
}
