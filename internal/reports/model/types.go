package model

import (
	"time"
)

type Report struct {
	ArchGoVersion string         `json:"version"`
	Summary       *Summary       `json:"summary"`
	Compliance    Compliance     `json:"compliance"`
	Coverage      Coverage       `json:"coverage"`
	Details       *ReportDetails `json:"details-old"`
}

type Summary struct {
	Pass     bool          `json:"pass"`
	Time     time.Time     `json:"timestamp"`
	Duration time.Duration `json:"duration"`
}

type Compliance struct {
	Pass      bool        `json:"pass"`
	Rate      int         `json:"rate"`
	Threshold *int        `json:"threshold"`
	Total     int         `json:"total"`
	Passed    int         `json:"passed"`
	Failed    int         `json:"failed"`
	Summary   []string    `json:"summary"`
	Details   interface{} `json:"details"`
}

type Coverage struct {
	Pass      bool              `json:"pass"`
	Rate      int               `json:"rate"`
	Threshold *int              `json:"threshold"`
	Uncovered []string          `json:"uncovered_packages"`
	Details   []CoverageDetails `json:"details"`
}

type CoverageDetails struct {
	Package           string `json:"package"`
	ContentsRules     int    `json:"contents_rules"`
	DependenciesRules int    `json:"dependencies_rules"`
	FunctionsRules    int    `json:"functions_rules"`
	NamingRules       int    `json:"naming_rules"`
	Covered           bool   `json:"covered"`
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
	Threshold  *int
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
