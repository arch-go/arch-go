package functions

import (
	"fmt"
	"github.com/fdaines/arch-go/internal/model"
)

func checkMaxLines(pkg *model.PackageInfo, mainPackage string, maxLines int) (bool, []string) {
	var details []string
	passes := true
	functions, _ := retrieveFunctions(pkg, mainPackage)
	for _, fn := range functions {
		if fn.NumLines > maxLines {
			passes = false
			details = append(details,
				fmt.Sprintf("Function %s in file %s has too many lines (%d)",
					fn.Name, fn.FilePath, fn.NumLines))
		}
	}
	return passes, details
}
