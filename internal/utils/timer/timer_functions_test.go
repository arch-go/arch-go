package timer_test

import (
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/fdaines/arch-go/internal/utils/timer"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

func TestExecuteWithTimer(t *testing.T) {
	t.Run("Calls ExecuteWithTimer function", func(t *testing.T) {
		var out string
		printPatch := gomonkey.ApplyFunc(fmt.Fprintf, func(w io.Writer, f string, a ...interface{}) (int, error) {
			out = fmt.Sprintf(f, a)
			return 0, nil
		})
		defer printPatch.Reset()
		patch := gomonkey.ApplyFuncReturn(time.Since, time.Duration(123456789))
		defer patch.Reset()
		innerFunctionCalls := 0
		innerFunction := func() { innerFunctionCalls++ }

		timer.ExecuteWithTimer(innerFunction)

		printPatch.Reset()
		assert.Equal(t, 1, innerFunctionCalls)
		assert.Equal(t, "Time: [0.123] seconds\n", out)
	})
}
