package model

import "go/build"

type ModuleInfo struct {
	MainPackage string
	Packages    []*PackageInfo
}

type PackageInfo struct {
	PackageData *build.Package
	Name        string
	Path        string
	Covered     bool
}
