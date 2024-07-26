package functions

import (
	"fmt"
)

func checkMaxReturnValues(functions []*FunctionDetails, maxReturnValues *int) (bool, []string) {
	var details []string

	passes := true

	if maxReturnValues == nil {
		return passes, details
	}

	for _, fn := range functions {
		if fn.NumReturns > *maxReturnValues {
			passes = false

			details = append(details,
				fmt.Sprintf("Function %s in file %s returns too many values (%d)",
					fn.Name, fn.FilePath, fn.NumReturns))
		}
	}

	return passes, details
}
