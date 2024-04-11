package values

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIntRef(t *testing.T) {
	t.Run("test cases", func(t *testing.T) {
		v1 := 10
		result1 := GetIntRef(v1)
		assert.Equal(t, result1, &v1)

		v2 := 0
		result2 := GetIntRef(v2)
		assert.Equal(t, result2, &v2)

		v3 := 12345
		result3 := GetIntRef(v3)
		assert.Equal(t, result3, &v3)
	})
}
