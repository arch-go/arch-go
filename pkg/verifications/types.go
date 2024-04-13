package verifications

import (
	"time"

	"github.com/fdaines/arch-go/pkg/verifications/contents"
	"github.com/fdaines/arch-go/pkg/verifications/functions"
	"github.com/fdaines/arch-go/pkg/verifications/naming"
)

type Result struct {
	Time                time.Time
	Duration            time.Duration
	Passes              bool
	FunctionsRuleResult *functions.RulesResult
	ContentsRuleResult  *contents.RulesResult
	NamingRuleResult    *naming.RulesResult
}
