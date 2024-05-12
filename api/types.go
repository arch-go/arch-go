package api

import (
	"time"

	"github.com/fdaines/arch-go/internal/verifications/contents"
	"github.com/fdaines/arch-go/internal/verifications/dependencies"
	"github.com/fdaines/arch-go/internal/verifications/functions"
	"github.com/fdaines/arch-go/internal/verifications/naming"
)

type Result struct {
	Time                   time.Time
	Duration               time.Duration
	Passes                 bool
	DependenciesRuleResult *dependencies.RulesResult
	FunctionsRuleResult    *functions.RulesResult
	ContentsRuleResult     *contents.RulesResult
	NamingRuleResult       *naming.RulesResult
}
