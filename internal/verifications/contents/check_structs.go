package contents

import (
	"github.com/arch-go/arch-go/api/configuration"
)

func checkStructs(pkg *PackageContents, rule *configuration.ContentsRule) (bool, []string) {
	var details []string

	if pkg.Structs > 0 {
		if rule.ShouldNotContainStructs {
			details = append(details, "contains structs and it should not")
		}

		if rule.ShouldOnlyContainInterfaces {
			details = append(details, "contains structs and should only contain interfaces")
		}

		if rule.ShouldOnlyContainMethods {
			details = append(details, "contains structs and should only contain methods")
		}

		if rule.ShouldOnlyContainFunctions {
			details = append(details, "contains structs and should only contain functions")
		}
	}

	return len(details) == 0, details
}
