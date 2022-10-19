package packages

import "strings"

func IsExternalPackage(pkg string) bool {
	if strings.HasPrefix(pkg, "golang.org/x") {
		return false
	}
	if strings.ContainsAny(pkg, ".") {
		return true
	}
	return false
}

func IsStandardPackage(pkg string) bool {
	if strings.HasPrefix(pkg, "golang.org/x") {
		return true
	}
	if strings.ContainsAny(pkg, ".") {
		return false
	}
	return true
}
