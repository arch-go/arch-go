package reports

import (
	"time"
)

type Report struct {
	Summary *ReportSummary
	Details *ReportDetails
}

type ReportSummary struct {
	Status              string
	Time                time.Time
	Duration            time.Duration
	Total               int
	Succeeded           int
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
