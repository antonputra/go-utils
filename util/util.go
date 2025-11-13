package util

import (
	"fmt"
	"log/slog"
	"math/rand/v2"
	"os"
	"slices"
	"time"

	"github.com/antonputra/go-utils/monitoring"
)

// Random generates a random number within a specified range.
func Random(min int, max int) int {
	return rand.IntN(max-min) + min
}

// Sleep pauses the execution of the program for the generated number of microsecond.
// Interval is generated randomly with the upper bound limit provided by you.
func Sleep(n int) {
	// Generate a random number where the upper bound is n.
	r := rand.IntN(n)

	// Suspend the program's execution.
	time.Sleep(time.Duration(r) * time.Microsecond)
}

// Annotate provides an additional context for the error.
func Annotate(err error, format string, args ...any) error {
	if err != nil {
		return fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), err)
	}
	return nil
}

// A Warn notifies you that something has gone wrong with the execution of the program.
func Warn(err error, format string, args ...any) {
	if err != nil {
		slog.Warn(fmt.Sprintf("%s: %s", fmt.Sprintf(format, args...), err))
	}
}

// A Fail prints the error message and then exits the program.
func Fail(err error, format string, args ...any) {
	if err != nil {
		slog.Error(fmt.Sprintf("%s: %s", fmt.Sprintf(format, args...), err))
		os.Exit(1)
	}
}

// Source for the string generator.
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// GenString generates a random string with length n
func GenString(n int) string {
	// Create a rune slice with length n.
	b := make([]rune, n)

	// Iterate over the rune slice to generate a random string.
	for i := range b {
		b[i] = letterRunes[rand.IntN(len(letterRunes))]
	}

	// Convert the runes to a string and return it.
	return string(b)
}

// ValidateStr checks if the string is allowed. Example:
//
//	ValidateStr("redis", []string{"redis", "memcached"}
func ValidateStr(key string, keys []string) {
	if !slices.Contains(keys, key) {
		slog.Error("Key is not supported", "key", key, "options", keys)
		os.Exit(1)
	}
}

// DoWork enables rate-limiting of client-provided work and accepts a random interval
// to pause between executions of the provided function, as well as a rate argument
// that defines the allowed number of operations per second.
func DoWork(work func(m *monitoring.Metrics), rate int, pauseMs int, m *monitoring.Metrics) {
	// Initialize the start time and the counter to track the number of operations.
	var start time.Time = time.Now()
	var count int = 0

	// Start an infinite loop and perform the work.
	for {
		// Keep track of the time elapsed between each operation.
		end := time.Now()
		elapsed := end.Sub(start)

		// Reset the number of operations each second.
		if elapsed >= time.Second {
			start = time.Now()
			count = 0
		}

		// If the number of operations equals or exceeds the rate, sleep for the remaining time until the next second.
		// Sleeping avoids wasting CPU cycles, allowing for more efficient use of resources.
		if count >= rate {
			next := time.Second - elapsed
			if next > time.Nanosecond {
				time.Sleep(next)
			}
		}

		// Execute the client-provided operation.
		work(m)

		// Increment the operation counter.
		count++

		// Pause execution to prevent overloading the target.
		Sleep(pauseMs)
	}
}
