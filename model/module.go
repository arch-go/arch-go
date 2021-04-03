package model

type ModuleInfo struct {
	MainPackage string
	Packages    []*PackageInfo
}
