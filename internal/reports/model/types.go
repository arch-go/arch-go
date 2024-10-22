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
	Summary   []string       `json:"summary,omitempty"`
	Details   *ReportDetails `json:"details,omitempty"`
}

type Coverage struct {
	Pass      bool              `json:"pass"`
	Rate      int               `json:"rate"`
	Threshold *int              `json:"threshold"`
	Uncovered []string          `json:"uncoveredPackages,omitempty"`
	Details   []CoverageDetails `json:"details,omitempty"`
}

type CoverageDetails struct {
	Package           string `json:"package"`
	ContentsRules     int    `json:"contentsRules"`
	DependenciesRules int    `json:"dependenciesRules"`
	FunctionsRules    int    `json:"functionsRules"`
	NamingRules       int    `json:"namingRules"`
	Covered           bool   `json:"covered"`
}

type ReportDetails struct {
	DependenciesVerificationDetails Verification `json:"dependenciesRules"`
	FunctionsVerificationDetails    Verification `json:"functionsRules"`
	ContentsVerificationDetails     Verification `json:"contentsRules"`
	NamingVerificationDetails       Verification `json:"namingRules"`
}

type Verification struct {
	Total   int                   `json:"total"`
	Passed  int                   `json:"passed"`
	Failed  int                   `json:"failed"`
	Details []VerificationDetails `json:"details,omitempty"`
}

type VerificationDetails struct {
	Rule           string           `json:"rule"`
	Pass           bool             `json:"pass"`
	Total          int              `json:"total"`
	Passed         int              `json:"passed"`
	Failed         int              `json:"failed"`
	PackageDetails []PackageDetails `json:"packageDetails,omitempty"`
}

type PackageDetails struct {
	Package string   `json:"package"`
	Pass    bool     `json:"pass"`
	Details []string `json:"details,omitempty"`
}

type ThresholdSummary struct {
	Rate       int
	Threshold  *int
	Pass       bool
	Violations []string
}
