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
		idx := strings.LastIndex(str, ".**")
		str = str[:idx] + strings.Replace(str[idx:], ".**", "(/[\\w-\\.]+)*$", 1)
	}

	if strings.HasSuffix(str, ".*") {
		idx := strings.LastIndex(str, ".*")
		str = str[:idx] + strings.Replace(str[idx:], ".*", "(/[\\w-\\.]+){0,1}$", 1)
	}

	str = strings.ReplaceAll(str, ".**.", "(/[\\w-\\.]+/)+")
	str = strings.ReplaceAll(str, ".*.", "/[\\w-\\.]+/")

	return str
}
