package packages_test

import (
	"testing"

	"github.com/arch-go/arch-go/v2/internal/utils/packages"
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
		t.Run("input: "+tt.input, func(t *testing.T) {
			ans := packages.IsExternalPackage(tt.input)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
