package utils

import (
	"fmt"
	"io"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func Test_ExecuteWithTimer(t *testing.T) {
	t.Run("Calls ExecuteWithTimer function", func(t *testing.T) {
		var out string
		printPatch := monkey.Patch(fmt.Fprintf, func(w io.Writer, f string, a ...interface{}) (int, error) {
			out = fmt.Sprintf(f, a)
			return 0, nil
		})
		defer printPatch.Unpatch()
		patch := monkey.Patch(time.Since, func(t time.Time) time.Duration {
			return 123456789
		})
		defer patch.Unpatch()
		innerFunctionCalls := 0
		innerFunction := func(){ innerFunctionCalls++ }

		ExecuteWithTimer(innerFunction)

		assert.Equal(t, 1, innerFunctionCalls)
		assert.Equal(t, "Time: [0.123] seconds\n", out)
	})
}
