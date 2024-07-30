package packages_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/internal/utils/packages"
)

func TestIsPublic(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"foobar", false},
		{"fooBar", false},
		{"Foobar", true},
		{"", false},
	}

	for _, tt := range tests {
		testCase := "input: " + tt.input
		t.Run(testCase, func(t *testing.T) {
			ans := packages.IsPublic(tt.input)
			assert.Equal(t, tt.want, ans)
		})
	}
}
