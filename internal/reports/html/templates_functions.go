package html

import (
	"fmt"
	"time"
)

func toHumanTime() func(d time.Duration) string {
	return func(d time.Duration) string {
		if d.Seconds() > 0.9 {
			return fmt.Sprintf("%v [s]", d.Seconds())
		}

		if d.Milliseconds() > 0 {
			return fmt.Sprintf("%v [ms]", d.Milliseconds())
		}

		if d.Microseconds() > 0 {
			return fmt.Sprintf("%v [μs]", d.Microseconds())
		}

		return fmt.Sprintf("%v [ns]", d.Nanoseconds())
	}
}

func formatTime() func(t time.Time) string {
	return func(t time.Time) string {
		return t.Format("15:04:05")
	}
}

func formatDate() func(t time.Time) string {
	return func(t time.Time) string {
		return t.Format("2006/01/02")
	}
}

func calculateRatio() func(num int, den int) int {
	return func(num int, den int) int {
		if den == 0 {
			return 100
		}

		return 100 * num / den
	}
}

func toYesNo() func(pass bool) string {
	return func(pass bool) string {
		if pass {
			return "YES"
		}

		return "NO"
	}
}

func toPassFail() func(pass bool) string {
	return func(pass bool) string {
		if pass {
			return "PASS"
		}

		return "FAIL"
	}
}
