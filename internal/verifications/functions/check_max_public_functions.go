package functions

import (
	"fmt"
)

func checkMaxPublicFunctions(functions []*FunctionDetails, maxPublicFunctions *int) (bool, []string) {
	var details []string

	passes := true

	if maxPublicFunctions == nil {
		return passes, details
	}

	publicFunctions := map[string]int{}

	for _, fn := range functions {
		if fn.IsPublic {
			current := publicFunctions[fn.FilePath]
			publicFunctions[fn.FilePath] = current + 1
		}
	}

	for key, value := range publicFunctions {
		if value > *maxPublicFunctions {
			passes = false

			details = append(details,
				fmt.Sprintf("File %s has too many public functions (%d)", key, value))
		}
	}

	return passes, details
}
