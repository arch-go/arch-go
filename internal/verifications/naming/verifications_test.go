package naming

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXXXXXXX(t *testing.T) {
	t.Run("getPatternComparator case 1", func(t *testing.T) {
		comparator, s := getPatternComparator("")

		assert.Equal(t, reflect.ValueOf(comparator), reflect.ValueOf(strings.EqualFold))
		assert.Equal(t, "", s)
	})
}
