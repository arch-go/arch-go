package functions

import (
	"fmt"
)

func checkMaxLines(functions []*FunctionDetails, maxLines *int) (bool, []string) {
	var details []string
	passes := true
	if maxLines == nil {
		return passes, details
	}
	for _, fn := range functions {
		if fn.NumLines > *maxLines {
			passes = false
			details = append(details,
				fmt.Sprintf("Function %s in file %s has too many lines (%d)",
					fn.Name, fn.FilePath, fn.NumLines))
		}
	}
	return passes, details
}
