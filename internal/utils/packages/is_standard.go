package packages

import "strings"

func IsStandardPackage(pkg string) bool {
	if strings.HasPrefix(pkg, "golang.org/x") {
		return true
	}

	if strings.ContainsAny(pkg, ".") {
		return false
	}

	return true
}
