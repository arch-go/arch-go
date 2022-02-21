package output_test

import (
	"fmt"
	"io"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"

	"github.com/fdaines/arch-go/internal/common"
	"github.com/fdaines/arch-go/internal/utils/output"
)

func Test_Output(t *testing.T) {
	t.Parallel()

	t.Run("test for Print", func(t *testing.T) {
		var out string
		printPatch := monkey.Patch(fmt.Fprintf, func(w io.Writer, f string, a ...interface{}) (int, error) {
			out = fmt.Sprint(f)
			return 0, nil
		})
		defer printPatch.Unpatch()
		expected := "Hello World"

		output.Print("Hello World")

		assert.Equal(t, out, expected)
	})

	t.Run("test for Printf", func(t *testing.T) {
		var out string
		printPatch := monkey.Patch(fmt.Fprintf, func(w io.Writer, f string, a ...interface{}) (int, error) {
			out = fmt.Sprintf(f, a)
			return 0, nil
		})
		defer printPatch.Unpatch()
		expected := "Hello [World]"

		output.Printf("Hello %s", "World")

		assert.Equal(t, expected, out)
	})

	t.Run("test for PrintVerbose when verbose mode is disabled", func(t *testing.T) {
		var out string
		printPatch := monkey.Patch(fmt.Fprintf, func(w io.Writer, f string, a ...interface{}) (int, error) {
			out = fmt.Sprintf(f, a)
			return 0, nil
		})
		defer printPatch.Unpatch()
		expected := ""

		common.Verbose = false
		output.PrintVerbose("Hello %s", "World")

		assert.Equal(t, expected, out)
	})

	t.Run("test for PrintVerbose when verbose mode is enabled", func(t *testing.T) {
		var out string
		printPatch := monkey.Patch(fmt.Fprintf, func(w io.Writer, f string, a ...interface{}) (int, error) {
			out = fmt.Sprintf(f, a)
			return 0, nil
		})
		defer printPatch.Unpatch()
		expected := "Hello [World]"

		common.Verbose = true
		output.PrintVerbose("Hello %s", "World")

		assert.Equal(t, expected, out)
	})

}

/*
*/