package api

import (
	"sync"
	"time"

	"github.com/arch-go/arch-go/internal/verifications/contents"
	"github.com/arch-go/arch-go/internal/verifications/dependencies"
	"github.com/arch-go/arch-go/internal/verifications/functions"
	"github.com/arch-go/arch-go/internal/verifications/naming"
)

// Result contains the result of an architecture analysis.
type Result struct {
	Time                   time.Time                 // the moment when the analysis was executed
	Duration               time.Duration             // the duration of the analysis
	Pass                   bool                      // if true, then the analysis was succeeded
	DependenciesRuleResult *dependencies.RulesResult // contains all the verifications of dependencies rules
	FunctionsRuleResult    *functions.RulesResult    // contains all the verifications of functions rules
	ContentsRuleResult     *contents.RulesResult     // contains all the verifications of contents rules
	NamingRuleResult       *naming.RulesResult       // contains all the verifications of naming rules
	mu                     sync.Mutex                // mutex to protect the result
}

func (r *Result) setPass(passes bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Pass = r.Pass && passes
}
