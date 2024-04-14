package packages_test

import (
	"fmt"
	"testing"

	"github.com/fdaines/arch-go/old/utils/packages"
)

func TestIsStandard(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"foobar", true},
		{"fooBar", true},
		{"golang.org/x/foobar", true},
		{"golang.org/x", true},
		{"foo.bar", false},
		{"github.com/foobar", false},
	}

	for _, tt := range tests {
		testCase := fmt.Sprintf("input: %s", tt.input)
		t.Run(testCase, func(t *testing.T) {
			ans := packages.IsStandardPackage(tt.input)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
