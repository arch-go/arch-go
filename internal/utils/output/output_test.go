package output_test

import (
	"testing"

	"github.com/fdaines/arch-go/internal/utils/output"

	"github.com/stretchr/testify/assert"
)

func TestNilWriter(t *testing.T) {
	t.Run("createNilWriter", func(t *testing.T) {
		writer := output.CreateNilWriter()

		n, err := writer.Write([]byte("foobar"))

		assert.Nil(t, err)
		assert.Zero(t, n)
	})
}
