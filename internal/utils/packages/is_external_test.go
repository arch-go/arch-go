package packages_test

import (
	"fmt"
	"testing"

	"github.com/fdaines/arch-go/old/utils/packages"
)

func TestIsExternal(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"foobar", false},
		{"fooBar", false},
		{"golang.org/x/foobar", false},
		{"golang.org/x", false},
		{"foo.bar", true},
	}

	for _, tt := range tests {
		testCase := fmt.Sprintf("input: %s", tt.input)
		t.Run(testCase, func(t *testing.T) {
			ans := packages.IsExternalPackage(tt.input)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
