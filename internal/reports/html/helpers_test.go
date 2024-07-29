package html

import (
	"bytes"
	"io"
	"os"
	"testing"

	monkey "github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

func TestReportHtmlHelpers(t *testing.T) {
	t.Run("writeReport", func(t *testing.T) {
		var (
			filename     string
			fileContents string
			permissions  os.FileMode
		)

		patch := monkey.ApplyFunc(os.WriteFile, func(name string, data []byte, perm os.FileMode) error {
			filename = name
			fileContents = string(data)
			permissions = perm

			return nil
		})

		defer patch.Reset()

		outputBuffer := bytes.NewBufferString("")
		content := "foobar"
		expectedOutput := "HTML report generated at: .arch-go/report.html\n"

		writeReport(content, outputBuffer)

		outputBytes, _ := io.ReadAll(outputBuffer)
		assert.Equal(t, expectedOutput, string(outputBytes))
		assert.Equal(t, ".arch-go/report.html", filename)
		assert.Equal(t, "foobar", fileContents)
		assert.Equal(t, os.FileMode(0o600), permissions)
	})
}
