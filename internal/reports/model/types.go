package model

import (
	"time"
)

type Report struct {
	ArchGoVersion string
	Summary       *ReportSummary
	Details       *ReportDetails
	CoverageInfo  []CoverageInfo
}

type ReportSummary struct {
	Status              string
	Time                time.Time
	Duration            time.Duration
	Total               int
	Passed              int
	Failed              int
	ComplianceThreshold *ThresholdSummary
	CoverageThreshold   *ThresholdSummary
}

type ThresholdSummary struct {
	Rate       int
	Threshold  int
	Status     string
	Violations []string
}

type ReportDetails struct {
	DependenciesVerificationDetails Verification
	FunctionsVerificationDetails    Verification
	ContentsVerificationDetails     Verification
	NamingVerificationDetails       Verification
}

type Verification struct {
	Total   int
	Passed  int
	Failed  int
	Details []VerificationDetails
}

type VerificationDetails struct {
	Rule           string
	Status         string
	Total          int
	Passed         int
	Failed         int
	PackageDetails []PackageDetails
}

type PackageDetails struct {
	Package string
	Status  string
	Details []string
}

type CoverageInfo struct {
	Package           string
	ContentsRules     int
	DependenciesRules int
	FunctionsRules    int
	NamingRules       int
	Covered           bool
}
