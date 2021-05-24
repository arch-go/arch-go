package packages_test

import (
	"fmt"
	"github.com/fdaines/arch-go/internal/utils/packages"
	"testing"
)

func TestIsPublic(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"foobar", false},
		{"fooBar", false},
		{"Foobar", true},
		{"", false},
	}

	for _, tt := range tests {
		testCase := fmt.Sprintf("input: %s", tt.input)
		t.Run(testCase, func(t *testing.T) {
			ans := packages.IsPublic(tt.input)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
