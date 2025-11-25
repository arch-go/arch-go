package contents

import (
	"github.com/arch-go/arch-go/v2/api/configuration"
)

func checkFunctions(pkg *PackageContents, rule *configuration.ContentsRule) (bool, []string) {
	var details []string

	if pkg.Functions > 0 {
		if rule.ShouldNotContainFunctions {
			details = append(details, "contains functions and it should not")
		}

		if rule.ShouldOnlyContainStructs {
			details = append(details, "contains functions and should only contain structs")
		}

		if rule.ShouldOnlyContainMethods {
			details = append(details, "contains functions and should only contain methods")
		}

		if rule.ShouldOnlyContainInterfaces {
			details = append(details, "contains functions and should only contain interfaces")
		}
	}

	return len(details) == 0, details
}
