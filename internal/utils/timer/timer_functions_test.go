package timer_test

import (
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"

	"github.com/arch-go/arch-go/internal/utils/timer"
)

func TestExecuteWithTimer(t *testing.T) {
	t.Run("Calls ExecuteWithTimer function", func(t *testing.T) {
		var out string

		printPatch := gomonkey.ApplyFunc(fmt.Fprintf, func(_ io.Writer, format string, args ...interface{}) (int, error) {
			out = fmt.Sprintf(format, args)

			return 0, nil
		})
		patch := gomonkey.ApplyFuncReturn(time.Since, time.Duration(123456789))

		defer printPatch.Reset()
		defer patch.Reset()

		innerFunctionCalls := 0
		innerFunction := func() { innerFunctionCalls++ }

		timer.ExecuteWithTimer(innerFunction)

		printPatch.Reset()
		assert.Equal(t, 1, innerFunctionCalls)
		assert.Equal(t, "Time: [0.123] seconds\n", out)
	})
}
