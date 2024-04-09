package describe

import (
	"bytes"
	monkey "github.com/agiledragon/gomonkey/v2"
	"github.com/fdaines/arch-go/old/config"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestDescribeCommand(t *testing.T) {

	t.Run("describe threshold", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		cp := 87
		cv := 34
		threshold := &config.Threshold{
			Compliance: &cp,
			Coverage:   &cv,
		}
		expectedOutput := `Threshold Rules
	* The module must comply with at least 87% of the rules described above.
	* The rules described above must cover at least 34% of the packages in this module.

`

		describeThresholdRules(threshold, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)

		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("empty thresold", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		expectedOutput := ``

		describeThresholdRules(nil, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)

		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})

	t.Run("invalid configuration", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		patch := monkey.ApplyFuncReturn(config.LoadConfig, &config.Config{}, nil)
		defer patch.Reset()
		patchExit := monkey.ApplyFunc(os.Exit, func(c int) {})
		defer patchExit.Reset()

		expectedOutput := `Invalid Configuration: configuration file should have at least one rule
`

		NewCommand(outputBuffer).Run()

		outputBytes, _ := io.ReadAll(outputBuffer)

		assert.Equal(t, expectedOutput, string(outputBytes), "Output doesn't match expected values.")
	})
}
