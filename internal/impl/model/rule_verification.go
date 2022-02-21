package model

import "github.com/fdaines/arch-go/internal/model"

type RuleVerification interface {
	Type() string
	Status() bool
	Name() string
	PrintResults()
	GetVerifications() []model.PackageVerification
	Verify() bool
	ValidatePatterns() bool
}