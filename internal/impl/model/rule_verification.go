package model

type RuleVerification interface {
	PrintResults()
	Verify() bool
}
