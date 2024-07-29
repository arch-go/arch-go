package describe

import (
	"fmt"
	"io"
	"strings"

	"github.com/fdaines/arch-go/api/configuration"
	"github.com/fdaines/arch-go/internal/common"
)

func describeContentRules(rules []*configuration.ContentsRule, out io.Writer) {
	fmt.Fprint(out, "Content Rules\n")

	if len(rules) == 0 {
		fmt.Fprint(out, common.NoRulesDefined)

		return
	}

	for _, r := range rules {
		fmt.Fprintf(out, "\t* Packages that match pattern '%s' %s\n", r.Package, resolveContentRule(r))
	}
}

func resolveContentRule(rule *configuration.ContentsRule) string {
	var shouldNotContain []string

	if rule.ShouldOnlyContainStructs {
		return "should only contain structs"
	}

	if rule.ShouldOnlyContainInterfaces {
		return "should only contain interfaces"
	}

	if rule.ShouldOnlyContainFunctions {
		return "should only contain functions"
	}

	if rule.ShouldOnlyContainMethods {
		return "should only contain methods"
	}

	if rule.ShouldNotContainStructs {
		shouldNotContain = append(shouldNotContain, "structs")
	}

	if rule.ShouldNotContainInterfaces {
		shouldNotContain = append(shouldNotContain, "interfaces")
	}

	if rule.ShouldNotContainFunctions {
		shouldNotContain = append(shouldNotContain, "functions")
	}

	if rule.ShouldNotContainMethods {
		shouldNotContain = append(shouldNotContain, "methods")
	}

	return "should not contain " + strings.Join(shouldNotContain, " or ")
}
