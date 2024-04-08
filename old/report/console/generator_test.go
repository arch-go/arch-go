package console_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fdaines/arch-go/old/model/result"
	"github.com/fdaines/arch-go/old/report/console"
)

func TestGenerateConsoleReport(t *testing.T) {
	t.Run("Calls GenerateConsoleReport function", func(t *testing.T) {
		summary := result.NewRulesSummary()
		reportData := result.Report{
			Summary: summary,
		}
		bak := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		defer func() { os.Stdout = bak }()

		console.GenerateConsoleReport(reportData)

		w.Close()
		var buf bytes.Buffer
		io.Copy(&buf, r)

		expected := `+---+-----------------+-------+-----------+--------+
| # | RULE TYPE       | TOTAL | SUCCEEDED | FAILED |
+---+-----------------+-------+-----------+--------+
| 1 | Content Rule    |     0 |         0 |      0 |
| 2 | Dependency Rule |     0 |         0 |      0 |
| 3 | Function Rule   |     0 |         0 |      0 |
| 4 | Naming Rule     |     0 |         0 |      0 |
+---+-----------------+-------+-----------+--------+
`

		assert.Equal(t, expected, buf.String())
	})
}
