package model

import "go/build"

// ModuleInfo contains information to describe the contents of a module.
type ModuleInfo struct {
	MainPackage string         // main package
	Packages    []*PackageInfo // a set of information about packages in the module
}

// PackageInfo contains information to describe a package.
type PackageInfo struct {
	PackageData *build.Package // build information for package
	Name        string         // package name
	Path        string         // package path
}
