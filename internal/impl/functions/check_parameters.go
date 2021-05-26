package functions

import (
	"fmt"
	"github.com/fdaines/arch-go/internal/model"
)

func checkMaxParameters(pkg *model.PackageInfo, mainPackage string, maxParameters int) (bool, []string) {
	var details []string
	passes := true
	functions, _ := retrieveFunctions(pkg, mainPackage)
	for _, fn := range functions {
		if fn.NumParams > maxParameters {
			passes = false
			details = append(details,
				fmt.Sprintf("Function %s in file %s receive too many parameters (%d)",
					fn.Name, fn.FilePath, fn.NumParams))
		}
	}
	return passes, details
}
