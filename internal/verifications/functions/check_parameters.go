package functions

import (
	"fmt"
)

func checkMaxParameters(functions []*FunctionDetails, maxParameters *int) (bool, []string) {
	var details []string
	passes := true
	if maxParameters == nil {
		return passes, details
	}
	for _, fn := range functions {
		if fn.NumParams > *maxParameters {
			passes = false
			details = append(details,
				fmt.Sprintf("Function %s in file %s receive too many parameters (%d)",
					fn.Name, fn.FilePath, fn.NumParams))
		}
	}
	return passes, details
}
