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
			name:         "root anchored without wildcard: matches with top level package",
			pattern:      "foo.bar",
			fullPath:     "github.com/mod/foo/bar",
			modulePrefix: "github.com/mod",
			want:         true,
		},
		{
			name:         "root anchored without wildcard: fails to match with subpackage",
			pattern:      "foo.bar",
			fullPath:     "github.com/mod/internal/foo/bar",
			modulePrefix: "github.com/mod",
			want:         false,
		},
		{
			name:         "root anchored with recursive wildcard suffix: matches with top level package",
			pattern:      "foo.bar.**",
			fullPath:     "github.com/mod/foo/bar",
			modulePrefix: "github.com/mod",
			want:         true,
		},
		{
			name:         "root anchored with recursive wildcard suffix: fails to match with subpackage",
			pattern:      "foo.bar.**",
			fullPath:     "github.com/mod/internal/foo/bar",
			modulePrefix: "github.com/mod",
			want:         false,
		},
		{
			name:         "root anchored with single wildcard suffix: matches with top level package",
			pattern:      "foo.bar.*",
			fullPath:     "github.com/mod/foo/bar",
			modulePrefix: "github.com/mod",
			want:         true,
		},
		{
			name:         "root anchored with single wildcard suffix: fails to match with subpackage",
			pattern:      "foo.bar.*",
			fullPath:     "github.com/mod/internal/foo/bar",
			modulePrefix: "github.com/mod",
			want:         false,
		},
		{
			name:         "root anchored with mid-pattern wildcard: matches with top level package",
			pattern:      "foo.*.baz",
			fullPath:     "github.com/mod/foo/bar/baz",
			modulePrefix: "github.com/mod",
			want:         true,
		},
		{
			name:         "root anchored with mid-pattern wildcard: fails to match with subpackage",
			pattern:      "foo.*.baz",
			fullPath:     "github.com/mod/internal/foo/bar/baz",
			modulePrefix: "github.com/mod",
			want:         false,
		},
		{
			name:         "prefix wildcard: matches package at root",
			pattern:      "**.foo.bar.**",
			fullPath:     "github.com/mod/foo/bar",
			modulePrefix: "github.com/mod",
			want:         true,
		},
		{
			name:         "prefix wildcard: matches package with arbitrary leading segments",
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
