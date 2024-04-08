package model

import "go/build"

type PackageInfo struct {
	PackageData *build.Package
	Name        string
	Path        string
	Covered     bool
}
