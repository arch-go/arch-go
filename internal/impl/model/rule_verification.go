package model

type RuleVerification interface {
	Type() string
	Status() bool
	Name() string
	PrintResults()
	Verify() bool
}
