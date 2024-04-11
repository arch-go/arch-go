package contents

import "github.com/fdaines/arch-go/pkg/config"

func checkInterfaces(pkg *PackageContents, rule *config.ContentsRule) (bool, []string) {
	var details []string
	if pkg.Interfaces > 0 {
		if rule.ShouldNotContainInterfaces {
			details = append(details, "contains interfaces and it should not")
		}
		if rule.ShouldOnlyContainStructs {
			details = append(details, "contains interfaces and should only contain structs")
		}
		if rule.ShouldOnlyContainMethods {
			details = append(details, "contains interfaces and should only contain methods")
		}
		if rule.ShouldOnlyContainFunctions {
			details = append(details, "contains interfaces and should only contain functions")
		}
	}

	return len(details) == 0, details
}
