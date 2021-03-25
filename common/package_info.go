package common

import "go/build"

type PackageInfo struct {
	PackageData *build.Package
	Name        string
	Path        string
}

type PackagesSummary struct {
	Packages []*PackageInfo
}
