package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fdaines/arch-go/internal/utils/values"
)

func TestValidatorsHelpers(t *testing.T) {
	t.Run("test trueValues function", func(t *testing.T) {
		testCases := []struct {
			input []bool
			want  int32
		}{
			{[]bool{true, true, false, true, false}, 3},
			{[]bool{}, 0},
			{[]bool{true}, 1},
			{[]bool{true, true, true, true, true, true, true}, 7},
			{[]bool{false, false, false, false, false, false, false}, 0},
			{[]bool{false, false, false, false, false, false, false, true}, 1},
		}

		for idx, tc := range testCases {
			result := trueValues(tc.input...)
			assert.Equal(t, tc.want, result, "case(%v), expected:%v, got:%v", idx+1, tc.want, result)
		}
	})

	t.Run("test countNotNil function", func(t *testing.T) {
		testCases := []struct {
			input []*int
			want  int32
		}{
			{[]*int{}, 0},
			{[]*int{nil, nil, nil, nil}, 0},
			{[]*int{values.GetIntRef(1), values.GetIntRef(0), values.GetIntRef(10)}, 3},
			{[]*int{values.GetIntRef(1), nil, values.GetIntRef(10)}, 2},
		}

		for idx, tc := range testCases {
			result := countNotNil(tc.input...)
			assert.Equal(t, tc.want, result, "case(%v), expected:%v, got:%v", idx+1, tc.want, result)
		}
	})
}
