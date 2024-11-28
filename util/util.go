package util

import (
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"slices"
	"time"
)

// Random generates a random number within a specified range.
func Random(min int, max int) int {
	return rand.Intn(max-min) + min
}

// Sleep pauses the execution of the program.
// Interval is generated randomly with the upper bound limit provided by you.
func Sleep(n int) {
	// Generate a random number where the upper bound is n.
	r := rand.Intn(n)

	// Suspend the program's execution.
	time.Sleep(time.Duration(r) * time.Millisecond)
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
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
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
