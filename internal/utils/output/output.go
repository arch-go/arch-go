package output

import (
	"fmt"
	"github.com/fdaines/arch-go/internal/common"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func Print(message string) {
	fmt.Fprintf(out,message)
}

func Printf(format string, a ...interface{}) {
	fmt.Fprintf(out, format, a...)
}

func PrintVerbose(format string, a ...interface{}) {
	if common.Verbose {
		fmt.Fprintf(out, format, a...)
	}
}

