package text_test

import (
	"regexp"
	"testing"

	"github.com/arch-go/arch-go/v2/internal/utils/text"
)

func TestMatchPath(t *testing.T) {
	tests := []struct {
		name         string
		pattern      string
		fullPath     string
		modulePrefix string
		want         bool
	}{
		{
			name:         "root: top level package should match",
			pattern:      "foo.bar.**",
			fullPath:     "github.com/mod/foo/bar",
			modulePrefix: "github.com/mod",
			want:         true,
		},
		{
			name:         "root: subpackage should NOT match",
			pattern:      "foo.bar.**",
			fullPath:     "github.com/mod/internal/foo/bar",
			modulePrefix: "github.com/mod",
			want:         false,
		},
		{
			name:         "prefix: top level package should match",
			pattern:      "**.foo.bar.**",
			fullPath:     "github.com/mod/foo/bar",
			modulePrefix: "github.com/mod",
			want:         true,
		},
		{
			name:         "prefix: subpackage should match",
			pattern:      "**.foo.bar.**",
			fullPath:     "github.com/mod/internal/foo/bar",
			modulePrefix: "github.com/mod",
			want:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := regexp.MustCompile(text.PreparePackageRegexp(tt.pattern))
			got := text.MatchPath(re, tt.fullPath, tt.modulePrefix)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
