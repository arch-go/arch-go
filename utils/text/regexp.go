package text

import "strings"

func PreparePackageRegexp(p string) string {
	str := p
	if strings.HasPrefix(str, "**.") {
		str = strings.Replace(str, "**.", "^([\\w-\\.]+/)+", 1)
	}
	if strings.HasPrefix(str, "*.") {
		str = strings.Replace(str, "*.", "^[\\w-\\.]+/", 1)
	}
	if strings.HasSuffix(str, ".**") {
		str = strings.Replace(str, ".**", "(/[\\w-\\.]+)*$", 1)
	}
	if strings.HasSuffix(str, ".*") {
		str = strings.Replace(str, ".*", "(/[\\w-\\.]+){0,1}$", 1)
	}
	str = strings.Replace(str, ".**.", "(/[\\w-\\.]+/)+", -1)
	str = strings.Replace(str, ".*.", "/[\\w-\\.]+/", -1)

	return str
}