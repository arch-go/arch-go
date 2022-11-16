package html

import "embed"

//go:embed templates/*.tmpl
var templateFiles embed.FS

//go:embed templates/*.css
var styles embed.FS
