package packages

import "unicode"

func IsPublic(name string) bool {
	return unicode.IsUpper([]rune(name)[0])
}
