package model

type RuleVerification interface {
	Type() string
	Status() bool
	Name() string
	PrintResults()
	GetVerifications() []PackageVerification
	Verify() bool
	ValidatePatterns() bool
}