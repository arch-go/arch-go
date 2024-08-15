package model

import (
	"time"
)

type Report struct {
	ArchGoVersion string     `json:"version"`
	Summary       *Summary   `json:"summary"`
	Compliance    Compliance `json:"compliance"`
	Coverage      Coverage   `json:"coverage"`
}

type Summary struct {
	Pass     bool          `json:"pass"`
	Time     time.Time     `json:"timestamp"`
	Duration time.Duration `json:"duration"`
}

type Compliance struct {
	Pass      bool           `json:"pass"`
	Rate      int            `json:"rate"`
	Threshold *int           `json:"threshold"`
	Total     int            `json:"total"`
	Passed    int            `json:"passed"`
	Failed    int            `json:"failed"`
	Summary   []string       `json:"summary"`
	Details   *ReportDetails `json:"details"`
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

type ReportDetails struct {
	DependenciesVerificationDetails Verification `json:"dependencies_rules"`
	FunctionsVerificationDetails    Verification `json:"functions_rules"`
	ContentsVerificationDetails     Verification `json:"contents_rules"`
	NamingVerificationDetails       Verification `json:"naming_rules"`
}

type Verification struct {
	Total   int                   `json:"total"`
	Passed  int                   `json:"passed"`
	Failed  int                   `json:"failed"`
	Details []VerificationDetails `json:"details"`
}

type VerificationDetails struct {
	Rule           string           `json:"rule"`
	Pass           bool             `json:"pass"`
	Total          int              `json:"total"`
	Passed         int              `json:"passed"`
	Failed         int              `json:"failed"`
	PackageDetails []PackageDetails `json:"package_details"`
}

type PackageDetails struct {
	Package string   `json:"package"`
	Pass    bool     `json:"pass"`
	Details []string `json:"details"`
}

type ThresholdSummary struct {
	Rate       int
	Threshold  *int
	Pass       bool
	Violations []string
}
