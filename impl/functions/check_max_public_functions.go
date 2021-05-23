package functions

import (
	"fmt"
	"github.com/fdaines/arch-go/model"
)

func check_max_public_functions(pkg *model.PackageInfo, mainPackage string, maxPublicFunctions int) (bool, []string) {
	var details []string
	passes := true
	functions, _ := retrieveFunctions(pkg, mainPackage)
	publicFunctions := map[string]int{}
	for _, fn := range functions {
		if fn.IsPublic {
			current, ok := publicFunctions[fn.FilePath]
			if !ok {
				current = 0
			}
			publicFunctions[fn.FilePath] = current + 1
		}
	}
	for key, value := range publicFunctions {
		if value > maxPublicFunctions {
			passes = false
			details = append(details,
				fmt.Sprintf("File %s has too many public functions (%d)", key, value))
		}
	}
	return passes, details
}
