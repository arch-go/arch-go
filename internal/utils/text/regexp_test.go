package text

import (
	"fmt"
	"testing"
)

func TestPreparePackageRegexp(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"foobar", "foobar"},
		{"*.foobar", "^[\\w-\\.]+/foobar"},
		{"**.foobar", "^([\\w-\\.]+/)+foobar"},
		{"foobar.*", "foobar(/[\\w-\\.]+){0,1}$"},
		{"foobar.**", "foobar(/[\\w-\\.]+)*$"},
		{"*.foobar.*", "^[\\w-\\.]+/foobar(/[\\w-\\.]+){0,1}$"},
		{"**.foobar.**", "^([\\w-\\.]+/)+foobar(/[\\w-\\.]+)*$"},
		{"foo.*.bar", "foo/[\\w-\\.]+/bar"},
		{"foo.**.bar", "foo(/[\\w-\\.]+/)+bar"},
		{"*.foo.**.bar.**", "^[\\w-\\.]+/foo(/[\\w-\\.]+/)+bar(/[\\w-\\.]+)*$"},
	}

	for _, tt := range tests {
		testCase := fmt.Sprintf("input: %s", tt.input)
		t.Run(testCase, func(t *testing.T) {
			ans := preparePackageRegexp(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
