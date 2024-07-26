package output_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/fdaines/arch-go/internal/utils/output"
)

func TestNilWriter(t *testing.T) {
	t.Run("createNilWriter", func(t *testing.T) {
		writer := output.CreateNilWriter()

		n, err := writer.Write([]byte("foobar"))

		require.NoError(t, err)
		assert.Zero(t, n)
	})
}
