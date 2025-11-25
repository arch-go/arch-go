package describe

import (
	"fmt"
	"io"

	"github.com/arch-go/arch-go/v2/api/configuration"
	"github.com/arch-go/arch-go/v2/internal/common"
)

func describeFunctionRules(rules []*configuration.FunctionsRule, out io.Writer) {
	fmt.Fprint(out, "Function Rules\n")

	if len(rules) == 0 {
		fmt.Fprint(out, common.NoRulesDefined)

		return
	}

	for _, rule := range rules {
		fmt.Fprintf(out, "\t* Packages that match pattern '%s' should comply with the following rules:\n",
			rule.Package)

		if rule.MaxLines != nil {
			fmt.Fprintf(out, "\t\t* Functions should not have more than %d lines\n",
				*rule.MaxLines)
		}

		if rule.MaxParameters != nil {
			fmt.Fprintf(out, "\t\t* Functions should not have more than %d parameters\n",
				*rule.MaxParameters)
		}

		if rule.MaxReturnValues != nil {
			fmt.Fprintf(out, "\t\t* Functions should not have more than %d return values\n",
				*rule.MaxReturnValues)
		}

		if rule.MaxPublicFunctionPerFile != nil {
			fmt.Fprintf(out, "\t\t* Files should not have more than %d public functions\n",
				*rule.MaxPublicFunctionPerFile)
		}
	}
}
