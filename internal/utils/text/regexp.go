package text

import (
	"regexp"
	"strings"
)

func PreparePackageRegexp(p string) string {
	str := p

	re := regexp.MustCompile(`(\w)+\*`)
	pkgs := re.FindAllString(str, -1)

	for i, pkg := range pkgs {
		pkg = strings.ReplaceAll(pkg, "*", `[\w-\.]*`)
		str = strings.ReplaceAll(str, pkgs[i], pkg)
	}

	re = regexp.MustCompile(`\*(\w)+`)
	pkgs = re.FindAllString(str, -1)

	for i, pkg := range pkgs {
		pkg = strings.ReplaceAll(pkg, "*", `[\w-\.]*`)
		str = strings.ReplaceAll(str, pkgs[i], pkg)
	}

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
