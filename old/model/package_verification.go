package model

type PackageVerification struct {
	Package *PackageInfo
	Details []string
	Passes  bool
}
