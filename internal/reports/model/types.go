package model

import (
	"time"
)

type Report struct {
	ArchGoVersion string         `json:"version"`
	Summary       *Summary       `json:"summary"`
	SummaryOld    *ReportSummary `json:"summary-old"`
	Details       *ReportDetails
	CoverageInfo  []CoverageInfo
}

type Summary struct {
	Pass     bool          `json:"pass"`
	Time     time.Time     `json:"timestamp"`
	Duration time.Duration `json:"duration"`
}

type ReportSummary struct {
	Pass                bool
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
	Pass       bool
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
	Pass           bool
	Total          int
	Passed         int
	Failed         int
	PackageDetails []PackageDetails
}

type PackageDetails struct {
	Package string
	Pass    bool
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
