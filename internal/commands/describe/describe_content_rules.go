package describe

import (
	"fmt"
	"io"
	"strings"

	"github.com/fdaines/arch-go/api/configuration"

	"github.com/fdaines/arch-go/internal/common"
)

func describeContentRules(rules []*configuration.ContentsRule, out io.Writer) {
	fmt.Fprintf(out, "Content Rules\n")

	if len(rules) == 0 {
		fmt.Fprintf(out, common.NoRulesDefined)
		return
	}

	for _, r := range rules {
		fmt.Fprintf(out, "\t* Packages that match pattern '%s' %s\n", r.Package, resolveContentRule(r))
	}
}

func resolveContentRule(r *configuration.ContentsRule) string {
	var shouldNotContain []string

	if r.ShouldOnlyContainStructs {
		return "should only contain structs"
	}

	if r.ShouldOnlyContainInterfaces {
		return "should only contain interfaces"
	}

	if r.ShouldOnlyContainFunctions {
		return "should only contain functions"
	}

	if r.ShouldOnlyContainMethods {
		return "should only contain methods"
	}

	if r.ShouldNotContainStructs {
		shouldNotContain = append(shouldNotContain, "structs")
	}

	if r.ShouldNotContainInterfaces {
		shouldNotContain = append(shouldNotContain, "interfaces")
	}

	if r.ShouldNotContainFunctions {
		shouldNotContain = append(shouldNotContain, "functions")
	}

	if r.ShouldNotContainMethods {
		shouldNotContain = append(shouldNotContain, "methods")
	}

	return fmt.Sprintf("should not contain %s", strings.Join(shouldNotContain, " or "))
}
