package values

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStringRef(t *testing.T) {
	t.Run("test cases", func(t *testing.T) {
		v1 := "10"
		result1 := GetStringRef(v1)
		assert.Equal(t, result1, &v1)

		v2 := ""
		result2 := GetStringRef(v2)
		assert.Equal(t, result2, &v2)
	})
}
