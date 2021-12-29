package aoc

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Setup makes sure that we are in the current directory for calling into the aoc project.
// Simplifying the pathing. The function returned should be deferred which will result in a log of
// the task and its execution time.
func Setup(year, day int) func() {
	startTime := time.Now()

	dayName := fmt.Sprintf("day%d", day)
	path, _ := os.Getwd()

	_ = os.Chdir(filepath.Join(path, fmt.Sprintf("%d", year), dayName))

	return func() {
		fmt.Printf("AOC (%d:%d): %v\n", year, day, time.Since(startTime))
	}
}

// PrintAnswer  will simply log the answer in a consistent format.
func PrintAnswer(part int, answer interface{}) {
	fmt.Printf("Part %d: %v\n", part, answer)
}
