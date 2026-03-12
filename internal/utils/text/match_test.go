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
			name:         "wildcard prefix matches full path",
			pattern:      "**.internal.foo.**",
			fullPath:     "github.com/mod/internal/foo/bar",
			modulePrefix: "github.com/mod",
			want:         true,
		},
		{
			name:         "wildcard prefix does not match unrelated path",
			pattern:      "**.internal.foo.**",
			fullPath:     "github.com/mod/pkg/bar",
			modulePrefix: "github.com/mod",
			want:         false,
		},
		{
			name:         "literal prefix matches after stripping module",
			pattern:      "internal.foo.**",
			fullPath:     "github.com/mod/internal/foo/bar",
			modulePrefix: "github.com/mod",
			want:         true,
		},
		{
			name:         "literal prefix does not match with wrong module",
			pattern:      "internal.foo.**",
			fullPath:     "github.com/other/internal/foo/bar",
			modulePrefix: "github.com/mod",
			want:         false,
		},
		{
			name:         "empty prefix matches full path directly",
			pattern:      "internal.foo.**",
			fullPath:     "internal/foo/bar",
			modulePrefix: "",
			want:         true,
		},
		{
			name:         "empty prefix no match on qualified path",
			pattern:      "internal.foo.**",
			fullPath:     "github.com/mod/internal/foo/bar",
			modulePrefix: "",
			want:         false,
		},
		{
			name:         "fullPath equals modulePrefix returns false",
			pattern:      "internal.**",
			fullPath:     "github.com/mod",
			modulePrefix: "github.com/mod",
			want:         false,
		},
		{
			name:         "exact relative package match",
			pattern:      "internal.foo",
			fullPath:     "github.com/mod/internal/foo",
			modulePrefix: "github.com/mod",
			want:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := regexp.MustCompile(text.PreparePackageRegexp(tt.pattern))
			got := text.MatchPath(re, tt.fullPath, tt.modulePrefix)
			if got != tt.want {
				t.Errorf("MatchPath(%q, %q, %q) = %v, want %v",
					tt.pattern, tt.fullPath, tt.modulePrefix, got, tt.want)
			}
		})
	}
}
