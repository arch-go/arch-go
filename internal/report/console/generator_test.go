package console_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fdaines/arch-go/internal/model/result"
	"github.com/fdaines/arch-go/internal/report/console"
)

func TestGenerateConsoleReport(t *testing.T) {
	t.Run("Calls GenerateConsoleReport function", func(t *testing.T) {
		summary := result.NewRulesSummary()
		bak := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		defer func () { os.Stdout = bak }()

		console.GenerateConsoleReport(summary)

		w.Close()
		var buf bytes.Buffer
		io.Copy(&buf, r)

		expected := `+---+------------------+-------+-----------+--------+
| # | RULE TYPE        | TOTAL | SUCCEEDED | FAILED |
+---+------------------+-------+-----------+--------+
| 1 | ContentRule      |     0 |         0 |      0 |
| 2 | CycleRule        |     0 |         0 |      0 |
| 3 | DependenciesRule |     0 |         0 |      0 |
| 4 | FunctionsRule    |     0 |         0 |      0 |
| 5 | NamingRule       |     0 |         0 |      0 |
+---+------------------+-------+-----------+--------+
`

		assert.Equal(t, expected, buf.String())
	})
}