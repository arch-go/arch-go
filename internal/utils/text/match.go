package text

import (
	"regexp"
	"strings"
)

func MatchPath(compiledRegex *regexp.Regexp, fullPath, modulePrefix string) bool {
	if compiledRegex.MatchString(fullPath) {
		return true
	}

	if modulePrefix != "" {
		relativePath := strings.TrimPrefix(fullPath, modulePrefix+"/")
		if relativePath != fullPath {
			return compiledRegex.MatchString(relativePath)
		}
	}

	return false
}
