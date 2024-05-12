package timer

import (
	"fmt"
	"time"
)

type command func()

func ExecuteWithTimer(fn command) {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	fmt.Printf("Time: %.3f seconds\n", elapsed.Seconds())
}
