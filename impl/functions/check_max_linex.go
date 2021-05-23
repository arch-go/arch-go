package functions

import (
	"fmt"
	"github.com/fdaines/arch-go/model"
)

func check_max_lines(pkg *model.PackageInfo, mainPackage string, maxLines int) (bool, []string) {
	var details []string
	passes := true
	functions, _ := retrieveFunctions(pkg, mainPackage)
	for _, fn := range functions {
		if fn.NumLines > maxLines {
			passes = false
			details = append(details,
				fmt.Sprintf("Function %s in file %s receive too many lines (%d)",
					fn.Name, fn.FilePath, fn.NumLines))
		}
	}
	return passes, details
}
