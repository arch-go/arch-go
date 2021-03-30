package utils

import (
	"github.com/fdaines/arch-go/utils/output"
	"time"
)

type command func()

func ExecuteWithTimer(fn command) {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	output.Printf("Time: %.3f seconds\n", elapsed.Seconds())
}
