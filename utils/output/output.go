package output

import (
	"fmt"
	"github.com/fdaines/arch-go/common"
)

func PrintError(message string, err error) {
	fmt.Printf("Error: %s - %s\n", message, err.Error())
}

func Print(message string) {
	fmt.Println(message)
}

func Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func PrintVerbose(format string, a ...interface{}) {
	if common.Verbose {
		fmt.Printf(format, a...)
	}
}

func PrintStep() {
	if common.Verbose {
		fmt.Print(".")
	}
}

