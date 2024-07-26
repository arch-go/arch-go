package contents

import (
	"github.com/fdaines/arch-go/api/configuration"
)

func checkMethods(pkg *PackageContents, rule *configuration.ContentsRule) (bool, []string) {
	var details []string

	if pkg.Methods > 0 {
		if rule.ShouldNotContainMethods {
			details = append(details, "contains methods and it should not")
		}

		if rule.ShouldOnlyContainInterfaces {
			details = append(details, "contains methods and should only contain interfaces")
		}

		if rule.ShouldOnlyContainStructs {
			details = append(details, "contains methods and should only contain structs")
		}

		if rule.ShouldOnlyContainFunctions {
			details = append(details, "contains methods and should only contain functions")
		}
	}

	return len(details) == 0, details
}
