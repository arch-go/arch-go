package values

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLessThanZero(t *testing.T) {
	t.Run("test cases", func(t *testing.T) {
		testCases := []struct {
			input    *int
			expected bool
		}{
			{nil, false},
			{GetIntRef(0), false},
			{GetIntRef(100), false},
			{GetIntRef(-1), true},
		}

		for idx, tt := range testCases {
			result := IsLessThanZero(tt.input)
			assert.Equal(
				t,
				tt.expected,
				result,
				"Case:",
				idx+1,
				"input:",
				tt.input,
				"expectedOutput:",
				tt.expected,
			)
		}
	})
}
