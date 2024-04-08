package functions

import (
	"fmt"
	"github.com/fdaines/arch-go/old/model"
)

func checkMaxReturnValues(pkg *model.PackageInfo, mainPackage string, maxReturnValues int) (bool, []string) {
	var details []string
	passes := true
	functions, _ := retrieveFunctions(pkg, mainPackage)
	for _, fn := range functions {
		if fn.NumReturns > maxReturnValues {
			passes = false
			details = append(details,
				fmt.Sprintf("Function %s in file %s returns too many values (%d)",
					fn.Name, fn.FilePath, fn.NumReturns))
		}
	}
	return passes, details
}
