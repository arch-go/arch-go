package text_test

import (
	"regexp"
	"testing"

	"github.com/arch-go/arch-go/v2/internal/utils/text"
)

func TestMatchPath(t *testing.T) {
	modulePrefix := "github.com/mod"

	tests := []struct {
		name    string
		pattern string

		matchingPaths    []string
		nonMatchingPaths []string
	}{
		{
			name:    "no wildcards",
			pattern: "foo.bar",
			matchingPaths: []string{
				"github.com/mod/foo/bar",
				"github.com/mod/foo.bar",
				"github.com/mod/fooXbar", // fixme this matches, but should it?
			},
			nonMatchingPaths: []string{
				"github.com/mod/foo/bar/baz",
				"github.com/mod/internal/foo/bar",
			},
		},
		{
			name:    "single wildcard name-suffix",
			pattern: "foo.bar*",
			matchingPaths: []string{
				"github.com/mod/foo/bar",
				"github.com/mod/fooXbar", // fixme this matches, but should it?
				"github.com/mod/foo/barX",
				"github.com/mod/foo/barXZ",
			},
			nonMatchingPaths: []string{
				"github.com/mod/foo/Xbar",
				"github.com/mod/foo/bar/baz",
				"github.com/mod/foo/bar/baz/qux",
				"github.com/mod/internal/foo/bar",
			},
		},
		{
			name:    "single wildcard name-prefix",
			pattern: "*foo.bar",
			matchingPaths: []string{
				"github.com/mod/foo/bar",
				"github.com/mod/Xfoo/bar",
				"github.com/mod/XZfoo/bar",
			},
			nonMatchingPaths: []string{
				"github.com/mod/foo/barX",
				"github.com/mod/Xfoo/bar/baz",
				"github.com/mod/internal/foo/bar",
			},
		},
		{
			name:    "single wildcard path-suffix",
			pattern: "foo.bar.*",
			matchingPaths: []string{
				"github.com/mod/foo/bar",
				"github.com/mod/foo/bar/baz",
			},
			nonMatchingPaths: []string{
				"github.com/mod/foo/bar/baz/qux",
				"github.com/mod/internal/foo/bar",
			},
		},
		{
			name:    "recursive wildcard path-suffix",
			pattern: "foo.bar.**",
			matchingPaths: []string{
				"github.com/mod/foo/bar",
				"github.com/mod/foo/bar/baz",
				"github.com/mod/foo/bar/baz/qux",
			},
			nonMatchingPaths: []string{
				"github.com/mod/internal/foo/bar",
			},
		},
		{
			name:    "single wildcard path-mid-pattern",
			pattern: "foo.*.baz",
			matchingPaths: []string{
				"github.com/mod/foo/bar/baz",
			},
			nonMatchingPaths: []string{
				"github.com/mod/foo/bar/qux/baz",
				"github.com/mod/internal/foo/bar/baz",
			},
		},
		{
			name:    "recursive wildcard path-mid-pattern",
			pattern: "foo.**.baz",
			matchingPaths: []string{
				"github.com/mod/foo/bar/baz",
				"github.com/mod/foo/bar/qux/baz",
			},
			nonMatchingPaths: []string{
				"github.com/mod/internal/foo/bar/baz",
			},
		},
		{
			name:    "single wildcard path-prefix",
			pattern: "*.foo",
			matchingPaths: []string{
				"github.com/mod/internal/foo",
			},
			nonMatchingPaths: []string{
				"github.com/mod/foo",
				"github.com/mod/foo/bar",
				"github.com/mod/internal/foo/bar",
			},
		},
		{
			name:    "recursive wildcard path-prefix",
			pattern: "**.foo",
			matchingPaths: []string{
				"github.com/mod/foo",
				"github.com/mod/bar/baz/foo",
				"github.com/mod/internal/foo",
			},
			nonMatchingPaths: []string{
				"github.com/mod/foo/bar",
				"github.com/mod/internal/foo/bar",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := regexp.MustCompile(text.PreparePackageRegexp(tt.pattern))

			for _, path := range tt.matchingPaths {
				if !text.MatchPath(re, path, modulePrefix) {
					t.Errorf("expected path '%s' to match pattern '%s'", path, tt.pattern)
				}
			}

			for _, path := range tt.nonMatchingPaths {
				if text.MatchPath(re, path, modulePrefix) {
					t.Errorf("expected path '%s' to NOT match pattern '%s'", path, tt.pattern)
				}
			}
		})
	}
}
