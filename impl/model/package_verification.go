package model

import "github.com/fdaines/arch-go/model"

type PackageVerification struct {
	Package *model.PackageInfo
	Details []string
	Passes  bool
}