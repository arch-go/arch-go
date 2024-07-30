package packages_test

import (
	"testing"

	"github.com/arch-go/arch-go/internal/utils/packages"
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
		t.Run("input: "+tt.input, func(t *testing.T) {
			ans := packages.IsStandardPackage(tt.input)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
